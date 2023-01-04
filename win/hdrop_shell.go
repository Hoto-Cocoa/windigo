//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to an internal drop structure.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdrop
type HDROP HANDLE

// This function is rather tricky. Prefer using HDROP.ListFilesAndFinish().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragfinish
func (hDrop HDROP) DragFinish() {
	syscall.Syscall(proc.DragFinish.Addr(), 1,
		uintptr(hDrop), 0, 0)
}

// This function is rather tricky. Prefer using HDROP.ListFilesAndFinish().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragqueryfilew
func (hDrop HDROP) DragQueryFile(
	iFile uint32, lpszFile *uint16, cch uint32) uint32 {

	ret, _, err := syscall.Syscall6(proc.DragQueryFile.Addr(), 4,
		uintptr(hDrop), uintptr(iFile), uintptr(unsafe.Pointer(lpszFile)),
		uintptr(cch), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// Returns true if dropped within client area.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragquerypoint
func (hDrop HDROP) DragQueryPoint() (POINT, bool) {
	var pt POINT
	ret, _, _ := syscall.Syscall(proc.DragQueryPoint.Addr(), 2,
		uintptr(hDrop), uintptr(unsafe.Pointer(&pt)), 0)
	return pt, ret != 0
}
