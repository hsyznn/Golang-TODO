package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Bu kod, Go programlama dilinde bir komut satırı uygulaması için bayrak (flag) işleme mekanizması oluşturuyor.
Özellikle bir "Todo" (yapılacaklar) listesi uygulaması için komut satırı arayüzü sağlıyor.
*/

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
	/*
	   Bu yapı, programın kabul ettiği komut satırı bayraklarını tanımlar:
	   Add: Yeni bir görev eklemek için string değer
	   Del: Silinecek görevin indeksi (int)
	   Edit: Düzenlenecek görevin indeksi ve yeni başlığı (string)
	   Toggle: Tamamlanma durumu değiştirilecek görevin indeksi (int)
	   List: Tüm görevleri listelemek için boolean değer
	*/
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "Add", "", "Add a new Todo")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a todo by index")
	flag.IntVar(&cf.Del, "del", -1, "Specify todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify todo by index to toggle complete true/false")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.Parse()
	return &cf
	/*
		Yeni bir CmdFlags nesnesi oluşturur
		Her bir komut satırı bayrağını tanımlar:
		-Add: Yeni bir görev eklemek için (varsayılan değer: boş string)
		-Edit: Bir görevi düzenlemek için (varsayılan değer: boş string)
		-del: Bir görevi silmek için (varsayılan değer: -1)
		-toggle: Bir görevin tamamlanma durumunu değiştirmek için (varsayılan değer: -1)
		-list: Tüm görevleri listelemek için (varsayılan değer: false)
		flag.Parse() ile komut satırı argümanlarını ayrıştırır
		Oluşturulan CmdFlags nesnesinin adresini döndürür
	*/

}

func (cf *CmdFlags) Execute(tasks *Tasks) {
	switch {
	case cf.List:
		tasks.print()
	case cf.Add != "":
		tasks.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit. Please use index:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
		}
		tasks.edit(index, parts[1])
	case cf.Toggle != -1:
		tasks.toggle(cf.Toggle)
	case cf.Del != -1:
		tasks.delete(cf.Del)
	default:
		fmt.Println("Invalid Command")
		/*
		   Bu metot, komut satırı bayraklarına göre uygun eylemi gerçekleştirir:
		   switch ifadesi, hangi bayrağın kullanıldığını kontrol eder
		   Bayraklara göre işlemler:
		   cf.List true ise: Tüm görevleri listeler
		   cf.Add boş değilse: Yeni bir görev ekler
		   cf.Edit boş değilse:
		   Girilen değeri ":" karakterine göre böler (indeks:yeni_başlık formatında)
		   Doğru format kontrolü yapar
		   İndeksi string'den int'e dönüştürür
		   Görevi düzenler
		   cf.Toggle -1 değilse: Belirtilen görevin tamamlanma durumunu değiştirir
		   cf.Del -1 değilse: Belirtilen görevi siler
		   Hiçbir bayrak belirtilmemişse: "Invalid Command" mesajı gösterir

		*/
	}
}
