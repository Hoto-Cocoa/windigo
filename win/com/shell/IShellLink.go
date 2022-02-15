package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishelllinkw
type IShellLink struct{ win.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IShellLink.Release().
//
// Example:
//
//  lnk := shell.NewIShellLink(
//      win.CoCreateInstance(
//          shellco.CLSID_ShellLink, nil,
//          co.CLSCTX_INPROC_SERVER,
//          shellco.IID_IShellLink),
//  )
//  defer lnk.Release()
func NewIShellLink(base win.IUnknown) IShellLink {
	return IShellLink{IUnknown: base}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getarguments
func (me *IShellLink) GetArguments() string {
	buf := make([]uint16, 1024) // arbitrary
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetArguments, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return win.Str.FromNativeSlice(buf)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getdescription
func (me *IShellLink) GetDescription() string {
	buf := make([]uint16, 1024) // arbitrary
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetDescription, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return win.Str.FromNativeSlice(buf)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-geticonlocation
func (me *IShellLink) GetIconLocation() (path string, index int32) {
	buf := make([]uint16, 256) // arbitrary
	iconIndex := int32(0)

	ret, _, _ := syscall.Syscall6(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetIconLocation, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1),
		uintptr(unsafe.Pointer(&iconIndex)), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return win.Str.FromNativeSlice(buf), iconIndex
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getpath
func (me *IShellLink) GetPath(
	fd *win.WIN32_FIND_DATA, flags shellco.SLGP) string {

	buf := make([]uint16, 256) // arbitrary
	ret, _, _ := syscall.Syscall6(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetPath, 5,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1),
		uintptr(unsafe.Pointer(fd)), uintptr(flags), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return win.Str.FromNativeSlice(buf)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getshowcmd
func (me *IShellLink) GetShowCmd() co.SW {
	cmd := co.SW(0)
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetShowCmd, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&cmd)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return cmd
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getworkingdirectory
func (me *IShellLink) GetWorkingDirectory() string {
	buf := make([]uint16, 256) // arbitrary
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetWorkingDirectory, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return win.Str.FromNativeSlice(buf)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-resolve
func (me *IShellLink) Resolve(hWnd win.HWND, flags shellco.SLR) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).Resolve, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(flags))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setarguments
func (me *IShellLink) SetArguments(args string) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetArguments, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(args))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setdescription
func (me *IShellLink) SetDescription(descr string) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetDescription, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(descr))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-seticonlocation
func (me *IShellLink) SetIconLocation(path string, index int32) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetDescription, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(path))),
		uintptr(index))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setpath
func (me *IShellLink) SetPath(path string) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetPath, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(path))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setrelativepath
func (me *IShellLink) SetRelativePath(path string) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetRelativePath, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(path))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setshowcmd
func (me *IShellLink) SetShowCmd(cmd co.SW) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetShowCmd, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(cmd), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setworkingdirectory
func (me *IShellLink) SetWorkingDirectory(path string) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetWorkingDirectory, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(path))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
