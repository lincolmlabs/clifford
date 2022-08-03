/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package main

import (
	"testing"
)

func TestMyId(t *testing.T) {

	id := myId()
	length := len([]rune(id))

	if length < 64 {
		t.Errorf("Invalid sha256sum: %q", id)
	}
}
