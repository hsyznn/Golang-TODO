package main

import (
	"encoding/json"
	"os"
)

/*
Go programlama dilinde jenerik bir veri depolama sistemi oluşturuyor. Dosya tabanlı bir JSON depolama mekanizması sağlıyor.
*/

type Storage[T any] struct {
	FileName string
	/*
		Storage[T any]: Bu, Go'nun 1.18 sürümüyle gelen jenerik yapıyı kullanıyor.
		T any: Herhangi bir veri tipini kabul eden bir tip parametresi.
		FileName: Verilerin kaydedileceği dosyanın adını tutan bir alan.
	*/
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}

	/*
		NewStorage: Yeni bir Storage nesnesi oluşturan bir yapıcı fonksiyon.
		Parametre olarak dosya adını alıyor ve belirtilen dosya adıyla bir Storage nesnesi döndürüyor.
		Döndürülen değer bir pointer (*Storage[T])
	*/
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, fileData, 0644)
	/*
		Save: Verilen veriyi JSON formatına dönüştürüp dosyaya kaydeden bir metot.
		json.MarshalIndent: Veriyi okunabilir bir JSON formatına dönüştürür.
		İlk parametre: Dönüştürülecek veri
		İkinci parametre: Başlangıç girintisi (burada boş)
		Üçüncü parametre: Her seviye için girinti (burada boş)
		os.WriteFile: JSON verisini dosyaya yazar.
		İlk parametre: Dosya adı
		İkinci parametre: Yazılacak veri
		Üçüncü parametre: Dosya izinleri (0644: sahibi okuyabilir ve yazabilir, diğerleri sadece okuyabilir)
	*/
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
		/*
			Load: Dosyadan veriyi okuyup, verilen nesneye yükleyen bir metot.
			os.ReadFile: Dosyadan veriyi byte dizisi olarak okur.
			json.Unmarshal: JSON verisini Go nesnesine dönüştürür.
			İlk parametre: JSON verisi (byte dizisi)
			İkinci parametre: Verinin yükleneceği nesne (pointer olarak)
		*/
	}

	return json.Unmarshal(fileData, data)
}
