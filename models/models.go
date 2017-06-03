package models

import "github.com/astaxie/beego/orm"

type Price struct {
	Id        int     `form:"-"`
	Starttime string  `form:"starttime,text,starttime:"`
	Endtime   string  `form:"endtime,text,endtime:"`
	Avgprice  float64 `form:"avgprice,text,avgprice:"`
}

func (p *Price) TableName() string {
	return "prices"
}

func init() {
	orm.RegisterModel(new(Price))
}
