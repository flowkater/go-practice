package main

// package oldmain

import "fmt"

var (
	start = rune(44032)
	end   = rune(55204)
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}

func hangul() {
	fruits := []string{"사과", "바나나", "토마토", "수박"}
	for _, fruit := range fruits {
		if HasConsonantSuffix(fruit) {
			fmt.Printf("%s은 맛있다.\n", fruit)
		} else {
			fmt.Printf("%s는 맛있다.\n", fruit)
		}

	}
}

func main() {
	const name string = "flowkater"
	var language string = "go"
	service := "Stoodi"
	fmt.Println(language)
	fmt.Println(service)
	fmt.Println(name)
	fmt.Println("I Love Go!")

	for i, r := range "가나다" {
		fmt.Println(i, r)
	}

	fmt.Println(len("가나다"))

	for _, r := range "가갛힣" {
		fmt.Println(string(r), r)
	}

	s := "abc"
	ps := &s

	fmt.Println(s)
	fmt.Println(ps)
	fmt.Println(*ps)
	s += "def"

	fmt.Println(s)
	fmt.Println(ps)
	fmt.Println(*ps)

	hangul()

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])

	// fmt.Println(nums[:-1]) // => X
	fmt.Println(nums[:len(nums)-1])

	nums = append(nums, 6, 7, 8)
	fmt.Println(nums)

	nums2 := []int{10, 11, 12, 13, 14, 15}
	nums3 := append(nums, nums2...)
	nums4 := append(nums[:5], nums2...)

	fmt.Println(nums2)
	fmt.Println(nums3)
	fmt.Println(nums4)

	ns := make([]int, 3, 5)

	fmt.Println("len(ns)", len(ns))
	fmt.Println("cap(ns)", cap(ns))

	ns2 := make([]int, 5)
	ns2 = ns2[:3]

	fmt.Println("len(ns2)", len(ns2))
	fmt.Println("cap(ns2)", cap(ns2))

	ns3 := make([]int, 0, 15)

	fmt.Println("len(ns3)", len(ns3))
	fmt.Println("cap(ns3)", cap(ns3))
}
