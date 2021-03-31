package main

import (
	"apiproject/models"
	_ "apiproject/routers"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	user := models.User{Id: 123, Username: "nikhil", Password: "nikhil", Email: "email"}
	id, err := models.AddUser(&user)
	fmt.Println("After adding ", id, err)

	user.Email = "email@domain.com"
	id, err = models.UpdateUser(&user)
	fmt.Println("After updating ", id, err)

	var uid int64 = 123
	user1, err := models.GetUser(uid)
	fmt.Println("After Reading", user1, err)

	num, ul, err := models.GetAllUsers()
	fmt.Println("After Reading All")
	for _, u := range ul {
		fmt.Println(num, u, err)
	}
	id, err = models.DeleteUser(&user)
	fmt.Println("After deleting ", id, err)
	beego.Run()
}
