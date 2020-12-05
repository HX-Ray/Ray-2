package main

import (
	"fmt"
	"net/http"
)

var users = make(map[string]*user)

type user struct {
	username    string
	password    string
	information string
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var auser user

	//fmt.Print(users)
	//username := r.FormValue("username")
	//password := r.FormValue("password")
	auser.username = r.FormValue("username")
	auser.password = r.FormValue("password")
	auser.information = r.FormValue("information")
	//fmt.Println(username, password, "---")

	if _, ok := users[auser.username]; ok {
		fmt.Fprintln(w, "该用户已存在")
	} else {
		users[auser.username] = &auser
		fmt.Fprintln(w, "注册成功")
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if _, ok := users[username]; ok {
		if password == users[username].password {
			fmt.Fprintln(w, "登录成功")
		} else {
			fmt.Fprintln(w, "密码错误")
		}
	} else {
		fmt.Fprintln(w, "该用户不存在")
	}
}

func Rewrite(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	newUsername := r.FormValue("newusername")
	newPassword := r.FormValue("newpassword")
	newInformation := r.FormValue("newinformation")

	if _, ok := users[username]; ok {
		if password == users[username].password {
			var newuser user
			newuser.username = newUsername
			newuser.password = newPassword
			newuser.information = newInformation

			delete(users, username)

			users[newUsername] = &newuser

			fmt.Fprintln(w, "修改成功")
		} else {
			fmt.Fprintln(w, "密码错误")
		}
	} else {
		fmt.Fprintln(w, "该用户不存在")
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:1024",
	}

	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/rewrite", Rewrite)

	server.ListenAndServe()
}
