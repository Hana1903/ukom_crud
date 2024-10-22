package utils

import "unicode"

func IsNumeric(s string) bool {
    for _, r := range s { //looping variabel r yang menyimpan hasil iterasi
        if !unicode.IsDigit(r) { //fungsi unicode mengembalikan true jika karakter merupakan angka, dan false jika bukan angka. Jika ditemukan satu karakter yang bukan angka, fungsi langsung mengembalikan false.
            return false
        }
    }
    return true
}

//IsNumeric menerima argumen berupa string s & akan mengembalikan
//true jika semua karakter dlm string adalah angka 0-9 dan false
//jika karakter bukan angka

