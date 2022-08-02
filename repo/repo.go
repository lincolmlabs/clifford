package repo

import (
	"bytes"
	"fmt"

	clet "internal/cmdlet"

	"github.com/google/uuid"
)

func Git(name string, branch string, readme bool) {
	var buffer bytes.Buffer
	var id = uuid.New()

	var command = clet.CmdLet()

	var gitFolder = fmt.Sprintf("/tmp/%s/%s", id.String(), name)

	buffer.WriteString(fmt.Sprintf("rm -Rf %s\n", gitFolder))
	buffer.WriteString(fmt.Sprintf("mkdir %s\ncd %s\n", gitFolder, gitFolder))
	buffer.WriteString(fmt.Sprintf("git init -b %s\n", branch))

	if readme {
		buffer.WriteString(fmt.Sprintf("echo '# Initial Content as Blank' > README.md\n", name))
	}
	fmt.Println(buffer.String())
}
