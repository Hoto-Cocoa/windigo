/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"windigo/co"
	proc "windigo/win/internal"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdc
type HDC HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatibledc
func (hdc HDC) CreateCompatibleDC() HDC {
	ret, _, lerr := syscall.Syscall(proc.CreateCompatibleDC.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "CreateCompatibleDC"))
	}
	return HDC(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawicon
func (hdc HDC) DrawIcon(x, y int32, hIcon HICON) {
	ret, _, lerr := syscall.Syscall6(proc.DrawIcon.Addr(), 4,
		uintptr(hdc), uintptr(x), uintptr(y), uintptr(hIcon), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "DrawIcon"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func (hdc HDC) EnumDisplayMonitors(
	lprcClip *RECT,
	lpfnEnum func(hMon HMONITOR, hdcMon HDC, rcMon uintptr, lp LPARAM) bool,
	dwData LPARAM) error {

	ret, _, lerr := syscall.Syscall6(proc.EnumDisplayMonitors.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(lprcClip)),
		syscall.NewCallback(
			func(hMon HMONITOR, hdcMon HDC, rcMon uintptr, lp LPARAM) uintptr {
				return boolToUintptr(lpfnEnum(hMon, hdcMon, rcMon, lp))
			}),
		0, 0, 0)
	if ret == 0 {
		return NewWinError(co.ERROR(lerr), "EnumDisplayMonitors")
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deletedc
func (hdc HDC) DeleteDC() {
	if hdc != 0 {
		syscall.Syscall(proc.DeleteDC.Addr(), 1,
			uintptr(hdc), 0, 0)
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillrgn
func (hdc HDC) FillRgn(hrgn HRGN, hbr HBRUSH) {
	ret, _, lerr := syscall.Syscall(proc.FillRgn.Addr(), 3,
		uintptr(hdc), uintptr(hrgn), uintptr(hbr))
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "FillRgn"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-framergn
func (hdc HDC) FrameRgn(hrgn HRGN, hbr HBRUSH, w, h int32) {
	ret, _, lerr := syscall.Syscall6(proc.FrameRgn.Addr(), 5,
		uintptr(hdc), uintptr(hrgn), uintptr(hbr), uintptr(w), uintptr(h),
		0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "FrameRgn"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdevicecaps
func (hdc HDC) GetDeviceCaps(index co.GDC) int32 {
	ret, _, _ := syscall.Syscall(proc.GetDeviceCaps.Addr(), 2,
		uintptr(hdc), uintptr(index), 0)
	return int32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getpolyfillmode
func (hdc HDC) GetPolyFillMode() co.POLYF {
	ret, _, lerr := syscall.Syscall(proc.GetPolyFillMode.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "GetPolyFillMode"))
	}
	return co.POLYF(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextextentpoint32w
func (hdc HDC) GetTextExtentPoint32(lpString string) SIZE {
	sz := SIZE{}
	ret, _, lerr := syscall.Syscall6(proc.GetTextExtentPoint32.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpString))),
		uintptr(len(lpString)), uintptr(unsafe.Pointer(&sz)), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "GetTextExtentPoint32"))
	}
	return sz
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextfacew
func (hdc HDC) GetTextFace() string {
	buf := [_LF_FACESIZE]uint16{}
	ret, _, lerr := syscall.Syscall(proc.GetTextFace.Addr(), 3,
		uintptr(hdc), uintptr(len(buf)), uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "GetTextFace"))
	}
	return Str.FromUint16Slice(buf[:])
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-invertrgn
func (hdc HDC) InvertRgn(hrgn HRGN) {
	ret, _, lerr := syscall.Syscall(proc.InvertRgn.Addr(), 2,
		uintptr(hdc), uintptr(hrgn), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "InvertRgn"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lineto
func (hdc HDC) LineTo(x, y int32) {
	ret, _, lerr := syscall.Syscall(proc.LineTo.Addr(), 3,
		uintptr(hdc), uintptr(x), uintptr(y))
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "LineTo"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-paintrgn
func (hdc HDC) PaintRgn(hrgn HRGN) {
	ret, _, lerr := syscall.Syscall(proc.PaintRgn.Addr(), 2,
		uintptr(hdc), uintptr(hrgn), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "PaintRgn"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polydraw
func (hdc HDC) PolyDraw(apt []POINT, aj []co.PT) {
	if len(apt) != len(aj) {
		panic(fmt.Sprintf("PolyDraw different slice sizes: %d, %d.",
			len(apt), len(aj)))
	}
	ret, _, lerr := syscall.Syscall6(proc.PolyDraw.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])),
		uintptr(unsafe.Pointer(&aj[0])), uintptr(len(apt)),
		0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr),
			fmt.Sprintf("PolyDraw, %d points", len(apt))))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polygon
func (hdc HDC) Polygon(apt []POINT) {
	ret, _, lerr := syscall.Syscall(proc.Polygon.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])), uintptr(len(apt)))
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr),
			fmt.Sprintf("Polygon, %d points", len(apt))))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polyline
func (hdc HDC) Polyline(apt []POINT) {
	ret, _, lerr := syscall.Syscall(proc.Polyline.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])), uintptr(len(apt)))
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr),
			fmt.Sprintf("Polyline, %d points", len(apt))))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polylineto
func (hdc HDC) PolylineTo(apt []POINT) {
	ret, _, lerr := syscall.Syscall(proc.PolylineTo.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(&apt[0])), uintptr(len(apt)))
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr),
			fmt.Sprintf("PolylineTo failed for %d points.", len(apt))))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-restoredc
func (hdc HDC) RestoreDC(nSavedDC int32) {
	syscall.Syscall(proc.RestoreDC.Addr(), 2,
		uintptr(hdc), uintptr(nSavedDC), 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-savedc
func (hdc HDC) SaveDC() int32 {
	ret, _, lerr := syscall.Syscall(proc.SaveDC.Addr(), 1,
		uintptr(hdc), 0, 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SaveDC"))
	}
	return int32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpolyfillmode
func (hdc HDC) SetPolyFillMode(iMode co.POLYF) co.POLYF {
	ret, _, lerr := syscall.Syscall(proc.SetPolyFillMode.Addr(), 2,
		uintptr(hdc), uintptr(iMode), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SelectPolyFillMode"))
	}
	return co.POLYF(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBitmap(b HBITMAP) HBITMAP {
	ret, _, lerr := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(b), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SelectObject, HBITMAP"))
	}
	return HBITMAP(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBrush(b HBRUSH) HBRUSH {
	ret, _, lerr := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(b), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SelectObject, HBRUSH"))
	}
	return HBRUSH(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectFont(f HFONT) HFONT {
	ret, _, lerr := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(f), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SelectObject, HFONT"))
	}
	return HFONT(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectPen(p HPEN) HPEN {
	ret, _, lerr := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(p), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SelectObject, HPEN"))
	}
	return HPEN(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectRgn(r HRGN) co.REGION {
	ret, _, lerr := syscall.Syscall(proc.SelectObject.Addr(), 2,
		uintptr(hdc), uintptr(r), 0)
	if ret == _HGDI_ERROR {
		panic(NewWinError(co.ERROR(lerr), "SelectObject, HRGN"))
	}
	return co.REGION(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkcolor
func (hdc HDC) SetBkColor(color COLORREF) COLORREF {
	ret, _, lerr := syscall.Syscall(proc.SetBkColor.Addr(), 2,
		uintptr(hdc), uintptr(color), 0)
	if ret == _CLR_INVALID {
		panic(NewWinError(co.ERROR(lerr), "SetBkColor"))
	}
	return COLORREF(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
func (hdc HDC) SetBkMode(mode co.BKMODE) co.BKMODE {
	ret, _, lerr := syscall.Syscall(proc.SetBkMode.Addr(), 2,
		uintptr(hdc), uintptr(mode), 0)
	if ret == 0 {
		panic(NewWinError(co.ERROR(lerr), "SetBkMode"))
	}
	return co.BKMODE(ret)
}
