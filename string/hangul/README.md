```Go
func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("고 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))
	// Output:
	// .
}
```
```Text
$ go test
--- FAIL: ExampleHasConsonantSuffix (0.00s)
got:
false
true
false
want:
.
FAIL
exit status 1
```
```Go
func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("고 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))
	// Output:
	// false
    // true
    // false
}
```
```Text
$ go test
PASS
ok
```
