package test

import (
	"bytes"
	_ "griddy/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"fmt"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestEnergyPrices(t *testing.T) {
	var jsonStr = []byte(`{"starttime":"201506031105","endtime":"201506031200"}`)
	r, _ := http.NewRequest("POST", "/prices", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	fmt.Println(bytes.NewBuffer(jsonStr))
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestEnergyPrices", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Gets Average Energy Prices\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
