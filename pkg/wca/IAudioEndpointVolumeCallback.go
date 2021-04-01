package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioEndpointVolumeCallback struct {
	ole.IUnknown
}

type IAudioEndpointVolumeCallbackVtbl struct {
	ole.IUnknownVtbl
	OnNotify uintptr
}

func (v *IAudioEndpointVolumeCallback) VTable() *IAudioEndpointVolumeCallbackVtbl {
	return (*IAudioEndpointVolumeCallbackVtbl)(unsafe.Pointer(v.RawVTable))
}
