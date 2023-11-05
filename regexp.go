package crare

import "github.com/grafana/regexp"

// cmdRx   = regexp.MustCompile(`^(/\w+)(@(\w+))?(\s|\n|$)(.+)?`)
var cbackRx = regexp.MustCompile(`^\f([-\w]+)(\|(.+))?$`)
