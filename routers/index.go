package routers

import (
	"fmt"
	"html/template"
	"net/http"
)

type index struct {
	Host string
	Path string
	Uri  string
	Up   string //繁体参数 ?l=0
	Ift  bool   //是否转繁体
}

func Newindex() index {
	return index{}
}
func Index(w http.ResponseWriter, req *http.Request) {
	rd := Newindex()
	rd.Host = "http://" + req.Host
	//统计
	//--组织模板数据
	TemplatesFiles := []string{
		"tpl/index.html",
		"tpl/pub/static.html",
		"tpl/pub/header.html",
		"tpl/pub/search.html",
		"tpl/pub/footer.html", // 多加的文件
	}
	funcMap := template.FuncMap{ //--需要注册的函数

	}
	t, _ := template.New("index.html").Funcs(funcMap).ParseFiles(TemplatesFiles...)
	//--New("index.html") 的 index.html必须是TemplatesFiles第一个文件名
	err := t.ExecuteTemplate(w, "index.html", rd)
	if err != nil {
		fmt.Println(req.URL.Path, err)
	}
}
