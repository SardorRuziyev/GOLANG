package main

import "fmt"

func main() {
	mp := map[string]int{
		"sam": 1,
		"cam": 2,
		"jam": 3,
		"bam": 4,
		"dam": 5,
		"ram": 6,
	}

	fmt.Println(mp)
	fmt.Printf("%#v\n", mp)
	fmt.Println(len(mp))
	fmt.Println("------ ----- ---- ")

	if v, ok := mp["xam"]; !ok {
		fmt.Println("key didnt exist")
	} else {
		fmt.Println("the valye prints", v)
	}

	for k, v := range mp {
		fmt.Println(k, v)

	}

	for _, v := range mp {
		fmt.Println(v)
	}

	for k := range mp {
		fmt.Println(k)
	}

	xi := []int{42, 43, 44, 45, 46}

	for k, v := range xi {
		fmt.Println(k, v)

	}

	for _, v := range xi {
		fmt.Println(v)
	}

	for k := range xi {
		fmt.Println(k)
	}

	delete(mp, "cam")

}
