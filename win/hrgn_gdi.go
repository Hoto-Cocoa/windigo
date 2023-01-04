//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a region.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hrgn
type HRGN HGDIOBJ

// ⚠️ You must defer HRGN.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createellipticrgn
func CreateEllipticRgn(boundTopLeft, boundBottomRight POINT) HRGN {
	ret, _, err := syscall.Syscall6(proc.CreateEllipticRgn.Addr(), 4,
		uintptr(boundTopLeft.X), uintptr(boundTopLeft.Y),
		uintptr(boundBottomRight.X), uintptr(boundBottomRight.Y),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// ⚠️ You must defer HRGN.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createrectrgnindirect
func CreateRectRgnIndirect(rc *RECT) HRGN {
	ret, _, err := syscall.Syscall(proc.CreateRectRgnIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(rc)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// ⚠️ You must defer HRGN.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createroundrectrgn
func CreateRoundRectRgn(topLeft, bottomRight POINT, szEllipse SIZE) HRGN {
	ret, _, err := syscall.Syscall6(proc.CreateRoundRectRgn.Addr(), 6,
		uintptr(topLeft.X), uintptr(topLeft.Y),
		uintptr(bottomRight.X), uintptr(bottomRight.Y),
		uintptr(szEllipse.Cx), uintptr(szEllipse.Cy))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hRgn HRGN) DeleteObject() {
	HGDIOBJ(hRgn).DeleteObject()
}

// Combines the two regions and stores the result in current region.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-combinergn
func (hRgn HRGN) CombineRgn(hrgnSrc1, hrgnSrc2 HRGN, mode co.RGN) co.REGION {
	ret, _, err := syscall.Syscall6(proc.CombineRgn.Addr(), 4,
		uintptr(hRgn), uintptr(hrgnSrc1), uintptr(hrgnSrc2), uintptr(mode), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-offsetrgn
func (hRgn HRGN) OffsetRgn(x, y int32) co.REGION {
	ret, _, err := syscall.Syscall(proc.OffsetRgn.Addr(), 3,
		uintptr(hRgn), uintptr(x), uintptr(y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}
