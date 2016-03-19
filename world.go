package main
import (
	"fmt"
	"time"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("The current time is", time.Now())
	fmt.Println("Here's the output of two numbers:", add(32, 4))
	a, b := swap("hello", "world")
	fmt.Println("Here's the output of two strings swapped:", a, b)
	i, j := 42, 99

	p := &i // Point to i
	fmt.Println(*p, j) // read i through the pointer
	p = &j // Point to j
	*p = *p - 10 // Subtract 10 from j through the pointer
	fmt.Println(j) // Print j's new value
	v := Vertex{1, 2} // v is a new Vertex struct
	g := &v // g is a pointer to type Vertex
	g.X = 32 // g accesses the field X and changes its value to 32
	fmt.Println(v)  // Print v
	var arr [2]string
	arr[0] = "hello"
	arr[1] = "world"
	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}