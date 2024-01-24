package scrapper

import "go_scrapper/logger"

var HelloFreshDishTitlePath = "h1.ceYciq"
var HelloFreshDishDescPath = "span.fivcnB>p"

type HelloFresh struct {
	WebSite WebSite
}

func initHelloFresh() HelloFresh {
	return HelloFresh{
		WebSite{
			DomainName: "https://www.hellofresh.fr/recipes/",
			Path: []string{
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
			HtmlToFind: "div.web-1nlafhw",
			RecipeUrl: ChildAttr{
				HtmlPath: "div.biCnbF>a",
				Attr:     "href",
			},
		},
	}
}

func ScrapHelloFresh() {
	// create struct with HelloFresh params to scrap
	hellofresh := initHelloFresh()
	recipeUrls, err := hellofresh.WebSite.Scrap()
	if err != nil {
		logger.Log(err.Error(), "")
		panic(err)
	}

	// prepare struct to be queued and enqueue
	for _, url := range recipeUrls {
		toQueue := ScrapQueued{
			Url: url,
		}
		queue = append(queue, toQueue)
	}
}
