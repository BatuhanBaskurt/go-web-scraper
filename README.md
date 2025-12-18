🕸️ Go Scraper & SS Tool
Verdiğin herhangi bir web sitesinin linklerini toplayan, HTML'ini indiren ve sayfanın tam boy ekran görüntüsünü alan basit bir Go aracı.

🛠️ Ne işe yarıyor?
Link Ayıklama: Sayfadaki bütün href linklerini bulur ve ekrana basar.

HTML Kayıt: Sayfanın kaynak kodunu .html olarak kaydeder.

Screenshot: Headless Chrome kullanarak sayfanın 1920x1080 çözünürlüğünde görselini alır (.png).

🚀 Kullanım
Bağımlılıkları yüklemek için:

Bash

go mod tidy
Çalıştırmak için:

Bash

go run main.go
Çalıştırdıktan sonra terminale URL'i (örn: https://google.com) yazman yeterli.

📦 Gereksinimler
Go

Chrome/Chromium (Screenshot alabilmesi için sistemde yüklü olmalı)

📝 Bilgi
Dosya isimlerini URL'den otomatik temizleyip oluşturur.

Bağlantı hataları, timeout veya yanlış domain gibi durumlarda terminale hata detayını basar.

Sayfa tam yüklensin diye ekran görüntüsü almadan önce 2 saniye bekler.