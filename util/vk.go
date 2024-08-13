package util

const (
	// ALPHABET
	VK_A = 0x41
	VK_B = 0x42
	VK_C = 0x43
	VK_D = 0x44
	VK_E = 0x45
	VK_F = 0x46
	VK_G = 0x47
	VK_H = 0x48
	VK_I = 0x49
	VK_J = 0x4A
	VK_K = 0x4B
	VK_L = 0x4C
	VK_M = 0x4D
	VK_N = 0x4E
	VK_O = 0x4F
	VK_P = 0x50
	VK_Q = 0x51
	VK_R = 0x52
	VK_S = 0x53
	VK_T = 0x54
	VK_U = 0x55
	VK_V = 0x56
	VK_W = 0x57
	VK_X = 0x58
	VK_Y = 0x59
	VK_Z = 0x5A

	// NUMBERS
	VK_0 = 0x30
	VK_1 = 0x31
	VK_2 = 0x32
	VK_3 = 0x33
	VK_4 = 0x34
	VK_5 = 0x35
	VK_6 = 0x36
	VK_7 = 0x37
	VK_8 = 0x38
	VK_9 = 0x39

	// FUNCTION KEYS
	VK_F1  = 0x70
	VK_F2  = 0x71
	VK_F3  = 0x72
	VK_F4  = 0x73
	VK_F5  = 0x74
	VK_F6  = 0x75
	VK_F7  = 0x76
	VK_F8  = 0x77
	VK_F9  = 0x78
	VK_F10 = 0x79
	VK_F11 = 0x7A
	VK_F12 = 0x7B

	// NUMPAD KEYS
	VK_NUMLOCK  = 0x90
	VK_DIVIDE   = 0x6F
	VK_MULTIPLY = 0x6A
	VK_SUBTRACT = 0x6D
	VK_ADD      = 0x6B
	VK_DECIMAL  = 0x6E
	VK_NUMPAD0  = 0x60
	VK_NUMPAD1  = 0x61
	VK_NUMPAD2  = 0x62
	VK_NUMPAD3  = 0x63
	VK_NUMPAD4  = 0x64
	VK_NUMPAD5  = 0x65
	VK_NUMPAD6  = 0x66
	VK_NUMPAD7  = 0x67
	VK_NUMPAD8  = 0x68
	VK_NUMPAD9  = 0x69

	// SYMBOLS
	VK_TILDE     = 0xC0
	VK_MINUS     = 0xBD
	VK_EQUAL     = 0xBB
	VK_LBRACKET  = 0xDB
	VK_RBRACKET  = 0xDD
	VK_BACKSLASH = 0xDC
	VK_SEMICOLON = 0xBA
	VK_QUOTE     = 0xDE
	VK_COMMA     = 0xBC
	VK_PERIOD    = 0xBE
	VK_SLASH     = 0xBF
	VK_BACKTICK  = 0xC0
	VK_BACKSPACE = 0x08
	VK_TAB       = 0x09
	VK_CAPSLOCK  = 0x14
	VK_SHIFT     = 0x10
	VK_CTRL      = 0x11
	VK_ALT       = 0x12
	VK_SPACE     = 0x20
	VK_ENTER     = 0x0D
	VK_ESCAPE    = 0x1B
	VK_PGUP      = 0x21
	VK_PGDN      = 0x22
	VK_CAPS      = 0x14
	VK_DEL       = 0x2E
	VK_INS       = 0x2D
	VK_HOME      = 0x24
	VK_END       = 0x23
	VK_BACK      = 0x08
	VK_MENU      = 0x5D
	VK_LWIN      = 0x5B
	VK_RWIN      = 0x5C
	VK_RETURN    = 0x0D
)

var KeyMap = map[string]uintptr{
	"0": VK_0,
	"1": VK_1,
	"2": VK_2,
	"3": VK_3,
	"4": VK_4,
	"5": VK_5,
	"6": VK_6,
	"7": VK_7,
	"8": VK_8,
	"9": VK_9,

	"a": VK_A,
	"b": VK_B,
	"c": VK_C,
	"d": VK_D,
	"e": VK_E,
	"f": VK_F,
	"g": VK_G,
	"h": VK_H,
	"i": VK_I,
	"j": VK_J,
	"k": VK_K,
	"l": VK_L,
	"m": VK_M,
	"n": VK_N,
	"o": VK_O,
	"p": VK_P,
	"q": VK_Q,
	"r": VK_R,
	"s": VK_S,
	"t": VK_T,
	"u": VK_U,
	"v": VK_V,
	"w": VK_W,
	"x": VK_X,
	"y": VK_Y,
	"z": VK_Z,

	"f1":  VK_F1,
	"f2":  VK_F2,
	"f3":  VK_F3,
	"f4":  VK_F4,
	"f5":  VK_F5,
	"f6":  VK_F6,
	"f7":  VK_F7,
	"f8":  VK_F8,
	"f9":  VK_F9,
	"f10": VK_F10,
	"f11": VK_F11,
	"f12": VK_F12,

	"numlock":  VK_NUMLOCK,
	"divide":   VK_DIVIDE,
	"multiply": VK_MULTIPLY,
	"subtract": VK_SUBTRACT,
	"add":      VK_ADD,
	"decimal":  VK_DECIMAL,
	"numpad0":  VK_NUMPAD0,
	"numpad1":  VK_NUMPAD1,
	"numpad2":  VK_NUMPAD2,
	"numpad3":  VK_NUMPAD3,
	"numpad4":  VK_NUMPAD4,
	"numpad5":  VK_NUMPAD5,
	"numpad6":  VK_NUMPAD6,
	"numpad7":  VK_NUMPAD7,
	"numpad8":  VK_NUMPAD8,
	"numpad9":  VK_NUMPAD9,

	"tilde":     VK_TILDE,
	"minus":     VK_MINUS,
	"equal":     VK_EQUAL,
	"lbracket":  VK_LBRACKET,
	"rbracket":  VK_RBRACKET,
	"backslash": VK_BACKSLASH,
	"semicolon": VK_SEMICOLON,
	"quote":     VK_QUOTE,
	"comma":     VK_COMMA,
	"period":    VK_PERIOD,
	"slash":     VK_SLASH,
	"backtick":  VK_BACKTICK,
	"backspace": VK_BACKSPACE,
	"tab":       VK_TAB,
	"capslock":  VK_CAPSLOCK,
	"shift":     VK_SHIFT,
	"ctrl":      VK_CTRL,
	"alt":       VK_ALT,
	"space":     VK_SPACE,
	"enter":     VK_ENTER,
	"escape":    VK_ESCAPE,
	"pgup":      VK_PGUP,
	"pgdn":      VK_PGDN,
	"caps":      VK_CAPS,
	"del":       VK_DEL,
	"ins":       VK_INS,
	"home":      VK_HOME,
	"end":       VK_END,
	"back":      VK_BACK,
	"menu":      VK_MENU,
	"lwin":      VK_LWIN,
	"rwin":      VK_RWIN,
	"return":    VK_RETURN,
}
