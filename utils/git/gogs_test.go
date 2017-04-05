package git

import (
	"testing"
	"log"
)

func TestCreateRepo(t *testing.T) {
	user := "luocheng"
	pwd := "lc08170819"
	appname := "test0320"
	desc := "test"
	giturl := CreateRepo(user, pwd, appname, desc)
	log.Println(giturl)
}

func TestCloneRepo(t *testing.T) {
	CloneRepo("")
}

func TestDeleteRepo(t *testing.T) {
	DeleteRepo("luocheng", "123456", "app1")
}