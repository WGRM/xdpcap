// Code generated by "stringer -type=XDPAction"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[XDPAborted-0]
	_ = x[XDPDrop-1]
	_ = x[XDPPass-2]
	_ = x[XDPTx-3]
}

const _XDPAction_name = "XDPAbortedXDPDropXDPPassXDPTx"

var _XDPAction_index = [...]uint8{0, 10, 17, 24, 29}

func (i XDPAction) String() string {
	if i < 0 || i >= XDPAction(len(_XDPAction_index)-1) {
		return "XDPAction(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _XDPAction_name[_XDPAction_index[i]:_XDPAction_index[i+1]]
}
