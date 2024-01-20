package scrapper

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
)

var baseAddr string = "https://www.hellofresh.fr/recipes"

func Scrap() {
	urls := getAddressCategories()
	for _, url := range urls {
		scrapCategory(&url)
		time.Sleep(1 * time.Second)
	}
}

func scrapCategory(categoryUrl *string) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("div.web-1nlafhw", func(e *colly.HTMLElement) {
		attr := e.ChildAttr("div.biCnbF>a", "href")
		if attr[0:4] == "http" {
			fmt.Println(attr)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	err := c.Visit("https://www.hellofresh.fr/recipes/" + *categoryUrl + "?page=30")
	if err != nil {
		panic(err)
	}
}

func getAddressCategories() [17]string {
	var urls [17]string

	urls[0] = "recettes-les-plus-populaires"
	urls[1] = "recettes-rapides"
	urls[2] = "recettes-faciles"
	urls[3] = "recettes-vegetariennes"
	urls[4] = "recettes-italiennes"
	urls[5] = "recettes-mexicaines"
	urls[6] = "recettes-fusion"
	urls[7] = "recettes-asiatiques"
	urls[8] = "recettes-vietnamiennes"
	urls[9] = "recettes-chinoises"
	urls[10] = "recettes-japonaises"
	urls[11] = "recettes-francaises"
	urls[12] = "recettes-indiennes"
	urls[13] = "recettes-espagnoles"
	urls[14] = "recettes-mediterraneennes"
	urls[15] = "recettes-indonesiennes"
	urls[16] = "recettes-nordiques"

	return urls
}
