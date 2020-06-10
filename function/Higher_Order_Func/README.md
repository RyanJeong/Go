* 고계함수(Higher-order functoin)는 함수를 넘기고 받는 함수를 나타냄<br>
```Go
package Higher_Order_Func

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

func ExampleReadFrom_Print() {
	r := strings.NewReader("bill\ntom\njane\n")
	err := ReadFrom(r, func(line string) {
		fmt.Println("(", line, ")")
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// ( bill )
	// ( tom )
	// ( jane )

	return
}
```
