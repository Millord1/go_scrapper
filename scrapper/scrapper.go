package scrapper

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
)

type webSite struct {
	domainName string
	path       []string
}

func ScrapHelloFresh() {
	helloFresh := webSite{
		domainName: "https://www.hellofresh.fr/recipes",
		path: []string{
			"recettes-les-plus-populaires",
			"recettes-rapides",
			"recettes-faciles",
			"recettes-vegetariennes",
			"recettes-italiennes",
			"recettes-mexicaines",
			"recettes-fusion",
			"recettes-asiatiques",
			"recettes-vietnamiennes",
			"recettes-chinoises",
			"recettes-japonaises",
			"recettes-francaises",
			"recettes-indiennes",
			"recettes-espagnoles",
			"recettes-mediterraneennes",
			"recettes-indonesiennes",
			"recettes-nordiques",
		},
	}

	scrap(helloFresh)
}

func scrap(toScrap webSite) {
	for _, url := range toScrap.path {
		err := scrapCategory(&url, &toScrap.domainName)
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}

func scrapCategory(categoryUrl *string, domain *string) error {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
		// return err
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("div.web-1nlafhw", func(e *colly.HTMLElement) {
		attr := e.ChildAttr("div.biCnbF>a", "href")
		if attr[0:4] == "http" {
			// do something here
			fmt.Println(attr)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	err := c.Visit("https://www.hellofresh.fr/recipes/" + *categoryUrl + "?page=30")
	return err
}
