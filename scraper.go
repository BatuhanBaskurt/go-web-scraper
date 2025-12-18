package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
	"net/http"
	"github.com/gocolly/colly" 
	"github.com/chromedp/chromedp"
)

var hreflinkleri []string
// temizledosyaadi fonksiyonu verilen domainden başını ve özel karakterleri silerek kaydedilicek olan dosyanın adını oluşturur
func temizledosyaadi(url string) string {
	url = strings.ReplaceAll(url, "https://", "")
	url = strings.ReplaceAll(url, "http://", "")
	url = strings.ReplaceAll(url, "/", "_")
	url = strings.ReplaceAll(url, "?", "_")
	url = strings.ReplaceAll(url, ":", "")
	return url
}


// ekrangoruntusual fonksiyonu önce girilen domaine gider ve oranın ss'ini alır. daha sonra import ettigimiz os pakedi ile bulunan dizine fotografı kaydeder
func ekrangoruntusual(url string, dosyaAdi string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var goruntuVerisi []byte 

	hata := chromedp.Run(ctx,
		chromedp.EmulateViewport(1920, 1080), 
		chromedp.Navigate(url),
		chromedp.Sleep(2*time.Second),
		chromedp.FullScreenshot(&goruntuVerisi, 90),
	)
	
	if hata != nil {
		fmt.Printf("\nekran görüntüsü alınamadı. oluşan hata : %v\n", hata)
		return
	}

	hata = os.WriteFile(dosyaAdi, goruntuVerisi, 0644)
	if hata != nil {
		fmt.Printf("\n ekran görüntüsü kaydedilemedi, oluşan hata : %v\n", hata)
		return
	}
	fmt.Printf("Fotograf başarıyla kaydedildi: %s\n", dosyaAdi)
}


func main() {

	fmt.Print("Lütfen bir domain girin örn: https://example.com : ")
	var domain string
	fmt.Scanln(&domain)

	// HasPrefix ile girilen domainin başlangıcı http ile başlamıyorsa otomatik olarak https eklemir
	if !strings.HasPrefix(domain, "http") {
		domain = "https://" + domain
	}
	
	dosyaAdi := temizledosyaadi(domain) + ".html"
	goruntuDosyaAdi := temizledosyaadi(domain) + "_SS.png"

	fmt.Printf("Girilen domain: %s , çekme işlemi başlatılıyor...\n", domain)

	toplayici := colly.NewCollector()
	var responseHTML []byte 


	toplayici.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		linkler := e.Request.AbsoluteURL(link)
		hreflinkleri = append(hreflinkleri, linkler)
	})


	toplayici.OnError(func(yanit *colly.Response, err error) {
		fmt.Print("\nVeri çekme başarısız oldu. ")
		
		
		if yanit != nil && yanit.StatusCode >= http.StatusBadRequest {
			fmt.Printf("Sunucu Hata Kodu. Kod: %d.\n", yanit.StatusCode)
			return
		}

		// hata mesajının içerdigi kelimelere göre farklı hata geri dönüşü yapıyor
		hataMesaji := err.Error()
		if strings.Contains(hataMesaji, "no such host") {
			fmt.Printf("Geçersiz Domain veya DNS Çözümleme Hatası.\n")
			return
		} else if strings.Contains(hataMesaji, "timeout") {
			fmt.Printf("Bağlantı Zaman Aşımı Hatası.\n")
			return
		} else if strings.Contains(hataMesaji, "connection refused") {
			fmt.Printf("Bağlantı Reddedildi Hatası.\n")
			return
		} else {
			fmt.Printf("Genel Ağ Hatası. Detay: %v\n", err)
			return
		}
	})


	toplayici.OnResponse(func(yanit *colly.Response) {
		responseHTML = yanit.Body 
	})


	toplayici.Visit(domain)
	toplayici.Wait() 

	// program burada çektigi verinin uzunlugunun 0 dan büyük olup olmadıgını kontrol ediyor. eğer  büyükse dosyaya yazıyor. degilse hata mesajı veriyor
	if len(responseHTML) > 0 {
		hata := os.WriteFile(dosyaAdi, responseHTML, 0644) 
		if hata != nil {
			fmt.Printf("\nbir hata oluştu ve dosya içerisine yazılamadı, oluşan hata :  %v\n", hata)
			return
		} else {
			fmt.Printf("Girilen domainin bilgileri başarıyla çekildi. bulundugunuz dizini kontrol edin\n")
			ekrangoruntusual(domain, goruntuDosyaAdi)
		}
	} else {
		fmt.Println("bir hata oluştu ve dosya içerisine yazılamadı. ")
		return
	}

	// hreflinkleri adlı domain listesinden bulunan domainleri çekiyor
	if len(hreflinkleri) > 0 {
		fmt.Printf("\ndomain içerisinde bulunan linkler:\n ")
		for i, link := range hreflinkleri {
			fmt.Printf(" %d. %s\n", i+1, link)
		}
	} else {
		fmt.Println(" domain içerisinde herhangi bir link bulunamadı.")
	}
}