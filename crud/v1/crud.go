package main

import "database/sql"
import _ "github.com/lib/pq" // 导入postgres数据库驱动并在init函数中注册

type Post struct {
	Id      int
	Content string
	Author  string
}

// Db 指向sql.DB结构的指针
// DB is a database handle representing a pool of zero or more underlying connections.
// It's safe for concurrent use by multiple goroutines.
// DB是一个数据库句柄，代表一个包含了零或多个数据库连接的连接池，这个连接池有sql包管理。
// 即DB只是一个句柄而不是实际的连接，这个句柄代表的是一个会自动对连接进行管理的连接池。
var Db *sql.DB

// func Open(driverName, dataSourceName string) (*DB, error)
// args: 数据库驱动名字 / 数据源名字
// Open函数在执行时并不会真正地与数据库进行连接，也不会检查用户给定的参数。
// 它的真正作用是设置好数据库所需的各个结构，并以惰性的方式，等到真正需要建立连接时才建立真正地数据库连接。
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// Posts 获取多篇帖子
// Db.Query方法返回一个Rows接口，它是一个迭代器，可以通过反复调用它的Next方法来对其进行迭代并获得相应的
// sql.Row, 当所有行都被迭代完毕，Next方法将返回io.EOF作为结果。
func Posts(limit int) (posts []Post, err error) {
	// func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}

		posts = append(posts, post)
	}

	rows.Close()
	return
}

// GetPost 获取单篇帖子
// Db.QueryRow方法返回一个指向sql.Row结构的引用。
// 如果被执行的SQL语句返回多于一个sql.Row，那么将会返回结果中的第一个sql.Row,并丢弃剩余的。
// 同时，它不会返回任何错误，只会返回一个指向sql.Row结构的引用，所以它可以跟Row结构的Scan方法搭配使用，
// 并由Scan方法把行中的值复制到给定的参数中。
func GetPost(id int) (post Post, err error) {
	post = Post{}
	// func (db *DB) QueryRow(query string, args ...interface{}) *Row
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// Create 创建帖子
// 通常，结构的Id字段值是由数据库的自增主键自动生成的
func (post *Post) Create() (err error) {
	// 定义一条SQL预处理语句，通常需要重复执行的SQL语句采用这这种方式
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	// Prepare creates a prepared statement for later queries or executions.
	// 使用Db.Prepare会创建一个指向sql.Stmt接口的引用，即上面的预处理语句
	// 否则，可以向GetPost方法中直接使用Db.Query("xxx")一样将整条SQL语句写入其中，两者都是可行的方式。
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

// Update 更新帖子
func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

// Delete 删除帖子
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
