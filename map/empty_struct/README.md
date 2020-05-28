# What’s an empty struct?<br>
* An empty struct is a struct type without fields struct{}. The cool thing about an empty structure is that **it occupies zero bytes of storage**. You can find an accurate description about the actual mechanism inside the golang compiler in [this](https://dave.cheney.net/2014/03/25/the-empty-struct) post by Dave Chaney.<br>

> The short version is:<br>
The size of a struct is the sum of the size of the types of its fields, since there are no fields: no size!<br>

> A natural question would be:<br>
What can I use an empty struct for, if it has no fields?<br>

* Basically an empty struct can be used every time you are only interested in a property of a data structure rather than its value (this will be more clear in the practical examples).<br>

```Go
// Empty Struct Example Code:
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
```
