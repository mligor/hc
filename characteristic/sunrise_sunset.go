// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	SunriseSunsetDay     int = 0
	SunriseSunsetSunrise int = 1
	SunriseSunsetSunset  int = 2
	SunriseSunsetNight   int = 3
)

const TypeSunriseSunset = "DCD72283"

type SunriseSunset struct {
	*Int
}

func NewSunriseSunset() *SunriseSunset {
	char := NewInt(TypeSunriseSunset)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &SunriseSunset{char}
}
