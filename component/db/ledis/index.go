package ledis

import (
	"os"

	"github.com/fatih/color"
	cfg "github.com/ledisdb/ledisdb/config"
	"github.com/ledisdb/ledisdb/ledis"
)

var DB *ledis.DB

func MakeDefault(prefix string) {
	cfg := cfg.NewConfigDefault()
	cfg.DataDir = prefix + "/db/Db"
	l, err := ledis.Open(cfg)
	if err != nil {
		color.Red("加载ledis错误：", err.Error())
		os.Exit(0)
	}
	DB, _ = l.Select(0)

}
