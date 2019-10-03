# Go
## 환경설정<br>
* Ubuntu<br>
```bash
$ vi ~/.bashrc

// 아래 세 줄 bashrc 하단에 추가
export GOPATH=$HOME/GO
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

$ cd $GOPATH/src/github.com/username/go
$ git init

// vim-go 설정
// 1. pathogen 설치
$ mkdir -p ~/.vim/autoload ~
/.vim/bundle
$ cd ~/.vim/autoload
$ -curl -LSso pathogen.vim https://tpo.pe/pathogen.vim

// 2. .vimrc 파일 수정
$ vi ~/.vimrc
execute pathogen#infect()
syntax on
filetype plugin indent on

// 3. vim-go 설치
$ cd ~/.vim/bundle
$ git clone https://github.com/fatih/vim-go.git

// 4. mercurial 설치
$ apt-get install mercurial

// 5. vim 실행 후 :GoInstallBinaries 수행
$ vim
:GoInstallBinaries

// 6. 자동완성 기능 추가
$ apt-get install cmake
$ apt-get install python-dev
$ cd ~/.vim/bundle
$ git clone https://github.com/Valloric/YouCompleteMe.git
$ cd YouCompleteMe
$ ./install.sh

// 7. TagBar 설치
$ apt-get install ctags
$ cd ~/.vim/bundle
$ git clone https://github.com/majutsushi/tagbar.git

// 8. 파일 네비게이션 설정
$ cd ~/.vim/bundle
$ git clone https://github.com/scrooloose/nerdtree.git 

// 9. 단축키 지정
$ vi ~/.vimrc
set ts=4
 
map <F8> :NERDTreeToggle<CR>
map <F2> :GoDef<CR>
map <F4> :TagbarToggle<CR>
```
* vim-go 주요 기능<br>

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

* Windows(추가예정)<br>

* Go 도구<br>
```text
$ godoc fmt		// fmt package 문서 열람
$ godoc fmt Printf	// fmt package 중에서 Printf 문서 열람
$ godoc cmd/go		// cmd/go 하위 directory 열람
$ godoc -src fmt	// fmt package의 소스 코드 열람
$ godoc -src fmt Printf	// fmt package 중에서 Printf 소스 코드 열람
$ godoc -q Reader	// Reader가 들어간 내용 검색
$ godoc -http=:6060	// 웹 서버 구동
http://localhost:6060/pkg/github.com/ryanjeong/gogo/seq/

$ oracle(추가 예정)
$ go tool vet github.com/ryanjeong/gogo/seq	// 소스 코드 검사
$ go tool vet *.go				// 소스 코드 검사
$ go tool fix github.com/ryanjeong/gogo/seq	// 옛 API 호출 등을 자동으로 수정, Go 버전 업그레이드될 때 한 번 실행
$ go test github.com/ryanjeong/gogo/seq		// 테스트 수행
```
## 컴파일<br>
```bash
$ gofmt -w file.go
```
* 작성한 파일로부터 잘못된 형식을 고쳐 다시 저장<br>
```bash
$ goimports -w file.go
```
* 작성한 파일에 필요한 Module을 소스 코드 상단에 추가함<br>
```bash
$ go get URL		// add dependencies to current module and install them
$ go install URL	// compile and install packages and dependencies
$ go run file.go	// compile and run Go program
```

## 기본 구조<br>
```go
package main	// Package 이름, 큰따옴표 사용하지 않음

import "fmt"	// Module 경로, 큰따옴표 사용

func main() {	// Main 함수
	fmt.Println("Hello, world!")	// Go 언어에서는 들여쓰기로 Tab 사용하는 것이 일반적
}
```
* Package: Go Package In the most basic terms, A package is nothing but a directory inside your Go workspace containing one or more Go source files, or other Go packages. <br>
	* Package 용도에 따라 main package 또는 library package로 사용할 수 있음<br>
	* go install 명령을 사용하면 해당 directory를 main package로써 사용하겠다는 의미<br>
	* 소스코드 내부에 import "directory"를 추가한다면 해당 directory를 library package로써 사용하겠다는 의미<br>
		* library package 내부 함수들의 첫 글자가 대문자인 경우 다른 package에서 사용할 수 있고, 소문자인 경우 다른 package에서 사용할 수 없음<br>	
	* Package 이름은 소문자로 간결하게 작성하는 것이 원칙<br>
		* util과 같이 너무 일반적인 이름이나 기본 library와 비슷한 이름은 피해서 작성<br>
* Module: A module is a Go package that can be versioned, updated, and managed independently.<br>
	* Library package는 일종의 module로써 사용
	* 기본으로 제공되는 library들과 다른 library들을 서로 구분하기 위해, 아래와 같이 중간에 빈칸을 삽입<br>
	* 이때, gofmt 명령을 사용할 경우 알파벳 순서로 정렬을 수행하는데, 아래와 같이 빈칸이 삽입되어 있으면 빈칸을 기준으로 따로 정렬 수행<br>
```go
import (
	"a"
	"b"
	
	"c"
	"d"
	"e"
)
```
* 세미콜론(;)을 각 행의 끝에 붙이지 않는 것처럼 보이지만, 구문 분석기가 소스 코드를 분석하는 과정에서 특정 규칙을 기준으로 세미콜론을 붙임<br>
	* ex) 함수의 여는 중괄호({)는 func 선언과 같은 줄에 있어야 함<br>

## 자료형 및 변수<br>
* Go 언어는 자료형을 정적으로 검사하므로, 변수에 자료형이 정해짐<br>
* Go 언어는 자료형 추론 기능이 있음<br>
	* **물론, 정적 자료형을 지원하기 때문에 해당 변수에 다른 자료형의 값을 담을 순 없음**<br>
```go
var x int	// 변수 x는 integer 자료형
var arr [5]int	// 변수 arr는 크기가 5인 integer 배열 자료형
func(int, int) {
	// 함수는 두 개의 integer 자료형 입력을 받음
}
func(int) int {
	// 함수는 한 개의 integer 자료형 입력을 받아 integer 자료형의 값을 반환함
}
func(int, func(int, int)) func(int) int {
	// 함수는 ...
}
var p *int	// 변수 p는 포인터이며 integer 자료형을 가리킴
var x int = 10	// 변수 x는 integer 자료형이며, 값은 10으로 초기화
var i = 10	// 변수 i는 integer 자료형(형추론)이며, 값은 10으로 초기화
var p = &i	// 변수 p는 포인터(형추론)이며 integer 자료형을 가리킴, 값은 변수 i의 주소로 초기화

i := 10		// var i = 10과 동일한 표현식
p := &i		// var p = &i와 동일한 표현식
```
```go
i = 10		// 불가능, i는 선언되지 않은 변수
var i = 10	// 가능, i는 새로운 integer 자료형 변수
j := 10		// 가능, j는 새로운 integer 자료형 변수
j = 20		// 가능, j는 이미 선언된 integer 자료형 변수
j = "test"	// 불가능, Go 언어는 정적 자료형 사용
```

## 제어 구조<br>
```go
// if
func facRcr(n int) int {
	if n <= 0 {	// 괄호 사용하지 않음

		return 1
	}

	return n * (fac(n - 1))
}

// for while-like
func facItr1(n int) int {
	result := 1
	for n > 0 {	// while처럼 동작하는 for, 괄호 사용하지 않음
		result *= n
		n--	// Go에는 전위연산자가 없으므로, 무조건 후위연산자를 사용해야 함
	}

	return result
}

// for with condition statement
func facItr2(n int) int {
	result := 1
	for i := 2; i <= n; i++ {	// 괄호 사용하지 않음
		result *= i
	}

	return result
}
```
## 참고자료
[https://godoc.org](https://godoc.org)
