package d2d1

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1vt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1rendertarget
type ID2D1RenderTarget struct{ ID2D1Resource }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ID2D1RenderTarget.Release().
func NewID2D1RenderTarget(base com.IUnknown) ID2D1RenderTarget {
	return ID2D1RenderTarget{ID2D1Resource: NewID2D1Resource(base)}
}

// ⚠️ You must defer ID2D1RenderTarget.EndDraw().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-begindraw
func (me *ID2D1RenderTarget) BeginDraw() {
	ret, _, _ := syscall.Syscall(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).BeginDraw, 1,
		uintptr(unsafe.Pointer(me.Ptr())), 0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-enddraw
func (me *ID2D1RenderTarget) EndDraw() (tag1, tag2 uint64) {
	ret, _, _ := syscall.Syscall(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).EndDraw, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&tag1)), uintptr(unsafe.Pointer(&tag2)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	} else {
		return
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-flush
func (me *ID2D1RenderTarget) Flush() (tag1, tag2 uint64) {
	ret, _, _ := syscall.Syscall(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).Flush, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&tag1)), uintptr(unsafe.Pointer(&tag2)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	} else {
		return
	}
}
