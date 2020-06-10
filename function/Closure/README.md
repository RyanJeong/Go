* 클로저(Closure)는 외부에서 선언한 변수를 함수 리터럴 내에서 마음대로 접근할 수 있는 코드<br>
```Go
package Closure

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {

		return err
	}

	return nil
}

func ExampleReadFrom_append() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	err := ReadFrom(r, func(line string) {
		lines = append(lines, line)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [bill tom jane]

	return
}
```
