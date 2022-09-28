package databaseconn

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Salary    int    `json:"salary"`
}
type MysqlConfig struct {
	User   string
	Pass   string
	Host   string
	Port   string
	DbName string
}
type Store struct {
	Db  *sql.DB
	Err error
}

func (k Store) Getdata(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("id")
	var d Employee
	if s == "" {
		w.WriteHeader(http.StatusBadRequest)
		b, err := json.Marshal(d)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(b)
		return
	}
	if k.Err != nil {
		log.Fatalln(k.Err)
	}
	name := k.Db.QueryRow("select * from Employee where id=?", s)
	err := name.Scan(&d.Id, &d.FirstName, &d.LastName, &d.Age, &d.Salary)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no result found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)

}
func (s Store) PostData(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var e Employee
	err = json.Unmarshal(data, &e)
	if e.Id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}
	s.Db.Exec("insert into Employee values(?,?,?,?,?)", e.Id, e.FirstName, e.LastName, e.Age, e.Salary)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("created"))

}
func (s Store) Deldata(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Id"))
		return
	}
	del, err := s.Db.Exec("delete from Employee where id=?", id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid data found"))
		return
	}
	count, err := del.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if count == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("data  found"))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Deleted"))
}
func (s Store) Putdata(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		var e Employee
		err = json.Unmarshal(data, &e)
		if e.Id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
			return
		}

		ins, err := s.Db.Exec("update Employee set id=?,firstname=?,lastname=?,age=?,salary=? where id=?", e.Id, e.FirstName, e.LastName, e.Age, e.Salary, id)

		if err != nil {
			fmt.Println("hii")
			fmt.Println(err, "hii")
		}

		count, err := ins.RowsAffected()

		if err != nil {
			fmt.Println(err)
		}

		if count == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no such id found"))
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("changed"))
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("no id present"))
}
