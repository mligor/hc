// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCurrentTime = "9B"

type CurrentTime struct {
	*String
}

func NewCurrentTime() *CurrentTime {
	char := NewString(TypeCurrentTime)
	char.Format = FormatString
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue("")

	return &CurrentTime{char}
}
