package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

type UserController struct {
	BaseController
}

func (this *UserController) ShowLogin() {
	this.TplName = "login.html"
}
//注释的这个传递json数据格式有误
//func (this *UserController) HandleLogin() {
//	defer this.ServeJSON()
//	username := this.GetString("username")
//	pwd := this.GetString("pwd")
//	db, err := bolt.Open("test.db", 0600, nil)
//	if err != nil {
//		panic(err)
//		return
//	}
//	defer db.Close()
//	var res ResponseJSON
//	db.View(func(tx *bolt.Tx) error {
//		bucket := tx.Bucket([]byte(username))
//		if bucket == nil {
//			res = ResponseJSON{403, nil, "用户名不存在"}
//			fmt.Println("不存在该用户")
//		} else {
//			password := bucket.Get([]byte(username))
//			if string(password) == pwd {
//				user:=User{time.Now().UnixNano(),username,pwd}
//				//bytes, _ := json.Marshal(user)
//				//jsonstring = &ResponseJSON{200, string(bytes), "用户登录成功"}
//				res = ResponseJSON{200, user, "用户登录成功"}
//				fmt.Println("登录成功")
//			} else {
//				res = ResponseJSON{403, nil, "密码有误"}
//				fmt.Println("密码有误")
//			}
//		}
//		return nil
//	})
//	bytes, _ := json.Marshal(res)
//	this.Data["json"] = string(bytes);
//}

//下面这个json传递方式正确
func (this *UserController) HandleLogin() {
	defer this.ServeJSON()
	//表单获取数据
	username := this.GetString("username")
	pwd := this.GetString("pwd")
	//获取json传递的参数  app.conf需要配置  copyrequestbody = true
	//var ob models.UserParam;
	//body:=this.Ctx.Input.RequestBody
	//fmt.Println("body:",body)
	//json.Unmarshal(body,&ob)
	//username:=ob.Username
	//pwd:=ob.Pwd
	fmt.Println("username:",username)
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
	var jsonstring *ResponseJSON
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(username))
		if bucket == nil {
			jsonstring = &ResponseJSON{403, nil, "用户名不存在"}
			//panic("不存在该用户")
			fmt.Println("不存在该用户")
		} else {
			password := bucket.Get([]byte(username))
			if string(password) == pwd {
				user:=User{time.Now().UnixNano(),username,pwd}
				//bytes, _ := json.Marshal(user)
				//jsonstring = &ResponseJSON{200, string(bytes), "用户登录成功"}
				jsonstring = &ResponseJSON{200, user, "用户登录成功"}
				fmt.Println("登录成功")
			} else {
				jsonstring = &ResponseJSON{403, nil, "密码有误"}
				fmt.Println("密码有误")
			}
		}
		return nil
	})
	this.Data["json"] = jsonstring;
}

func (this *UserController) ShowRegister() {
	this.TplName = "register.html"
}

func (this *UserController) HandleRegister() {
	defer this.ServeJSON()
	username := this.GetString("username")
	pwd := this.GetString("pwd")

	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
	var jsonstring *ResponseJSON
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(username))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(username))
			if err != nil {
				jsonstring = &ResponseJSON{200, nil, "用户注册失败"}
				return err
			}
		}
		user:=User{time.Now().UnixNano(),username,pwd}
		bytes, _ := json.Marshal(user)
		jsonstring = &ResponseJSON{200, string(bytes), "用户注册成功"}
		bucket.Put([]byte(username), []byte(pwd))
		return nil
	})
	this.Data["json"] = jsonstring;
}
