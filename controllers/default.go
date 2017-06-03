package controllers

import (
	"encoding/json"
	"fmt"
	"griddy/models"

	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (main *MainController) Get() {
	main.Data["Form"] = &models.Price{}
	main.Layout = "basic-layout.tpl"
	main.LayoutSections = make(map[string]string)
	main.LayoutSections["Header"] = "header.tpl"
	main.LayoutSections["Footer"] = "footer.tpl"
	main.TplName = "index.tpl"
}

func (main *MainController) View() {
	main.Layout = "basic-layout.tpl"
	main.LayoutSections = make(map[string]string)
	main.LayoutSections["Header"] = "header.tpl"
	main.LayoutSections["Footer"] = "footer.tpl"
	main.TplName = "prices.tpl"

	// var p models.Price
	// json.Unmarshal(main.Ctx.Input.RequestBody, &p)
	// starttime = p.Starttime
	// endtime = p.Endtime

	starttime := main.GetString("starttime")
	endtime := main.GetString("endtime")

	avgprice, err := GetAvgPrice(starttime, endtime)
	if err != nil {
		fmt.Println(err)
	}

	price := models.Price{Starttime: starttime, Endtime: endtime, Avgprice: avgprice}

	o := orm.NewOrm()
	o.Using("default")
	o.Insert(&price)

	var prices []*models.Price
	num, err := o.QueryTable("prices").All(&prices)
	if err != orm.ErrNoRows && num > 0 {
		main.Data["records"] = prices
	}
}

func GetAvgPrice(starttime, endtime string) (float64, error) {
	response, err := http.Get("https://hourlypricing.comed.com/api?type=5minutefeed&datestart=" + starttime + "&dateend=" + endtime)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	// to simplify parsing floats given as string values in JSON, use json.Number
	var prices []map[string]json.Number

	// 1. response.Body is an io.ReadCloser interface
	// 2. json.NewDecoder takes in io.Reader and returns Decoder
	// 3. Decode reads the next JSON-encoded value from its input and
	// stores it in the value pointed to by prices.

	// This also applies to handling this array of JSON objects:
	// https://golang.org/pkg/encoding/json/#Decoder.Decode

	if err := json.NewDecoder(response.Body).Decode(&prices); err != nil {
		return 0, err
	}

	fmt.Println("\nEnergy prices between", starttime, "and", endtime, "\n")
	fmt.Println(prices)

	sum := 0.0
	for _, p := range prices {
		f, err := p["price"].Float64()
		if err != nil {
			return 0, err
		}
		sum += f
	}

	return Round(sum/float64(len(prices)), 0.1), nil
}

//Rounds to nearest "unit"
func Round(x, unit float64) float64 {
	if x > 0 {
		return float64(int64(x/unit+0.5)) * unit
	}
	//handles negative energy prices
	return float64(int64(x/unit-0.5)) * unit
}

func (main *MainController) HelloSitepoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
	main.TplName = "default/hello.tpl"
}
