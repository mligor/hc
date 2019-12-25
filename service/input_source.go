// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeInputSource = "D9"

type InputSource struct {
	*Service

	ConfiguredName         *characteristic.ConfiguredName
	InputSourceType        *characteristic.InputSourceType
	IsConfigured           *characteristic.IsConfigured
	CurrentVisibilityState *characteristic.CurrentVisibilityState
}

func NewInputSource() *InputSource {
	svc := InputSource{}
	svc.Service = New(TypeInputSource)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddCharacteristic(svc.ConfiguredName.Characteristic)

	svc.InputSourceType = characteristic.NewInputSourceType()
	svc.AddCharacteristic(svc.InputSourceType.Characteristic)

	svc.IsConfigured = characteristic.NewIsConfigured()
	svc.AddCharacteristic(svc.IsConfigured.Characteristic)

	svc.CurrentVisibilityState = characteristic.NewCurrentVisibilityState()
	svc.AddCharacteristic(svc.CurrentVisibilityState.Characteristic)

	return &svc
}
