package models

import "appMng/utils/db"

type App struct {
	Id string `orm:"pk;column(id)" json:"id"`
	Name string `orm:"column(name)" json:"name"`
	User string `orm:"column(user)" json:"user"`
	Description string `orm:"column(description)" json:"description"`
	CreatedTime string `orm:"column(createdTime)" json:"createdTime"`
	Lang string `orm:"column(lang)" json:"lang"`
	Type string `orm:"column(type)" json:"type"`
	Services string `orm:"column(services)" json:"services"`
	Git string `orm:"column(git)" json:"git"`
	State string `orm:"column(state)" json:"state"`
}

//查询所有app
func GetApps(userId string) ([]App, error) {
	o := db.GetOrmer()
	var apps []App
	o.QueryTable("app").Filter("user", userId).All(&apps)
	return apps, nil
}

//增加app
func AddApp(app *App) error {
	o := db.GetOrmer()
	_, err := o.Insert(app)
	if err != nil {
		return err
	}
	return nil
}

//删除app
func DeleteApp(appId string) error {
	o := db.GetOrmer()
	_, err := o.Delete(&App{Id:appId})
	if err != nil {
		return err
	}
	return nil
}

//query a app
func GetaApp(appId string) (*App, error) {
	o := db.GetOrmer()
	var app App
	err := o.QueryTable("app").Filter("id", appId).One(&app)
	if err != nil {
		return nil, err
	}
	return &app, nil
}

//query app name is used
func IsNameUsed(user string, name string) bool {
	o := db.GetOrmer()
	var app App
	err := o.QueryTable("app").Filter("user", user).Filter("name", name).One(&app)
	if err != nil {
		return false
	}
	return true
}
