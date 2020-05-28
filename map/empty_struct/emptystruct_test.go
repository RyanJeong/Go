package empty_struct

import (
	"fmt"
	"unsafe"
)

func Example_Empty_Struct() {
	m := map[int]struct{}{}
	i := []int{1, 2, 1, 2, 3}

	fmt.Println("m size:", unsafe.Sizeof(m))
	for _, r := range i {
		if _, exists := m[r]; exists {
			fmt.Printf("m[%d] is exists!\n", r)
			fmt.Printf("m[%d] size:%d\n", r, unsafe.Sizeof(m[r]))
		} else {
			fmt.Printf("m[%d] is not exists!\n", r)
			fmt.Printf("m[%d] size:%d\n", r, unsafe.Sizeof(m[r]))
			m[r] = struct{}{}
		}
	}

	// Output:
	// m size: 8
	// m[1] is not exists!
	// m[1] size:0
	// m[2] is not exists!
	// m[2] size:0
	// m[1] is exists!
	// m[1] size:0
	// m[2] is exists!
	// m[2] size:0
	// m[3] is not exists!
	// m[3] size:0
}
