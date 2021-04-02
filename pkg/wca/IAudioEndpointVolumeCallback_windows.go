// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func NewIAudioEndpointVolumeCallback(callback func(fMasterVolume float32, bMuted bool, nChannels uint32, afChannelVolumes [1]float32, guidEventContext ole.GUID) error) *IAudioEndpointVolumeCallback {
	vTable := &IAudioEndpointVolumeCallbackVtbl{}

	// IUnknown methods
	vTable.QueryInterface = syscall.NewCallback(aevcQueryInterface)
	vTable.AddRef = syscall.NewCallback(aevcAddRef)
	vTable.Release = syscall.NewCallback(aevcRelease)

	// IAudioSessionNotification methods
	vTable.OnNotify = syscall.NewCallback(aevcOnNotify)

	aevc := &IAudioEndpointVolumeCallback{}

	aevc.vTable = vTable
	aevc.callback = callback

	return aevc
}

func aevcQueryInterface(this uintptr, riid *ole.GUID, ppInterface *uintptr) int64 {
	*ppInterface = 0

	if ole.IsEqualGUID(riid, ole.IID_IUnknown) ||
		ole.IsEqualGUID(riid, IID_IAudioSessionNotification) {
		aevcAddRef(this)
		*ppInterface = this

		return ole.S_OK
	}

	return ole.E_NOINTERFACE
}

func aevcAddRef(this uintptr) int64 {
	aevc := (*IAudioSessionNotification)(unsafe.Pointer(this))

	aevc.refCount += 1

	return int64(aevc.refCount)
}

func aevcRelease(this uintptr) int64 {
	aevc := (*IAudioSessionNotification)(unsafe.Pointer(this))

	aevc.refCount -= 1

	return int64(aevc.refCount)
}

func aevcOnNotify(this uintptr, pNotifyPtr uintptr) int64 {
	aevc := (*IAudioEndpointVolumeCallback)(unsafe.Pointer(this))

	if aevc.callback == nil {
		return ole.S_OK
	}

	pNotify := (*AudioVolumeNotificationData)(unsafe.Pointer(pNotifyPtr))

	err := aevc.callback(pNotify.fMasterVolume, pNotify.bMuted, pNotify.nChannels, pNotify.afChannelVolumes, pNotify.guidEventContext)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}
