package climacell

type unit int

const (
	Si unit = iota
	Us
)

// String returns the string value of the unit
func (u unit) String() string {
	return [...]string{"si", "us"}[u]
}
