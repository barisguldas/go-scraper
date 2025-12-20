# Go Web Scraper 

Bu proje, Go (Golang) ve `chromedp` kütüphanesi kullanılarak geliştirilmiş, komut satırı üzerinden çalışan bir web scraping (veri kazıma) aracıdır. Siber Tehdit İstihbaratı (CTI) süreçlerinde temel keşif ve veri toplama amacıyla tasarlanmıştır.

##  Özellikler

- **HTML Çekme:** Hedef sitenin kaynak kodunu (HTML) `output.txt` olarak kaydeder.
- **Ekran Görüntüsü:** Sitenin tam boy ekran görüntüsünü `screenshot.png` olarak alır.
- **Link Analizi:** Sayfadaki tüm URL'leri ayıklayıp `links.txt` dosyasına listeler.

## 🛠️ Gereksinimler

- [Go](https://go.dev/dl/) (1.20 veya üzeri)
- Google Chrome veya Chromium tarayıcı

##  Kurulum

Projeyi klonlayın ve gerekli bağımlılıkları indirin:

```bash
git clone [https://github.com/KULLANICI_ADIN/go-scraper.git](https://github.com/KULLANICI_ADIN/go-scraper.git)
cd go-scraper
go mod tidy
