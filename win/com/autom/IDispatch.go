package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/autom/automvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// IDispatch COM interface.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
type IDispatch struct{ win.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IDispatch.Release().
func NewIDispatch(base win.IUnknown) IDispatch {
	return IDispatch{IUnknown: base}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-getidsofnames
func (me *IDispatch) GetIDsOfNames(
	lcid win.LCID, names []string) ([]MEMBERID, error) {

	memberIds := make([]MEMBERID, len(names))
	oleStrs := make([]*uint16, 0, len(names))
	for _, name := range names {
		oleStrs = append(oleStrs, win.Str.ToNativePtr(name))
	}

	ret, _, _ := syscall.Syscall6(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).GetIDsOfNames, 6,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.GuidFromIid(co.IID_NULL))),
		uintptr(unsafe.Pointer(&oleStrs[0])), uintptr(len(names)),
		uintptr(lcid), uintptr(unsafe.Pointer(&memberIds[0])))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return memberIds, nil
	} else if hr == errco.DISP_E_UNKNOWNNAME || hr == errco.DISP_E_UNKNOWNLCID {
		return nil, hr
	} else {
		panic(hr)
	}
}

// ⚠️ You must defer ITypeInfo.Release().
//
// Example:
//
//  var iDisp autom.IDispatch // initialized somewhere
//
//  tyInfo := iDisp.GetTypeInfo(win.LCID_SYSTEM_DEFAULT)
//  defer tyInfo.Release()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfo
func (me *IDispatch) GetTypeInfo(lcid win.LCID) ITypeInfo {
	var ppQueried win.IUnknown
	ret, _, _ := syscall.Syscall6(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).GetTypeInfo, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, uintptr(lcid),
		uintptr(unsafe.Pointer(&ppQueried)), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewITypeInfo(ppQueried)
	} else {
		panic(hr)
	}
}

// If the object provides type information, this number is 1; otherwise the
// number is 0.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfocount
func (me *IDispatch) GetTypeInfoCount() int {
	var pctInfo uint32
	ret, _, _ := syscall.Syscall(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).GetTypeInfoCount, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pctInfo)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(pctInfo)
	} else {
		panic(hr)
	}
}
