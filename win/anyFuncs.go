/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

// Returns *uint16.
// Wrapper to syscall.UTF16PtrFromString().
// Panics on error.
func StrToPtr(s string) *uint16 {
	// We won't return an uintptr right away because it has no pointer semantics,
	// it's just a number, so pointed memory can be garbage-collected.
	// https://stackoverflow.com/a/51188315
	pstr, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		panic(fmt.Sprintf("StrToPtr failed \"%s\": %s",
			s, err))
	}
	return pstr
}

// Returns a null-terminated []uint16.
// Wrapper to syscall.UTF16FromString().
// Panics on error.
func StrToSlice(s string) []uint16 {
	sli, err := syscall.UTF16FromString(s)
	if err != nil {
		panic(fmt.Sprintf("StrToSlice failed \"%s\": %s",
			s, err))
	}
	return sli
}

// Returns *uint16, or nil of empty string.
// Wrapper to syscall.UTF16PtrFromString().
// Panics on error.
func StrToPtrBlankIsNil(s string) *uint16 {
	if s != "" {
		return StrToPtr(s)
	}
	return nil
}

//------------------------------------------------------------------------------

func EnumWindows(
	lpEnumFunc func(hwnd HWND, lParam LPARAM) bool,
	lParam LPARAM) {

	ret, _, lerr := syscall.Syscall(proc.EnumWindows.Addr(), 2,
		syscall.NewCallback(
			func(hwnd HWND, lParam LPARAM) int32 {
				return boolToInt32(lpEnumFunc(hwnd, lParam))
			}),
		uintptr(lParam), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("EnumWindow failed."))
	}
}

func InitCommonControls() {
	syscall.Syscall(proc.InitCommonControls.Addr(), 0, 0, 0, 0)
}

func GetAsyncKeyState(virtKeyCode co.VK) uint16 {
	ret, _, _ := syscall.Syscall(proc.GetAsyncKeyState.Addr(), 1,
		uintptr(virtKeyCode), 0, 0)
	return uint16(ret)
}

func GetCurrentThreadId() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetCurrentThreadId.Addr(), 0,
		0, 0, 0)
	return uint32(ret)
}

func GetCursorPos() *POINT {
	pt := &POINT{}
	ret, _, lerr := syscall.Syscall(proc.GetCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(pt)), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("GetCursorPos failed."))
	}
	return pt
}

// Available in Windows 10, version 1607.
func GetDpiForSystem() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetDpiForSystem.Addr(), 0,
		0, 0, 0)
	return uint32(ret)
}

func GetSystemMetrics(index co.SM) int32 {
	ret, _, _ := syscall.Syscall(proc.GetSystemMetrics.Addr(), 1,
		uintptr(index), 0, 0)
	return int32(ret)
}

// Multiplies two 32-bit values and then divides the 64-bit result by a third
// 32-bit value. The final result is rounded to the nearest integer.
func MulDiv(number, numerator, denominator int32) int32 {
	ret, _, _ := syscall.Syscall(proc.MulDiv.Addr(), 3,
		uintptr(number), uintptr(numerator), uintptr(denominator))
	return int32(ret)
}

func PostQuitMessage(exitCode int32) {
	syscall.Syscall(proc.PostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
}

func PostThreadMessage(
	idThread uint32, Msg co.WM, wParam WPARAM, lParam LPARAM) {

	ret, _, lerr := syscall.Syscall6(proc.PostThreadMessage.Addr(), 4,
		uintptr(idThread), uintptr(Msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("PostThreadMessage failed."))
	}
}

func RegisterWindowMessage(lpString string) uint32 {
	ret, _, lerr := syscall.Syscall(proc.RegisterWindowMessage.Addr(), 1,
		uintptr(unsafe.Pointer(StrToPtr(lpString))), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("RegisterWindowMessage failed."))
	}
	return uint32(ret)
}

func ReplyMessage(lResult uintptr) bool {
	ret, _, _ := syscall.Syscall(proc.ReplyMessage.Addr(), 1,
		lResult, 0, 0)
	return ret != 0
}

// Available in Windows 10, version 1703.
func SetProcessDpiAwarenessContext(value co.DPI_AWARE_CTX) {
	ret, _, lerr := syscall.Syscall(proc.SetProcessDpiAwarenessContext.Addr(), 1,
		uintptr(value), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("SetProcessDpiAwarenessContext failed."))
	}
}

// Available in Windows Vista.
func SetProcessDPIAware() {
	ret, _, _ := syscall.Syscall(proc.SetProcessDPIAware.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic("SetProcessDPIAware failed.")
	}
}

func SetWindowsHookEx(idHook co.WH,
	lpfn func(code int, wp WPARAM, lp LPARAM) uintptr,
	hmod HINSTANCE, dwThreadId uint32) HHOOK {

	ret, _, lerr := syscall.Syscall6(proc.SetWindowsHookEx.Addr(), 4,
		uintptr(idHook), syscall.NewCallback(lpfn),
		uintptr(hmod), uintptr(dwThreadId), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("SetWindowsHookEx failed."))
	}
	return HHOOK(ret)
}

// Depends of CoInitializeEx().
func SHGetFileInfo(pszPath string, dwFileAttributes co.FILE_ATTRIBUTE,
	uFlags co.SHGFI) *SHFILEINFO {

	shfi := &SHFILEINFO{}
	ret, _, _ := syscall.Syscall6(proc.SHGetFileInfo.Addr(), 5,
		uintptr(unsafe.Pointer(StrToPtr(pszPath))),
		uintptr(dwFileAttributes), uintptr(unsafe.Pointer(shfi)),
		unsafe.Sizeof(*shfi), uintptr(uFlags), 0)

	if (uFlags&co.SHGFI_EXETYPE) == 0 || (uFlags&co.SHGFI_SYSICONINDEX) == 0 {
		if ret == 0 {
			panic("SHGetFileInfo failed.")
		}
	}

	if (uFlags & co.SHGFI_EXETYPE) != 0 {
		if ret == 0 {
			panic("SHGetFileInfo failed.")
		}
	}

	return shfi
}

func Sleep(dwMilliseconds uint32) {
	syscall.Syscall(proc.Sleep.Addr(), 1,
		uintptr(dwMilliseconds), 0, 0)
}

func SystemParametersInfo(uiAction co.SPI, uiParam uint32,
	pvParam unsafe.Pointer, fWinIni uint32) {

	ret, _, lerr := syscall.Syscall6(proc.SystemParametersInfo.Addr(), 4,
		uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni),
		0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("SystemParametersInfo failed."))
	}
}
