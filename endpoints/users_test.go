package endpoints

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/5sigma/go-echo-api/dao"
	"github.com/Jeffail/gabs"
	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type epFunc func(echo.Context) error

func TestListUsers(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gDB, _ := gorm.Open("postgres", db)
	h := Handler{DB: &dao.DAO{DB: gDB}}

	mockRows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email_address"})
	mockRows.AddRow(1, "alice", "alisson", "test@test.com")
	mock.ExpectQuery("SELECT").WillReturnRows(mockRows)

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/api/v1/users", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := h.ListUsers(c)
	if err != nil {
		t.Errorf("Error calling handler: %s", err.Error())
	}

	resString := rec.Body.String()
	if err != nil {
		t.Errorf("Error calling endpoint: %s", err.Error())
	}

	json, err := gabs.ParseJSON([]byte(resString))
	if err != nil {
		t.Errorf("Error reading JSON: %s", err.Error())
	}
	if !json.Exists("users") {
		t.Errorf("Users node not present in body\nResponse was: %s\n", resString)
		return
	}

	usersArray, _ := json.Path("users").Children()
	if len(usersArray) != 1 {
		t.Errorf("Users array length was %d", len(usersArray))
	}

}

func TestGetUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gDB, _ := gorm.Open("postgres", db)
	h := Handler{DB: &dao.DAO{DB: gDB}}

	mockRows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email_address"})
	mockRows.AddRow(1, "alice", "alisson", "test@test.com")
	mock.ExpectQuery("id=1").WillReturnRows(mockRows)

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/api/v1/users", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := h.GetUser(c)
	if err != nil {
		t.Errorf("Error calling handler: %s", err.Error())
	}

	resString := rec.Body.String()

	json, err := gabs.ParseJSON([]byte(resString))
	if err != nil {
		t.Errorf("Error reading JSON: %s", err.Error())
	}

	if !json.Path("firstName").Exists() {
		t.Errorf("first name not present in response.\n%s\n", resString)
	}

}
