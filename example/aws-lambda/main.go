package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/logs"
)

type RequestBody struct {
	PublicSrcUrl string `json:"src_image_name"`
	DestUrl      string `json:"dest_image_name"`
	DestUsername string `json:"dest_username"`
	DestPassword string `json:"dest_password"`
}

type simpleKeychain struct {
	Url      string
	Username string
	Password string
}

func (kc simpleKeychain) Resolve(resource authn.Resource) (authn.Authenticator, error) {
	reg := resource.RegistryStr()
	if strings.Contains(kc.Url, reg) {
		log.Println("simple resolve reg", reg, kc.Username, kc.Password)
		return &authn.Basic{
			Username: kc.Username,
			Password: kc.Password,
		}, nil
	}
	log.Println("simple resolve Anonymous", reg)
	return authn.Anonymous, nil
}

func init() {
	logs.Progress.SetOutput(os.Stdout)
}

// copy from public registry with Anonymous auth
// dest registry supports username and password auth
func handler(ctx context.Context, event json.RawMessage) (events.LambdaFunctionURLResponse, error) {
	fmt.Printf("Request msg: %s\n", event)

	var requestBody RequestBody
	if err := json.Unmarshal(event, &requestBody); err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: 400,
			Body:       "Invalid event json",
		}, nil
	}

	log.Printf("COPY EVENT: Copy %s to %s", requestBody.PublicSrcUrl, requestBody.DestUrl)

	if err := crane.Copy(requestBody.PublicSrcUrl, requestBody.DestUrl, crane.WithAuthFromKeychain(
		authn.NewMultiKeychain(
			simpleKeychain{
				Url:      requestBody.DestUrl,
				Username: requestBody.DestUsername,
				Password: requestBody.DestPassword,
			},
		),
	)); err != nil {
		// cancel()
		return events.LambdaFunctionURLResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("COPY EVENT: Copy %s to %s", requestBody.PublicSrcUrl, requestBody.DestUrl),
	}, nil
}

func main() {
	lambda.Start(handler)
}
