# file-server

file-server는 간단한 http 파일 서버 프로그램입니다.

아래는 모든 옵션을 포함한 file-server의 실행 예입니다.
기본적으로 :80 포트로 설정되어 있으며, 서비스를 실행하기 위해서는 root 권한으로 실행해야 합니다.

```
# file-server -addr :80 -dir /some/dir
```

이 명령은 localhost에 /some/dir 경로를 바인딩한 파일 서버를 만듭니다.

## 다운로드
[Download](https://github.com/studio2l/file-server/releases/tag/v0.1)

## Go를 활용한 설치

Go 언어 설치 후 터미널에서 다음처럼 타이핑해주세요.

```bash
$ go get github.com/studio2l/file-server
```