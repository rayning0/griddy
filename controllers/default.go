package controllers

import (
	"encoding/json"
	"fmt"
	"griddy/models"
	"io/ioutil"
	"strconv"

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

	starttime := main.GetString("starttime")
	endtime := main.GetString("endtime")
	avgprice := getAvgPrice(starttime, endtime)

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

type arrayOfMaps []map[string]string

func getAvgPrice(starttime, endtime string) float64 {
	response, err := http.Get("https://hourlypricing.comed.com/api?type=5minutefeed&datestart=" + starttime + "&dateend=" + endtime)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	energyJSON, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var energyPrices arrayOfMaps
	err = json.Unmarshal(energyJSON, &energyPrices)

	fmt.Println("Energy prices between", starttime, "and", endtime)
	fmt.Println(energyPrices)

	var sum float64
	var size int
	for _, p := range energyPrices {
		price, _ := strconv.ParseFloat(p["price"], 64)
		sum += price
		size++
	}
	avg := Truncate(sum / float64(size))
	fmt.Println("Average price:", avg)
	return avg
}

//Truncate a float to 2 levels of precision
func Truncate(some float64) float64 {
	return float64(int(some*10)) / 10
}

func (main *MainController) HelloSitepoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
	main.TplName = "default/hello.tpl"
}
