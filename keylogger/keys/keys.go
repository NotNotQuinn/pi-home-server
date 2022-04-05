package keys

// KeyMode represents whether this key was pressed, or released.
type KeyMode int

const (
	// The key was released
	Released KeyMode = iota
	// The key was pressed
	Pressed
)

// Key is a string representation of a key that has been pressed
//
// Only 28 key constants are defined.
type Key string

const (
	// Number keys

	NUMPAD_0 Key = "KEY_NUMPAD_0"
	NUMPAD_1 Key = "KEY_NUMPAD_1"
	NUMPAD_2 Key = "KEY_NUMPAD_2"
	NUMPAD_3 Key = "KEY_NUMPAD_3"
	NUMPAD_4 Key = "KEY_NUMPAD_4"
	NUMPAD_5 Key = "KEY_NUMPAD_5"
	NUMPAD_6 Key = "KEY_NUMPAD_6"
	NUMPAD_7 Key = "KEY_NUMPAD_7"
	NUMPAD_8 Key = "KEY_NUMPAD_8"
	NUMPAD_9 Key = "KEY_NUMPAD_9"

	// Math keys

	PLUS     Key = "KEY_PLUS"
	DASH     Key = "KEY_DASH"
	ASTERISK Key = "KEY_ASTERISK"
	SLASH    Key = "KEY_SLASH"
	EQUALS   Key = "KEY_EQUALS"

	// Arrow keys

	UP    Key = "KEY_UP"
	DOWN  Key = "KEY_DOWN"
	LEFT  Key = "KEY_LEFT"
	RIGHT Key = "KEY_RIGHT"

	// Navigation keys

	HOME     Key = "KEY_HOME"
	END      Key = "KEY_END"
	PAGEUP   Key = "KEY_PAGEUP"
	PAGEDOWN Key = "KEY_PAGEDOWN"
	INSERT   Key = "KEY_INSERT"
	DELETE   Key = "KEY_DELETE"

	// Other keys

	ENTER     Key = "KEY_ENTER"
	BACKSPACE Key = "KEY_BACKSPACE"
	DOT       Key = "KEY_DOT"
)
