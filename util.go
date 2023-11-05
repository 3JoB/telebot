package crare

import (
	"strings"

	"github.com/3JoB/ulib/pool"
	"github.com/jamiealquiza/fnv"
)

var (
	ReleaseBuffer = pool.ReleaseBuffer
)

func process(input string) (command, bot, payload string) {
	if !strings.HasPrefix(input, "/") {
		return
	}
	atIdx := strings.Index(input, "@")
	if atIdx != -1 {
		command = input[:atIdx]
		botPayload := input[atIdx+1:]

		spaceIdx := strings.Index(botPayload, " ")
		if spaceIdx != -1 {
			bot = botPayload[:spaceIdx]
			payload = botPayload[spaceIdx+1:]
		} else {
			bot = botPayload
		}
	} else {
		splits := strings.SplitN(input, " ", 2)
		command = splits[0]
		if len(splits) > 1 {
			payload = splits[1]
		}
	}

	payload = strings.ReplaceAll(payload, "\\n", "\n")

	return
}

func hash32(e *Error) uint32 {
	return fnv.Hash32(e.String())
}
