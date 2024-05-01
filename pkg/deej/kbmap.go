package deej

import (
	"github.com/micmonay/keybd_event"
)

var KeybindingMap = map[string]int{

	"SHIFT":  0x10 + 0xFFF,
	"CTRL":   0x11 + 0xFFF,
	"ALT":    0x12 + 0xFFF,
	"LSHIFT": 0xA0 + 0xFFF,
	"RSHIFT": 0xA1 + 0xFFF,
	"LCTRL":  0xA2 + 0xFFF,
	"RCTRL":  0xA3 + 0xFFF,
	"LWIN":   0x5B + 0xFFF,
	"RWIN":   0x5C + 0xFFF,
	"WIN":    0x5B + 0xFFF,

	"SP1":  keybd_event.VK_SP1,
	"SP2":  keybd_event.VK_SP2,
	"SP3":  keybd_event.VK_SP3,
	"SP4":  keybd_event.VK_SP4,
	"SP5":  keybd_event.VK_SP5,
	"SP6":  keybd_event.VK_SP6,
	"SP7":  keybd_event.VK_SP7,
	"SP8":  keybd_event.VK_SP8,
	"SP9":  keybd_event.VK_SP9,
	"SP10": keybd_event.VK_SP10,
	"SP11": keybd_event.VK_SP11,
	"SP12": keybd_event.VK_SP12,

	"1":   keybd_event.VK_1,
	"2":   keybd_event.VK_2,
	"3":   keybd_event.VK_3,
	"4":   keybd_event.VK_4,
	"5":   keybd_event.VK_5,
	"6":   keybd_event.VK_6,
	"7":   keybd_event.VK_7,
	"8":   keybd_event.VK_8,
	"9":   keybd_event.VK_9,
	"0":   keybd_event.VK_0,
	"A":   keybd_event.VK_A,
	"B":   keybd_event.VK_B,
	"C":   keybd_event.VK_C,
	"D":   keybd_event.VK_D,
	"E":   keybd_event.VK_E,
	"F":   keybd_event.VK_F,
	"G":   keybd_event.VK_G,
	"H":   keybd_event.VK_H,
	"I":   keybd_event.VK_I,
	"J":   keybd_event.VK_J,
	"K":   keybd_event.VK_K,
	"L":   keybd_event.VK_L,
	"M":   keybd_event.VK_M,
	"N":   keybd_event.VK_N,
	"O":   keybd_event.VK_O,
	"P":   keybd_event.VK_P,
	"Q":   keybd_event.VK_Q,
	"R":   keybd_event.VK_R,
	"S":   keybd_event.VK_S,
	"T":   keybd_event.VK_T,
	"U":   keybd_event.VK_U,
	"V":   keybd_event.VK_V,
	"W":   keybd_event.VK_W,
	"X":   keybd_event.VK_X,
	"Y":   keybd_event.VK_Y,
	"Z":   keybd_event.VK_Z,
	"ESC": keybd_event.VK_ESC,
	//f keys
	"F1":  keybd_event.VK_F1,
	"F2":  keybd_event.VK_F2,
	"F3":  keybd_event.VK_F3,
	"F4":  keybd_event.VK_F4,
	"F5":  keybd_event.VK_F5,
	"F6":  keybd_event.VK_F6,
	"F7":  keybd_event.VK_F7,
	"F8":  keybd_event.VK_F8,
	"F9":  keybd_event.VK_F9,
	"F10": keybd_event.VK_F10,
	"F11": keybd_event.VK_F11,
	"F12": keybd_event.VK_F12,
	"F13": keybd_event.VK_F13,
	"F14": keybd_event.VK_F14,
	"F15": keybd_event.VK_F15,
	"F16": keybd_event.VK_F16,
	"F17": keybd_event.VK_F17,
	"F18": keybd_event.VK_F18,
	"F19": keybd_event.VK_F19,
	"F20": keybd_event.VK_F20,
	"F21": keybd_event.VK_F21,
	"F22": keybd_event.VK_F22,
	"F23": keybd_event.VK_F23,
	"F24": keybd_event.VK_F24,

	"NUMLOCK":    keybd_event.VK_NUMLOCK,
	"SCROLLLOCK": keybd_event.VK_SCROLLLOCK,
	"RESERVED":   keybd_event.VK_RESERVED,
	"MINUS":      keybd_event.VK_MINUS, // - on main keyboard
	"EQUAL":      keybd_event.VK_EQUAL, // = on main keyboard
	"BACKSPACE":  keybd_event.VK_BACKSPACE,
	"TAB":        keybd_event.VK_TAB,
	"LEFTBRACE":  keybd_event.VK_LEFTBRACE,  // [ on main keyboard
	"RIGHTBRACE": keybd_event.VK_RIGHTBRACE, // ] on main keyboard
	"ENTER":      keybd_event.VK_ENTER,
	"SEMICOLON":  keybd_event.VK_SEMICOLON,  // ; on main keyboard
	"APOSTROPHE": keybd_event.VK_APOSTROPHE, // ' on main keyboard
	"GRAVE":      keybd_event.VK_GRAVE,      // ` on main keyboard
	"BACKSLASH":  keybd_event.VK_BACKSLASH,  // \ on main keyboard
	"COMMA":      keybd_event.VK_COMMA,      // , on main keyboard
	"DOT":        keybd_event.VK_DOT,        // . on main keyboard
	"SLASH":      keybd_event.VK_SLASH,      // / on main keyboard
	"KPASTERISK": keybd_event.VK_KPASTERISK,
	"SPACE":      keybd_event.VK_SPACE,
	"CAPSLOCK":   keybd_event.VK_CAPSLOCK,
	//special keys
	"PAGEUP":   keybd_event.VK_PAGEUP,
	"PAGEDOWN": keybd_event.VK_PAGEDOWN,
	"END":      keybd_event.VK_END,
	"HOME":     keybd_event.VK_HOME,
	"LEFT":     keybd_event.VK_LEFT,
	"UP":       keybd_event.VK_UP,
	"RIGHT":    keybd_event.VK_RIGHT,
	"DOWN":     keybd_event.VK_DOWN,
	"SNAPSHOT": keybd_event.VK_SNAPSHOT, // print screen
	"INSERT":   keybd_event.VK_INSERT,
	"DELETE":   keybd_event.VK_DELETE,
	//keypad keys
	"KP0":     keybd_event.VK_KP0,
	"KP1":     keybd_event.VK_KP1,
	"KP2":     keybd_event.VK_KP2,
	"KP3":     keybd_event.VK_KP3,
	"KP4":     keybd_event.VK_KP4,
	"KP5":     keybd_event.VK_KP5,
	"KP6":     keybd_event.VK_KP6,
	"KP7":     keybd_event.VK_KP7,
	"KP8":     keybd_event.VK_KP8,
	"KP9":     keybd_event.VK_KP9,
	"KPMINUS": keybd_event.VK_KPMINUS,
	"KPPLUS":  keybd_event.VK_KPPLUS,
	"KPDOT":   keybd_event.VK_KPDOT,

	"SCROLL":            keybd_event.VK_SCROLL, // scroll lock
	"LMENU":             keybd_event.VK_LMENU,
	"RMENU":             keybd_event.VK_RMENU,
	"BROWSER_BACK":      keybd_event.VK_BROWSER_BACK,
	"BROWSER_FORWARD":   keybd_event.VK_BROWSER_FORWARD,
	"BROWSER_REFRESH":   keybd_event.VK_BROWSER_REFRESH,
	"BROWSER_STOP":      keybd_event.VK_BROWSER_STOP,
	"BROWSER_SEARCH":    keybd_event.VK_BROWSER_SEARCH,
	"BROWSER_FAVORITES": keybd_event.VK_BROWSER_FAVORITES,
	"BROWSER_HOME":      keybd_event.VK_BROWSER_HOME,
	"VOLUME_MUTE":       keybd_event.VK_VOLUME_MUTE,
	"VOLUME_DOWN":       keybd_event.VK_VOLUME_DOWN,
	"VOLUME_UP":         keybd_event.VK_VOLUME_UP,
	"MEDIA_NEXT_TRACK":  keybd_event.VK_MEDIA_NEXT_TRACK,
	"MEDIA_PREV_TRACK":  keybd_event.VK_MEDIA_PREV_TRACK,
	"MEDIA_STOP":        keybd_event.VK_MEDIA_STOP,
	"MEDIA_PLAY_PAUSE":  keybd_event.VK_MEDIA_PLAY_PAUSE,

	"LBUTTON":    keybd_event.VK_LBUTTON,
	"RBUTTON":    keybd_event.VK_RBUTTON,
	"CANCEL":     keybd_event.VK_CANCEL,
	"MBUTTON":    keybd_event.VK_MBUTTON,
	"XBUTTON1":   keybd_event.VK_XBUTTON1,
	"XBUTTON2":   keybd_event.VK_XBUTTON2,
	"BACK":       keybd_event.VK_BACK,
	"CLEAR":      keybd_event.VK_CLEAR,
	"PAUSE":      keybd_event.VK_PAUSE,
	"CAPITAL":    keybd_event.VK_CAPITAL,
	"KANA":       keybd_event.VK_KANA,
	"HANGUEL":    keybd_event.VK_HANGUEL,
	"HANGUL":     keybd_event.VK_HANGUL,
	"JUNJA":      keybd_event.VK_JUNJA,
	"FINAL":      keybd_event.VK_FINAL,
	"HANJA":      keybd_event.VK_HANJA,
	"KANJI":      keybd_event.VK_KANJI,
	"CONVERT":    keybd_event.VK_CONVERT,
	"NONCONVERT": keybd_event.VK_NONCONVERT,
	"ACCEPT":     keybd_event.VK_ACCEPT,
	"MODECHANGE": keybd_event.VK_MODECHANGE,

	"SELECT":  keybd_event.VK_SELECT,
	"PRINT":   keybd_event.VK_PRINT,
	"EXECUTE": keybd_event.VK_EXECUTE,
	"HELP":    keybd_event.VK_HELP,

	"LAUNCH_MAIL":         keybd_event.VK_LAUNCH_MAIL,
	"LAUNCH_MEDIA_SELECT": keybd_event.VK_LAUNCH_MEDIA_SELECT,
	"LAUNCH_APP1":         keybd_event.VK_LAUNCH_APP1,
	"LAUNCH_APP2":         keybd_event.VK_LAUNCH_APP2,
	"OEM_1":               keybd_event.VK_OEM_1,
	"OEM_PLUS":            keybd_event.VK_OEM_PLUS,
	"OEM_COMMA":           keybd_event.VK_OEM_COMMA,
	"OEM_MINUS":           keybd_event.VK_OEM_MINUS,
	"OEM_PERIOD":          keybd_event.VK_OEM_PERIOD,
	"OEM_2":               keybd_event.VK_OEM_2,
	"OEM_3":               keybd_event.VK_OEM_3,
	"OEM_4":               keybd_event.VK_OEM_4,
	"OEM_5":               keybd_event.VK_OEM_5,
	"OEM_6":               keybd_event.VK_OEM_6,
	"OEM_7":               keybd_event.VK_OEM_7,
	"OEM_8":               keybd_event.VK_OEM_8,
	"OEM_102":             keybd_event.VK_OEM_102,
	"PROCESSKEY":          keybd_event.VK_PROCESSKEY,
	"PACKET":              keybd_event.VK_PACKET,
	"ATTN":                keybd_event.VK_ATTN,
	"CRSEL":               keybd_event.VK_CRSEL,
	"EXSEL":               keybd_event.VK_EXSEL,
	"EREOF":               keybd_event.VK_EREOF,
	"PLAY":                keybd_event.VK_PLAY,
	"ZOOM":                keybd_event.VK_ZOOM,
	"NONAME":              keybd_event.VK_NONAME,
	"PA1":                 keybd_event.VK_PA1,
	"OEM_CLEAR":           keybd_event.VK_OEM_CLEAR,
}