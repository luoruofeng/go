package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
)
import _ "github.com/go-sql-driver/mysql"

/*
DOCKER operation:
-------------------------------------------------------------------------------
server:
	docker run --name some-mysql -p 8888:3306 -e MYSQL_ROOT_PASSWORD=luoruofeng -d mysql
if some-mysql container is existed:
	docker start some-mysql
create network:
	docker network create mysql-net
add server to network:
	docker network connect mysql-net some-mysql
check network connect complete:
	docker inspect some-mysql

client:
	docker run -it --network mysql-net --rm mysql mysql -hsome-mysql -uroot -p
or:
	docker run -it --rm mysql mysql -h172.17.0.2 -uroot -p

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
type Service struct {
	db *sql.DB
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := s.db
	switch r.URL.Path {
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	case "/luo":
		// This is a short SELECT. Use the request context as the base of the context timeout.
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		name := "luo"
		var birth mysql.NullTime
		var birthTime time.Time
		err := db.QueryRowContext(ctx, `
select
	birth
from
	student
where
	name = ?
;`,
			name,
		).Scan(&birth)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		birthTime = birth.Time
		io.WriteString(w, birthTime.Format("2006-01-02 15:04:05"))
		return
	case "/names":
		// This is a long SELECT. Use the request context as the base of
		// the context timeout, but give it some time to finish. If
		// the client cancels before the query is done the query will also
		// be canceled.
		ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
		defer cancel()

		var names []string
		rows, err := db.QueryContext(ctx, "select name from student where sex = false;")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for rows.Next() {
			var name string
			err = rows.Scan(&name)
			if err != nil {
				break
			}
			names = append(names, name)
		}
		// Check for errors during rows "Close".
		// This may be more important if multiple statements are executed
		// in a single batch and rows were written as well as read.
		if closeErr := rows.Close(); closeErr != nil {
			http.Error(w, closeErr.Error(), http.StatusInternalServerError)
			return
		}

		// Check for row scan error.
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check for errors during row iteration.
		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(names)
		return
		// case "/async-action":
		// 	//async-action:
		// 	// This action has side effects that we want to preserve
		// 	// even if the client cancels the HTTP request part way through.
		// 	// For this we do not use the http request context as a base for
		// 	// the timeout.
		// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// 	defer cancel()

		// 	var orderRef = "ABC123"
		// 	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
		// 	//存储过程
		// 	_, err = tx.ExecContext(ctx, "stored_proc_name", orderRef)

		// 	if err != nil {
		// 		tx.Rollback()
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}
		// 	err = tx.Commit()
		// 	if err != nil {
		// 		http.Error(w, "action in unknown state, check state before attempting again", http.StatusInternalServerError)
		// 		return
		// 	}
		// 	w.WriteHeader(http.StatusOK)
		// 	return
	}
}

func MysqlApp() {

	db, err := sql.Open("mysql", "root:luoruofeng@tcp(localhost:8888)/school")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	s := &Service{db: db}

	http.ListenAndServe(":8080", s)

	var student Student
	row := db.QueryRow("SELECT name,sex,birth FROM student WHERE name=?", "ruo")
	fmt.Println(row)
	row.Scan(&student)
	fmt.Println(student)
}
