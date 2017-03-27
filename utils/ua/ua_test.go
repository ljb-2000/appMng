package ua

import (
	"testing"
	"log"
)

func TestGetUserPwd(t *testing.T) {
	name, pwd := GetUserNamePwd("18022222222")
	log.Println(name + ": " + pwd)
}
