package endpoints

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/5sigma/go-echo-api/dao"
	"github.com/5sigma/go-echo-api/models"
	"github.com/Jeffail/gabs"
	"github.com/labstack/echo"
)

func TestListWidgets(t *testing.T) {
	h := Handler{DB: dao.NewMemory()}

	u := h.DB.CreateUser(&models.User{})
	h.DB.CreateWidget(&models.Widget{CreatorID: u.ID})
	h.DB.CreateWidget(&models.Widget{CreatorID: u.ID})

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/api/widgets", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("CurrentUser", u)
	err := h.ListWidgets(c)
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
	if !json.Exists("widgets") {
		t.Errorf("Widgets node not present in body\nResponse was: %s\n", resString)
		return
	}

	widgetsArray, _ := json.Path("widgets").Children()
	if len(widgetsArray) != 2 {
		t.Errorf("Widgets array length was %d", len(widgetsArray))
	}
}

func TestGetWidget(t *testing.T) {
	h := Handler{DB: dao.NewMemory()}

	h.DB.CreateWidget(&models.Widget{Name: "test"})

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/api/widgets", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/widgets/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := h.GetWidget(c)
	if err != nil {
		t.Errorf("Error calling handler: %s", err.Error())
	}

	resString := rec.Body.String()

	json, err := gabs.ParseJSON([]byte(resString))
	if err != nil {
		t.Errorf("Error reading JSON: %s\n%s\n", err.Error(), resString)
		return
	}

	if !json.Path("name").Exists() {
		t.Errorf("name not present in response.\n%s\n", resString)
	} else {
		if v, _ := json.Path("name").Data().(string); v != "test" {
			t.Errorf("name not valid in response: %s", v)
		}
	}
}

func TestCreateWidget(t *testing.T) {
	h := Handler{DB: dao.NewMemory()}
	creator := h.DB.CreateUser(&models.User{})

	payload := `
		{
			"name": "My Widget",
			"description": "This is my widget",
			"price": 1.99
		}
	`

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/api/widgets", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/widgets")
	c.Set("user", creator)

	err := h.CreateWidget(c)
	if err != nil {
		t.Errorf("Error calling handler: %s", err.Error())
		return
	}

	resString := rec.Body.String()
	json, err := gabs.ParseJSON([]byte(resString))

	if v, _ := json.Path("name").Data().(string); v != "My Widget" {
		t.Errorf("name not valid in response: %s", v)
	}

	if v, _ := json.Path("description").Data().(string); v != "This is my widget" {
		t.Errorf("description not valid in response: %s", v)
	}

	if v, _ := json.Path("price").Data().(float64); v != 1.99 {
		t.Errorf("price not valid in response: %f", v)
	}

	if v, _ := json.Path("creator.id").Data().(float64); uint(v) != creator.ID {
		t.Errorf("creator.id not valid in response: %d", uint(v))
	}

	widgets := h.DB.AllWidgets()
	if len(widgets) != 1 {
		t.Errorf("Widget count in database incorrect: %d", len(widgets))
	}
}

func TestUpdateWidget(t *testing.T) {
	h := Handler{DB: dao.NewMemory()}
	h.DB.CreateWidget(&models.Widget{})
	payload := `
		{
			"id": 1,
			"name": "My Widget",
			"description": "This is my widget",
			"price": 1.99
		}
	`

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/api/widgets", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/widgets")

	err := h.UpdateWidget(c)
	if err != nil {
		t.Errorf("Error calling handler: %s", err.Error())
		return
	}

	resString := rec.Body.String()
	json, err := gabs.ParseJSON([]byte(resString))

	if v, _ := json.Path("name").Data().(string); v != "My Widget" {
		t.Errorf("name not valid in response: %s", v)
	}

	if v, _ := json.Path("description").Data().(string); v != "This is my widget" {
		t.Errorf("description not valid in response: %s", v)
	}

	if v, _ := json.Path("price").Data().(float64); v != 1.99 {
		t.Errorf("price not valid in response: %f", v)
	}

	widget := h.DB.GetWidgetByID(1)
	if widget.Name != "My Widget" {
		t.Errorf("Name not updated: '%s'", widget.Name)
	}
}

func TestDeleteWidget(t *testing.T) {
	h := Handler{DB: dao.NewMemory()}

	h.DB.CreateWidget(&models.Widget{Name: "test"})

	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/api/widgets", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/widgets/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := h.DeleteWidget(c)
	if err != nil {
		t.Errorf("Error calling handler: %s", err.Error())
	}

	widgets := h.DB.AllWidgets()
	if len(widgets) > 0 {
		t.Errorf("Widget not deleted from database")
	}
}

func TestWidgetIntegration(t *testing.T) {
	var (
		client   = &http.Client{}
		err      error
		req      *http.Request
		res      *http.Response
		resBytes []byte
	)

	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Server setup
	e := echo.New()
	h := Handler{
		DB: dao.NewMemory(),
	}
	Setup(e, h)
	ts := httptest.NewServer(e)
	defer ts.Close()

	//Data setup
	user := h.DB.CreateUser(&models.User{})
	apiKey := h.DB.CreateAPIKeyForUser(user)
	h.DB.CreateWidget(&models.Widget{CreatorID: user.ID})
	h.DB.CreateWidget(&models.Widget{CreatorID: user.ID})
	h.DB.CreateWidget(&models.Widget{CreatorID: user.ID})
	h.DB.CreateWidget(&models.Widget{CreatorID: user.ID + 1})

	// Make request
	req, err = http.NewRequest("GET", fmt.Sprintf("%s/widgets", ts.URL), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey.Key))
	if err != nil {
		t.Errorf("Request error: %s", err.Error())
	}

	res, err = client.Do(req)
	if err != nil {
		t.Errorf("Request error: %s", err.Error())
	}
	resBytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading result: %s", resBytes)
	}
	if res.StatusCode != 200 {
		t.Errorf("Non 200 response: %d\n%s", res.StatusCode, string(resBytes))
	}

	json, err := gabs.ParseJSON(resBytes)
	if err != nil {
		t.Errorf("Error reading JSON: %s", err.Error())
	}
	if !json.Exists("widgets") {
		t.Errorf("Widgets node not present in body\nResponse was: %s\n",
			string(resBytes))
		return
	}

	widgetsArray, _ := json.Path("widgets").Children()
	if len(widgetsArray) != 3 {
		t.Errorf("Widgets array length was %d", len(widgetsArray))
	}
}
