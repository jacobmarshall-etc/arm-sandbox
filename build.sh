GOOS=linux GOARCH=arm go build -o ./bin/sandbox ./src/cli/sandbox
docker build -t arm-sandbox .