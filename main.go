package main

import (
	"fmt"
	"sync"
)

/*
1. Deklarasikan variabel global stokTiket = 1
2. Di dalam fungsi main, buat variabel WaitGroup (misal: var wg sync.WaitGroup)
3. Buat perulangan (for loop) sebanyak 50 kali.
4. Di dalam loop, beritahu WaitGroup ada 1 goroutine baru (wg.Add(1)).
5. Masih di dalam loop, panggil sebuah goroutine anonim (pakai kata kunci `go func()`).
6. Di dalam goroutine tersebut:
     a. Tulis pengecekan: JIKA stokTiket > 0
     b. Cetak ke terminal "Tiket berhasil dibeli!"
     c. Kurangi stokTiket sebanyak 1 (stokTiket--)
     d. Akhiri tugas goroutine (wg.Done())
7. Di luar loop, di akhir fungsi main, perintahkan WaitGroup untuk menunggu (wg.Wait())
8. Cetak hasil akhir stokTiket ke terminal.

*/

var stokTiket = 1
var mutex  sync.Mutex

func main() {
	var wg sync.WaitGroup
	for index :=  1; index <= 50; index++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			if stokTiket > 0 {
				fmt.Println("Tiket berhasil di beli")
				stokTiket--
			}
		}()
	}
	wg.Wait()
	fmt.Println(stokTiket)

}