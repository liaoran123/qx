package routers

import (
	"fmt"
	"strings"

	"xbdb"
)

var Table map[string]*xbdb.Table
var xb *xbdb.Xb

func Ini() {

	xb = xbdb.NewDb("db")

	//建表
	dbinfo := xbdb.NewTableInfoNil(xb.Db)

	if dbinfo.GetInfo("c").FieldType == nil {
		createc(dbinfo)
	}
	if dbinfo.GetInfo("meta").FieldType == nil {
		createmeta(dbinfo)
	}
	if dbinfo.GetInfo("kw").FieldType == nil {
		createkw(dbinfo)
	}

	Table = xb.GetTables()

	//Table["kws"].Select.ForTb(Pr)

	//Table["ca"].Select.FindPrefixFun([]byte("ca."), false, Pr) //ca.fid-\x00\x00\x00\x01-
	/*
		for i := 10; i < 49; i++ {
			id := strconv.Itoa(i)
			xbdb.Tables["record"].Del(id)
		}*/

}
func Pr(k, v []byte) bool {
	ks, kv := string(k), string(v)
	if strings.Contains(ks, "太玄经") {
		fmt.Println(ks, kv)
	}
	//fmt.Println(ks, kv)
	return true
}

// 创建文章内容表，该表是全文搜索，故而名称尽量短，可以减少文件大小。
// 带全文搜索索引的内容表c
func createc(tbifo *xbdb.TableInfo) {
	name := "user"                                                       //用户表，
	fields := []string{"id", "nc", "psw", "dt", "llog"}                  //字段 id，昵称，密码，注册日期,最后登录
	fieldType := []string{"int", "string", "string", "string", "string"} //字段
	idxs := []string{"1"}                                                //索引字段,fields的下标对应的字段。支持组合查询，字段之间用,分隔
	fullText := []string{}                                               //考据级全文搜索索引字段的下标。
	ftlen := "7"                                                         //全文搜索的长度，中文默认是7
	patterns := []string{}                                               //搜索词模型。 1,中文;2字母;3，数字；4，标点符号；5，自定义。不符合的字被过滤。可以组合。
	diychar := ""
	r := tbifo.Create(name, ftlen, diychar, fields, fieldType, idxs, fullText, patterns)
	fmt.Printf("r: %v\n", r)
}

func createmeta(tbifo *xbdb.TableInfo) {
	name := "diary"                                            //日记表，
	fields := []string{"id", "title", "text", "dt"}            //字段 id,标题，内容，日期
	fieldType := []string{"int", "string", "string", "string"} //字段
	idxs := []string{}                                         //索引字段,fields的下标对应的字段。支持组合查询，字段之间用,分隔
	fullText := []string{}                                     //考据级全文搜索索引字段的下标。
	ftlen := "7"                                               //全文搜索的长度，中文默认是7
	patterns := []string{}                                     //搜索词模型。 1,中文;2字母;3，数字；4，标点符号；5，自定义。不符合的字被过滤。可以组合。
	diychar := ""
	r := tbifo.Create(name, ftlen, diychar, fields, fieldType, idxs, fullText, patterns)
	fmt.Printf("r: %v\n", r)
}
func createkw(tbifo *xbdb.TableInfo) {
	name := "qx"                                                                //目录表，
	fields := []string{"id", "did", "emot", "front", "level", "dt"}             //字段 id，日记的id，情绪词，否则正面，程度，日期
	fieldType := []string{"int", "int", "string", "string", "string", "string"} //字段
	idxs := []string{"1", "2", "3", "4", "5"}                                   //索引字段,fields的下标对应的字段。支持组合查询，字段之间用,分隔
	fullText := []string{}                                                      //考据级全文搜索索引字段的下标。
	ftlen := "7"                                                                //全文搜索的长度，中文默认是7
	patterns := []string{}                                                      //搜索词模型。 1,中文;2字母;3，数字；4，标点符号；5，自定义。不符合的字被过滤。可以组合。
	diychar := ""
	r := tbifo.Create(name, ftlen, diychar, fields, fieldType, idxs, fullText, patterns)
	fmt.Printf("r: %v\n", r)
}
