package wca

import "github.com/go-ole/go-ole"

type IAudioSessionNotificationVtbl struct {
	ole.IUnknownVtbl
	OnSessionCreated uintptr
}

type IAudioSessionNotificationCallback struct {
	OnSessionCreated func(*IAudioSessionControl) error
}

type IAudioSessionNotification struct {
	vTable   *IAudioSessionNotificationVtbl
	refCount int
	callback IAudioSessionNotificationCallback
}
