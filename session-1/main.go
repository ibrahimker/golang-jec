package main

import "fmt"

// ini function main
/*
	ini function main sample komen berbeda
*/
func main() {
	fmt.Println("hello world")
	fmt.Println("hello", "world")

	// contoh variabel panjang
	var nama = "ibam"
	var nama2 string = "budi"
	fmt.Println(nama, nama2)

	// contoh variabel golang short declaration
	orangPertama := "saya" // camelcase
	orang_kedua := "dia"   // snake case
	var tiga int = 3
	tigax := 3

	malamHari := true

	fmt.Println(orangPertama, orang_kedua)
	fmt.Println(tiga, tigax)
	fmt.Printf("Orang pertama adalah %s dan Orang kedua adalah %s\n", orangPertama, orang_kedua)
	fmt.Printf("%T \n", tiga)

	fmt.Println(malamHari)

	// multiple declaration
	orangKetiga, orangKeempat := "aku", "kamu"
	fmt.Println(orangKetiga, orangKeempat)

	// underscore variable
	orangKelima, _, _ := daftarOrang()
	fmt.Println(orangKelima)

	// constanta
	const nama3 string = "ibam"

	// contoh if
	fmt.Println(len(nama))
	if len(nama) > 5 {
		fmt.Println(nama)
	}

	fmt.Println("==============================")
	//x, _ := daftarOrang()
	//if x == "aku" {
	//	fmt.Println(x)
	//}

	if x, _, err := daftarOrang(); err == nil {
		fmt.Println(x)
	} else {
		fmt.Println(err)
	}

	score := 7
	switch score {
	case 0:
		fmt.Println("jelek")
	case 9:
		fmt.Println("bagus")
	default:
		fmt.Println("mantap")
	}

	fmt.Println("===========START LOOPING===========")
	//for (int i = 0; i < 9; i++) {
	//	// isinya
	//}
	for i := 0; i < 9; i++ {
		fmt.Println(i)
	}

	for {
		if score == 10 {
			break
		}
		fmt.Println("Score", score)
		score++
	}
}

// function sample
func daftarOrang() (string, string, error) {
	return "aku", "kamu", nil
}
