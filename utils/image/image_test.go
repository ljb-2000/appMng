package image

import "testing"

func TestPs(t *testing.T) {
	Ps()
}

func TestListImage(t *testing.T) {
	ListImage()
}

func TestBuildImage(t *testing.T) {
	BuildImage("luocheng", "gg", "3.0", "giturl", "go")
}

func TestPushImage(t *testing.T) {
	PushImage("registry.time-track.cn:8443/luocheng/gg:3.0")
}

func TestPushImageBack(t *testing.T) {
	PushImageBack("registry.time-track.cn:8443/luocheng/gg:2.0")
}