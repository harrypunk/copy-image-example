package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
)

type ImageInfo struct {
	SrcUrl       string `json:"src_image_name"`
	SrcUsername  string `json:"src_username"`
	SrcPassword  string `json:"src_password"`
	DestUrl      string `json:"dest_image_name"`
	DestUsername string `json:"dest_username"`
	DestPassword string `json:"dest_password"`
}

func doCopy(inputEvent []byte) error {
	var info ImageInfo
	err := json.Unmarshal(inputEvent, &info)
	if err != nil {
		return err
	}

	srcOption := crane.WithAuth(&authn.Basic{
		Username: info.SrcUsername,
		Password: info.SrcPassword,
	})
	destOption := crane.WithAuth(&authn.Basic{
		Username: info.DestUsername,
		Password: info.DestPassword,
	})

	err = crane.Copy(info.SrcUrl, info.DestUrl, srcOption, destOption)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	log.Println("copy start")
	filePath := os.Args[0]
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln("read file failed", err.Error())
	}
	err = doCopy(bytes)
	if err != nil {
		log.Fatalln("copy failed", err.Error())
	}

	log.Println("copy ok")
}
