package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/harrypunk/copy-image"
)

type CopyImageRequest struct {
	SrcImageName  string `json:"src_image_name"`
	SrcUsername   string `json:"src_username"`
	SrcPassword   string `json:"src_password"`
	DestImageName string `json:"dest_image_name"`
	DestUsername  string `json:"dest_username"`
	DestPassword  string `json:"dest_password"`
}

func handleRequest(ctx context.Context, request CopyImageRequest) (string, error) {
	srcInfo := copy.ImageInfo{
		ImageName: request.SrcImageName,
		Username:  request.SrcUsername,
		Password:  request.SrcPassword,
	}

	destInfo := copy.ImageInfo{
		ImageName: request.DestImageName,
		Username:  request.DestUsername,
		Password:  request.DestPassword,
	}

	err := copy.CopyImage(srcInfo, destInfo)
	if err != nil {
		log.Printf("Error copying image: %v", err)
		return "Error", err
	}

	return "Image copied successfully", nil
}

func main() {
	lambda.Start(handleRequest)
}
