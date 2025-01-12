# Copy container image with google/go-containerregistry

An example of copying docker image, with simple username password, between two registries with `crane`.

## TL,DR
[Use multiple keychain](autnmulti) and implement `KeyChain` interface
```go
func (obj MyOjbect) Resolve(target authn.Resource) (authn.Authenticator, error) {
	reg := target.RegistryStr()
	if strings.Contains(src url, reg) {
		return &authn.Basic{
			Username: src username,
			Password: src password,
		}, nil
	} else {
		return &authn.Basic{
			Username: dest username,
			Password: dest password,
		}, nil
	}
}

func main() {
    	_ = crane.Copy(src url, dest url, crane.WithAuthFromKeychain(obj))
}

```
# Credits
* `authn` readme: [google/go-containerregistry/pkg/authn](authnmulti)
* Exmaple of copying to AWS ECR with helper: [ekirmayer/aws-lambda-copy-container-image-to-ecr](ekirmayer)

[authnmulti]:https://github.com/google/go-containerregistry/tree/main/pkg/authn#using-multiple-keychains
[ekirmayer]:https://github.com/ekirmayer/aws-lambda-copy-container-image-to-ecr