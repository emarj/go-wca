package wca

import "github.com/go-ole/go-ole"

type IAudioSessionNotification struct {
	vTable   *IAudioSessionNotificationVtbl
	refCount uint
	callback IAudioSessionNotificationCallback
}

type IAudioSessionNotificationVtbl struct {
	ole.IUnknownVtbl
	OnSessionCreated uintptr
}

type IAudioSessionNotificationCallback struct {
	OnSessionCreated func(*IAudioSessionControl) error
}
