//go:build tinygo

package keyboard

import (
	"machine"
)

type SquaredMatrixKeyboard struct {
	State    []State
	Keys     [][]Keycode
	callback Callback

	Pins         []machine.Pin
	cycleCounter []uint8
}

const squaredMatrixCyclesToPreventChattering = uint8(4)

func (d *Device) AddSquaredMatrixKeyboard(pins []machine.Pin, keys [][]Keycode) *SquaredMatrixKeyboard {
	state := make([]State, len(pins)*(len(pins)-1))
	cycleCnt := make([]uint8, len(state))

	for i := range pins {
		pins[i].Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}

	keydef := make([][]Keycode, LayerCount)
	for l := 0; l < len(keydef); l++ {
		keydef[l] = make([]Keycode, len(state))
	}
	for l := 0; l < len(keys); l++ {
		for kc := 0; kc < len(keys[l]); kc++ {
			keydef[l][kc] = keys[l][kc]
		}
	}

	k := &SquaredMatrixKeyboard{
		Pins:         pins,
		State:        state,
		Keys:         keydef,
		callback:     func(layer, index int, state State) {},
		cycleCounter: cycleCnt,
	}

	d.kb = append(d.kb, k)
	return k
}

func (d *SquaredMatrixKeyboard) SetCallback(fn Callback) {
	d.callback = fn
}

func (d *SquaredMatrixKeyboard) Get() []State {
	c := int(0)
	cols := []int{}
	for i := range d.Pins {
		for j := range d.Pins {
			d.Pins[j].Configure(machine.PinConfig{Mode: machine.PinInputPullup})
			if i == j {
				c = j
			} else {
				cols = append(cols, j)
			}
		}

		for r, j := range cols {
			d.Pins[j].Configure(machine.PinConfig{Mode: machine.PinOutput})
			d.Pins[j].Low()
			current := !d.Pins[c].Get()
			idx := r*(len(cols)+1) + c

			switch d.State[idx] {
			case None:
				if current {
					if d.cycleCounter[idx] >= squaredMatrixCyclesToPreventChattering {
						d.State[idx] = NoneToPress
						d.cycleCounter[idx] = 0
					} else {
						d.cycleCounter[idx]++
					}
				} else {
					d.cycleCounter[idx] = 0
				}
			case NoneToPress:
				if current {
					d.State[idx] = Press
					d.callback(0, idx, Press)
				} else {
					d.State[idx] = PressToRelease
					d.callback(0, idx, Press)
					d.callback(0, idx, PressToRelease)
				}
			case Press:
				if current {
					d.cycleCounter[idx] = 0
				} else {
					if d.cycleCounter[idx] >= squaredMatrixCyclesToPreventChattering {
						d.State[idx] = PressToRelease
						d.callback(0, idx, PressToRelease)
						d.cycleCounter[idx] = 0
					} else {
						d.cycleCounter[idx]++
					}
				}
			case PressToRelease:
				if current {
					d.State[idx] = NoneToPress
					d.callback(0, idx, Press)
				} else {
					d.State[idx] = None
				}
			}
			d.Pins[j].Configure(machine.PinConfig{Mode: machine.PinInputPullup})
		}
		cols = cols[:0]
	}

	return d.State
}

func (d *SquaredMatrixKeyboard) Key(layer, index int) Keycode {
	if layer >= LayerCount {
		return 0
	}
	if index >= len(d.Keys[layer]) {
		return 0
	}
	return d.Keys[layer][index]
}

func (d *SquaredMatrixKeyboard) SetKeycode(layer, index int, key Keycode) {
	if layer >= LayerCount {
		return
	}
	if index >= len(d.Keys[layer]) {
		return
	}
	d.Keys[layer][index] = key
}

func (d *SquaredMatrixKeyboard) GetKeyCount() int {
	return len(d.State)
}

func (d *SquaredMatrixKeyboard) Init() error {
	return nil
}
