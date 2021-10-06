package controller

import (
	"chat-im/params/bo"
	"chat-im/pkg/response"
	"fmt"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	fmt.Println(mobile, passwd)

	if mobile == "15626168049" && passwd == "123" {
		loginBo := bo.UserLoginBo{
			UserId: "111",
			Token: "test_token",
		}
		response.Success(writer, loginBo, "登录成功")
	} else {
		response.Fail(writer, "账号/密码不正确")
	}
}
