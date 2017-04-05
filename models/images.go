package models

import "appMng/utils/db"

type Image struct {
	Id string `orm:"pk;column(id)" json:"id"`
	Name string `orm:"column(name)" json:"name"`
	Tag string `orm:"column(tag)" json:"tag"`
	User string `orm:"column(user)" json:"user"`
	AppId string `orm:"column(appId)" json:"appId"`
	Lang string `orm:"column(lang)" json:"lang"`
	Git string `orm:"column(git)" json:"git"`
	Img string `orm:"column(img)" json:"img"`
	CreatedTime string `orm:"column(createdTime)" json:"createdTime"`
	State string `orm:"column(state)" json:"state"`
}


//query all images of an app
func GetImages(appId string) ([]Image, error) {
	o := db.GetOrmer()
	var imgs []Image
	o.QueryTable("image").Filter("appId", appId).All(&imgs)
	return imgs, nil
}

//add an image
func AddImage(img *Image) error {
	o := db.GetOrmer()
	_, err := o.Insert(img)
	if err != nil {
		return err
	}
	return nil
}

//delete an image
func DeleteImage(imgId string) error {
	o := db.GetOrmer()
	_, err := o.Delete(&Image{Id:imgId})
	if err != nil {
		return err
	}
	return nil
}

//query an image
func GetaImage(imgId string) (*Image, error) {
	o := db.GetOrmer()
	var img Image
	err := o.QueryTable("image").Filter("id", imgId).One(&img)
	if err != nil {
		return nil, err
	}
	return &img, nil
}


// set image build error state
func SetImageBuildStatus(imageId string, state string) (error) {
	o := db.GetOrmer()
	img := Image{Id:imageId}
	errRead := o.Read(&img)
	if errRead != nil {
		return errRead
	}
	img.State = state
	_, err := o.Update(&img, "state")
	if err != nil {
		return err
	}
	return nil
}



