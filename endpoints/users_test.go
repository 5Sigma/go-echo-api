package endpoints

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/5sigma/go-echo-api/dao"
	"github.com/5sigma/go-echo-api/models"
	"github.com/Jeffail/gabs"
	"github.com/labstack/echo"
)

func TestListUsers(t *testing.T) {
	h := Handler{DB: dao.NewMemory()}

	h.DB.CreateUser(models.User{})
	h.DB.CreateUser(models.User{})

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/api/users", strings.NewReader(""))
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
	if len(usersArray) != 2 {
		t.Errorf("Users array length was %d", len(usersArray))
	}
}

func TestGetUser(t *testing.T) {
	h := Handler{DB: dao.NewMemory()}

	h.DB.CreateUser(models.User{FirstName: "test"})

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/api/users", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := h.GetUser(c)
	if err != nil {
		t.Errorf("Error calling handler: %s", err.Error())
	}

	resString := rec.Body.String()

	json, err := gabs.ParseJSON([]byte(resString))
	if err != nil {
		t.Errorf("Error reading JSON: %s\n%s\n", err.Error(), resString)
		return
	}

	if !json.Path("firstName").Exists() {
		t.Errorf("first name not present in response.\n%s\n", resString)
	} else {
		if v, _ := json.Path("firstName").Data().(string); v != "test" {
			t.Errorf("first name not valid in response: %s", v)
		}
	}

}
