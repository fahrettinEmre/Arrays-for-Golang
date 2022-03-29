package main

import "fmt"

func main() {

	/* var colors = [4]int{1, 2, 3, 4}
	fmt.Println(colors)

	color := [4]int{1, 2, 3, 4}
	fmt.Println(color)

	clrs := []int{1, 2, 3, 4}
	fmt.Println(clrs)

	fmt.Println(len(clrs))
	fmt.Println(cap(clrs))

	var cars [3]string
	cars[0] = "bmw"
	cars[1] = "mercedes"
	cars[2] = "renualt"

	fmt.Println(cars)
	fmt.Println(len(cars))
	fmt.Println(cap(cars))

	i := 0
	for i <= len(cars)-1 {
		fmt.Println(cars[i])
		i++
	}
	*/

	/* a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	var s []int = a[3:6]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s)) // slice' lar cap için başladığı yerden sona kadar alır. Slice'lar veriyi kopyalamaz. Adreslerini alır.
	var d []int = a[:]
	fmt.Println(d)
	s[2] = 5
	fmt.Println(s)
	a[1] = 28
	fmt.Println(s)
	fmt.Println(a) */

	/* a := [4]int{1, 2, 3, 4}
	b := a
	fmt.Println(b)
	a[1] = 11
	fmt.Println(b)

	a1 := []int{1, 2, 3, 4}
	b1 := a1
	fmt.Println(b1)
	a1[1] = 11
	fmt.Println(b1)
	fmt.Println(a1) // slice larda işlem yapılırsa array da etkilenir.
	// çünkü slice lar veri tutmaz . adresteki veiriyi alır.
	fmt.Println(&b1[1])
	a1 = append(a1, 5)
	fmt.Println(a1)
	fmt.Println(b1)
	b1 = append(b1, 90)
	fmt.Println(a1)
	fmt.Println(b1) */

	/* var colors = []string{"mavi", "beyaz", "kırmızı"}
	fmt.Println(colors)
	colors = append(colors, "mor")
	fmt.Println(colors)
	b := append(colors[1:len(colors)])
	fmt.Println(b) */

	/* var colors = []int{1, 2, 3}
	fmt.Println(colors)
	colors = append(colors, 5)
	fmt.Println(colors)
	colors = append(colors)
	colors = append(colors[1:len(colors)]) */

	numbers := make([]int, 5, 5) // uzunluk 5 kapasite 10
	fmt.Println(numbers)
	numbers = append(numbers, 90)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

}
