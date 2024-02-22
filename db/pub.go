package db

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 打开词库文件转为map
func openKwFile(ph string) map[string]bool {
	ks := make(map[string]bool)
	var err error
	var file *os.File
	file, err = os.Open(ph)
	if err != nil {
		return ks
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var line []byte
	var kw string
	for {
		line, _, err = reader.ReadLine()
		if err == io.EOF {
			return ks
		}
		kw = strings.Split(string(line), " ")[0]
		ks[kw] = true
	}
}

// 将内容追加文件末尾
func appendToFile(file, str string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", file)
		return
	}
	defer f.Close()
	f.WriteString("\n" + str)

}
