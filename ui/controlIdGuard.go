/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"wingows/api"
)

var (
	globalBaseCtrlId = api.ID(1000) // arbitrary, taken from Visual Studio resource editor
)

// Encapsulates the control ID and, if not initialized, uses an auto-incremented value.
type controlIdGuard struct {
	ctrlId api.ID // defaults to zero
}

// Optional; returns a ctrlIdGuard with a custom control ID, not using the
// default auto-incremented one.
func makeCtrlIdGuard(initialId api.ID) controlIdGuard {
	return controlIdGuard{
		ctrlId: initialId, // properly initialized
	}
}

// Returns the ID of this child window control.
func (me *controlIdGuard) CtrlId() api.ID {
	if me.ctrlId == 0 { // not initialized yet?
		globalBaseCtrlId++ // increments sequential global ID
		me.ctrlId = globalBaseCtrlId
	}
	return me.ctrlId
}
