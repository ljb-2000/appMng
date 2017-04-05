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

)

type Build struct {
	Git string
	Reg string
	User string
	App string
	Img string
	Tag string
}

var gGitUrl string
var gRegUrl string

func init()  {
	gGitUrl = beego.AppConfig.String("gogsurl")
	gRegUrl = beego.AppConfig.String("regurl")
}


func BuildImg(img models.Image, appName string) {

	dir := GetExecutableDir()
	beego.Debug(dir)

	buildFile := dir + "/buildgo.sh"
	if img.Lang == "Go" {
		buildFile = dir + "/buildgo.sh"
	} else if img.Lang == "Python" {
		buildFile = dir + "/buildpython.sh"
	} else if img.Lang == "Javascript" {
		buildFile = dir + "/buildjs.sh"
	}

	buff, err := ioutil.ReadFile(buildFile)
	if err != nil {
		log.Println(err.Error())
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

	filename := dir + "buildfiles/build" + img.User + img.Name + img.Tag + ".sh"
	f, err := os.Create(filename)

	err = t.Execute(f, b)
	if err != nil {
		log.Println("executing template:", err)
	}

	cmd := exec.Command("chmod", "755", filename)
	out, err := cmd.Output()
	if err != nil {
		log.Println(string(out))
	}

	var state string
	cmd = exec.Command("/bin/sh", filename)
	out, err = cmd.Output()
	if err != nil {
		log.Println(string(out))
		state = "build error: " + string(out)
	} else {
		state = "succeed"
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