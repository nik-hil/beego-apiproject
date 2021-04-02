package test

import (
	_ "apiproject/routers"
	"bytes"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/user", nil)
	response := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(response, request)

	logs.Info("testing", "TestGet", "Code[%d]\n%s", response.Code, response.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(response.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(response.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestPost(t *testing.T) {
	var jsonStr = []byte(`{"Username": "nikhil1", "Password": "nikhil1", "Email": "email1@domain.com"}`)
	request, err := http.NewRequest("POST", "/v1/user", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(response, request)
	logs.Info("testing", "TestPost", "Code[%d]\n%s", response.Code, response.Body.String())
	if status := response.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
