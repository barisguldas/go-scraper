package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	var url string
	url = os.Args[1]
	fmt.Printf("Hedef URL: %s\n", url)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var htmlContent string
	var screenshot []byte
	var links []string

	fmt.Println("Siteye baglaniliyor.")
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),                   //Siteyi actik
		chromedp.Sleep(2*time.Second),            // 2 saniye bekleme komutu. Sayfa tam yüklensin diye ekledim.
		chromedp.OuterHTML("html", &htmlContent), //htmlContent degiskenini pointer kullanarak degistirdim. GO'da bir fonksiyonun bir degiskeni degistirmesini istiyorsam pointer ile bellege ulasmam gerekiyor.
		chromedp.FullScreenshot(&screenshot, 90),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(n => n.href)`, &links), //javascript kullanarak site içerisindeki URL'leri cektim.
	)

	if err != nil {
		log.Fatalf("Hata oluştu: %v", err)
	}

	// Variable'lari kaydettigimiz kisim

	err = os.WriteFile("output.txt", []byte(htmlContent), 0644)
	if err != nil {
		log.Fatalf("HTML yazilamadi: %v", err)
	}

	err = os.WriteFile("screenshot.png", screenshot, 0644)
	if err != nil {
		log.Fatalf("Resim yazilamadi: %v", err)
	}

	linksText := strings.Join(links, "\n")
	err = os.WriteFile("links.txt", []byte(linksText), 0644)
	if err != nil {
		log.Fatalf("Link dosyasi yazilamadi: %v", err)
	}

	fmt.Printf("HTML Kaydedildi: output.txt\n")
	fmt.Println("Resim Kaydedildi: screenshot.png")
	fmt.Printf("Linkler Kaydedildi: links.txt (%d adet link bulundu)\n", len(links))

}
