package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB
type Body struct {//定义一个结构体来接受POST请求的数据
	Name string `json:"name"`
	Passwd string `json:"passwd"`
}

func InitDB() (err error){
	dns := "root:admin@tcp(127.0.0.1:3306)/class"
	db, err = sql.Open("mysql",dns)
	if err != nil {
		fmt.Printf("Open failed,err:%v",err)
		return err
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Ping failed,err:%v",err)
		return err
	}
	return nil
}

func Reg (name string,passwd string) (err error){
	sqlStr := "INSERT INTO user (name, passwd) VALUES (?, ?);"
	_, err = db.Exec(sqlStr,name,passwd)
	if err != nil {
		fmt.Println("INSERT failed,err:%v",err)
		return err
	}
	return nil
}

func Login (name string, passwd string) (err error){
	sqlStr := "SELECT name, passwd FROM user where name=?;"
	err = db.QueryRow(sqlStr,name).Scan(&name, &passwd)
	if err != nil {
		fmt.Println("SELECT failed,err:%v",err)
		return
	}
	return nil
}

func main() {
	_ = InitDB()
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.Status(200)
	})
	r.POST("/login", func(c *gin.Context) {
		var usr Body
		err := c.ShouldBindJSON(&usr)//解析JSON数据，注意：若多次请求，应该用ShouldBindBodyWith()
		if err != nil {
			fmt.Printf("Bind failed,err:%v\n", err)
		}
			err = Login(usr.Name, usr.Passwd)
			if err != nil {
				c.String(http.StatusOK, "登录失败")
			}else{
				c.String(http.StatusOK,"登录成功")
		}
	})
	r.GET("/register", func(c *gin.Context) {
		c.Status(200)
	})
	r.POST("/register", func(c *gin.Context) {
		var usr Body
		err := c.ShouldBindJSON(&usr)
		if err != nil {
			fmt.Printf("Bind failed,err:%v\n",err)
		}
		if usr.Name == "" || usr.Passwd == ""{
			c.String(200,"用户名或密码不能为空")
			return
		}
		err = Reg(usr.Name,usr.Passwd)
		if err != nil {
			fmt.Printf("Register failed,err:%v\n",err)
			c.String(200,"注册失败",)
		}else{
			c.String(200,"注册成功")
		}
	})
	err := r.Run("127.0.0.1:2359")
	if err != nil {
		fmt.Printf("Run failed,err:%v",err)
	}
}