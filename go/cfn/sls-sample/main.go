package main

import (
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Envs map[string]string `json:"environments"`
}

func Handler() (Response, error) {
	res := &Response{map[string]string{}}
	for _, pair := range os.Environ() {
		facts := strings.SplitN(pair, "=", 2)
		name, val := facts[0], facts[1]
		if strings.HasPrefix(name, "APP_") {
			res.Envs[name] = val
		}
	}
	return *res, nil
}

func main() {
	lambda.Start(Handler)
}
