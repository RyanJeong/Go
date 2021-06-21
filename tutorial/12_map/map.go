package main

func main() {
	/* -------- map -------- */
	// maps keys to values
	// zero value of map is nil
	// nil map has no keys, nor can keys be added
	// make function returns a map of the given type, initialized and ready for use

}

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
