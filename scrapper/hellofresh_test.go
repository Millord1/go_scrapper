package scrapper_test

import (
	"go_scrapper/scrapper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initHelloFresh() scrapper.HelloFresh {
	return scrapper.HelloFresh{
		scrapper.WebSite{
			DomainName: "https://www.hellofresh.fr/recipes/",
			Path: []string{
				"recettes-les-plus-populaires",
			},
			HtmlToFind: "div.web-1nlafhw",
			RecipeUrl: scrapper.ChildAttr{
				HtmlPath: "div.biCnbF>a",
				Attr:     "href",
			},
		},
	}
}

func TestScrapHelloFresh(t *testing.T) {

	toTest := initHelloFresh()

	recipeUrls, err := toTest.WebSite.Scrap()
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, len(recipeUrls), 32)
}

func TestGetDishFromWeb(t *testing.T) {
	toTest := scrapper.ScrapQueued{
		Url: "https://www.hellofresh.fr/recipes/rigatoni-and-sauce-onctueuse-au-bleu-6544a9e93357e0637a82bad1",
		// PathToDishName: "h1.ceYciq",
		PathToDishData: "div.gpdTEW",
	}

	scrapper.GetRecipeFromWebPage(toTest)
	// t.Logf("%v", dish)
}
