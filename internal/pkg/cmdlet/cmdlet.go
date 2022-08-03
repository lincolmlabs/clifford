/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package cmdlet

import "encoding/json"

type cmdletStruct struct {
	hash    string `help: The hash signing this command to be execute by the remote API.`
	token   string `help: The token used to authenticate the request.`
	command string `help: The command to be execute by the remote API.`
}

func CmdLet() {
	var content = cmdletStruct{}
	json.Marshal(content)
}
