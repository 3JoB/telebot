package telebot

import (
	"strings"

	"github.com/3JoB/ulib/pool"
	"github.com/jamiealquiza/fnv"
)

var (
	ReleaseBuffer = pool.ReleaseBuffer
	mod           = 1000000007
)

func process(input string) (command, bot, payload string) {
	/*Benchmark
	RE2: wasm, no cgo
	REG: github.com/grafana/regexp
	Strings: process()
	cpu: 12th Gen Intel(R) Core(TM) i7-12700H
	----
	Benchmark_RE2-20                            	  377274	      3039 ns/op	     592 B/op	      13 allocs/op
	Benchmark_REG-20                            	 1746291	       685.1 ns/op	     436 B/op	       3 allocs/op
	Benchmark_Strings-20                        	28667738	        43.29 ns/op	      32 B/op	       1 allocs/op*/
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

func hash32p(e string) uint32 {
	return fnv.Hash32(e)
}

func hash32(e *Error) uint32 {
	return fnv.Hash32(e.String())
}
