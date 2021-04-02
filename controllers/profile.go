package controllers

import (
	"apiproject/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
)

// ProfileController operations for Profile
type ProfileController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProfileController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Profile
// @Param	body		body 	models.Profile	true		"body for Profile content"
// @Success 201 {object} models.Profile
// @Failure 403 body is empty
// @router / [post]
func (c *ProfileController) Post() {
	var profile models.Profile
	defer c.ServeJSON()
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &profile)
	validated, err := valid.Valid(&profile)
	if err != nil {
		log.Println(err)
	}
	if !validated {
		errors := make(map[string]string)
		for _, err := range valid.Errors {
			errors[err.Field] = err.Message
		}
		c.Data["json"] = errors
		c.Ctx.Output.Status = http.StatusBadRequest
	}
	uid, err := models.AddProfile(&profile)
	if err != nil {
		c.Data["json"] = err
		c.Ctx.Output.Status = http.StatusInternalServerError
	}
	c.Data["json"] = map[string]interface{}{"id": uid}
	c.Ctx.Output.Status = http.StatusCreated

}

// GetOne ...
// @Title GetOne
// @Description get Profile by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Profile
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProfileController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Profile
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Profile
// @Failure 403
// @router / [get]
func (c *ProfileController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Profile
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Profile	true		"body for Profile content"
// @Success 200 {object} models.Profile
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProfileController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Profile
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProfileController) Delete() {

}
