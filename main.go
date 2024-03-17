package main

import "fmt"

func main() {
	redisClone := NewRedisClone()

	redisClone.AddKeyValue("a", "b")

	aVal, err := redisClone.GetValue("a")

	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Printf("A Val: %s\n", aVal)
	}

	redisClone.AddKeyValue("c", "e")

	cVal, err := redisClone.GetValue("c")

	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Printf("C Val: %s\n", cVal)
	}
}
