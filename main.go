package main

import (
	"encoding/json"
	"log"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
)

type ImageInfo struct {
	SrcUrl       string `json:"src-url"`
	SrcUsername  string `json:"src-username"`
	SrcPassword  string `json:"src-password"`
	DestUrl      string `json:"dest-url"`
	DestUsername string `json:"dest-username"`
	DestPassword string `json:"dest-password"`
}

func doCopy(inputEvent string) error {
	log.Println("copy start")
	var info ImageInfo
	err := json.Unmarshal([]byte(inputEvent), &info)
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
	log.Println("copy ok")

	return nil
}

func main() {
}
