# Hash

hash adalah algoritma enkipsi untuk mengubah text menjadi deretan karakter acak, jumplah karakter hasil hashing selalu sama, hash termasuk _one-way encryption_ yang membuat hasil dari hash tidak bisa dikembalikan ke bentuk asli-nya.

## penerapa hash sha256

go menyediakan package `crypto/sha256` yang berisi library untuk keperluan _hashing_, cara penerapannya seperti di bawah ini.

```go
package main

import (
  "crypto/sha256"
  "encoding/hex"
  "fmt"
)

func Hash1(s string) string {
  sha := sha256.New()
	sha.Write([]byte(s))
	encrypted := sha.Sum(nil)
	return fmt.Sprintf("%x", encrypted)
}

func Hash2(s string) string {
  rahasia := sha256.Sum256([]byte(s))
	return hex.EncodeToString(rahasia[:])
}

func main(){
  text := "ini rahasia"
  hash1 := Hash1(text)
  fmt.Println(hash1)

  hash2 := Hash2(text)
  fmt.Println(hash2)
}
```

pada contoh di atas, terdapat dua cara melakukan enkripsi data. 
cara pertama pada function pertama `Hash1()` object bertipe `hash.Hash` dibuat menggunakan function `sha256.New()`, struct `hash.Hash` tersebut memiliki 2 method `Write()` dan `Sum()`.

- method `Write()` digunakan untuk men-set data yang akan di enkripsi, datanya harus dalam bentuk `[]byte`.
- method `Sum()`digunakan untuk mengeksekusi proses hashing, yang nantinya akan menghasilkan data yang sudah di enkripsi dalam bentuk `[]byte`, method ini membutuhkan parameter, isi dengan nil.

## metode salting pada hash sha256

Salt adalah nilai acak yang dihasilkan secara unik dan ditambahkan ke kata sandi sebelum di-hash. 
Karena hash adalah enkripsi satu arah dengan lebar data yang sudah pasti, sangat mungkin sekali hasil dari hash untuk beberapa data adalah sama.
ketika data yang akan mau di hash digabungkan dengan salt, maka hasil dari hashing dari gabungan dari data dan salt akan selalu berbeda meskipun data yang mau di hash adalah sama, karena setiap salt yang digunakan harus berbeda dengan data lain yang akan di hash.

```go
package main

import (
  "crypto/sha256"
	"encoding/hex"
  "fmt"
	"strconv"
	"time"
)

func main(){
  text := "ini rahasia"

  salt := strconv.FormatInt(time.Now().UnixNano(), 10)
	saltedPassword := text + salt

	sha := sha256.Sum256([]byte(saltedPassword))

	fmt.Println(hex.EncodeToString(sha[:]))
}
```

dari contoh diatas salt yang digunakan adalah hasil dari expresi `time.Now().UnixNano()`, hasilnya akan selalu unix untuk setiap detik-nya, karena scope terendah dari waktu tersebut adalah `nano second`.