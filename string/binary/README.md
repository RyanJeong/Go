```Go
package binary

import "fmt"

func Example_printBytes() {
	s := "가나다"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output:
	// b0:80:eb:82:98:eb:8b:a4:
}

func Example_printBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	// Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	// s := "가나다"
	// s[2]++
	// -> not working

	// string: 읽기 전용, bytes: 읽기 쓰기 가능
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	// Output:
	// 각나다
}
```

```Text
$ go test
--- FAIL: Example_printBytes (0.00s)
got:
ea:b0:80:eb:82:98:eb:8b:a4:
want:
b0:80:eb:82:98:eb:8b:a4:
FAIL
exit status 1
FAIL
```

```Go
package binary

import "fmt"

func Example_printBytes() {
	s := "가나다"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	// Output:
	// eab080eb8298eb8ba4
	// b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	// s := "가나다"
	// s[2]++
	// -> not working

	// string: 읽기 전용, bytes: 읽기 쓰기 가능
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	// Output:
	// 각나다
}
```

```Text
$ go test
--- FAIL: Example_printBytes2 (0.00s)
got:
eab080eb8298eb8ba4
ea b0 80 eb 82 98 eb 8b a4
want:
eab080eb8298eb8ba4
b0 80 eb 82 98 eb 8b a4
FAIL
exit status 1
FAIL
```

```Go
package binary

import "fmt"

func Example_printBytes() {
	s := "가나다"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	// Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	// s := "가나다"
	// s[2]++
	// -> not working

	// string: 읽기 전용, bytes: 읽기 쓰기 가능
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	// Output:
	// 각나다
}
```

```Text
$ go test
PASS
ok
```
