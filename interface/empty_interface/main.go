package main

import "fmt"

func main() {
	lista := []interface{}{10, 7.5, true, "texto"}

	for _, v := range lista {
		switch v.(type) {
		case string:
			fmt.Println(v, "string")
		default:
			fmt.Println(v)
		}
	}
}
