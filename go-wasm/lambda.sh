GOOS=linux GOARCH=amd64 go build -o lambda/main lambda/main.go
zip -j lambda/deployment.zip lambda/main  index.html main.wasm wasm_exec.js
#aws lambda update-function-code --function-name go-wasm --zip-file fileb://lambda/deployment.zip


