package greeting

import (
	"fmt"
)

func greeting() {
	time := 14

	if time <= 12 {
		fmt.Println("good morning")
		time = 13
	}
	if time <= 18 {
		fmt.Println("good afternoon")
		time = 19
	}
	if time <= 24 {
		fmt.Println("good evening")
		time = 0
	}
}
