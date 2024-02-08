package routers

import (
	"fmt"
	"html/template"
	"net/http"
)

type emotions struct {
	Host string
	Path string
	Uri  string
	Up   string //繁体参数 ?l=0
	Ift  bool   //是否转繁体
}

func Newemotions() emotions {
	return emotions{}
}
func Emotions(w http.ResponseWriter, req *http.Request) {
	rd := Newemotions()
	rd.Host = "http://" + req.Host
	//统计
	//--组织模板数据
	TemplatesFiles := []string{
		"tpl/emotions.html",
		"tpl/pub/static.html",
		"tpl/pub/header.html",
		"tpl/pub/search.html",
		"tpl/pub/footer.html", // 多加的文件
	}
	funcMap := template.FuncMap{ //--需要注册的函数

	}
	t, _ := template.New("emotions.html").Funcs(funcMap).ParseFiles(TemplatesFiles...)
	//--New("emotions.html") 的 emotions.html必须是TemplatesFiles第一个文件名
	err := t.ExecuteTemplate(w, "emotions.html", rd)
	if err != nil {
		fmt.Println(req.URL.Path, err)
	}
}
