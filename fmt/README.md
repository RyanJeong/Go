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

### Formatting to a string<br>
