package main

import "fmt"

func main() {
	var i, n_passed, n_failed, n int
	var n1, n2, n3, avg float64

	fmt.Print("berapa jumlah siswa yang nilainya akan diproses")
	fmt.Scan(&n)
	n_passed = 0
	n_failed = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&n1, &n2, &n3)
		avg = (n1 + n2 + n3) / 3
		if avg > 80.0 {
			fmt.Print("Memenuhi syarat administrasi")
			n_passed = n_passed + 1
		} else {
			fmt.Print("Tidak memenuhi syarat administarsi")
			n_failed = n_failed + 1

		}

	}
	fmt.Print("Jumlah siswa lolos seleksi administrasi", n_passed)
	fmt.Print("Jumlah siswa tidak lolos seleksi administrasi", n_failed)

}
