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
