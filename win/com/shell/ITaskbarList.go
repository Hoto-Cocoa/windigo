package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _ITaskbarListVtbl struct {
	win.IUnknownVtbl
	HrInit       uintptr
	AddTab       uintptr
	DeleteTab    uintptr
	ActivateTab  uintptr
	SetActiveAlt uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
type ITaskbarList struct {
	win.IUnknown // Base IUnknown.
}

// Calls CoCreateInstance(), typically with CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func NewITaskbarList(dwClsContext co.CLSCTX) ITaskbarList {
	iUnk := win.CoCreateInstance(
		shellco.CLSID_TaskbarList, nil, dwClsContext,
		shellco.IID_ITaskbarList)
	return ITaskbarList{IUnknown: iUnk}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
func (me *ITaskbarList) ActivateTab(hwnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).ActivateTab, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-addtab
func (me *ITaskbarList) AddTab(hwnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).AddTab, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-deletetab
func (me *ITaskbarList) DeleteTab(hwnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).DeleteTab, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-hrinit
func (me *ITaskbarList) HrInit() {
	syscall.Syscall(
		(*_ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).HrInit, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
func (me *ITaskbarList) SetActiveAlt(hwnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).SetActiveAlt, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
}
