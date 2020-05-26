#   String<br>
##  string type은 읽기 전용이며, []byte는 읽기 쓰기 둘 다 가능<br>

* Go 언어 소스 코드는 UTF-8 기반임<br>
* 문자열 수정(예를 들어, +=, strconv, Sprint 등)이 발생하면, 대상체를 기반으로 새로운 문자열이 생성되며, 기존 문자열을 가리키던 변수는 새로 생성된 문자열을 가리킴<br>

```Text
# May 26 2020, Working directory:
./
├── README.md
├── binary
│   ├── README.md
│   ├── binary_test.go
│   └── go.sh
├── concat
│   ├── README.md
│   ├── concat_test.go
│   └── go.sh
├── convert
│   ├── README.md
│   ├── convert_test.go
│   └── go.sh
├── hangul
│   ├── REAMDE.md
│   ├── go.sh
│   ├── hangul.go
│   └── hangul_test.go
└── unicode
    ├── README.md
    ├── go.sh
    └── unicode.go
```

