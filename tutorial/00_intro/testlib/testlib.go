package testlib

import "fmt"

var rock map[string]string

func init() {
	rock = make(map[string]string)
	rock["Muse"] = "Time Is Running Out"
	rock["Metalica"] = "Master of Puppets"
	rock["Green Day"] = "Basket Case"
}

// public
func GetMusic(artist string) string {
	return rock[artist]
}

// non-public
func getKeys() {
	for _, v := range rock {
		fmt.Println(v)
	}
}
