package util

import (
	"github.com/davecgh/go-spew/spew"
)

func Dump(a ...interface{}) {
	spew.Dump(a)
}
