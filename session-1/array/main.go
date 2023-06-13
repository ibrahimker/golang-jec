package main

import (
	"fmt"
)

func main() {
	//siswa := make([]string, 0)
	//var siswa []string
	siswa := []string{"anggi", "faqih", "budi"}
	//siswa2 := []string{"jbl", "bose", "sony"}
	for i := 0; i < len(siswa); i++ {
		fmt.Println(siswa[i])
	}
	fmt.Println("========")
	siswa = append(siswa, "roland")
	//siswa = append(siswa, siswa2...)
	for _, value := range siswa {
		fmt.Println(value)
	}

	siswa2 := []any{"anggi", 1, "budi"}
	fmt.Println(siswa2)

	fmt.Println("==============")
	mesin := make(map[string]any)
	mesin["ford"] = 1500
	mesin["toyota"] = "1200"
	for _, value := range mesin {
		fmt.Println(value)
	}
}
