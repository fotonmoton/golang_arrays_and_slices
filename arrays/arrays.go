package arrays

import "fmt"

/*
1. All elements in an array are automatically assigned the zero value of the array type
2. The size of the array is a part of the type
3. Arrays are value types (copy by value)
*/

func InitToZeroValues() {
	var ints [2]int
	var strings [3]string
	var structs [2]struct {
		a int
		b string
	}

	fmt.Printf("Here arrays: ints: %+v, strings: %+v, structs: %+v", ints, strings, structs)
}

func InitWithValues() {
	grades := [2]int{1, 2}

	fmt.Printf("%+v", grades)
}

func CompileTimeLength() {
	grades := [...]int{1, 2, 3, 4, 5, 5, 5}

	fmt.Printf("%+v", grades)

}

func SeparateType() {
	// var grades [2]int = [3]int{1, 2}

	// fmt.Printf("%+v", grades)

}

func Indexing() {
	var animals = [2]string{"giraffe", "monkey"}

	fmt.Printf("%+v", animals[1])

	var outOfRange = 8

	animals[outOfRange] = "lion"

	fmt.Printf("%+v", animals)

}

func Counting() {

	var animals = [2]string{"giraffe", "monkey"}

	fmt.Printf("%d", len(animals))
}

func NotAppendable() {
	var animals = [2]string{"giraffe", "monkey"}

	// animals = append(animals, "lions")
	fmt.Printf("%+v", animals)
}

func FixedSize() {
	var ints [2]int

	// Compiler check
	// ints[3] = 2

	fmt.Printf("ints: %+v", ints)
}

func ForRangeArrays() {
	var animals = [2]string{"giraffe", "monkey"}

	for _, animal := range animals {
		fmt.Printf("%d - %s\n", 0, animal)
	}
}

func ArrayCopyByValue() {
	var animals = [2]string{"giraffe", "monkey"}

	var glassAnimals = animals

	glassAnimals[0] = "lion"

	fmt.Println(animals)
	fmt.Println(glassAnimals)
}

func EscapeAnalysis() *[2]string {
	var animals = [2]string{"giraffe", "monkey"}

	return &animals
}
