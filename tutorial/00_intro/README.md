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
* 모든 Go 프로그램은 패키지로 생성됨
* `main` 패키지는 실행 프로그램을 가리킴
* 패키지명이 `main`인 Go 파일 내 존재하는 `main()`이 프로그램의 시작점
* 컴파일러는 `main` 패키지를 실행 프로그램으로 변환
* Package 용도에 따라 main package 또는 library package로 사용할 수 있음
* `go install` 명령을 사용하면 해당 directory를 main package로써 사용하겠다는 의미
* 소스코드 내부에 `import` "directory"를 추가한다면 해당 directory를 library package로써 사용하겠다는 의미
    * library package 내부 함수들의 첫 글자가 대문자인 경우 다른 package에서 사용할 수 있고, 소문자인 경우 다른 package에서 사용할 수 없음
* Package 이름은 소문자로 간결하게 작성하는 것이 원칙
    * util과 같이 너무 일반적인 이름 또는 기본 library와 비슷한 이름은 피해야 함 

### Module
#### A module is a Go package that can be versioned, updated, and managed independently.
* Library package는 일종의 module로써 사용
* 기본으로 제공되는 library(library package)들과 다른 library들을 서로 구분하기 위해, 아래와 같이 중간에 빈칸을 삽입
* 다른 프로그램에서 library를 사용하고자 할 때, Go 컴파일러는 GOROOT/pkg에 해당 library가 존재하는지 확인
* GOROOT/pkg 내에는 Go 언어에서 기본으로 제공하는 library들이 들어있으며, 만약 해당 경로에 사용하고자 하는 library가 없을 경우 GOPATH/pkg 탐색
* GOPATH/pkg 내에는 3rd-party library들이 들어있음
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

* Library package에서 이름이 대문자로 시작하는 경우, 다른 패키지에서 접근해 사용할 수 있음
    * Package scope 규칙에 의해 동작
    * 이름(identifier)이 대문자로 시작한다면, 이는 `public` 객체임
    * 이름(identifier)이 소문자로 시작한다면, 이는 `non-public` 객체임
     
```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10)) // math/rand 패키지의 Init() 함수 호출
}
```

#### `math` 패키지 안에 정의된 Pi 변수를 사용하고자 할 때, 아래와 같이 사용하면 오류 발생

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.pi) // error, rename math.pi to math.Pi
}
```

* library는 `import`됨과 동시에 초기화할 수 있음

```go
package testlib
 
var pop map[string]string
 
func init() {   // Init. when the package importing
    pop = make(map[string]string)
}
```

* `import ALIAS library` 형태로 추가될 library에 별명을 사용할 수 있음
* library 자체를 main package에 가져와 사용할 경우 library 안에 있는 이름이 서로 충돌할 수 있음
* 이름을 서로 구분하기 위해 별명을 사용해서 해당 library를 세분할 수 있음

```go
import (
    mongo "other/mongo/db"
    mysql "other/mysql/db"
)
func main() {
    mondb := mongo.Get()
    mydb := mysql.Get()
    //...
}
```

* 만약 library의 초기화 함수만 사용하는 것을 원하고, 그 이후에는 해당 library가 필요 없을 경우 아래와 같이 사용할 수 있음

```go
package main
import _ "testlib"
```

* 패키지 사용 예
##### main.go

```go
package main

import "testlib"

func main() {
    song := testlib.GetMusic("Metalica")
    println(song)
}
```

##### testlib/testlib.go
```go
// 1. mv testlib $GOPATH/src
// 2. go build testlib
// 3. go install testlib
package testlib

import "fmt"

var rock map[string]string

func init() {
    rock = make(map[string]string)
    rock["Muse"] = "Time Is Running Out"
    rock["Metalica"] = "Master of Puppets"
    rock["Green Day"] = "Basket Case"
}

// public
func GetMusic(artist string) string {
    return rock[artist]
}

// non-public
func getKeys() {
    for _, v := range rock {
        fmt.Println(v)
    }
}
```

### 세미콜론`;`을 각 행의 끝에 붙이지 않는 것처럼 보이지만, 구문 분석기가 소스 코드를 분석하는 과정에서 특정 규칙을 기준으로 세미콜론을 붙임
* ex) 함수의 여는 중괄호`{`는 `func` 선언과 같은 줄에 있어야 함
