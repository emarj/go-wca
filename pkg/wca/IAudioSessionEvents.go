package wca

import "github.com/go-ole/go-ole"

type IAudioSessionEvents struct {
	vTable   *IAudioSessionEventsVtbl
	refCount int
	callback IAudioSessionEventsCallback
}

type IAudioSessionEventsVtbl struct {
	ole.IUnknownVtbl

	OnDisplayNameChanged   uintptr
	OnIconPathChanged      uintptr
	OnSimpleVolumeChanged  uintptr
	OnChannelVolumeChanged uintptr
	OnGroupingParamChanged uintptr
	OnStateChanged         uintptr
	OnSessionDisconnected  uintptr
}

type IAudioSessionEventsCallback struct {
	OnDisplayNameChanged   func(newDisplayName string, eventCtx *ole.GUID) error
	OnIconPathChanged      func(newIconPath string, eventCtx *ole.GUID) error
	OnSimpleVolumeChanged  func(newVolume float32, newMute bool, eventCtx *ole.GUID) error
	OnChannelVolumeChanged func(newChannelVolumeArray []float32, changedChannel uint32, eventCtx *ole.GUID) error
	OnGroupingParamChanged func(newGroupingParam *ole.GUID, eventCtx *ole.GUID) error
	OnStateChanged         func(newState uint32) error
	OnSessionDisconnected  func(reason uint32) error
}
