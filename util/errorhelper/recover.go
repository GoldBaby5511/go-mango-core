package errorhelper

import (
	"github.com/GoldBaby5511/go-mango-core/log"
	"runtime/debug"
)

func Recover() {
	if err := recover(); err != nil {
		log.Error("", "Recover err=%v,stack=%v\r\n", err, string(debug.Stack()))
	}
}

func RecoverWarn() {
	if err := recover(); err != nil {
		log.Debug("", "Recover Warn:err=%v", err)
	}
}
