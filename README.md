# file-server

file-server는 간단한 http 파일 서버 생성 프로그램입니다.

아래는 모든 옵션을 포함한 file-server의 실행 예입니다.
(root 권한으로 실행해야 합니다.)

```
file-server -addr :80 -dir /some/dir
```

이 명령은 localhost에 /some/dir을 바인딩한 파일 서버를 만듭니다.

## 설치

Go를 설치하신 후 다음을 실행하세요.

```
go get github.com/studio2l/file-server
```

명령이 잘 실행되었다면 $GOPATH/bin 아래에 file-server 프로그램이 설치됩니다.
