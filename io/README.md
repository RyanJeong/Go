#   I/O<br>
##  입출력은 io.Reader와 io.Writer 인터페이스, 그리고 파생된 다른 인터페이스들을 이용함<br>
*   바이트들을 읽고 쓸 수 있는 인터페이스<br>

### 파일 읽기<br>
```Go
f, err := os.Open(filename)
if err != nil {

    return err // 혹은 다른 에러 처리
}

// defer: 해당 함수를 벗어날 때 호출할 함수를 등록하는 역할
defer f.Close()
var num int
if _, err := fmt.Fscanf(f, "%d\n", &num); err == nil {
    // 읽은 num 값 사용
}
```

### 파일 쓰기<br>
```Go
f, err := os.Create(filename)
if err != nil {

    return err // 혹은 다른 에러 처리
}

defer f.Close()
if _, err := fmt.Fprintf(f, "%d\n", num); err != nil {
    
    return err // 혹은 다른 에러 처리
}
```

```Go
package io

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

func WriteTo(w io.Writer, lines []string) error {
    for _, line := range lines {
        if _, err := fmt.Fprintln(w, line); err != nil {

            return err
        }
    }

    return nil
}

func ReadFrom(r io.Reader, lines *[]string) error {
    scanner := bufio.NewScanner(r)
    for scanner.Scan() {
        *lines = append(*lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {

        return err
    }

    return nil
}

func ExampleWriteTo() {
    lines := []string{
        "bill@mail.com",
        "tom@mail.com",
        "jane@mail.com",
    }
    if err := WriteTo(os.Stdout, lines); err != nil {
        fmt.Println(err)
    }
    // Output:
    // bill@mail.com
    // tom@mail.com
    // jane@mail.com

    return
}

func ExampleReadFrom() {
    r := strings.NewReader("bill\ntom\njane\n")
    var lines []string
    if err := ReadFrom(r, &lines); err != nil {
        fmt.Println(err)
    }
    fmt.Println(lines)
    // Output:
    // [bill tom jane]

    return
}
```

```Go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.OpenFile(
        "hello.txt",
        os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 파일이 없으면 생성,
                                          // 읽기/쓰기, 파일을 연 뒤 내용 삭제
        os.FileMode(0644),                // 파일 권한은 644
    )
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()             // main 함수가 끝나기 직전에 파일을 닫음

    w := bufio.NewWriter(file)     // io.Writer 인터페이스를 따르는 file로
                                   // io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성
    w.WriteString("Hello, world!") // 쓰기 인스턴스로 버퍼에 Hello, world! 쓰기
    w.Flush()                      // 버퍼의 내용을 파일에 저장

    r := bufio.NewReader(file)     // io.Reader 인터페이스를 따르는 file로
                                   // io.Reader 인터페이스를 따르는 읽기 인스턴스 r 생성
    fi, _ := file.Stat()           // 파일 정보 구하기
    b := make([]byte, fi.Size())   // 파일 크기만큼 바이트 슬라이스 생성

    file.Seek(0, os.SEEK_SET)      // 파일 읽기 위치를 파일의 맨 처음(0)으로 이동
    r.Read(b)                      // 읽기 인스턴스로 파일의 내용을 읽어서 b에 저장

    fmt.Println(string(b)) // 문자열로 변환하여 b의 내용 출력
}
```
