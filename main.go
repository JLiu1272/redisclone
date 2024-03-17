package main

import "fmt"

func redisClonePlayground() {
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

func listPlayground() {
	list := NewLinkList()

	// list.Insert(1)
	// list.Insert(2)
	// list.Insert(3)
	// list.Insert(4)
	// list.Insert(5)

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Display()
}

func main() {
	redisClonePlayground()
	listPlayground()
}
