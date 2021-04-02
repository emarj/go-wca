// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func NewIAudioSessionNotification(callback IAudioSessionNotificationCallback) *IAudioSessionNotification {
	vTable := &IAudioSessionNotificationVtbl{}

	// IUnknown methods
	vTable.QueryInterface = syscall.NewCallback(asnQueryInterface)
	vTable.AddRef = syscall.NewCallback(asnAddRef)
	vTable.Release = syscall.NewCallback(asnRelease)

	// IAudioSessionNotification methods
	vTable.OnSessionCreated = syscall.NewCallback(asnOnSessionCreated)

	asn := &IAudioSessionNotification{}

	asn.vTable = vTable
	asn.callback = callback

	return asn
}

func asnQueryInterface(this uintptr, riid *ole.GUID, ppInterface *uintptr) int64 {
	*ppInterface = 0

	if ole.IsEqualGUID(riid, ole.IID_IUnknown) ||
		ole.IsEqualGUID(riid, IID_IAudioSessionNotification) {
		asnAddRef(this)
		*ppInterface = this

		return ole.S_OK
	}

	return ole.E_NOINTERFACE
}

func asnAddRef(this uintptr) int64 {
	asn := (*IAudioSessionNotification)(unsafe.Pointer(this))

	asn.refCount += 1

	return int64(asn.refCount)
}

func asnRelease(this uintptr) int64 {
	asn := (*IAudioSessionNotification)(unsafe.Pointer(this))

	asn.refCount -= 1

	return int64(asn.refCount)
}

func asnOnSessionCreated(this uintptr, audioSessionControlPtr uintptr) int64 {
	asn := (*IAudioSessionNotification)(unsafe.Pointer(this))

	if asn.callback.OnSessionCreated == nil {
		return ole.S_OK
	}

	audioSessionControl := (*IAudioSessionControl)(unsafe.Pointer(audioSessionControlPtr))

	err := asn.callback.OnSessionCreated(audioSessionControl)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}
