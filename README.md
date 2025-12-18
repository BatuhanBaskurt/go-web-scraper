# Go Web Scraper

Go dilinde yazılmış hafif ve hızlı bir web scraping aracı. Herhangi bir web sitesinin içeriğini analiz etmek, kaynak kodunu indirmek ve görsel dokümantasyon oluşturmak için geliştirilmiştir.

## Ne yapıyor bu?

Basit bir web scraper. Verdiğin URL'e gidip:

- 📄 Sayfanın HTML'ini indiriyor
- 📸 Tam sayfa ekran görüntüsü alıyor (1920x1080)
- 🔗 İçindeki bütün linkleri buluyor

## Kurulum

```bash
git clone https://github.com/BatuhanBaskurt/go-web-scraper.git
cd go-web-scraper
go mod tidy
go run scraper.go
```

**Gerekli şeyler:** Go 1.16+ ve Chrome/Chromium

## Nasıl kullanılıyor?

Çalıştır, domain gir, enter bas.

```bash
$ go run scraper.go
Lütfen bir domain girin örn: https://example.com : github.com
```

Program şunları oluşturacak:

- `example.com.html` - Sayfanın kaynak kodu
- `example.com_SS.png` - Ekran görüntüsü
- Terminal'de link listesi

## Örnek çıktı

```bash
Girilen domain: https://github.com , çekme işlemi başlatılıyor...
Girilen domainin bilgileri başarıyla çekildi. bulundugunuz dizini kontrol edin
Fotograf başarıyla kaydedildi: github.com_SS.png

domain içerisinde bulunan linkler:
 1. https://github.com/features
 2. https://github.com/enterprise
 3. https://github.com/pricing
```

## Teknik detaylar

**Kullanılan paketler:**

- [Colly](https://github.com/gocolly/colly) - Scraping için
- [ChromeDP](https://github.com/chromedp/chromedp) - Screenshot için

## Hatalar

Bir şeyler ters giderse (yanlış domain, timeout, DNS hatası vs.) program sana söylüyor. Sıkıntı yok.

## 

**Batuhan Başkurt**  
[@BatuhanBaskurt](https://github.com/BatuhanBaskurt)

---

*Not: Bu araç eğitim amaçlı. Scraping yaparken sitelerin kurallarına dikkat et.*