// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeTelevision = "D8"

type Television struct {
	*Service

	Active             *characteristic.Active
	ActiveIdentifier   *characteristic.ActiveIdentifier
	ConfiguredName     *characteristic.ConfiguredName
	SleepDiscoveryMode *characteristic.SleepDiscoveryMode
}

func NewTelevision() *Television {
	svc := Television{}
	svc.Service = New(TypeTelevision)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.ActiveIdentifier = characteristic.NewActiveIdentifier()
	svc.AddCharacteristic(svc.ActiveIdentifier.Characteristic)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddCharacteristic(svc.ConfiguredName.Characteristic)

	svc.SleepDiscoveryMode = characteristic.NewSleepDiscoveryMode()
	svc.AddCharacteristic(svc.SleepDiscoveryMode.Characteristic)

	return &svc
}
