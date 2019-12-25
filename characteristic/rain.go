// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeRain = "56E59150"

type Rain struct {
	*Bool
}

func NewRain() *Rain {
	char := NewBool(TypeRain)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(false)

	return &Rain{char}
}
