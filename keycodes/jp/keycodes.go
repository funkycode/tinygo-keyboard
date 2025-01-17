package jp

import (
	"github.com/sago35/tinygo-keyboard/keycodes"
)

// for Japanese Keyboard
// based on machine/usb/hid/keyboard/keycode.go
const (
	KeyA           = 0xF000 | 0x04
	KeyB           = 0xF000 | 0x05
	KeyC           = 0xF000 | 0x06
	KeyD           = 0xF000 | 0x07
	KeyE           = 0xF000 | 0x08
	KeyF           = 0xF000 | 0x09
	KeyG           = 0xF000 | 0x0A
	KeyH           = 0xF000 | 0x0B
	KeyI           = 0xF000 | 0x0C
	KeyJ           = 0xF000 | 0x0D
	KeyK           = 0xF000 | 0x0E
	KeyL           = 0xF000 | 0x0F
	KeyM           = 0xF000 | 0x10
	KeyN           = 0xF000 | 0x11
	KeyO           = 0xF000 | 0x12
	KeyP           = 0xF000 | 0x13
	KeyQ           = 0xF000 | 0x14
	KeyR           = 0xF000 | 0x15
	KeyS           = 0xF000 | 0x16
	KeyT           = 0xF000 | 0x17
	KeyU           = 0xF000 | 0x18
	KeyV           = 0xF000 | 0x19
	KeyW           = 0xF000 | 0x1A
	KeyX           = 0xF000 | 0x1B
	KeyY           = 0xF000 | 0x1C
	KeyZ           = 0xF000 | 0x1D
	Key1           = 0xF000 | 0x1E
	Key2           = 0xF000 | 0x1F
	Key3           = 0xF000 | 0x20
	Key4           = 0xF000 | 0x21
	Key5           = 0xF000 | 0x22
	Key6           = 0xF000 | 0x23
	Key7           = 0xF000 | 0x24
	Key8           = 0xF000 | 0x25
	Key9           = 0xF000 | 0x26
	Key0           = 0xF000 | 0x27
	KeyEnter       = 0xF000 | 0x28
	KeyEsc         = 0xF000 | 0x29
	KeyBackspace   = 0xF000 | 0x2A
	KeyTab         = 0xF000 | 0x2B
	KeySpace       = 0xF000 | 0x2C
	KeyMinus       = 0xF000 | 0x2D
	KeyHat         = 0xF000 | 0x2E
	KeyAt          = 0xF000 | 0x2F
	KeyLeftBrace   = 0xF000 | 0x30
	KeyRightBrace  = 0xF000 | 0x32
	KeySemicolon   = 0xF000 | 0x33
	KeyColon       = 0xF000 | 0x34
	KeyHankaku     = 0xF000 | 0x35
	KeyComma       = 0xF000 | 0x36
	KeyPeriod      = 0xF000 | 0x37
	KeySlash       = 0xF000 | 0x38
	KeyCapsLock    = 0xF000 | 0x39
	KeyF1          = 0xF000 | 0x3A
	KeyF2          = 0xF000 | 0x3B
	KeyF3          = 0xF000 | 0x3C
	KeyF4          = 0xF000 | 0x3D
	KeyF5          = 0xF000 | 0x3E
	KeyF6          = 0xF000 | 0x3F
	KeyF7          = 0xF000 | 0x40
	KeyF8          = 0xF000 | 0x41
	KeyF9          = 0xF000 | 0x42
	KeyF10         = 0xF000 | 0x43
	KeyF11         = 0xF000 | 0x44
	KeyF12         = 0xF000 | 0x45
	KeyPrintscreen = 0xF000 | 0x46
	KeyScrollLock  = 0xF000 | 0x47
	KeyPause       = 0xF000 | 0x48
	KeyInsert      = 0xF000 | 0x49
	KeyHome        = 0xF000 | 0x4A
	KeyPageUp      = 0xF000 | 0x4B
	KeyDelete      = 0xF000 | 0x4C
	KeyEnd         = 0xF000 | 0x4D
	KeyPageDown    = 0xF000 | 0x4E
	KeyRight       = 0xF000 | 0x4F
	KeyLeft        = 0xF000 | 0x50
	KeyDown        = 0xF000 | 0x51
	KeyUp          = 0xF000 | 0x52
	KeyNumLock     = 0xF000 | 0x53
	KeypadSlash    = 0xF000 | 0x54
	KeypadAsterisk = 0xF000 | 0x55
	KeypadMinus    = 0xF000 | 0x56
	KeypadPlus     = 0xF000 | 0x57
	KeypadEnter    = 0xF000 | 0x58
	Keypad1        = 0xF000 | 0x59
	Keypad2        = 0xF000 | 0x5A
	Keypad3        = 0xF000 | 0x5B
	Keypad4        = 0xF000 | 0x5C
	Keypad5        = 0xF000 | 0x5D
	Keypad6        = 0xF000 | 0x5E
	Keypad7        = 0xF000 | 0x5F
	Keypad8        = 0xF000 | 0x60
	Keypad9        = 0xF000 | 0x61
	Keypad0        = 0xF000 | 0x62
	KeypadPeriod   = 0xF000 | 0x63
	KeyNonUSBS     = 0xF000 | 0x64
	KeyMenu        = 0xF000 | 0x65
	KeyF13         = 0xF000 | 0x68
	KeyF14         = 0xF000 | 0x69
	KeyF15         = 0xF000 | 0x6A
	KeyF16         = 0xF000 | 0x6B
	KeyF17         = 0xF000 | 0x6C
	KeyF18         = 0xF000 | 0x6D
	KeyF19         = 0xF000 | 0x6E
	KeyF20         = 0xF000 | 0x6F
	KeyF21         = 0xF000 | 0x70
	KeyF22         = 0xF000 | 0x71
	KeyF23         = 0xF000 | 0x72
	KeyF24         = 0xF000 | 0x73
	KeyBackslash   = 0xF000 | 0x87 // \ |
	KeyHiragana    = 0xF000 | 0x88
	KeyBackslash2  = 0xF000 | 0x89 // \ _
	KeyHenkan      = 0xF000 | 0x8A
	KeyMuhenkan    = 0xF000 | 0x8B
	KeyKana        = 0xF000 | 0x90
	KeyEisu        = 0xF000 | 0x91
	KeyLeftCtrl    = 0xF000 | 0xE0
	KeyLeftShift   = 0xF000 | 0xE1
	KeyLeftAlt     = 0xF000 | 0xE2
	KeyWindows     = 0xF000 | 0xE3
	KeyRightCtrl   = 0xF000 | 0xE4
	KeyRightShift  = 0xF000 | 0xE5
)

const (
	KeyMediaPlay        = 0xB0 | 0xE400
	KeyMediaPause       = 0xB1 | 0xE400
	KeyMediaRecord      = 0xB2 | 0xE400
	KeyMediaFastForward = 0xB3 | 0xE400
	KeyMediaRewind      = 0xB4 | 0xE400
	KeyMediaNextTrack   = 0xB5 | 0xE400
	KeyMediaPrevTrack   = 0xB6 | 0xE400
	KeyMediaStop        = 0xB7 | 0xE400
	KeyMediaEject       = 0xB8 | 0xE400
	KeyMediaRandomPlay  = 0xB9 | 0xE400
	KeyMediaPlayPause   = 0xCD | 0xE400
	KeyMediaPlaySkip    = 0xCE | 0xE400
	KeyMediaMute        = 0xE2 | 0xE400
	KeyMediaVolumeInc   = 0xE9 | 0xE400
	KeyMediaVolumeDec   = 0xEA | 0xE400
)

const (
	MouseLeft    = 0xD000 | 0x01 // mouse.Left
	MouseRight   = 0xD000 | 0x02 // mouse.Right
	MouseMiddle  = 0xD000 | 0x04 // mouse.Middle
	MouseBack    = 0xD000 | 0x08 // mouse.Back
	MouseForward = 0xD000 | 0x10 // mouse.Forward
	WheelDown    = 0xD000 | 0x20
	WheelUp      = 0xD000 | 0x40
)

const (
	KeyMod0 = keycodes.KeyMod0
	KeyMod1 = keycodes.KeyMod1
	KeyMod2 = keycodes.KeyMod2
	KeyMod3 = keycodes.KeyMod3
	KeyMod4 = keycodes.KeyMod4
	KeyMod5 = keycodes.KeyMod5

	KeyTo0 = keycodes.KeyTo0
	KeyTo1 = keycodes.KeyTo1
	KeyTo2 = keycodes.KeyTo2
	KeyTo3 = keycodes.KeyTo3
	KeyTo4 = keycodes.KeyTo4
	KeyTo5 = keycodes.KeyTo5
)
