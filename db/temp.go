package db

import "strings"

//将所有词典文件的词合并到dict/dict.txt
func AppendKws() {
	jieba_sentiment := openKwFile("dict/very.txt")
	dict := openKwFile("dict/dict.txt")
	var newkw strings.Builder
	newkw.WriteString("\n")
	for k, _ := range jieba_sentiment {
		if _, ok := dict[k]; !ok {
			newkw.WriteString(k + " 49 n \n") //结巴词格式==词 词频 词性 如 人们 49 n.
		}
	}
	appendToFile("dict/dict.txt", newkw.String())
}
