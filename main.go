package main

import "fmt"

func main() {

	size := 0
	arr := make([]int, size)
	input := 0

	fmt.Println("Array size\t: ", "Please input size of array")
	fmt.Println("EXIT\t\t: ", "Ctrl + c")
	fmt.Println()
	fmt.Print("Size\t\t: ")
	fmt.Scanf("%v", &size)

	for len(arr) != size {
		fmt.Println("=======================================")
		fmt.Println("Insert\t:", "dont input 0 and character")
		fmt.Println("Pop\t:", "just input 0")
		fmt.Println("EXIT\t:", "Ctrl + c")
		fmt.Print("Input\t: ")
		fmt.Scanf("%v", &input)
		fmt.Println("len", len(arr))
		if input == 0 && len(arr) >= 1 {
			fmt.Print("\033[H\033[2J")
			arr = arr[:len(arr)-1]
			fmt.Println("Arr", arr, "size", size-(len(arr)))
		} else {
			fmt.Print("\033[H\033[2J")
			arr = append(arr, input)
			fmt.Println("Arr", arr, "size", size-(len(arr)))
		}

	}
	fmt.Print("\033[H\033[2J")
	fmt.Println("Thank you.")
	fmt.Println("Result", arr)

}
