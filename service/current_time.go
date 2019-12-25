// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeCurrentTime = "670899F9"

type CurrentTime struct {
	*Service

	CurrentTime *characteristic.CurrentTime
}

func NewCurrentTime() *CurrentTime {
	svc := CurrentTime{}
	svc.Service = New(TypeCurrentTime)

	svc.CurrentTime = characteristic.NewCurrentTime()
	svc.AddCharacteristic(svc.CurrentTime.Characteristic)

	return &svc
}
