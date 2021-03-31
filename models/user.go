package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"

	// beego "github.com/beego/beego/v2/server/web"
	// _ "github.com/mattn/go-sqlite3" // import your required driver
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64
	Username string `orm:"size(100)"`
	Password string `orm:"size(100)"`
	Email    string `orm:"unique"`
}

type Profile struct {
	Id      int64  `orm:"auto"`
	Gender  string `orm:"size(10)"`
	Age     int64
	Address string `orm:"size(100)"`

	User *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Profile))
	// set default database
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "my_db")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8")
	orm.RunSyncdb("default", false, true)

	fmt.Println("In models.init")
}

func GetUser(uid int64) (*User, error) {
	o := orm.NewOrm()
	user := User{Id: uid}
	err := o.Read(&user)
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}
func AddUser(user *User) (int64, error) {
	o := orm.NewOrm()

	id, err := o.Insert(user)

	return id, err

}

func GetAllUsers() (num int64, users []*User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	num, err = qs.All(&users)
	if err != nil {
		return num, nil, err
	} else {
		return num, users, nil
	}
}

func UpdateUser(user *User) (int64, error) {
	o := orm.NewOrm()

	id, err := o.Update(user)

	return id, err

}
func DeleteUser(user *User) (int64, error) {
	o := orm.NewOrm()

	id, err := o.Delete(user)

	return id, err

}

// var UserList map[int]*User

// func GetAllUsers() map[int]*User {
// 	o := orm.NewOrm()
// 	_, err := o.QueryTable("users").Values(&UserList)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return UserList
// }
