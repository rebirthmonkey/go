package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	user_title	string
	user_author	string
}

func main() {
	db, err := sql.Open("mysql", "root:CCVg5TwyS8@tcp(127.0.0.1:3307)/db1")
	defer db.Close()

	if err != nil {
		panic(err)
	}
	fmt.Println("the DB is:", db)


	// INSERT
	sql1 := "insert into tb1 (user_title,user_author) values (?,?)"
	value1 := [3]string{"aaa", "aaa"}

	res, err := db.Exec(sql1, value1[0], value1[1])
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	lastId, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The last inserted row id: %d\n", lastId)


	// SELECT
	sql2 := "select user_title, user_author from tb1 where user_title=?"
	res2, err2 := db.Query(sql2, "aaa")
	defer res2.Close()
	if err2 != nil {
		fmt.Println("exec failed, ", err2)
		return
	}

	for res2.Next() {
		var user User
		err := res2.Scan(&user.user_title, &user.user_title)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v, %v\n", user.user_title, user.user_title)
	}

	// DELETE
	sql3 := "DELETE FROM tb1 WHERE user_title=?"
	res3, err3 := db.Exec(sql3, "aaa")

	if err3 != nil {
		panic(err3.Error())
	}

	affectedRows, err4 := res3.RowsAffected()

	if err != nil {
		fmt.Println(err4)
	}

	fmt.Printf("The statement affected %d rows\n", affectedRows)

}