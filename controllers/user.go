package controllers

import (
	"apiproject/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 201 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	valid := validation.Validation{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	validated, err := valid.Valid(&user)
	if err != nil {
		log.Println(err)
	}
	if !validated {
		errors := make(map[string]string)
		for _, err := range valid.Errors {
			errors[err.Field] = err.Message
		}
		u.Data["json"] = errors
		u.Ctx.Output.Status = http.StatusBadRequest
		u.ServeJSON()
	}
	uid, err := models.AddUser(&user)
	if err != nil {
		u.Data["json"] = err
		u.Ctx.Output.Status = http.StatusInternalServerError
		u.ServeJSON()
	}
	u.Data["json"] = map[string]interface{}{"uid": uid}
	u.Ctx.Output.Status = http.StatusCreated
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	_, users, _ := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	id := u.GetString(":uid")
	if id != "" {
		uid, _ := strconv.ParseInt(id, 10, 64)
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	id := u.GetString(":uid")
	if id != "" {
		uid, _ := strconv.ParseInt(id, 10, 64)
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		user.Id = uid
		uu, err := models.UpdateUser(&user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	id := u.GetString(":uid")
	uid, _ := strconv.ParseInt(id, 10, 64)
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	type LoginUser struct {
		Username string `valid:"Required"`
		Password string `valid:"Required"`
	}
	defer u.ServeJSON()
	var loginUser LoginUser
	json.Unmarshal(u.Ctx.Input.RequestBody, &loginUser)
	valid := validation.Validation{}
	validated, err := valid.Valid(&loginUser)
	if err != nil {
		log.Println(err)
	}
	if !validated {
		errors := make(map[string]string)
		for _, err := range valid.Errors {
			errors[err.Field] = err.Message
		}
		u.Data["json"] = errors
		u.Ctx.Output.Status = http.StatusBadRequest
	}

	o := orm.NewOrm()
	qs := o.QueryTable("user")
	qs = qs.Filter("Username", loginUser.Username).Filter("Password", loginUser.Password)
	exist := qs.Exist()
	type Response struct {
		Success bool
		Message string
	}
	var data Response
	if exist {
		data.Success = true
		data.Message = "User logged in successfully"
		var user models.User
		qs.One(&user)
		// successfully login, put the user into session
		u.SetSession("User", user.Username)
		log.Printf("Logged in user %s", user.Username)
	} else {
		data.Success = false
		data.Message = "Incorrect username or password"
	}

	u.Data["json"] = data
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	username := u.GetSession("User")
	log.Printf("Logged out user %s", username)
	u.DestroySession()
	u.ServeJSON()
}
