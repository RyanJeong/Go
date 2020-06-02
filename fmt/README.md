#   Formatted I/O Functions in Golang<br>
##  Verbs in Golang<br>
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

##  Width and Precision<br>
```Text
%8.0f ➡ "     101"
%8.1f ➡ "   100.6"
%8.2f ➡ "  100.57"
%8.3f ➡ " 100.567"
```

##  Left Alignment<br>
```Text
%-8.0f ➡ "101     "
%-8.1f ➡ "100.6   "
%-8.2f ➡ "100.57  "
%-8.3f ➡ "100.567 "
```

##  Zero Padding<br>
```Text
%08d ➡ "00000123"
```

##  Spacing<br>
```Text
// For example, formatting []byte{1,2,3,4} with and without the space flag:
%x  ➡ "01020304"
% x ➡ "01 02 03 04"
```

##  Other Verbs and Flags:<br>
* the [Printing](https://golang.org/pkg/fmt/#hdr-Printing) section of the [fmt](https://golang.org/pkg/fmt/) godoc.

##  Printing<br>
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

	fmt.Println(list1, list2)
	fmt.Print(list1, list2)
	fmt.Printf("%d%d", list1, list2)
	fmt.Printf("%d %d", list1, list2)
	// Output:
	// [1 2 3] [4 5 6]
	// [1 2 3] [4 5 6][1 2 3][4 5 6][1 2 3] [4 5 6]

	return
}
```
> I typically use Printf() when I need specific formatting options and Println() when I want default options. I almost always want a new line appended so I don’t personally use Print() much. One exception is if I am requesting interactive input from a user and I want the cursor immediately after what I print. For example, this line:<br>
```Go
fmt.Print("What is your name? ")
```
> Will output to the terminal with the cursor immediately after the last space:<br>
```Go
What is your name? █
```
>> https://medium.com/go-walkthrough/go-walkthrough-fmt-55a14bbbfc53<br>
