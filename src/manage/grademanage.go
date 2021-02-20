package manage

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//表结构
type student struct {
	id    int     `Db:"id"`
	name  string  `Db:"name"`
	grade float32 `Db:"grade"`
}

//连接数据库初始化
func init() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@/student?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
}

func query() error {
	//查询表
	rows, err := Db.Query("SELECT * FROM studentGrade")
	//用完关闭
	defer rows.Close()
	//遍历打印
	for rows.Next() {
		var s student
		err = rows.Scan(&s.id, &s.name, &s.grade)
		fmt.Println(s)
	}
	return err
}

func insert() error {
	var s student
	fmt.Println("请依次输入要插入的学号，姓名和成绩")
	fmt.Scan(&s.id, &s.name, &s.grade)
	//执行MySql语句
	_, err := Db.Exec("INSERT INTO studentGrade(id,name,grade)VALUES (?,?,?)", s.id, s.name, s.grade)
	if err != nil {
		return errors.New("插入错误")
	}
	fmt.Println("插入成功")
	return nil
}

func set() error {
	var s student
	fmt.Println("请输入要改的学号")
	fmt.Scan(&s.id)
	fmt.Println("输入更改后的姓名")
	fmt.Scan(&s.name)
	fmt.Println("输入更改后的成绩")
	fmt.Scan(&s.grade)
	//执行MySql语句
	_, err := Db.Exec("UPDATE studentGrade SET name=?,grade=? WHERE id=?", s.name, s.grade, s.id)
	if err != nil {
		return errors.New("更新错误")
	}
	fmt.Println("更新成功")
	return nil
}

func delete() error {
	var s student
	fmt.Println("请输入删除的学号")
	fmt.Scan(&s.id)
	Db.Exec("DELETE FROM studentGrade where id=?", s.id)
	query()
	fmt.Println("删除成功")
	return nil
}

func sort() error {
	//查询表
	rows, err := Db.Query("SELECT * FROM studentGrade ORDER BY grade")
	//关闭
	defer rows.Close()
	//遍历打印
	for rows.Next() {
		var s student
		err = rows.Scan(&s.id, &s.name, &s.grade)
		fmt.Println(s)
	}
	return err
}

func Test() {
	fmt.Println("|***************************|\n1.插入(insert)\n2.查看所有成绩(query)\n3.更改数据(set)\n4.删除数据\n5.退出(quit)\n|***************************|")
	//进入循环判断
	for true {
		cm := ""
		fmt.Scan(&cm)
		switch cm {
		case "insert":
			insert()
		case "query":
			query()
		case "set":
			set()
		case "sort":
			sort()
		case "delete":
			delete()
		case "quit":
			return
		}
	}
}
