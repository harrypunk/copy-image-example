package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/logs"
)

type imageInfo struct {
	SrcUrl       string `json:"src_image_name"`
	SrcUsername  string `json:"src_username"`
	SrcPassword  string `json:"src_password"`
	DestUrl      string `json:"dest_image_name"`
	DestUsername string `json:"dest_username"`
	DestPassword string `json:"dest_password"`
}

const DOCKER_IO = "docker.io"

func (o imageInfo) Resolve(resource authn.Resource) (authn.Authenticator, error) {
	reg := resource.RegistryStr()
	if strings.Contains(reg, DOCKER_IO) {
		log.Println("docker.io resolve", reg)
		return &authn.Basic{
			Username: o.SrcUsername,
			Password: o.SrcPassword,
		}, nil
	}
	log.Println("not docker.io", reg)
	return &authn.Basic{
		Username: o.DestUsername,
		Password: o.DestPassword,
	}, nil
}

func init() {
	// logs.Debug.SetOutput(os.Stdout)
	logs.Progress.SetOutput(os.Stdout)
}

func main() {
	log.Println("copy start")
	for i, v := range os.Args {
		log.Println(i, v)
	}
	// pass file path as an arg
	filePath := os.Args[1]
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln("read file failed", err.Error())
	}

	var info imageInfo
	err = json.Unmarshal(bytes, &info)
	if err != nil {
		log.Fatalln("json unmarshal failed", err.Error())
	}

	keych := authn.NewMultiKeychain(info)
	err = crane.Copy(info.SrcUrl, info.DestUrl, crane.WithAuthFromKeychain(keych))
	if err != nil {
		log.Fatalln("copy failed", err.Error())
	}
	log.Println("copy finished")
}
