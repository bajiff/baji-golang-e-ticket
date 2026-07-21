package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
1. Di dalam fungsi main, buat Context dengan batas waktu (timeout) 2 detik.
   (Gunakan: ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second))
2. WAJIB: Pasang jaring pengaman `defer cancel()` tepat di bawahnya.
3. Buat sebuah pipa komunikasi (channel) untuk menampung pesan sukses:
   chSukses := make(chan string)

4. Jalankan sebuah goroutine anonim (`go func()`) untuk mensimulasikan API Bank yang lambat:
   a. Di dalam goroutine, buat program tertidur selama 5 detik: time.Sleep(5 * time.Second)
   b. Setelah bangun, kirimkan teks ke dalam pipa: chSukses <- "Checkout di Midtrans Berhasil!"

5. Di luar goroutine (kembali ke jalur utama fungsi main), buat blok `select` untuk mengadu dua kondisi:
   a. case pesan := <-chSukses:
      (Artinya: Jika pipa Bank merespons duluan)
      Cetak isi variabel `pesan` ke terminal.

   b. case <-ctx.Done():
      (Artinya: Jika stopwatch Context berbunyi duluan sebelum Bank merespons)
      Cetak teks peringatan: "Timeout! API Bank terlalu lambat, transaksi dibatalkan."
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


	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)	
	defer cancel()
	chSukses := make(chan string)

	go func(){
		time.Sleep(5 * time.Second)
		chSukses <- "Sukses di Midtrans Berhasil"
	}()
	
	select {
	case pesan := <- chSukses:
		fmt.Println("Berhasil cuy", pesan)
	case <- ctx.Done():
		fmt.Printf("Timeout! API Bank terlalu lambat, transaksi dibatalkan: %v\n", ctx.Done())
	}

}