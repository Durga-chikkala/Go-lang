package Driver

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db  *sql.DB
	err error
}
type Car struct {
	id         int
	name       string
	model      string
	engineType string
}

func (s Store) Set(c Car) bool {
	if c.id == 0 {
		return false
	} else {

		insert, err := s.db.Exec("INSERT INTO CARS VALUES(%d, '%v' ,'%v','%v')", c.id, c.name, c.model, c.engineType)
		//s.db.Exec()
		if err != nil {
			return false
		}
		count, err := insert.RowsAffected()
		if count > 0 {
			return true
		}
		return false
	}
}

func (s Store) Get(id int) Car {
	var c Car
	k := 0

	sel, err := s.db.Query("select * from CARS where id=%v", id)
	if err != nil {
		return Car{}
	}
	for sel.Next() {
		k = 1
		err := sel.Scan(&c.id, &c.name, &c.model, &c.engineType)
		if err != nil {
			c.id = -1
		}
		return c
	}
	if k == 0 {
		return Car{}
	}
	return c
}
func (s Store) Delete(id int) bool {

	delete, err := s.db.Exec("delete from CARS where id=%v", id)
	if err != nil {

		return false
	}
	count, err := delete.RowsAffected()

	if count > 0 {
		return true
	}
	fmt.Println(count)
	return false
}
