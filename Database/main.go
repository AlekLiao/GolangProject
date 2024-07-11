//https://ithelp.ithome.com.tw/m/articles/10234657

package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "root"
	PASSWORD = "00000"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "test"
)

type User struct {
	ID       string
	Username string
	Password string
}

type User_ORM struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Username string `json:"username"`
	Password string `json:""`
}

const DbStyle = "ORM"

func main() {
	fmt.Println("Go MySQL Tutorial")

	if DbStyle == "ORM" {
		db := OpenDB_ORM()
		fmt.Println("Hello, World!", db)

		CreateTabel_ORM(db)

		for index := 0; index < 10; index++ {
			sUserName := "test_ORM" + strconv.Itoa(index+1)
			sPassword := strconv.Itoa(index)
			user := &User_ORM{
				Username: sUserName,
				Password: sPassword,
			}
			InsertData_ORM(db, user)
		}

		if user, err := QueryData_ORM(db, 23); err == nil {
			log.Println("查詢到 User 為 ", user)
		} else {
			panic("查詢 user 失敗，原因為 " + err.Error())
		}
	} else {
		// 基本資料庫操作
		var db *sql.DB
		db = ConnectDatabase()
		CreateTabel(db)

		for index := 0; index < 10; index++ {
			sUserName := "test" + strconv.Itoa(index+1)
			sPassword := strconv.Itoa(index)
			InsertData(db, sUserName, sPassword)
		}

		QueryData(db, "test")
		CloseDatabase(db)
	}

}

func OpenDB_ORM() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
	}
	return db
}

func CreateTabel_ORM(db *gorm.DB) {
	if err := db.AutoMigrate(new(User_ORM)); err != nil {
		panic("資料庫 Migrate 失敗，原因為 " + err.Error())
	}
}

func InsertData_ORM(db *gorm.DB, model interface{}) error {
	return db.Create(model).Error
}

func QueryData_ORM(db *gorm.DB, id int64) (*User_ORM, error) {
	user := new(User_ORM)
	user.ID = id
	err := db.First(&user).Error
	return user, err
}

func ConnectDatabase() *sql.DB {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("開啟 MySQL 連線發生錯誤，原因為：", err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("資料庫連線錯誤，原因為：", err.Error())
		panic(err)
	}
	return db

}

func CloseDatabase(db *sql.DB) {
	defer db.Close()
}

func CreateTabel(db *sql.DB) error {
	// users is the table name
	sql := `CREATE TABLE IF NOT EXISTS users(	
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64)
	); `

	if _, err := db.Exec(sql); err != nil {
		fmt.Println("建立 Table 發生錯誤:", err)
		return err
	}

	fmt.Println("建立 Table 成功！")
	return nil
}

func InsertData(db *sql.DB, sUserName, sPassword string) error {
	_, err := db.Exec("insert INTO users(username,password) values(?,?)", sUserName, sPassword)
	if err != nil {
		fmt.Printf("建立使用者失敗，原因是：%v", err)
		return err
	}

	fmt.Println("建立使用者成功！")
	return nil
}

func QueryData(db *sql.DB, sUserName string) {
	/*
		user := new(User)
		row := db.QueryRow("select * from users where username=?", sUserName)
		if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			fmt.Printf("映射使用者失敗，原因為：%v\n", err)
			return
		}
		fmt.Println("查詢使用者成功", *user)
	*/
}
