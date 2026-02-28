//go:build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	openClipboard    = user32.NewProc("OpenClipboard")
	emptyClipboard   = user32.NewProc("EmptyClipboard")
	setClipboardData = user32.NewProc("SetClipboardData")
	closeClipboard   = user32.NewProc("CloseClipboard")
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	globalAlloc      = kernel32.NewProc("GlobalAlloc")
	globalLock       = kernel32.NewProc("GlobalLock")
	globalUnlock     = kernel32.NewProc("GlobalUnlock")
	lstrcpy          = kernel32.NewProc("lstrcpyW")
)

func (a *App) writeNativeClipboard(text string) error {
	utf16, _ := syscall.UTF16PtrFromString(text)
	size := uintptr(len(text)*2 + 2)

	hMem, _, _ := globalAlloc.Call(0x0002, size)
	ptr, _, _ := globalLock.Call(hMem)
	lstrcpy.Call(ptr, uintptr(unsafe.Pointer(utf16)))
	globalUnlock.Call(hMem)

	openClipboard.Call(0)
	defer closeClipboard.Call()
	emptyClipboard.Call()

	res, _, _ := setClipboardData.Call(13, hMem)
	if res == 0 {
		return fmt.Errorf("falha ao definir clipboard")
	}
	return nil
}
