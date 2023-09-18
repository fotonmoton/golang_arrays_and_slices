package slices

import "fmt"

/*
1. All elements in a slice are automatically assigned the zero value of the array type
2. The size of the slice is dynamic
3. Slices are pointer types (passed by pointer)
*/

func Example() {
	var slice = make([]byte, 5)
	sliceLiteral := []byte{}

	sliceLiteral = append(sliceLiteral, 0)
	slice[0] = sliceLiteral[0]

}

func Declaration() {
	var animals = []string{"giraffe"}

	fmt.Printf("%+v", animals)
}

func Appendable() {
	var animals = []string{"giraffe"}

	animals = append(animals, "lion")

	fmt.Printf("%+v", animals)
}

func CanAllocateMemory() {
	var animals = make([]string, 0, 10)

	fmt.Printf("animals: %+v, len: %d, cap: %d\n", animals, len(animals), cap(animals))

	animals = append(animals, "lion", "giraffe")

	fmt.Printf("animals: %+v, len: %d, cap: %d\n", animals, len(animals), cap(animals))
}

func CanGrow() {
	var animals = make([]string, 0, 1)

	fmt.Printf("animals: %+v, len: %d, cap: %d\n", animals, len(animals), cap(animals))

	animals = append(animals, "lion")

	fmt.Printf("animals: %+v, len: %d, cap: %d\n", animals, len(animals), cap(animals))

	animals = append(animals, "giraffe")

	fmt.Printf("animals: %+v, len: %d, cap: %d\n", animals, len(animals), cap(animals))

	animals = append(animals, "zebra")

	fmt.Printf("animals: %+v, len: %d, cap: %d\n", animals, len(animals), cap(animals))

	for range animals {
		animals = append(animals, "zebra", "zebra")
	}

	fmt.Printf("animals: %+v, len: %d, cap: %d\n", animals, len(animals), cap(animals))

}

func ArrayIntoSlice() {
	var animals [1]string

	// convert array into slice
	animalsSlice := animals[0:1]

	fmt.Printf(
		"animals: %+v, len: %d, cap: %d\nanimalsSlice: %+v, len: %d, cap: %d\n",
		animals,
		len(animals),
		cap(animals),
		animalsSlice,
		len(animalsSlice),
		cap(animalsSlice),
	)

	animalsSlice[0] = "giraffe"

	fmt.Printf(
		"animals: %+v, len: %d, cap: %d\nanimalsSlice: %+v, len: %d, cap: %d\n",
		animals,
		len(animals),
		cap(animals),
		animalsSlice,
		len(animalsSlice),
		cap(animalsSlice),
	)

	// // Creates new copy of underlying array and do not modifies original one
	animalsSlice = append(animalsSlice, "lion")

	animalsSlice[0] = "monkey"
	fmt.Printf(
		"animals: %+v, len: %d, cap: %d\nanimalsSlice: %+v, len: %d, cap: %d\n",
		animals,
		len(animals),
		cap(animals),
		animalsSlice,
		len(animalsSlice),
		cap(animalsSlice),
	)

	sliceIntoLion := animalsSlice[1:]

	fmt.Printf(
		"animals: %+v, len: %d, cap: %d\nanimalsSlice: %+v, len: %d, cap: %d\nsliceIntoLion: %+v, len: %d, cap: %d",
		animals,
		len(animals),
		cap(animals),
		animalsSlice,
		len(animalsSlice),
		cap(animalsSlice),
		sliceIntoLion,
		len(sliceIntoLion),
		cap(sliceIntoLion),
	)

	sliceIntoLion[0] = "Dead lion"

	fmt.Printf(
		"animals: %+v, len: %d, cap: %d\nanimalsSlice: %+v, len: %d, cap: %d\nsliceIntoLion: %+v, len: %d, cap: %d",
		animals,
		len(animals),
		cap(animals),
		animalsSlice,
		len(animalsSlice),
		cap(animalsSlice),
		sliceIntoLion,
		len(sliceIntoLion),
		cap(sliceIntoLion),
	)
}

func RemoveElement() {
	i := 2
	animals := []string{"giraffe", "dead lion", "monkey"}
	fmt.Printf("animals: %+v, len: %d, cap: %d\n", animals, len(animals), cap(animals))
	animals = append(animals[:i], animals[i+1:]...)
	fmt.Printf("animals: %+v, len: %d, cap: %d\n", animals, len(animals), cap(animals))

}

func MultiDimensional() {

	chessBoard := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8}}

	fmt.Printf("%v\n", chessBoard)

	chessBoard[0][1] = 9

	fmt.Printf("%v\n", chessBoard)

}

func CanBeCopied() {
	animals := []string{"giraffe", "lion"}

	copyAnimals := make([]string, 10000)

	smallAnimals := make([]string, 3)

	copy(smallAnimals, copyAnimals[1:4])

	copyAnimals = nil

	copy(copyAnimals, animals)

	animals[0] = "monkey"

	fmt.Printf("animals: %+v, len: %d, cap: %d\n", animals, len(animals), cap(animals))
	fmt.Printf("copyAnimals: %+v, len: %d, cap: %d\n", copyAnimals, len(copyAnimals), cap(copyAnimals))
}

func StrangeSyntax() {
	// The followings slice composite literals
	// are equivalent to each other.
	a := []string{"break", "continue", "fallthrough"}
	b := []string{0: "break", 1: "continue", 2: "fallthrough"}
	c := []string{2: "fallthrough", 1: "continue", 0: "break"}
	d := []string{2: "fallthrough", 0: "break", "continue"}

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
}
