/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"wingows/api"
)

// Modal popup window.
// Allows message and notification handling.
type WindowModal struct {
	windowBase
	prevFocus api.HWND // child control last focused on parent
	setup     windowModalSetup
}

// Parameters that will be used to create the window.
func (me *WindowModal) Setup() *windowModalSetup {
	me.setup.initOnce() // guard
	return &me.setup
}

// Creates the modal window and disables the parent. This function will return
// only after the modal is closed.
func (me *WindowModal) Show(parent Window) {
	me.setup.initOnce() // guard
	hInst := parent.Hwnd().GetInstance()
	me.windowBase.registerClass(me.setup.genWndClassEx(hInst))

	me.windowBase.OnMsg().WmClose(func() { // default WM_CLOSE handling
		me.windowBase.Hwnd().GetWindow(api.GW_OWNER).EnableWindow(true) // re-enable parent
		me.windowBase.Hwnd().DestroyWindow()                            // then destroy modal
		me.prevFocus.SetFocus()
	})

	me.prevFocus = api.GetFocus()     // currently focused control
	parent.Hwnd().EnableWindow(false) // https://devblogs.microsoft.com/oldnewthing/20040227-00/?p=40463

	_, _, cx, cy := multiplyByDpi(0, 0, me.setup.Width, me.setup.Height)

	me.windowBase.createWindow("WindowModal", me.setup.ExStyle,
		me.setup.ClassName, me.setup.Title, me.setup.Style,
		0, 0, // initially anchored at zero
		cx, cy, parent, api.HMENU(0), hInst)

	rc := me.windowBase.Hwnd().GetWindowRect()
	rcParent := parent.Hwnd().GetWindowRect() // both rc relative to screen

	me.windowBase.Hwnd().SetWindowPos(api.SWP_HWND(0), // center modal over parent (warning: happens after WM_CREATE processing)
		rcParent.Left+(rcParent.Right-rcParent.Left)/2-(rc.Right-rc.Left)/2,
		rcParent.Top+(rcParent.Bottom-rcParent.Top)/2-(rc.Bottom-rc.Top)/2,
		0, 0, api.SWP_NOZORDER|api.SWP_NOSIZE)
}

//------------------------------------------------------------------------------

type windowModalSetup struct {
	wasInit bool // default to false

	ClassName        string      // Optional, defaults to a hash generated by WNDCLASSEX parameters. Passed to RegisterClassEx.
	ClassStyle       api.CS      // Window class style, passed to RegisterClassEx.
	HCursor          api.HCURSOR // Window cursor, passed to RegisterClassEx.
	HBrushBackground api.HBRUSH  // Window background brush, passed to RegisterClassEx.

	Title   string    // The title of the window, passed to CreateWindowEx.
	Width   uint32    // Initial width of the window, passed to CreateWindowEx.
	Height  uint32    // Initial height of the window, passed to CreateWindowEx.
	Style   api.WS    // Window style, passed to CreateWindowEx.
	ExStyle api.WS_EX // Window extended style, passed to CreateWindowEx.
}

func (me *windowModalSetup) initOnce() {
	if !me.wasInit {
		me.wasInit = true

		me.ClassStyle = api.CS_DBLCLKS

		me.Width = 500 // arbitrary dimensions
		me.Height = 400
		me.Style = api.WS_CAPTION | api.WS_SYSMENU | api.WS_CLIPCHILDREN | api.WS_BORDER | api.WS_VISIBLE
		me.ExStyle = api.WS_EX(0)
	}
}

func (me *windowModalSetup) genWndClassEx(hInst api.HINSTANCE) *api.WNDCLASSEX {
	wcx := api.WNDCLASSEX{}

	wcx.CbSize = uint32(unsafe.Sizeof(wcx))
	wcx.HInstance = hInst
	wcx.Style = me.ClassStyle

	if me.HCursor != 0 {
		wcx.HCursor = me.HCursor
	} else {
		wcx.HCursor = api.HINSTANCE(0).LoadCursor(api.IDC_ARROW)
	}

	if me.HBrushBackground != 0 {
		wcx.HbrBackground = me.HBrushBackground
	} else {
		wcx.HbrBackground = api.CreateSysColorBrush(api.COLOR_BTNFACE)
	}

	if me.ClassName == "" {
		me.ClassName = wcx.Hash() // generate hash after all other fields are set
	}
	wcx.LpszClassName = api.StrToUtf16Ptr(me.ClassName)

	return &wcx
}
