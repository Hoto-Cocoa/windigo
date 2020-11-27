/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package shell

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type (
	// ITaskbarList > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
	ITaskbarList struct{ win.IUnknown }

	ITaskbarListVtbl struct {
		win.IUnknownVtbl
		HrInit       uintptr
		AddTab       uintptr
		DeleteTab    uintptr
		ActivateTab  uintptr
		SetActiveAlt uintptr
	}
)

// Typically uses CLSCTX_INPROC_SERVER.
//
// You must defer Release().
func CoCreateITaskbarList(dwClsContext co.CLSCTX) *ITaskbarList {
	iUnk, err := win.CoCreateInstance(
		win.NewGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a, 0x006097c9a090), // CLSID_TaskbarList
		nil,
		dwClsContext,
		win.NewGuid(0x56fdf342, 0xfd6d, 0x11d0, 0x958a, 0x006097c9a090), // IID_ITaskbarList
	)
	if err != nil {
		panic(err)
	}
	return &ITaskbarList{
		IUnknown: *iUnk,
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-hrinit
func (me *ITaskbarList) HrInit() {
	syscall.Syscall(
		(*ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).HrInit, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-addtab
func (me *ITaskbarList) AddTab(hwnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).AddTab, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.AddTab"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-deletetab
func (me *ITaskbarList) DeleteTab(hwnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).DeleteTab, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.DeleteTab"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
func (me *ITaskbarList) ActivateTab(hwnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).ActivateTab, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.ActivateTab"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
func (me *ITaskbarList) SetActiveAlt(hwnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).SetActiveAlt, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.SetActiveAlt"))
	}
}
