package castles

type Season int32

const (
	Summer = iota
	Autumn
	Winter
	Spring
)

func (s Season) String() string {
	strings := []string{"Summer", "Autumn", "Winter", "Fall"}

	if s < Summer || s > Spring {
		return "Unknown"
	}

	return strings[s]
}
