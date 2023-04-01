package main

import "fmt"

func main() {
	hesap := 10000
	for true {
		fmt.Println(`**********İŞLEMLER**********
		1. Para Çekme 
		2. Para Yatırma
		3. Bakiye Sorma
		4. Çıkış
		
		Lütfen İşlem Tipinizi Seçiniz: `)

		var islem int
		fmt.Scanln(&islem)

		if islem == 1 {
			fmt.Println("Çekilecek Tutar Giriniz:")

			var cekilecek_tutar int
			fmt.Scanln(&cekilecek_tutar)

			if cekilecek_tutar <= hesap {
				fmt.Println("Paranızı Kutudan Alınız")
				hesap -= cekilecek_tutar
				fmt.Println("Bakiyeniz: ", hesap)
			} else {
				fmt.Println("Yetersiz Bakiye")
			}
		} else if islem == 2 {
			fmt.Println("Yatırılacak Tutarı Giriniz: ")

			var yatırılacak_tutar int
			fmt.Scanln(&yatırılacak_tutar)
			hesap += yatırılacak_tutar
			fmt.Println("Yeni Bakiye: ", hesap)
		} else if islem == 3 {
			fmt.Println("Hesap Bakiyeniz: ", hesap)
		} else if islem == 4 {
			break
		}
	}
}
