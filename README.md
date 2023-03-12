# GOLANG HASH PASSWORD

Bu yazıda bcrypt paketi kullanarak go dilinde şifrelerin nasıl hash edileceğini ve hash edilmiş şifreyi nasıl karşılaştırabileceğinizden bahsedeceğim.

Aşağıda vereceğim örnekleri https://github.com/mkaganm/GolangPasswordHashing linkteki repo da bulabilirsiniz.

Öncelikle bcrypt kütüphasini indirmemiz gerekiyor. Aşağıdaki komutu kullanarak paketi indirebilirsiniz.

```bash
go get golang.org/x/crypto/bcrypt
```

Bundan sonra yazacağımız uygulama bu kütüphaneden faydalanabilecektir.

### Şifreyi hash e çevirmek için
```go
func PasswordToHash(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
```
Yukarıdaki fonksiyonu kullanarak girdiğiniz şifreyi hash e çevirebilirsiniz. Burada bcrypt paketinde bulunan GenerateFromPassword fonksiyonunu kullandık.

GenerateFromPassword, verilen int değer karşılığında parolanın hash değerini döndürür. Verilen int değeri MinCost’tan azsa, bunun yerine DefaultCost olarak ayarlanacaktır. GenerateFromPassword, bcrypt’in çalışacağı en uzun parola olan 72 bayttan uzun parolaları kabul etmez.

### Hash e çevirilen şifreyi karşılaştırmak için

```go
func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

Yukarıdaki fonksiyonu kullanarak verdiğiniz hash değerini ve şifreyi karşılaştırabilir doğru veya yanlış olduğunu teyit edebilirsiniz.

CompareHashAndPassword, bir hash parolasını olası düz metin eşdeğeriyle karşılaştırır. Başarı durumunda nil veya başarısızlık durumunda hata döndürür.

---

```go
func main() {

    password := "123456"

    hash, _ := PasswordToHash(password)

    match := CheckPassword(password, hash)

    fmt.Println("Password\t:", password)
    fmt.Println("Hash\t\t:", hash)
    fmt.Println("Match\t\t:", match)
}
```

Yukarıda verdiğim main fonksiyounu çalıştırdığınızda aşağıdaki sonuç ile karşılaşacaksınız.

```bash
go run main.go
```
```
Password	: 123456
Hash		: $2a$14$tcma..Er/KDgluAHgGR10.8RJxKyT.qOPajPbMJnStUwrLqiDOg8i
Match		: true
```

