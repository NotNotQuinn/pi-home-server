package keys

// Event represents a single keypress
type Event struct {
	Key               Key     `json:"key"`
	Mode              KeyMode `json:"mode"`
	KeyRepresentation string  `json:"keyRepresentation"`
}

// Pressed is true on a key press, and false on a key release.
//
// Equivalent to e.Mode == keys.Pressed
func (e *Event) Pressed() bool {
	return e.Mode == Pressed
}
