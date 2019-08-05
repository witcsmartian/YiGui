package main

import(
	"database/sql"
	_"github.com/Go-SQL-Driver/MySQL"//就是你下载的文件地址，如果是自己拷贝的，那么就写自己创建的路径
	"fmt"
 )

 func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

 func main() {
	 db,err:=sql.Open("mysql", "root:123456@/mysql?charset=utf8")
	 checkErr(err)

	 stmt, err := db.Prepare("INSERT user_info SET id=?,name=?")
	 checkErr(err)
 
	 res, err := stmt.Exec(1, "zhangcheng")
	 checkErr(err)

	 stmt, err = db.Prepare("update user_info set name=? where id=?")
	 checkErr(err)
 
	 res, err = stmt.Exec("zhancheng_update", 1)
	 checkErr(err)

	 affect,err:=res.RowsAffected()
	 checkErr(err)

	 fmt.Println(affect)

	 rows, err := db.Query("SELECT * FROM user_info")
	 checkErr(err)

	 for rows.Next() {
        var uid int
        var username string

        err = rows.Scan(&uid, &username)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
	}
	
	  // delete
	  stmt, err = db.Prepare("delete from user_info where id=?")
	  checkErr(err)
  
	  res, err = stmt.Exec(1)
	  checkErr(err)
  
	  // query
	  rows, err = db.Query("SELECT * FROM user_info")
	  checkErr(err)
  
	  for rows.Next() {
		  var uid int
		  var username string
  
		  err = rows.Scan(&uid, &username)
		  checkErr(err)
		  fmt.Println(uid)
		  fmt.Println(username)
	  }
  
	  db.Close()
 }