// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func NewIAudioSessionEvents(callback IAudioSessionEventsCallback) *IAudioSessionEvents {
	vTable := &IAudioSessionEventsVtbl{}

	// IUnknown methods
	vTable.QueryInterface = syscall.NewCallback(aseQueryInterface)
	vTable.AddRef = syscall.NewCallback(aseAddRef)
	vTable.Release = syscall.NewCallback(aseRelease)

	// IMMNotificationClient methods
	vTable.OnDisplayNameChanged = syscall.NewCallback(aseOnDisplayNameChanged)
	vTable.OnIconPathChanged = syscall.NewCallback(aseOnIconPathChanged)
	vTable.OnSimpleVolumeChanged = syscall.NewCallback(aseOnSimpleVolumeChanged)
	vTable.OnChannelVolumeChanged = syscall.NewCallback(aseOnChannelVolumeChanged)
	vTable.OnGroupingParamChanged = syscall.NewCallback(aseOnGroupingParamChanged)
	vTable.OnStateChanged = syscall.NewCallback(aseOnStateChanged)
	vTable.OnSessionDisconnected = syscall.NewCallback(aseOnSessionDisconnected)

	ase := &IAudioSessionEvents{}

	ase.vTable = vTable
	ase.callback = callback

	return ase
}

func aseQueryInterface(this uintptr, riid *ole.GUID, ppInterface *uintptr) int64 {
	*ppInterface = 0

	if ole.IsEqualGUID(riid, ole.IID_IUnknown) ||
		ole.IsEqualGUID(riid, IID_IMMNotificationClient) {
		aseAddRef(this)
		*ppInterface = this

		return ole.S_OK
	}

	return ole.E_NOINTERFACE
}

func aseAddRef(this uintptr) int64 {
	ase := (*IMMNotificationClient)(unsafe.Pointer(this))

	ase.refCount += 1

	return int64(ase.refCount)
}

func aseRelease(this uintptr) int64 {
	ase := (*IMMNotificationClient)(unsafe.Pointer(this))

	ase.refCount -= 1

	return int64(ase.refCount)
}

func aseOnDisplayNameChanged(this uintptr, newDisplayNamePtr uintptr, eventContext *ole.GUID) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnDisplayNameChanged == nil {
		return ole.S_OK
	}

	// device := syscall.UTF16ToString(*(*[]uint16)(unsafe.Pointer(pwstrDeviceId)))
	newDisplayName := LPCWSTRToString(newDisplayNamePtr, 1024)

	err := ase.callback.OnDisplayNameChanged(newDisplayName, eventContext)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK

}
func aseOnIconPathChanged(this uintptr, newIconPathPtr uintptr, eventCtx *ole.GUID) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnIconPathChanged == nil {
		return ole.S_OK
	}

	// device := syscall.UTF16ToString(*(*[]uint16)(unsafe.Pointer(pwstrDeviceId)))
	newIconPath := LPCWSTRToString(newIconPathPtr, 1024)

	err := ase.callback.OnIconPathChanged(newIconPath, eventCtx)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK

}
func aseOnSimpleVolumeChanged(this uintptr, newVolume float32, newMute bool, eventCtx *ole.GUID) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnSimpleVolumeChanged == nil {
		return ole.S_OK
	}

	err := ase.callback.OnSimpleVolumeChanged(newVolume, newMute, eventCtx)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK

}

func aseOnChannelVolumeChanged(this uintptr, channelCount uint32, newChannelVolumeArrayPtr uintptr, changedChannel uint32, eventCtx *ole.GUID) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnChannelVolumeChanged == nil {
		return ole.S_OK
	}

	newChannelVolumeArray := pointerToSliceFloat32(newChannelVolumeArrayPtr, channelCount)

	err := ase.callback.OnChannelVolumeChanged(newChannelVolumeArray, changedChannel, eventCtx)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK

}

func aseOnGroupingParamChanged(this uintptr, newGroupingParam *ole.GUID, eventCtx *ole.GUID) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnGroupingParamChanged == nil {
		return ole.S_OK
	}

	err := ase.callback.OnGroupingParamChanged(newGroupingParam, eventCtx)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func aseOnStateChanged(this uintptr, newState uint32) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnStateChanged == nil {
		return ole.S_OK
	}

	err := ase.callback.OnStateChanged(newState)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK

}
func aseOnSessionDisconnected(this uintptr, reason uint32) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnSessionDisconnected == nil {
		return ole.S_OK
	}

	err := ase.callback.OnSessionDisconnected(reason)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK

}

func pointerToSliceFloat32(ptr uintptr, len uint32) []float32 {

	result := []float32{}
	if ptr == 0 || len == 0 {
		return result
	}

	for i := uint32(0); i < len; i++ {
		u := *(*float32)(unsafe.Pointer(ptr + uintptr(i)))

		if u == 0 {
			break
		}

		result = append(result, u)
	}

	return result
}
