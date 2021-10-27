// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func pcvSetDefaultEndpoint(pcv *IPolicyConfigVista, deviceID string, eRole ERole) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
