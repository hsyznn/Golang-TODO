package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Task struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Tasks []Task

/*
Burada yapılacaklar todo.go yapısını ve işlevselliğini tanımlayacağız.
Bunlar bir yapılacaklar listesi için ihtiyaç duyacağımız tüm alanlardır.
Ayrıca tüm yapılacaklar listemizi tutacak bir dilime de ihtiyacımız var.
*/

func (tasks *Tasks) add(title string) {
	task := Task{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*tasks = append(*tasks, task)
}

/*
Bu satır, yeni oluşturulan görevi tasks listesine ekliyor.
/*tasks ifadesi, pointer'ın gösterdiği asıl Tasks nesnesine erişiyor.
append fonksiyonu, bir dilime (slice) yeni bir eleman ekler ve yeni dilimi döndürür.
Sonuç tekrar *tasks'a atanarak, orijinal Tasks nesnesi güncelleniyor.
*/

/*
kaldırma, düzenleme veya geçiş gibi işlemler için sağlanan dizinin geçerli olup olmadığını
kontrol eden bir yardımcı yöntem ekleyeceğiz
*/

func (tasks *Tasks) validateIndex(index int) error {
	if index < 0 || index >= len(*tasks) {
		err := errors.New("task not found")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

/*
Dizin sınırların dışındaysa, 'geçersiz dizin' hatası yazdırır ve döndürür.
Aksi takdirde, dizinin geçerli olduğunu belirten nil döndürür.
*/

func (tasks *Tasks) delete(index int) error {
	t := *tasks
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*tasks = append(t[:index], t[index+1:]...)

	return nil
}

/*
Bu, silinecek öğeden önceki tüm öğeleri (t[:index]) ve silinecek öğeden
sonraki tüm öğeleri (t[index+1:]) birleştirir.
Yöntemde delete, geçerli aralıkta olduğundan emin olmak için önce yardımcı yöntemimizi kullanarak sağlanan dizini doğrularız.
Dizin geçerliyse, listeyi belirtilen dizinde bölerek yapılacaklar öğesini kaldırırız.
Daha sonra dizinden önceki ve sonraki iki bölümü birleştiririz, bu da öğeyi listeden kaldırır.
*/

func (tasks *Tasks) toggle(index int) error {

	//Değişken Tanımlaması ve İndeks Doğrulama
	t := *tasks
	if err := t.validateIndex(index); err != nil {
		return err
	}

	//Görevin Mevcut Durumunu Kontrol Etme
	isCompleted := t[index].Completed

	//Tamamlanma Durumuna Göre İşlem Yapma

	if !isCompleted {

		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	//Görevin Durumunu Değiştirme
	t[index].Completed = !isCompleted

	//İşlem başarıyla tamamlandığında nil (hata yok) döndürülür.
	return nil
}

/*
Yöntemde , işaretçiyi referanssızlaştırarak ve doğru aralıkta
olduğundan emin olmak için verilen dizini doğrulayarak toggle başlıyoruz.
Bu dizindeki öğe henüz tamamlanmamışsa, geçerli zamanı tamamlanma zamanı olarak işaretliyoruz.
Daha sonra durumu tersine çeviririz.
Eğer (tamamlanmamışsa) (tamamlanmış) olur ve tam tersi.
*/

func (tasks *Tasks) edit(index int, title string) error {
	t := *tasks
	if err := t.validateIndex(index); err != nil {
		return err
	}

	//Görev Başlığını Güncelleme
	t[index].Title = title

	return nil
}

func (tasks *Tasks) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created AT", "Completed AT")

	for index, t := range *tasks {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
