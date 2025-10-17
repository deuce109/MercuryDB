package commands

import (
	"io"
	"strings"
)

const PING string = "PING"

func (h *CommandHandler) ping() io.Reader {
	return strings.NewReader("pong")
}
