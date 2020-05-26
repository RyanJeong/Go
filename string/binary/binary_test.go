package binary

import "fmt"

func Example_printBytes() { // byte: uint8의 별칭
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
