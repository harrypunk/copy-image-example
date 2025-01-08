package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

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

// Error: fetching "registry.example.com/namespace123/hello-world:latest": GET https://registry.example.com/service/token?scope=repository%namespace123%2Fhello-world%3Apull&service=token-service: unexpected status code 401 Unauthorized
//
// func (info ImageInfo) Get(serverURL string) (string, string, error) {
// 	log.Println("ImageInfo helper serverURL", serverURL)
// 	if strings.Contains(info.SrcUrl, serverURL) {
// 		return info.SrcPassword, info.SrcPassword, nil
// 	} else {
// 		return info.DestPassword, info.DestPassword, nil
// 	}
// }

// Resolve implements Keychain.
func (info ImageInfo) Resolve(target authn.Resource) (authn.Authenticator, error) {
	reg := target.RegistryStr()
	log.Println("ImageInfo Resolve", reg)
	if strings.Contains(info.SrcUrl, reg) {
		return &authn.Basic{
			Username: info.SrcUsername,
			Password: info.SrcPassword,
		}, nil
	} else {
		return &authn.Basic{
			Username: info.DestUsername,
			Password: info.DestPassword,
		}, nil
	}
}

func doCopy(inputEvent []byte) error {
	var info ImageInfo
	err := json.Unmarshal(inputEvent, &info)
	if err != nil {
		return err
	}

	keych := authn.NewMultiKeychain(info)

	err = crane.Copy(info.SrcUrl, info.DestUrl, crane.WithAuthFromKeychain(keych))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	log.Println("copy start")
	for i, v := range os.Args {
		log.Println(i, v)
	}
	filePath := os.Args[1]
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
