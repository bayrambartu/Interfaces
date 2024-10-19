// Real-World Example: Data Storage Layer

package main

import (
	"fmt"
)

type Storage interface {
	VeriKaydet(veri string) error
	VeriGetir(id int) (string, error)
}

// MySQL veri depolama yapısı
type MySQLDepo struct{}

func (m MySQLDepo) VeriKaydet(veri string) error {
	fmt.Println("MySQL'e veri kaydediliyor:", veri)
	return nil
}

func (m MySQLDepo) VeriGetir(id int) (string, error) {
	fmt.Println("MySQL'den veri getiriliyor with id:", id)
	return "MySQL Veri", nil
}

// MongoDB veri depolama yapısı
type MongoDepo struct{}

func (m MongoDepo) VeriKaydet(veri string) error {
	fmt.Println("MongoDB'ye veri kaydediliyor:", veri)
	return nil
}

func (m MongoDepo) VeriGetir(id int) (string, error) {
	fmt.Println("MongoDB'den veri getiriliyor with id:", id)
	return "MongoDB Veri", nil
}

// Depolama işlemleri yapan fonksiyon
// func DepolamaIslemleri(depo Storage) {
// 	depo.VeriKaydet("Yeni Veri")
// 	depo.VeriGetir(1)
// }

func DepolamaIslemleri(depo Storage, data string, id int) {

	depo.VeriKaydet(data)
	depo.VeriGetir(id)
}

func main() {
	// mysqlDepo := MySQLDepo{}
	// DepolamaIslemleri(mysqlDepo)

	// mongoDepo := MongoDepo{}
	// DepolamaIslemleri(mongoDepo)

	mysqlDepo := MySQLDepo{}
	DepolamaIslemleri(mysqlDepo, "merhaba", 1)

	mangoDepo := MongoDepo{}
	DepolamaIslemleri(mangoDepo, "bartu", 2)

}
