package main

import (
	"eq/db"
	"eq/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {

	db.Ini()
	http.HandleFunc("/static/", routers.Static)            //静态文件服务器
	http.HandleFunc("/", routers.Index)                    //首页
	http.HandleFunc("/jiage/", routers.Jiage)              //定价
	http.HandleFunc("/fabu/", routers.Fabu)                //发布
	http.HandleFunc("/diary/insert/", routers.Diaryinsert) //添加日记
	fmt.Println("http://127.0.0.1:8999")
	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}
}
