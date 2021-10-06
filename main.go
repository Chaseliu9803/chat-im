package main

import (
	"chat-im/controller"
	"html/template"
	"log"
	"net/http"
)

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", controller.Login)

	//http.Handle("/", http.FileServer(http.Dir(".")))
	//http.Handle("/asset/", http.FileServer(http.Dir(".")))

	//// /user/login.shtml
	//http.HandleFunc("/user/login.shtml", func(w http.ResponseWriter, r *http.Request) {
	//	//解释
	//	tpl, err := template.ParseFiles("view/user/login.html")
	//	if err != nil {
	//		//打印并直接退出
	//		log.Fatal(err.Error())
	//	}
	//	tpl.ExecuteTemplate(w, "/user/login.shtml", nil)
	//})

	//解释所有模板，而不需要一个个解释
	RegisterAllViews()

	http.ListenAndServe(":8080", nil)
}

func RegisterAllViews() {
	//解释所有模板
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			err := tpl.ExecuteTemplate(writer, tplName, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}
}