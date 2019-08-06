/**
*@date 2019_8_5
*@author martian
*@function:处理用户信息表
*/
package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
//用户注册，用户信息存入数据库
/**
	uID int not null primary key auto_increment,
	nickName char(50) not null DEFAULT 'leiyou',
	password char(20) not null DEFAULT '123456',
	birthday char(20) DEFAULT '1994-8-20',
	type char DEFAULT '0',
	avatar char(50) not null DEFAULT ''
*/
func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}

//db *sql.DB,tableName string,uID int,nickName [50]char,password [20]char,birthday [20]char,userType char,avatar [50]char
func userRegister(db *sql.DB,tableName string,uID int,nickName []byte,password []byte,birthday []byte,userType byte,avatar []byte){
	sql:=`INSERT into `+tableName + `(
		uID,
		nickName,
		password,
		birthday,
		userType,
		avatar
	) `+`values(?,?,?,?,?,?)`
	//` ,`+uID+','+' '+nickName+','+' '+password+','+' '+birthday+','+' '+userType+','+' '+avatar+';'
	fmt.Println("\n"+sql+"\n")
	n,err:=db.Exec(sql,uID,nickName,password,birthday,userType,avatar)
	if(err==nil){
		fmt.Printf("update success, affected rows:%d\n", n)
	}else{
		checkErr(err)
	}
}

func main(){
	//打开数据库
	db,err:=sql.Open("mysql","root:123456@/zhangcheng")
	checkErr(err)
	// var data []byte = []byte("userInfo")
	// userRegister(db,"userInfo",1,[]byte ("hehe"),[]byte ("123456"),[]byte ("1994-8-20"),0,[]byte ("https:www.baidu.com"))

	for i:=1;i<50;i++ {
		userRegister(db,"userInfo",i+1,[]byte ("hehe"),[]byte ("123456"),[]byte ("1994-8-20"),0,[]byte ("https:www.baidu.com"))
	}
}