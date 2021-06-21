package main

import "fmt"

func main() {
	/* -------- array -------- */
	// array: 'specifies a length' and an element type
	// The type [n]T is an array of n values of type T
	// ex. type [2]int is defferent than type [4]int

	var a1 [3]int // declaration
	// By default an array is zero-valued
	fmt.Println(a1)
	a1[0] = 1
	a1[1] = 2
	a1[2] = 4
	fmt.Println(a1)
	var a2 = [3]int{1, 2, 4} // declare and initialize
	fmt.Println(a2)
	var a3 = [...]int{1, 2, 4}
	fmt.Println(a3)
	var a4 [2][3]int
	a4[0][0] = 5
	a4[0][1] = 4
	a4[0][2] = 3
	a4[1][0] = 2
	a4[1][1] = 1
	a4[1][2] = 0
	fmt.Println(a4)
	var a5 = [2][3]int{
		{5, 4, 3},
		{2, 1, 0}, // MUST have comma
		// if it wouldn’t be there, Go would have added a semicolon(;)
		// by the rule which would have CRASHED the program
	}
	fmt.Println(a5)

	/* -------- slice -------- */
	// slice: array-like type, but slice has no specified length

	s := []int{0, 1, 2, 3, 4, 5}
	s = s[2:5]
	fmt.Println(s) // [2 3 4]

	s = s[1:]
	fmt.Println(s) // [3 4]

	s = append(s, 2)
	fmt.Println(s) // [3 4 2]

	s = append(s, 8, 9, 10)
	fmt.Println(s) // [3 4 2 8 9 10]

	// declare a slice, len = 0, cap = 3
	//
	// cap tells you the capacity of the underlying array
	// len tells you how many items are in the array.
	//
	// ex.
	// s := []int{0, 1, 2, 3, 4, 5}
	//
	// [ptr][len][cap]
	//   |  [ 6 ][ 6 ]
	//   |
	//   | [0][1][2][3][4][5]
	//   +--^
	//  len -  -  -  -  -  - (6)
	//  cap -  -  -  -  -  - (6)
	//
	// s = s[1:3]
	//
	// [ptr][len][cap]
	//   |  [ 2 ][ 5 ]
	//   |
	//   | [0][1][2][3][4][5]
	//   +-----^
	//  len    -  -          (2)
	//  cap    -  -  -  -  - (5)
	fmt.Println(len(s), cap(s)) // 6 6
	s = s[1:3]                  // includes the first element, but excludes the last one, [1,3)
	fmt.Println(len(s), cap(s)) // 2 5

	slice := make([]int, 0, 3)
	for i := 1; i <= 15; i++ {
		slice = append(slice, i)
		fmt.Println(len(slice), cap(slice))
		/*
			1 3
			2 3
			3 3
			4 6
			5 6
			6 6
			7 12
			8 12
			9 12
			10 12
			11 12
			12 12
			13 24
			14 24
			15 24
		*/
	}
	fmt.Println(slice)

	// slice... -> replace to [1 2 3 ... 13 14 15]
	fmt.Println(len(s), cap(s))         // 6 6
	fmt.Println(len(slice), cap(slice)) // 15 24
	s = append(s, slice...)
	fmt.Println(len(s), cap(s)) // 21 22
	fmt.Println(s)              // [3 4 2 8 9 10 1 2 3 ... 13 14 15]

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)

	// function copy returns the number of elements copied,
	// which will be the minimum of len(dst) and len(src)
	copy(target, source)
	fmt.Println(len(target), cap(target)) // 3 6
	fmt.Println(target)                   // [0 1 2]
	copy(target, s)
	fmt.Println(len(target), cap(target)) // 3 6
	fmt.Println(target)                   // [3 4 2]

	/* -------- map -------- */
	// maps keys to values
	// zero value of map is nil
	// nil map has no keys, nor can keys be added
	// make function returns a map of the given type, initialized and ready for use

}

/*
package main

import "fmt"

func main() {
		s := []int{2, 3, 5, 7, 11, 13}
			printSlice(s)

				// Slice the slice to give it zero length.
					s = s[:0]
						printSlice(s)

							// Extend its length.
								s = s[:4]
									printSlice(s)

										// Drop its first two values.
											s = s[2:]
												printSlice(s)
											}

											func printSlice(s []int) {
													fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
												}

*/

/* nil slice
// By default ptr pointer to nil, len and cap is zero
package main

import "fmt"

func main() {
		var s []int
			fmt.Println(s, len(s), cap(s))
				if s == nil {
							fmt.Println("nil!")
								}
							}

*/

/* slice with make

package main

import "fmt"

func main() {
		a := make([]int, 5)
			printSlice("a", a)

				b := make([]int, 0, 5)
					printSlice("b", b)

						c := b[:2]
							printSlice("c", c)

								d := c[2:5]
									printSlice("d", d)
								}

								func printSlice(s string, x []int) {
										fmt.Printf("%s len=%d cap=%d %v\n",
												s, len(x), cap(x), x)
											}

*/

/* slices of slices
// tic-tac-toe
package main

import (
		"fmt"
			"strings"
		)

		func main() {
				// Create a tic-tac-toe board.
					board := [][]string{
								[]string{"_", "_", "_"},
										[]string{"_", "_", "_"},
												[]string{"_", "_", "_"},
													}

														// The players take turns.
															board[0][0] = "X"
																board[2][2] = "O"
																	board[1][2] = "X"
																		board[1][0] = "O"
																			board[0][2] = "X"

																				for i := 0; i < len(board); i++ {
																							fmt.Printf("%s\n", strings.Join(board[i], " "))
																								}
																							}
*/

/* append
package main

import "fmt"

func main() {
		var s []int
			printSlice(s)

				// append works on nil slices.
					s = append(s, 0)
						printSlice(s)

							// The slice grows as needed.
								s = append(s, 1)
									printSlice(s)

										// We can add more than one element at a time.
										s = append(s, 2, 3, 4) // cap: 2, If the backing array of s is too small to fit all the given values a bigger array will be allocated. (capXn) => 6
										// The returned slice will point to the newly allocated array.
												printSlice(s)
											}

											func printSlice(s []int) {
													fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
												}

*/

/* Map intro
ickage main

import "fmt"

type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex

	func main() {
			m = make(map[string]Vertex)
				m["Bell Labs"] = Vertex{
							40.68433, -74.39967,
								}
									fmt.Println(m["Bell Labs"])
								}

*/

/* Map
package main

import "fmt"

type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
			"Bell Labs": Vertex{
						40.68433, -74.39967,
							},
								"Google": Vertex{
											37.42202, -122.08408,
												},
											}

											func main() {
													fmt.Println(m)
												}

*/

/* Map 2
package main

import "fmt"

type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
			"Bell Labs": {40.68433, -74.39967},
				"Google":    {37.42202, -122.08408},
			}

			func main() {
					fmt.Println(m)
				}

*/

/* Mutating Map
package main

import "fmt"

func main() {
		m := make(map[string]int)

			m["Answer"] = 42
				fmt.Println("The value:", m["Answer"])

					m["Answer"] = 48
						fmt.Println("The value:", m["Answer"])

							delete(m, "Answer")
								fmt.Println("The value:", m["Answer"])

									v, ok := m["Answer"]
										fmt.Println("The value:", v, "Present?", ok)
									}

*/

/* Mutating Map 2
package main

func main() {
	    tickers := map[string]string{
			        "GOOG": "Google Inc",
					        "MSFT": "Microsoft",
							        "FB":   "FaceBook",
									        "AMZN": "Amazon",
											    }

												    // map 키 체크
													    val, exists := tickers["MSFT"]
														    if !exists {
																        println("No MSFT ticker")
																		    }
																		}
*/

/* Mutating Map 3
package main

import "fmt"

func main() {
	    myMap := map[string]string{
			        "A": "Apple",
					        "B": "Banana",
							        "C": "Charlie",
									    }

										    // for range 문을 사용하여 모든 맵 요소 출력
											    // Map은 unordered 이므로 순서는 무작위
												    for key, val := range myMap {
														        fmt.Println(key, val)
																    }
																}
*/
