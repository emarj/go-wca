// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func pcvSetDefaultEndpoint(pcv *IPolicyConfigVista, deviceID string, eRole ERole) (err error) {
	deviceIDPtr, err := syscall.UTF16PtrFromString(deviceID)
	if err != nil {
		return
	}
	hr, _, _ := syscall.Syscall(
		pcv.VTable().SetDefaultEndpoint,
		3,
		uintptr(unsafe.Pointer(pcv)),
		uintptr(unsafe.Pointer(deviceIDPtr)),
		uintptr(uint32(eRole)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
