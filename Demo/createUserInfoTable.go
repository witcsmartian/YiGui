/**
*@date 2019_8_5
*@author martian
*@function:创建用户信息表
*/
package main

import (
	// "util"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}

/**
*param
*/
func createUserInfoTable(db *sql.DB,tableName string){
		sqlFir:=`DROP TABLE IF EXISTS `+tableName+`;`
		fmt.Println("\n"+sqlFir+"\n")
		n,err:=db.Exec(sqlFir)
		checkErr(err)

		if(err==nil){
			fmt.Printf("update success, affected rows:%d\n", n)
		}else{
			return
		}

		sqlSec:=`CREATE TABLE `+tableName + `(
			uID int not null primary key auto_increment,
			nickName char(50) not null DEFAULT 'leiyou',
			password char(20) not null DEFAULT '123456',
			birthday char(20) DEFAULT '1994-8-20',
			userType char DEFAULT '0',
			avatar char(50) not null DEFAULT 'xxx'
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

		fmt.Println("\n"+sqlSec+"\n")
		n1,err1:=db.Exec(sqlSec)
		checkErr(err1);
		if(err1==nil){
			fmt.Printf("update success, affected rows:%d\n", n1)
		}else{
			return
		}
}

func main(){
	//打开数据库
	db,err:=sql.Open("mysql","root:123456@/zhangcheng")
	checkErr(err)
	//创建数据库表
	createUserInfoTable(db,"userInfo")
}