package main

import "fmt"

func main() {
	//testScores := map[string]float32{
	//	"Henry":    87.3,
	//	"Hermione": 105,
	//	"Ronald":   63.5,
	//	"Neville":  27,
	//}
	//
	//c := clone(testScores)
	//
	//fmt.Println(c)
	generic()
}

//func clone[V any](s []V) []V {
//	result := make([]V, len(s))
//	for i, v := range s {
//		result[i] = v
//	}
//	return result
//}

func clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func generic() {
	a1 := []int{1, 2, 3}
	a2 := []float64{3.14, 6.02}
	a3 := []string{"foo", "bar", "baz"}

	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)

	fmt.Printf("Sum of %v: %v\n", a1, s1)
	fmt.Printf("Sum of %v: %v\n", a2, s2)
	fmt.Printf("Sum of %v: %v\n", a3, s3)
}

type addable interface {
	int | float64 | string
}

func add[V addable](s []V) V {
	var result V
	for _, v := range s {
		result += v
	}
	return result
}
