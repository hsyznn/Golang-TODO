package main

func main() {

	tasks := Tasks{}
	/*
		Bu satır, boş bir Tasks yapısı oluşturur.
		Tasks yapısı, muhtemelen bir görev listesini temsil eden ve daha önce tanımlanmış bir yapıdır.
	*/
	storage := NewStorage[Tasks]("Tasks.json")

	/*
		/*
			Storage yapısını kullanarak bir depolama nesnesi oluşturur:
			NewStorage[Tasks]: Jenerik Storage yapısını Tasks tipi için özelleştirir
			"Tasks.json": Görevlerin kaydedileceği JSON dosyasının adını belirtir
	*/

	storage.Load(&tasks)

	/*
		Bu satır, "Tasks.json" dosyasından önceden kaydedilmiş görevleri yükler:
		&tasks: tasks değişkeninin bellek adresini (pointer) Load fonksiyonuna geçirir
		Load fonksiyonu, dosyadan okunan JSON verisini bu adresteki tasks nesnesine dönüştürür
		Eğer dosya yoksa veya boşsa, tasks değişkeni boş kalır
	*/
	cmdFlags := NewCmdFlags()
	/*
		Bu satır,NewCmdFlags() fonksiyonunu çağırarak komut satırı bayraklarını işler:
		Kullanıcının girdiği komut satırı argümanlarını ayrıştırır
		Hangi işlemin yapılacağını belirleyen bir CmdFlags nesnesi oluşturur
	*/
	cmdFlags.Execute(&tasks)
	/*
		Bu satır, kullanıcının istediği komutu yürütür:
		&tasks: tasks değişkeninin bellek adresini Execute fonksiyonuna geçirir
		Execute fonksiyonu, komut satırı bayraklarına göre uygun işlemi gerçekleştirir:
		Görev ekleme
		Görev silme
		Görev düzenleme
		Görev durumunu değiştirme
		Görevleri listeleme
	*/
	storage.Save(tasks)
	/*
		Bu satır, yapılan değişiklikleri "Tasks.json" dosyasına kaydeder:
		tasks: Güncellenmiş görev listesi
		Save fonksiyonu, görev listesini JSON formatına dönüştürüp dosyaya yazar
	*/
}
