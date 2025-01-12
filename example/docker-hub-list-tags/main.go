package main

import (
	"fmt"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/logs"
)

func init() {
	logs.Debug.SetOutput(os.Stdout)
	logs.Progress.SetOutput(os.Stdout)
}

func main() {
	imageUrl := os.Getenv("DH_IMAGE_URL")
	uname := os.Getenv("DH_USERNAME")
	pwd := os.Getenv("DH_PASSWORD")

	opts := crane.WithAuth(&authn.Basic{
		Username: uname,
		Password: pwd,
	})

	tags, err := crane.ListTags(imageUrl, opts)
	if err != nil {
		fmt.Println("list tags failed", err.Error())
	}
	fmt.Println("tags", tags)
}
