package databaseconn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHand(t *testing.T) {
	testcases := []struct {
		desc       string
		method     string
		additional string
		expected   Employee
	}{
		{"name is given", "GET", "1", Employee{"1", "Durga", "Chikkala", 21, 70000}},
		{"no data present", "GET", "4", Employee{}},
		{"empty name is given", "GET", "", Employee{}},
	}
	var s Store
	db, mock, err := sqlmock.New()
	s.Db = db
	s.Err = err
	for i, v := range testcases {
		rows := sqlmock.NewRows([]string{"Id", "firstName", "lastName", "age", "salary"}).
			AddRow(v.expected.Id, v.expected.FirstName, v.expected.LastName, v.expected.Age, v.expected.Salary)
		mock.ExpectQuery("select (.+) from Employee where id=?").WithArgs(v.additional).WillReturnRows(rows).WillReturnError(err)
		req := httptest.NewRequest("GET", "localhost:8080/?id="+v.additional, nil)
		res := httptest.NewRecorder()
		s.Getdata(res, req)
		result := res.Result()
		e, err := io.ReadAll(result.Body)
		if err != nil {
			fmt.Println(err)
		}
		var out Employee
		err = json.Unmarshal(e, &out)
		if err != nil {
			fmt.Println(err)
		}

		if out.Id != v.expected.Id {
			t.Errorf("testcase failed [%v] got %v expected %v", i, out.Id, v.expected.Id)
		}

	}

}
func TestStore_PostData(t *testing.T) {
	db, mock, err := sqlmock.New()
	var s Store
	testcases := []struct {
		desc     string
		details  Employee
		expected int
	}{
		{"inserting data", Employee{"1", "sujatha", "chikkala", 42, 200000}, http.StatusCreated},
		{"inserting data", Employee{}, http.StatusBadRequest},
	}
	s.Db = db
	s.Err = err
	for i, v := range testcases {
		mock.ExpectExec("insert into Employee values ?,?,?,?").WithArgs(v.details.Id, v.details.FirstName, v.details.LastName, v.details.Age, v.details.Salary).WillReturnResult(sqlmock.NewResult(1, 1))
		data, err := json.Marshal(v.details)
		if err != nil {
			fmt.Println(err)
		}
		rdata := bytes.NewReader(data)
		req := httptest.NewRequest("POST", "/post", rdata)
		res := httptest.NewRecorder()
		s.PostData(res, req)
		status := res.Result().StatusCode
		if status != v.expected {
			t.Errorf("Test cases failed [%v] got %v expected %v", i, status, v.expected)
		}

	}
}
func TestStore_Deldata(t *testing.T) {
	db, mock, err := sqlmock.New()
	var s Store
	testcases := []struct {
		desc     string
		id       string
		expected int
		err      error
	}{
		{"deleting present data", "1", http.StatusAccepted, nil},
		{"deleteing not present data", "10", http.StatusBadRequest, errors.New("err")},
	}
	s.Db = db
	s.Err = err
	for i, v := range testcases {
		mock.ExpectExec("delete from Employee where id=?").WithArgs(v.id).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(v.err)
		req := httptest.NewRequest("DELETE", "localhost:8080/delete?id="+v.id, nil)
		res := httptest.NewRecorder()
		s.Deldata(res, req)
		status := res.Result().StatusCode
		if status != v.expected {
			t.Errorf("Test cases failed [%v] got %v expected %v", i, status, v.expected)
		}

	}
}
func TestStore_Putdata(t *testing.T) {
	db, mock, err := sqlmock.New()
	var s Store
	testcases := []struct {
		desc     string
		id       string
		details  Employee
		expected int
	}{
		{"inserting data", "1", Employee{"1", "sujatha", "chikkala", 42, 200000}, http.StatusAccepted},
		{"id not present", "", Employee{"5", "durga", "chikkala", 21, 2000}, http.StatusBadRequest},
	}
	s.Db = db
	s.Err = err
	for i, v := range testcases {

		mock.ExpectExec("update Employee set").WithArgs(v.details.Id, v.details.FirstName, v.details.LastName, v.details.Age, v.details.Salary, v.id).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(err)
		data, err := json.Marshal(v.details)

		if err != nil {
			fmt.Println(err)
		}
		rdata := bytes.NewReader(data)
		req := httptest.NewRequest("PUT", "localhost:8080/putdata?id="+v.id, rdata)
		res := httptest.NewRecorder()

		s.Putdata(res, req)

		status := res.Result().StatusCode
		if status != v.expected {
			t.Errorf("Test cases failed [%v] got %v expected %v", i, status, v.expected)
		}

	}
}
