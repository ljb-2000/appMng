package tpl

import "testing"
import "appMng/models"

func TestBuildImg(t *testing.T) {
	img := models.Image{}
	img.User = "luocheng"
	img.Name = "testapp4"
	img.Tag = "3.0"
	img.Id = "abc123456"
	appName := "testapp"
	BuildImg(img, appName)
}
