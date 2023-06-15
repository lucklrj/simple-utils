package runtime

import (
	"runtime/debug"

	"github.com/lucklrj/simple-utils/component/system/log"
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
		log.Logger.Error(p)
		log.Logger.Error(string(debug.Stack()))
	}
}
