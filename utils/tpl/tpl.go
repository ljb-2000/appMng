package tpl

import (
	"log"
	"text/template"
	"io/ioutil"
	"os"
	"os/exec"
	"appMng/models"
	"github.com/astaxie/beego"
	"github.com/kardianos/osext"
	"github.com/astaxie/beego/logs"
)

type Build struct {
	Git string
	Reg string
	User string
	App string
	Img string
	Tag string
}

var gRegUrl string

func init()  {
	gRegUrl = beego.AppConfig.String("regurl")
}

func BuildImg(img models.Image, appName string) {

	path := GetExecutableDir()
	beego.Debug(path)

	buildFile := path + "/buildgo.sh"
	if img.Lang == "Go" {
		buildFile = path + "/buildgo.sh"
	} else if img.Lang == "Python" {
		buildFile = path + "/buildpython.sh"
	} else if img.Lang == "Javascript" {
		buildFile = path + "/buildjs.sh"
	}

	buff, err := ioutil.ReadFile(buildFile)
	if err != nil {
		logs.Error(err.Error())
	}
	buildString := string(buff)

	b := Build{}
	b.Git = img.Git
	b.Reg = gRegUrl
	b.User = img.User
	b.App = appName
	b.Img = img.Name
	b.Tag = img.Tag

	t := template.Must(template.New("templates").Parse(buildString))
	filename := path + "/buildfiles/build" + img.User + img.Name + img.Tag + ".sh"
	f, err := os.Create(filename)

	err = t.Execute(f, b)
	if err != nil {
		logs.Error("execute template error: %s", err.Error())
	}

	cmd := exec.Command("chmod", "755", filename)
	out, err := cmd.Output()
	if err != nil {
		logs.Error(string(out))
	}

	var state string
	cmd = exec.Command("/bin/sh", filename)
	out, err = cmd.Output()
	if err != nil {
		logs.Debug(string(out))
		logs.Error("execute file: %s failed, error: %s, output: %s", filename, string(out), err.Error())
		state = "Error"
	} else {
		state = "Succeed"
	}

	models.SetImageBuildStatus(img.Id, state)
}

func GetExecutableDir() (dir string) {
	dir, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	return
}