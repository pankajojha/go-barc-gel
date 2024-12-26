package main

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// readStaticFile reads a file and returns its content and MIME type.
func readStaticFile(path string) ([]byte, string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, "", err
	}

	// Determine content type
	contentType := http.DetectContentType(data)
	return data, contentType, nil
}

// handler serves static files (main.wasm, wasm_exec.js, and index.html).
func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Determine file to serve
	var filePath string
	switch req.Path {
	case "/":
		filePath = "index.html"
	case "/main.wasm":
		filePath = "main.wasm"
	case "/wasm_exec.js":
		filePath = "wasm_exec.js"
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "404 - Not Found",
		}, nil
	}

	// Read the file
	data, contentType, err := readStaticFile(filePath)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error reading file",
		}, nil
	}

	// Encode the file as base64 for binary content
	isBinary := contentType == "application/wasm"
	body := string(data)
	if isBinary {
		body = base64.StdEncoding.EncodeToString(data)
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      http.StatusOK,
		Headers:         map[string]string{"Content-Type": contentType},
		Body:            body,
		IsBase64Encoded: isBinary,
	}, nil
}

func main() {
	// Lambda entry point
	lambda.Start(handler)
}
