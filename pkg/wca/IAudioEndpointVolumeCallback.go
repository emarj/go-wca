package wca

import (
	"github.com/go-ole/go-ole"
)

type IAudioEndpointVolumeCallbackVtbl struct {
	ole.IUnknownVtbl
	OnNotify uintptr
}

type IAudioEndpointVolumeCallback struct {
	vTable   *IAudioEndpointVolumeCallbackVtbl
	refCount int
	callback func(fMasterVolume float32, bMuted bool, nChannels uint32, afChannelVolumes [1]float32, guidEventContext ole.GUID) error
}

type AudioVolumeNotificationData struct {
	guidEventContext ole.GUID
	bMuted           bool
	fMasterVolume    float32
	nChannels        uint32
	afChannelVolumes [1]float32
}
