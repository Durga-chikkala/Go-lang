package Driver

import (
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}
	testcases := []struct {
		car      Car
		expected bool
		desc     string
		err      error
	}{
		{Car{1, "vitara", "330", "petrol"}, true, "sending car details", nil},
		{Car{2, "Tata", "5440", "Diesel"}, true, "sending tata details", nil},
		{Car{0, "hhhj", "330", "petrol"}, false, "with no car name", errors.New("e")},
	}
	var s Store
	s.db = db
	s.err = err
	if s.err != nil {
		panic(s.err)
	}
	for i, v := range testcases {
		mock.ExpectExec("INSERT INTO CARS VALUES ?,?,?,?").WithArgs(v.car.id, v.car.name, v.car.model, v.car.engineType).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(v.err)
		output := s.Set(v.car)
		if output != v.expected {
			t.Errorf("test cases failed[%v] expected %v got %v", i, v.expected, output)
		}
	}
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}
	var s Store
	s.db = db
	s.err = err
	testcases := []struct {
		id       int
		expected Car
		desc     string
	}{
		{id: 1, expected: Car{1, "vitara", "330", "petrol"}, desc: "sending car details"},
		{id: 2, expected: Car{2, "Tata", "5440", "Diesel"}, desc: "sending tata details"},
		{id: 4, expected: Car{}, desc: "with no car name"},
	}

	for i, v := range testcases {
		rows := sqlmock.NewRows([]string{"Id", "Name", "Model", "EngineType"}).
			AddRow(v.expected.id, v.expected.name, v.expected.model, v.expected.engineType)
		mock.ExpectQuery("select (.+) from CARS where id=?").WithArgs(v.id).WillReturnRows(rows).WillReturnError(err)
		output := s.Get(v.id)
		if output != v.expected {
			t.Errorf("test case failed [%v] expected %v got %v", i, v.expected, output)
		}

	}

}

func TestDelete(t *testing.T) {
	testcases := []struct {
		id       int
		expected bool
		desc     string
		err      error
	}{
		{1, true, "deleting car ", nil},
		{2, true, "deleting car details", nil},
		{5, false, "with out of range", errors.New("err")},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}
	var s Store
	s.db = db
	s.err = err
	for i, v := range testcases {

		mock.ExpectExec("delete from CARS where id=?").WithArgs(v.id).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(v.err)
		output := s.Delete(v.id)
		if output != v.expected {
			t.Errorf("test case failed [%v] expected %v got %v %v", i, v.expected, output, v.err)

		}

	}
}
