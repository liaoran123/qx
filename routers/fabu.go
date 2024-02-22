package routers

import (
	"fmt"
	"html/template"
	"net/http"
)

type fabu struct {
	Host string
	Path string
	Uri  string
	Up   string //繁体参数 ?l=0
	Ift  bool   //是否转繁体
}

func Newfabu() fabu {
	return fabu{}
}
func Fabu(w http.ResponseWriter, req *http.Request) {
	rd := Newfabu()
	rd.Host = "http://" + req.Host
	//统计
	//--组织模板数据
	TemplatesFiles := []string{
		"tpl/fabu.html",
		"tpl/pub/static.html",
		"tpl/pub/header.html",
		"tpl/pub/quill.html",
		"tpl/pub/footer.html", // 多加的文件
	}
	funcMap := template.FuncMap{ //--需要注册的函数

	}
	t, _ := template.New("fabu.html").Funcs(funcMap).ParseFiles(TemplatesFiles...)
	//--New("fabu.html") 的 fabu.html必须是TemplatesFiles第一个文件名
	err := t.ExecuteTemplate(w, "fabu.html", rd)
	if err != nil {
		fmt.Println(req.URL.Path, err)
	}
}
