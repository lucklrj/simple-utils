package runtime

import (
	"runtime/debug"
	"strings"

	"github.com/fatih/color"
)

func GoSafe(fn func()) {
	go RunSafe(fn)
}
func RunSafe(fn func()) {
	defer Recover()
	fn()
}

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		sp := strings.Repeat("*", 100)
		color.Red(sp)
		color.Red("panic:%v", p)
		color.Red(string(debug.Stack()))
	}
}
