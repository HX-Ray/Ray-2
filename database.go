package main

import (
	"fmt"
	"net/http"
	"database/sql"

	_"github.com/go-sql-driver/mysql"
)

//连接MySQL
var db, err = sql.Open("mysql", "root:********@tcp(127.0.0.1:3306)/test?charset=utf8")

//注册
func Signup(w http.ResponseWriter, r *http.Request) {

	var name string
	username := r.FormValue("username")
	password := r.FormValue("password")
	information := r.FormValue("information")

	err = db.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&name)
	if err == nil {
		fmt.Fprintln(w, "该用户已存在")
	} else {
		query := "INSERT INTO users (username, password, information) VALUES ('" + username + "','" + password + "','" + information + "')"
		db.Exec(query)

		fmt.Fprintln(w, "注册成功")
	}
}

//登录
func Signin(w http.ResponseWriter, r *http.Request) {

	var name, pass string
	username := r.FormValue("username")
	password := r.FormValue("password")

	err = db.QueryRow("SELECT username, password FROM users WHERE username = ?", username).Scan(&name, &pass)
	if err == nil {
		if pass == password {
			fmt.Fprintln(w, "登录成功")
		} else {
			fmt.Fprintln(w, "密码错误")
		}
	} else {
		fmt.Fprintln(w, "用户不存在")
	}
}

//修改
func Rewrite(w http.ResponseWriter, r *http.Request) {

	var name, pass string
	username := r.FormValue("username")
	password := r.FormValue("password")
	newUsername := r.FormValue("newusername")
	newPassword := r.FormValue("newpassword")
	newInformation := r.FormValue("newinformation")

	err = db.QueryRow("SELECT username, password FROM users WHERE username = ?", username).Scan(&name, &pass)
	if err == nil {
		if pass == password {
			stmt, err := db.Prepare("UPDATE users SET username=?, password=?, information=? WHERE username=?")
			if err != nil {
				panic(err)
			}

			_, err = stmt.Exec(newUsername, newPassword, newInformation, username)
			if err != nil {
				panic(err)
			}

			fmt.Fprintln(w, "修改成功")
		} else {
			fmt.Fprintln(w, "密码错误")
		}
	} else {
		fmt.Fprintln(w, "用户不存在")
	}
}

//删除
func Delete(w http.ResponseWriter, r *http.Request) {

	var name, pass string
	username := r.FormValue("username")
	password := r.FormValue("password")

	err = db.QueryRow("SELECT username, password FROM users WHERE username = ?", username).Scan(&name, &pass)
	if err == nil {
		if pass == password {
			stmt, err := db.Prepare("DELETE FROM users WHERE username=?")
			if err != nil {
				panic(err)
			}

			_, err = stmt.Exec(username)
			if err != nil {
				panic(err)
			}

			fmt.Fprintln(w, "已删除该用户")
		} else {
			fmt.Fprintln(w, "密码错误")
		}
	} else {
		fmt.Fprintln(w, "用户不存在")
	}
}

func main() {

	if err != nil {
		panic(err)
	}
	defer  db.Close()

	server := http.Server{
		Addr: "127.0.0.1:1024",
	}

	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/rewrite", Rewrite)
	http.HandleFunc("/delete", Delete)

	server.ListenAndServe()
}
