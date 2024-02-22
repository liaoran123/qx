package db

import (
	"github.com/wangbin/jiebago"
)

var seg jiebago.Segmenter

func Ini() {
	seg.LoadDictionary("dict/dict.txt")
}
