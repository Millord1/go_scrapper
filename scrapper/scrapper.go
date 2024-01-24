package scrapper

import (
	"fmt"
	"go_scrapper/entity"
	"go_scrapper/logger"
	"log"
	"time"

	"github.com/gocolly/colly"
)

var queue []ScrapQueued

type Recipe interface {
	GetDishFromWeb() entity.Dish
}

type ScrapQueued struct {
	Url               string
	PathToDishData    string
	PathToSteps       string
	PathToIngredients string
}

type WebSite struct {
	DomainName string
	Path       []string
	// first HTML element to find when scrapping
	HtmlToFind string
	RecipeUrl  ChildAttr
}

type ChildAttr struct {
	// path to the HTML child element containing recipe url
	HtmlPath string
	// HTML attribute with url (usually "href")
	Attr string
}

func (toScrap WebSite) Scrap() ([]string, error) {

	// return urls: one for each recipe found
	var recipesUrls []string

	for _, url := range toScrap.Path {
		// loop on categories url (which is containing recipes urls)
		recipesUrl, err := scrapCategory(&toScrap, &url)
		if err != nil {
			logger.Log(err.Error(), "")
		}
		recipesUrls = append(recipesUrls, recipesUrl...)
		time.Sleep(1 * time.Second)
	}

	return recipesUrls, nil
}

func scrapCategory(toScrap *WebSite, url *string) ([]string, error) {
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

	var urls []string
	c.OnHTML(toScrap.HtmlToFind, func(e *colly.HTMLElement) {
		attr := e.ChildAttr(toScrap.RecipeUrl.HtmlPath, toScrap.RecipeUrl.Attr)
		// get one url for each recipe
		if attr[0:4] == "http" {
			urls = append(urls, attr)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	err := c.Visit(*&toScrap.DomainName + *url)
	return urls, err
}

func GetRecipeFromWebPage(toScrap ScrapQueued) entity.Dish {
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

	var dish entity.Dish
	c.OnHTML(toScrap.PathToDishData, func(e *colly.HTMLElement) {

		// get Dish title
		titleTxt := e.ChildText(HelloFreshDishTitlePath)
		if len(titleTxt) > 1 {
			dish.Name = titleTxt
		}
		// get Dish description
		descTxt := e.ChildText(HelloFreshDishDescPath)
		if len(descTxt) > 1 {
			dish.Description = descTxt
		}
	})

	// c.OnHTML()

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	err := c.Visit(toScrap.Url)
	if err != nil {
		logger.Log(err.Error()+" for url "+toScrap.Url, "")
	}

	return entity.Dish{}
}
