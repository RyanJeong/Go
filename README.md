# Go
## 환경설정<br>
* Ubuntu<br>
1. .bashrc 설정
```bash
$ vi ~/.bashrc

# 아래 세 줄 bashrc 하단에 추가
export GOPATH=$HOME/Go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

2. workspace 설정
```bash
$ mkdir -p $GOPATH/src/github.com/username/go
$ cd $GOPATH/src/github.com/username/go
$ git init
```

3. Go 코드 작성에 도움이 되는 툴 설정<br>

> vim 플러그인들을 관리하는 패키지 매니저인 pathogen 설치<br>
```bash
$ mkdir -p ~/.vim/autoload
$ cd ~/.vim/autoload
$ curl -LSso pathogen.vim https://tpo.pe/pathogen.vim

$ vi ~/.vimrc

// 아래 세 줄 vimrc 하단에 추가
execute pathogen#infect()
syntax on
filetype plugin indent on
```
> vim-go 설치<br>
```bash
$ mkdir -p  ~/.vim/bundle
$ cd ~/.vim/bundle
$ git clone https://github.com/fatih/vim-go.git
```
> VCS(Version Control System)인 mercurial 설치<br>
```bash
// 자세한 내용은 https://www.mercurial-scm.org/wiki/QuickStart 방문하여 참조할 것
$ apt-get install mercurial
```
> vim 실행 후 GoInstallBinaries 명령 실행<br>
```bash
$ vim
> :GoInstallBinaries
```
> 자동완성 기능 추가(YouCompleteMe)<br>
* Ubuntu 환경에서 오류가 있어 동작하지 않음을 확인하였음. 업데이트를 통해 해당 내용을 수정할 예정
```bash
$ apt-get update
$ apt-get install -y build-essential cmake
$ apt-get install -y python-dev python3-dev
$ git clone https://github.com/Valloric/YouCompleteMe ~/.vim/bundle/YouCompleteMe
$ cd ~/.vim/bundle/YouCompleteMe
$ git submodule update --init --recursive
$ ~/.vim/bundle/YouCompleteMe/install.py

# 자동완성 단축키는 ctrl + 'y'
```
> TagBar 설치<br>
```bash
$ apt-get install ctags
$ cd ~/.vim/bundle
$ git clone https://github.com/majutsushi/tagbar.git
```
> 파일 네비게이션 설정<br>
```bash
$ cd ~/.vim/bundle
$ git clone https://github.com/scrooloose/nerdtree.git 
```
> 단축키 지정<br>
```bash
$ vi ~/.vimrc
set ts=4
 
map <F8> :NERDTreeToggle<CR>
map <F2> :GoDef<CR>
map <F4> :TagbarToggle<CR>
```

> NERDTreeToggle 기능 사용 중에, 아래와 같은 오류가 나타날 경우 대처 방안<br>
```Text
/root/.vim/bundle/nerdtree/syntax/nerdtree.vim 수행중 에러 발견:
   3 줄:
E121: 정의 안 된 변수: g:NERDTreeGlyphReadOnly
E15: 잘못된 표현식: 'syn match NERDTreeIgnore #\['.g:NERDTreeGlyphReadOnly.'\]#'
  25 줄:
E121: 정의 안 된 변수: g:NERDTreeDirArrowCollapsible
E116: 함수 escape(g:NERDTreeDirArrowCollapsible, '~') . '\ze .*/# containedin=NERDTreeDir,NERDTreeFile'(으)로 잘못된 인자>
가 넘겨졌습니다
E15: 잘못된 표현식: 'syn match NERDTreeClosable #' . escape(g:NERDTreeDirArrowCollapsible, '~') . '\ze .*/# containedin=NE
RDTreeDir,NERDTreeFile'
  26 줄:
E121: 정의 안 된 변수: g:NERDTreeDirArrowExpandable
E116: 함수 escape(g:NERDTreeDirArrowExpandable, '~') . '\ze .*/# containedin=NERDTreeDir,NERDTreeFile'(으)로 잘못된 인자가
 넘겨졌습니다
E15: 잘못된 표현식: 'syn match NERDTreeOpenable #' . escape(g:NERDTreeDirArrowExpandable, '~') . '\ze .*/# containedin=NER
DTreeDir,NERDTreeFile'
  28 줄:
E121: 정의 안 된 변수: g:NERDTreeDirArrowCollapsible
E116: 함수 escape(g:NERDTreeDirArrowCollapsible, '~]\-').escape(g:NERDTreeDirArrowExpandable, '~]\-')(으)로 잘못된 인자가 
넘겨졌습니다
E15: 잘못된 표현식: escape(g:NERDTreeDirArrowCollapsible, '~]\-').escape(g:NERDTreeDirArrowExpandable, '~]\-')
  29 줄:
E121: 정의 안 된 변수: s:dirArrows
E15: 잘못된 표현식: 'syn match NERDTreeDir #[^'.s:dirArrows.' ].*/#'
  31 줄:
E121: 정의 안 된 변수: s:dirArrows
E15: 잘못된 표현식: 'syn match NERDTreeFile  #^[^"\.'.s:dirArrows.'] *[^'.s:dirArrows.']*# contains=NERDTreeLink,NERDTreeR
O,NERDTreeBookmark,NERDTreeExecFile'
  34 줄:
E121: 정의 안 된 변수: g:NERDTreeGlyphReadOnly
E15: 잘못된 표현식: 'syn match NERDTreeRO # *\zs.*\ze \['.g:NERDTreeGlyphReadOnly.'\]# contains=NERDTreeIgnore,NERDTreeBoo
kmark,NERDTreeFile'
  41 줄:
E121: 정의 안 된 변수: g:NERDTreeNodeDelimiter
E116: 함수 char2nr(g:NERDTreeNodeDelimiter) . '# conceal containedin=ALL'(으)로 잘못된 인자가 넘겨졌습니다
E15: 잘못된 표현식: 'syn match NERDTreeNodeDelimiters #\%d' . char2nr(g:NERDTreeNodeDelimiter) . '# conceal containedin=AL
L'
계속하려면 엔터 혹은 명령을 입력하십시오
```
```Bash
$ vi ~/.vimrc

# 하단에 추가
let g:NERDTreeDirArrows = 1
let g:NERDTreeDirArrowExpandable = '▸'
let g:NERDTreeDirArrowCollapsible = '▾'
let g:NERDTreeGlyphReadOnly = "RO"
let g:NERDTreeWinSize = 40
let NERDTreeNodeDelimiter = "\t"
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

## Workspace<br>
*	Workspace($GOPATH)는 세 개의 하위 디렉토리를 포함하는 계층 구조를 나타냄<br>
	*	src<br>
		*	Go 소스 파일들을 보관하는 디렉토리<br>
	*	lib<br>
		*	패키지 객체들을 보관하는 디렉토리<br>
	*	bin<br>
		*	실행 가능한 바이너리 파일들을 보관하는 디렉토리<br>

*	go tool은 소스 패키지를 빌드한 결과들을 pkg 디렉토리와 bin 디렉토리에 저장함<br>
*	src 내부에 보관되는 디렉토리는 보통 여러 개의 버전 관리 저장소(git 또는 Mercurial과 같은)들을 포함함<br>
*	Workspace 내부에 어떠한 형태로 디렉토리가 관리되는지 보여주는 예:
```Text
$ cd $GOPATH
~/Go

$ tree bin
bin
├── authtest
├── benchcmp
...
├── stringer
└── toolstash

$ tree pkg
pkg
└── linux_amd64
    └── golang.org
        └── x
            └── tools
                └── cmd
                    ├── guru
                    │   └── serial.a
                    └── splitdwarf
                        └── internal
                            └── macho.a

$ tree src
src
├── github.com
│   └── ryanjeong
│       └── go
├── golang.org
│   └── x
│       ├── lint
│       │   ├── ...
│       │   └── ...
│       ├── mod
│       │   ├── ...
│       │   └── ...
│       ├── net
│       │   ├── ...
│       │   └── ...
│       ├── tools
│       │   ├── ...
│       │   └── ...
│       └── xerrors
│           ├── ...
│           └── ...
└── honnef.co
    └── go
        └── tools
				├── LICENSE
        ├── LICENSE-THIRD-PARTY
        ├── README.md
        ├── _benchmarks
        │   ├── bench.sh
        │   └── silent-staticcheck.sh
        ├── analysis
        │   ├── code
        │   │   ├── code.go
        │   │   ├── loops.go
        │   │   ├── stub.go
        │   │   ├── terminates.go
        │   │   └── visit.go
        │   ├── edit
        │   │   └── edit.go
        │   ├── facts
        │   │   ├── deprecated.go
        │   │   ├── directives.go
        │   │   ├── facts_test.go
        │   │   ├── generated.go
        │   │   ├── purity.go
        │   │   ├── testdata
        │   │   │   └── src
        │   │   │       ├── Deprecated
        │   │   │       │   └── Deprecated.go
        │   │   │       └── Purity
        │   │   │           └── CheckPureFunctions.go
        │   │   └── token.go
        │   ├── lint
        │   │   └── lint.go
        │   └── report
        │       └── report.go
        ...
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
