package main

import "fmt"

func main() {

	nama := "Fazrul"
	age := 20
	weight := 64.5
	address := "Jl. Manado"

	var nama1 string = "Fazrul 2"
	var age1 int = 20
	var weight1 float32 = 64.5
	var address1 = "Jl. Manado"

	var (
		nama2    string  = "Fazrul 3"
		age2     int     = 20
		weight2  float32 = 64.5
		address2         = "Jl. Manado"
	)

	const name3 string = "Fazrul 4"
	const age3 int = 20
	const weight3 float32 = 64.5
	const address3 = "Jl. Manado"

	const (
		name4    string  = "Fazrul 3"
		age4     int     = 20
		weight4  float32 = 64.5
		address4         = "Jl. Manado"
	)

	fmt.Println("Nama :", nama)
	fmt.Println("Umur :", age)
	fmt.Println("Berat Badan :", weight)
	fmt.Println("Alamat :", address)

	fmt.Println("=======================================")

	fmt.Println("Nama :", nama1)
	fmt.Println("Umur :", age1)
	fmt.Println("Berat Badan :", weight1)
	fmt.Println("Alamat :", address1)

	fmt.Println("=======================================")

	fmt.Println("Nama :", nama2)
	fmt.Println("Umur :", age2)
	fmt.Println("Berat Badan :", weight2)
	fmt.Println("Alamat :", address2)

	fmt.Println("=======================================")

	fmt.Println("Nama :", name3)
	fmt.Println("Umur :", age3)
	fmt.Println("Berat Badan :", weight3)
	fmt.Println("Alamat :", address3)

	fmt.Println("=======================================")

	fmt.Println("Nama :", name4)
	fmt.Println("Umur :", age4)
	fmt.Println("Berat Badan :", weight4)
	fmt.Println("Alamat :", address4)
}
