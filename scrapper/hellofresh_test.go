package scrapper_test

import (
	"go_scrapper/scrapper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrapHelloFresh(t *testing.T) {

	toTest := scrapper.HelloFresh{
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

	recipeUrls, err := toTest.WebSite.Scrap()
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, len(recipeUrls), 32)
}
