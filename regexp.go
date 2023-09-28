package telebot

import "github.com/grafana/regexp"

var (
	// cmdRx   = regexp.MustCompile(`^(/\w+)(@(\w+))?(\s|\n|$)(.+)?`)
	cbackRx = regexp.MustCompile(`^\f([-\w]+)(\|(.+))?$`)
)
