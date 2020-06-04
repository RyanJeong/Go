#   Formatted I/O Functions in Golang<br>
##  Templates<br>
### Verbs in Golang<br>
*   **%v** is a generic placeholder. It will automatically convert the variable into the string with some default options. This is generally useful when printing the primitives such as the strings or numbers, and you don’t need any other specific options.<br>
*   **%#v** prints your variable in Golang syntax. This means you could copy the output and paste it into your code, and it will syntactically correct. It is useful when working with the data structure like structs and slices because it will print out the types and field names.
*   **%T** prints your variable’s type. This placeholder is useful for debugging if your data is passed as an interface{}, and you want to see what it’s concrete type.
*   **%d** prints the integer in base-10. You can do the same with the %v, but this way is more explicit.
*   **%x** and **%X** print the integer in base-16. One helpful trick, though, is that you can pass in a byte slice, and it’ll print each byte as a two-digit hex number.
*   **%o** and **%O** print the integer in base-8.
*   **%b** print the integer in binary.
*   **%f** prints the floating-point number without an exponent. You can do the same with the %v, but this becomes more useful when we add the width and precision flags.
*   **%q** prints the quoted string. It is useful when your data may have invisible characters (such as zero width space) because the quoted-string will print them as escape sequences.
*   **%p** prints the pointer address of your variable. This one is beneficial when you are debugging code, and you want to check if the different pointer variables reference the same data.

### Width and Precision<br>
```Text
%8.0f ➡ "     101"
%8.1f ➡ "   100.6"
%8.2f ➡ "  100.57"
%8.3f ➡ " 100.567"
```

### Left Alignment<br>
```Text
%-8.0f ➡ "101     "
%-8.1f ➡ "100.6   "
%-8.2f ➡ "100.57  "
%-8.3f ➡ "100.567 "
```

### Zero Padding<br>
```Text
%08d ➡ "00000123"
```

### Spacing<br>
```Text
// For example, formatting []byte{1,2,3,4} with and without the space flag:
%x  ➡ "01020304"
% x ➡ "01 02 03 04"
```

### Other Verbs and Flags:<br>
* the [Printing](https://golang.org/pkg/fmt/#hdr-Printing) section of the [fmt](https://golang.org/pkg/fmt/) godoc.

##  Printing<br>
### Printing to stdout(Standard Output)<br>
```Text
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```
* The **Print()** function simply prints a list of variables to STDOUT with default formatting.
* The **Printf()** function allows you to specify the formatting using a format template.<br>
* The **Println()** function works like Print() except **it inserts spaces between the variables and appends a new line at the end.**<br>
```Go
package main

import "fmt"

func ExampleFunctions_print() {
	list1 := [...]int{1, 2, 3}
	list2 := [...]int{4, 5, 6}
	str1 := "STR1"
	str2 := "STR2"

	fmt.Println(list1, list2)
	fmt.Println(str1, str2)
	fmt.Print(list1, list2)
	fmt.Print(str1, str2)
	fmt.Printf("%d%d", list1, list2)
	fmt.Printf("%s%s", str1, str2)
	fmt.Printf("%d %d", list1, list2)
	fmt.Printf("%s %s", str1, str2)

	// Output:
	// [1 2 3] [4 5 6]
	// STR1 STR2
	// [1 2 3] [4 5 6]STR1STR2[1 2 3][4 5 6]STR1STR2[1 2 3] [4 5 6]STR1 STR2

	return
}
```
> fmt.Print() 함수와 fmt.Println() 함수의 차이는 출력 후 개행을 하는가와 문자열 가변인자를 인자값으로 받았을 때, 서로 띄어쓰기를 하는가에 있다. fmt.Printf() 함수는 사용자의 서식에 맞게 데이터를 출력할 수 있다. 이때 fmt.Printf()는 서식 분석에 따른 cost가 발생하므로, 간단한 출력에서는 fmt.Print() 또는 fmt.Println()을 사용하고, 엄격하게 출력 서식을 제어해야 할 경우에는 fmt.Printf()를 사용한다.<br>

### Printing to io.Writer<br>
* If you need to print to a non-STDOUT output (such as STDERR or a buffer) then you can use the Fprint functions. The **“F”** in these functions comes FILE which was the argument type used in C’s fprintf() function.<br>
```Text
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
```
* These functions are just like the Print functions except you specify the writer as the first argument. In fact, **the Print functions are just small wrappers around the Fprint functions.**<br>

### Formatting to a string<br>
* Sometimes you need to work with strings instead of writers. You could use the the Fprint functions to write to a buffer and convert it to a string but that’s a lot of work. Fortunately there are the Sprint convenience functions:<br>
```Text
func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
```
* The **“S”** here stands for “String”. These functions take the same arguments as the Print() functions except they return a string.<br>
* **While these functions are convenient, they can be a bottleneck if you’re frequently generating strings.** If you profile your application and find that you need to optimize it then reusing a bytes. **Buffer with the Fprint() functions can be much faster.**<br>

### Error formatting<br>
```Text
func Errorf(format string, a ...interface{}) error
```

* This is literally just a wrapper for errors.New() and Sprintf():<br>
```Go
func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}
```

### Print functions benchmark<br>
```Go
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	w := bufio.NewWriter(os.Stdout)

	s1 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Printf("P")
	}
	t1 := time.Since(s1)

	s2 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Fprintf(w, "F")
	}
	t2 := time.Since(s2)

	s3 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Sprintf("S")
	}
	t3 := time.Since(s3)

	s4 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Println("P")
	}
	t4 := time.Since(s4)

	s5 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Fprintln(w, "F")
	}
	t5 := time.Since(s5)

	s6 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Sprintln("S")
	}
	t6 := time.Since(s6)

	s7 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Print("P")
	}
	t7 := time.Since(s7)

	s8 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Fprint(w, "F")
	}
	t8 := time.Since(s8)

	s9 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Sprint("S")
	}
	t9 := time.Since(s9)

	fmt.Printf("fmt.Printf repeat 100,000 times: %v\n", t1)
	fmt.Printf("fmt.Fprintf repeat 100,000 times: %v\n", t2)
	fmt.Printf("fmt.Sprintf repeat 100,000 times: %v\n", t3)
	fmt.Printf("fmt.Println repeat 100,000 times: %v\n", t4)
	fmt.Printf("fmt.Fprintln repeat 100,000 times: %v\n", t5)
	fmt.Printf("fmt.Sprintln repeat 100,000 times: %v\n", t6)
	fmt.Printf("fmt.Print repeat 100,000 times: %v\n", t7)
	fmt.Printf("fmt.Fprint repeat 100,000 times: %v\n", t8)
	fmt.Printf("fmt.Sprint repeat 100,000 times: %v\n", t9)

	return
}
```

| Function     | Repeat  | Speed        |
| :----------- | :-----: | -----------: |
| fmt.Fprint   | 100,000 | 4.785735ms   |
| fmt.Fprintf  | 100,000 | 3.447175ms   |
| fmt.Fprintln | 100,000 | 50.9891ms    |
| fmt.Print    | 100,000 | 137.10993ms  |
| fmt.Printf   | 100,000 | 133.877017ms |
| fmt.Println  | 100,000 | 148.975918ms |
| fmt.Sprint   | 100,000 | 4.963986ms   |
| fmt.Sprintf  | 100,000 | 3.814547ms   |
| fmt.Sprintln | 100,000 | 5.130816ms   |

> 1. Print 계열의 함수는 Fprintf 계열의 함수로부터 첫 번째 인자가 STDOUT으로 설정된 상태로 래핑(wrapping)된 함수이다.<br>
2. Sprint 계열의 함수는 Print 계열의 함수에서 입력된 인자들을 하나의 문자열로 묶어 이를 반환하는 함수이다.<br>
3. 버퍼를 사용하는 Fprint 계열의 함수는 퍼버를 사용하지 않는 Print 계열의 함수 또는 입력된 인자들을 문자열로 변환하는 작업을 추가로 수행하는 Sprint 계열의 함수보다 빠르게 동작한다.<br>
4. 출력하고자 하는 인자가 하나이면서 동시에 형식 지정자를 따로 지정하지 않는 상황에서는 접미사로 f가 붙는 함수들이 성능이 가장 좋게 나타나는데, 이는 접두사와 접미사가 붙지 않는 함수들이 입력받은 인자가 문자열인지 문자열이 아닌지에 대한 행동을 구분하는 동작에서 시간 소모가 발생하기 때문이다.<br>
5. 반대로, 형식 지정자를 사용하는 형태로 인자를 넘겨주거나 입력된 인자의 갯수가 늘어나면, 접미사로 f가 붙는 함수의 시간 소모는 그만큼 늘어난다.<br>
6. 일반적인 상황에서 Fprint 함수를 사용하는 것이 적절하며, Print 함수 계열은 이미 래핑된 함수이면서 버퍼를 사용하지 않기에 좋은 성능을 기대할 수 없다. 여러 타입의 값들을 하나의 문자열로 묶어주고자 할 경우 Sprint 계열 함수를 사용하면 좋은 방안이 될 수 있다.<br>


##  Scanning<br>
### Scanning from STDIN<br>
```Text
func Scan(a ...interface{}) (n int, err error)
func Scanf(format string, a ...interface{}) (n int, err error)
func Scanln(a ...interface{}) (n int, err error)
```
* The Scan() function reads in **space-delimited** values into variables references passed into the function. It treats newlines as spaces.<br>
* The Scanf() does the same but **lets you specify formatting options** using a format string.<br>
* The Scanln() function works Scan() except that **it does not treat newlines as spaces.**<br>
* *These functions aren’t terribly useful because you will likely pass in data via flags, environment variables, or configuration files.*<br>
```Go
// Scan() example:
var name string
var age int
if _, err := fmt.Scan(&name, &age); err != nil {
        fmt.Println(err)
        os.Exit(1)
}
fmt.Printf("Your name is: %s\n", name)
fmt.Printf("Your age is: %d\n", age)
```

### Scanning from io.Reader<br>
```Text
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
```

### Scanning from a string<br>
```Text
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
```
* These functions can use to scan from an in-memory string.<br>
