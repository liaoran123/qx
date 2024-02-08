package main

import (
	"eq/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/static/", routers.Static) //静态文件服务器
	http.HandleFunc("/", routers.Index)         //首页
	http.HandleFunc("/jiage/", routers.Jiage)   //定价
	http.HandleFunc("/fabu/", routers.Fabu)     //发布
	//http.HandleFunc("/fabu/", routers.Fabu)
	fmt.Println("http://127.0.0.1:8999")
	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}
}
