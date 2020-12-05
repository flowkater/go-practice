package hangul

import "fmt"

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))
	// Output:
	// false
	// true
	// false
}

func Example_array() {
	fruits := [...]string{"사과", "바나나", "토마토", "수박"}
	for _, fruit := range fruits {
		if HasConsonantSuffix(fruit) {
			fmt.Println("$s는 맛있다.\n", fruit)
		} else {
			fmt.Println("$s은 맛있다.\n", fruit)
		}

	}
	// Output
	// .
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	// 수박는 맛있다.
}
