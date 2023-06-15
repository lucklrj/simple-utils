package runtime

import "github.com/zeromicro/go-zero/core/rescue"

func GoSafe(fn func()) {
	go RunSafe(fn)
}
func RunSafe(fn func()) {
	defer rescue.Recover()
	fn()
}
