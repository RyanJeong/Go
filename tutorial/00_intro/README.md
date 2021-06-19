# Intro.

## vim-go 주요 기능

| Command | 설명 |
|:--------|:--------|
| :GoRun | 코드 실행 |
| :GoBuild | 코드 빌드 |
| :GoErrCheck | 에러 체크 |
| :GoImport fmt | fmt package 추가 |
| :GoImports | package 자동 추가 |
| :GoDef | go 변수 정의 |
| :GoDoc | go 문서 열람 |
| :GoFmt | go 포맷 확인 |

## Go 도구
```text
$ godoc fmt             // fmt package 문서 열람
$ godoc fmt Printf      // fmt package 중에서 Printf 문서 열람
$ godoc cmd/go          // cmd/go 하위 directory 열람
$ godoc -src fmt        // fmt package의 소스 코드 열람
$ godoc -src fmt Printf // fmt package 중에서 Printf 소스 코드 열람
$ godoc -q Reader       // Reader가 들어간 내용 검색
$ godoc -http=:6060     // 웹 서버 구동
http://localhost:6060/pkg/github.com/ryanjeong/gogo/seq/

$ oracle(추가 예정)
$ go tool vet github.com/ryanjeong/gogo/seq // 소스 코드 검사
$ go tool vet *.go                          // 소스 코드 검사
$ go tool fix github.com/ryanjeong/gogo/seq // 옛 API 호출 등을 자동으로 수정, Go 버전 업그레이드될 때 한 번 실행
$ go test github.com/ryanjeong/gogo/seq     // 테스트 수행
```

## Workspace

### Workspace($GOPATH)는 세 개의 하위 디렉토리를 포함하는 계층 구조
### go tool은 소스 패키지를 빌드한 결과들을 pkg 디렉토리와 bin 디렉토리에 저장
### src 내부에 보관되는 디렉토리는 보통 여러 개의 버전 관리 저장소(git 또는 Mercurial과 같은)를 포함
* src
    * Go 소스 파일들을 보관하는 디렉토리

* lib
    * 패키지 객체들을 보관하는 디렉토리

* bin
    * 실행 가능한 바이너리 파일들을 보관하는 디렉토리


## 컴파일

### 작성한 파일에 잘못된 형식을 수정함
```bash
$ gofmt -w file.go
```

### 작성한 파일에 필요한 module을 소스 코드 상단에 추가함
```bash
$ goimports -w file.go
```

### 외부 패키지를 가져오고자 할 때 `go get` 사용 
```bash
# Add dependencies to current module and install them.
# Get downloads the packages named by the import paths,
#   along with their dependencies. It then installs the
#   named packages, like go install. 
$ go get URL        
```

### 외부 패키지를 다운로드만 받고, 설치는 하지 않고자 할 때, `go get -d` 사용
```bash
$ go get -d URL
```

### 패키지 소스코드를 이미 가지고 있거나, `go get -d` 명령을 통해 코드를 가져오기만 한 경우, `go install` 사용
```bash
# Compile and install packages and dependencies.
$ go install URL
```

### Go 프로그램을 컴파일하고자 할 경우, `go build [argument]` 사용
```bash
# compile packages and dependencies
go build file.go
```

### Go 프로그램을 컴파일함과 동시에 실행하고자 할 경우, `go run [argument]` 사용
```bash
# compile and run Go program
$ go run file.go    
```

### Go에서 컴파일을 할 경우, $GOROOT, $GOPATH 순으로 컴파일에 필요한 패키지 검사
> When the go command is looking for packages, it always looks in $GOROOT first.

### 기타 명령들:
```bash
go <command> [argument]

bug         start a bug report
build       compile packages and dependencies
clean       remove object files and cached files
doc         show documentation for package or symbol
env         print Go environment information
fix         update packages to use new APIs
fmt         gofmt (reformat) package sources
generate    generate Go files by processing source
get         add dependencies to current module and install them
install     compile and install packages and dependencies
list        list packages or modules
mod         module maintenance
run         compile and run Go program
test        test packages
tool        run specified go tool
version     print Go version
vet         report likely mistakes in packages
```

## Go 프로그램 기본 구조
```go
package main    // Package 이름, 큰따옴표 사용하지 않음

import "fmt"    // Module 경로, 큰따옴표 사용

func main() {   // Main 함수
    fmt.Println("Hello, world!")    // Go 언어에서는 들여쓰기로 Tab 사용하는 것이 일반적
}
```
### Package
#### In the most basic terms, A package is nothing but a directory inside your Go workspace containing one or more Go source files, or other Go packages.
* Package 용도에 따라 main package 또는 library package로 사용할 수 있음
* go install 명령을 사용하면 해당 directory를 main package로써 사용하겠다는 의미
* 소스코드 내부에 import "directory"를 추가한다면 해당 directory를 library package로써 사용하겠다는 의미
    * library package 내부 함수들의 첫 글자가 대문자인 경우 다른 package에서 사용할 수 있고, 소문자인 경우 다른 package에서 사용할 수 없음
* Package 이름은 소문자로 간결하게 작성하는 것이 원칙
    * util과 같이 너무 일반적인 이름 또는 기본 library와 비슷한 이름은 피해야 함 

### Module
#### A module is a Go package that can be versioned, updated, and managed independently.
* Library package는 일종의 module로써 사용
* 기본으로 제공되는 library들과 다른 library들을 서로 구분하기 위해, 아래와 같이 중간에 빈칸을 삽입
* 이때, gofmt 명령을 사용할 경우 알파벳 순서로 정렬을 수행하는데, 아래와 같이 빈칸이 삽입되어 있으면 빈칸을 기준으로 따로 정렬 수행

```go
import (
    "a"
    "b"
    
    "c"
    "d"
    "e"
)
```
### 세미콜론`;`을 각 행의 끝에 붙이지 않는 것처럼 보이지만, 구문 분석기가 소스 코드를 분석하는 과정에서 특정 규칙을 기준으로 세미콜론을 붙임
* ex) 함수의 여는 중괄호`{`는 `func` 선언과 같은 줄에 있어야 함
