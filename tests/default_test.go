package test

import (
	_ "griddy/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/astaxie/beego"

	"griddy/controllers"

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

func TestGetAvgEnergyPrice(t *testing.T) {
	//Use this code to test HTTP calls. We don't need it here.

	//var jsonStr = []byte(`{"starttime":"201506031105","endtime":"201506031200"}`)
	// r, err := http.NewRequest("POST", "/prices", bytes.NewBuffer(jsonStr))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// r.Header.Set("Content-Type", "application/json")
	// w := httptest.NewRecorder()

	// beego.BeeApp.Handlers.ServeHTTP(w, r)
	// beego.Trace("testing", "TestGetAvgEnergyPrice", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Gets Average Energy Prices\n", t, func() {
		avg1, _ := controllers.GetAvgPrice("201506031105", "201506031200")
		avg2, _ := controllers.GetAvgPrice("201606031105", "201606031200")
		avg3, _ := controllers.GetAvgPrice("201706021200", "201706021220")

		Convey("Average energy price", func() {
			So(strconv.FormatFloat(avg1, 'f', 1, 64), ShouldEqual, "2.7")
			So(strconv.FormatFloat(avg2, 'f', 1, 64), ShouldEqual, "3.8")
			So(strconv.FormatFloat(avg3, 'f', 1, 64), ShouldEqual, "8.7")
		})
	})
}
