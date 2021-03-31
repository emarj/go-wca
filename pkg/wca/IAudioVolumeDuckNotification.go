package wca

type IAudioVolumeDuckNotification struct {
	VTable *IAudioVolumeDuckNotificationVtbl
}

type IAudioVolumeDuckNotificationVtbl struct {
	QueryInterface             uintptr
	AddRef                     uintptr
	Release                    uintptr
	OnVolumeDuckNotification   uintptr
	OnVolumeUnduckNotification uintptr
}
