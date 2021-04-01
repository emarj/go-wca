package wca

import (
	"github.com/go-ole/go-ole"
)

type IAudioEndpointVolumeCallback struct {
	VTable *IAudioEndpointVolumeCallbackVtbl
}

type IAudioEndpointVolumeCallbackVtbl struct {
	ole.IUnknownVtbl
	OnNotify uintptr
}
