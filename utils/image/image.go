package image

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
	"os/exec"
	"io/ioutil"
	"github.com/jhoonb/archivex"
	"encoding/json"
	"encoding/base64"
	"log"
	//"appMng/utils/git"

	"appMng/utils/git"
	"appMng/utils/dir"
	"appMng/utils/tpl"
)

func Ps() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}

func ListImage()  {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Printf("%s %s\n", image.ID, image.RepoTags)
	}
}

func BuildImage(user, name, tag, giturl, lang string) (imgAddr string) {
	//first clone project
	git.CloneRepo(giturl)

	curPwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(curPwd)

	tags := "registry.time-track.cn:8443/" + user + "/" + name + ":" + tag
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	tar := new(archivex.TarFile)
	tar.Create("/Users/luocheng/build/final/archieve")
	tar.AddAll("/Users/luocheng/build/source/graph", false)
	tar.Close()

	dockerBuildContext, err := os.Open("/Users/luocheng/build/final/archieve.tar")
	defer dockerBuildContext.Close()

	options := types.ImageBuildOptions{
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     false,
		Tags:   []string{tags},
	}

	buildResponse, err := cli.ImageBuild(context.Background(), dockerBuildContext, options)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("********* %s **********", buildResponse.OSType)
	response, err := ioutil.ReadAll(buildResponse.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Println(string(response))

	imgAddr = tags
	return
}


func PushImage(imgAddr string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	auth := types.AuthConfig{
		Username: "innovation",
		Password: "0p;/(OL>",
	}
	authBytes, _ := json.Marshal(auth)
	authBase64 := base64.URLEncoding.EncodeToString(authBytes)

	options := types.ImagePushOptions{
		RegistryAuth: authBase64,
	}

	pr, ierr := cli.ImagePush(context.Background(), imgAddr, options)
	if ierr != nil {
		fmt.Printf("%s", ierr.Error())
	} else {
		log.Println("image push succeed.")
	}
	pr.Close()
	log.Println("finishi push.")
}

func PushImageBack(imgAddr string) error {
	cmd := "docker"
	args := []string{"push", imgAddr}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("finishi push.")
	return nil
}



