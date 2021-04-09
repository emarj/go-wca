package wca

import (
	"github.com/go-ole/go-ole"
)

type IAudioEndpointVolumeCallback struct {
	vTable   *IAudioEndpointVolumeCallbackVtbl
	refCount uint
	callback func(fMasterVolume float32, bMuted bool, nChannels uint32, afChannelVolumes [1]float32, guidEventContext ole.GUID) error
}

type IAudioEndpointVolumeCallbackVtbl struct {
	ole.IUnknownVtbl

	OnNotify uintptr
}

type AudioVolumeNotificationData struct {
	guidEventContext ole.GUID
	bMuted           bool
	fMasterVolume    float32
	nChannels        uint32
	afChannelVolumes [1]float32
}
