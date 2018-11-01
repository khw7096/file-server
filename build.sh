# OS별로 빌드함.
APP="file-server"
GOOS=linux GOARCH=amd64 go build -o ./bin/linux/$APP main.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows/$APP.exe main.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/$APP main.go

# Github Release에 업로드 하기위해 압축
cd ./bin/linux/ && tar -zcvf ../$APP.linux.tgz . && cd -
cd ./bin/windows/ && tar -zcvf ../$APP.windows.tgz . && cd -
cd ./bin/darwin/ && tar -zcvf ../$APP.darwin.tgz . && cd -
