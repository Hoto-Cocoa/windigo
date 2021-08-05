package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IBaseFilterVtbl struct {
	_IMediaFilterVtbl
	EnumPins        uintptr
	FindPin         uintptr
	QueryFilterInfo uintptr
	JoinFilterGraph uintptr
	QueryVendorInfo uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ibasefilter
type IBaseFilter struct {
	IMediaFilter // Base IMediaFilter > IPersist > IUnknown.
}

// Calls CoCreateInstance(), typically with CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/medfound/using-the-directshow-evr-filter
func NewEnhancedVideoRenderer(dwClsContext co.CLSCTX) IBaseFilter {
	iUnk := win.CoCreateInstance(
		dshowco.CLSID_EnhancedVideoRenderer, nil, dwClsContext,
		dshowco.IID_IBaseFilter)
	return IBaseFilter{
		IMediaFilter{
			IPersist{IUnknown: iUnk},
		},
	}
}

// Calls CoCreateInstance(), typically with CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/directshow/video-mixing-renderer-filter-9
func NewVideoMixingRenderer9(dwClsContext co.CLSCTX) IBaseFilter {
	iUnk := win.CoCreateInstance(
		dshowco.CLSID_VideoMixingRenderer9, nil, dwClsContext,
		dshowco.IID_IBaseFilter)
	return IBaseFilter{
		IMediaFilter{
			IPersist{IUnknown: iUnk},
		},
	}
}

// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-enumpins
func (me *IBaseFilter) EnumPins() IEnumPins {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IBaseFilterVtbl)(unsafe.Pointer(*me.Ppv)).EnumPins, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
	return IEnumPins{
		win.IUnknown{Ppv: ppQueried},
	}
}

// ⚠️ You must defer Release() if true.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-findpin
func (me *IBaseFilter) FindPin(id string) (IPin, bool) {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IBaseFilterVtbl)(unsafe.Pointer(*me.Ppv)).FindPin, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(id))),
		uintptr(unsafe.Pointer(&ppQueried)))

	if err := errco.ERROR(ret); err == errco.VFW_E_NOT_FOUND {
		return IPin{}, false
	} else if err == errco.S_OK {
		return IPin{
			win.IUnknown{Ppv: ppQueried},
		}, true
	} else {
		panic(err)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-joinfiltergraph
func (me *IBaseFilter) JoinFilterGraph(
	pGraph *IFilterGraph, pName string) error {

	ret, _, _ := syscall.Syscall(
		(*_IBaseFilterVtbl)(unsafe.Pointer(*me.Ppv)).JoinFilterGraph, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pGraph.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(pName))))

	if err := errco.ERROR(ret); err == errco.S_OK {
		return nil
	} else {
		return err
	}
}

// Calls IUnknown.QueryInterface() to return IMFGetService.
//
// ⚠️ You must defer Release().
func (me *IBaseFilter) QueryIMFGetService() IMFGetService {
	iUnk := me.QueryInterface(dshowco.IID_IMFGetService)
	return IMFGetService{IUnknown: iUnk}
}
