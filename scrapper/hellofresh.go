package scrapper

type HelloFresh struct {
	WebSite WebSite
}

func (HelloFresh) GetRecipeFromWebPage() {

}

// func ScrapHelloFresh() {
// 	// create struct with HelloFresh params to scrap
// 	hellofresh := HelloFresh{
// 		WebSite{
// 			DomainName: "https://www.hellofresh.fr/recipes/",
// 			Path: []string{
// 				"recettes-les-plus-populaires",
// 				"recettes-rapides",
// 				"recettes-faciles",
// 				"recettes-vegetariennes",
// 				"recettes-italiennes",
// 				"recettes-mexicaines",
// 				"recettes-fusion",
// 				"recettes-asiatiques",
// 				"recettes-vietnamiennes",
// 				"recettes-chinoises",
// 				"recettes-japonaises",
// 				"recettes-francaises",
// 				"recettes-indiennes",
// 				"recettes-espagnoles",
// 				"recettes-mediterraneennes",
// 				"recettes-indonesiennes",
// 				"recettes-nordiques",
// 			},
// 			HtmlToFind: "div.web-1nlafhw",
// 			RecipeUrl: ChildAttr{
// 				HtmlPath: "div.biCnbF>a",
// 				Attr:     "href",
// 			},
// 		},
// 	}
// }
