package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
DOCKER operation:
-------------------------------------------------------------------------------
server:
	docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=luoruofeng -d mysql
create network:
	docker network create mysql-net
add server to network:
	docker network connect mysql-net some-mysql
check network connect complete:
	docker inspect some-mysql

client:
	docker run -it --network mysql-net --rm mysql mysql -hsome-mysql -uroot -p


SQL operation:
-------------------------------------------------------------------------------
	SHOW DATABASES;
	CREATE DATABASE school;
	USE school;

	SHOW TABLES;
	CREATE TABLE student (name VARCHAR(20), sex CHAR(1), birth DATE);
	DESCRIBE student;
	INSERT INTO student (name,sex,birth) VALUES("luo",0,"1983-11-23 00:00:00"),("ruo",1,"1988-4-6 22:14:04"),("feng",0,"1990-12-3 04:00:00");
	SELECT * FROM student;
*/

type Student struct {
	Name  string
	Sex   bool
	Birth time.Time
}

func MysqlApp() {

	db, err := sql.Open("mysql", "root:luoruofeng@/school")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var student Student
	row := db.QueryRow("SELECT name,sex,birth FROM student WHERE name=?", "ruo")
	fmt.Println(row)
	row.Scan(&student)
	fmt.Println(student)
}
