package wca

import (
	"github.com/go-ole/go-ole"
)

type IMMNotificationClient struct {
	vTable   *IMMNotificationClientVtbl
	refCount uint
	callback IMMNotificationClientCallback
}

type IMMNotificationClientVtbl struct {
	ole.IUnknownVtbl

	OnDeviceStateChanged   uintptr
	OnDeviceAdded          uintptr
	OnDeviceRemoved        uintptr
	OnDefaultDeviceChanged uintptr
	OnPropertyValueChanged uintptr
}

type IMMNotificationClientCallback struct {
	OnDefaultDeviceChanged func(flow EDataFlow, role ERole, pwstrDeviceId string) error
	OnDeviceAdded          func(pwstrDeviceId string) error
	OnDeviceRemoved        func(pwstrDeviceId string) error
	OnDeviceStateChanged   func(pwstrDeviceId string, dwNewState uint32) error
	OnPropertyValueChanged func(pwstrDeviceId string, key uint64) error
}
