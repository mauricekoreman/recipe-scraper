package server

import (
	"fmt"
	"os"
	"testing"
)

// TODO: finish tests
func TestFindJSONLD(t *testing.T) {
	file, err := os.Open("mocks/mockdata.html")
	if err != nil {
		t.Fatalf("Failed to open mock HTML file: %v", err)
	}
	defer file.Close()

	contents, err := findJSONLD(file)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(contents)
}

func TestGetRecipeJSON(t *testing.T) {
	// place mock json in
	mockJSON := `{
	"@context": "https://schema.org",
	"@graph": [
		{
			"@type": "Organization",
			"@id": "https://miljuschka.nl/#organization",
			"name": "miljuschka.nl",
			"url": "https://miljuschka.nl",
			"logo": {
				"@type": "ImageObject",
				"@id": "https://miljuschka.nl/#logo",
				"url": "https://miljuschka.nl/wp-content/uploads/2021/02/Miljuschka.svg",
				"contentUrl": "https://miljuschka.nl/wp-content/uploads/2021/02/Miljuschka.svg",
				"caption": "miljuschka.nl",
				"inLanguage": "nl-NL",
				"width": "545",
				"height": "109"
			}
		},
		{
			"@type": "Recipe",
			"name": "Het beste bananenbrood recept ooit",
			"author": { "@type": "Person", "name": "Miljuschka Witzenhausen" },
			"description": "Het allerlekkerste bananenbrood dat je ooit hebt geproefd!",
			"datePublished": "2020-03-21T17:24:55+00:00",
			"image": [
				"https://miljuschka.nl/wp-content/uploads/2020/03/Bananenbrood-3.jpg",
				"https://miljuschka.nl/wp-content/uploads/2020/03/Bananenbrood-3-500x500.jpg",
				"https://miljuschka.nl/wp-content/uploads/2020/03/Bananenbrood-3-500x375.jpg",
				"https://miljuschka.nl/wp-content/uploads/2020/03/Bananenbrood-3-480x270.jpg"
			],
			"recipeYield": ["10", "10 portie(s)"],
			"prepTime": "PT10M",
			"cookTime": "PT60M",
			"totalTime": "PT70M",
			"recipeIngredient": [
				"125 gr boter (gesmolten)",
				"200 gr suiker",
				"2  grote eieren",
				"60 ml melk",
				"1 tl vanille-extract",
				"3  rijpe bananen",
				"250 gr bloem",
				"1 tl baking soda",
				"snufje zout",
				"handvol gehakte chocola of noten",
				"1  extra banaan voor de bovenkant (optioneel)"
			],
			"recipeInstructions": [
				{
					"@type": "HowToStep",
					"text": "Verwarm de oven voor op 160\u02daC voor een heteluchtoven. Bekleed een cakeblik van 25 cm met bakpapier of vet het goed in.",
					"name": "Verwarm de oven voor op 160\u02daC voor een heteluchtoven. Bekleed een cakeblik van 25 cm met bakpapier of vet het goed in.",
					"url": "https://miljuschka.nl/het-beste-bananenbrood-recept-ooit/#wprm-recipe-39443-step-0-0"
				},
				{
					"@type": "HowToStep",
					"text": "Meng als eerste de suiker met de boter en klop deze goed door elkaar. Voeg de eieren een voor een toe en meng even door.",
					"name": "Meng als eerste de suiker met de boter en klop deze goed door elkaar. Voeg de eieren een voor een toe en meng even door.",
					"url": "https://miljuschka.nl/het-beste-bananenbrood-recept-ooit/#wprm-recipe-39443-step-0-1"
				},
				{
					"@type": "HowToStep",
					"text": "Roer nu de melk en het vanille extract erdoor. Doe de bloem, de baking soda en het zout erbij en roer het om. Je hoeft dit niet heel lang te kloppen. Als het goed gemengd is, is het voldoende.",
					"name": "Roer nu de melk en het vanille extract erdoor. Doe de bloem, de baking soda en het zout erbij en roer het om. Je hoeft dit niet heel lang te kloppen. Als het goed gemengd is, is het voldoende.",
					"url": "https://miljuschka.nl/het-beste-bananenbrood-recept-ooit/#wprm-recipe-39443-step-0-2"
				},
				{
					"@type": "HowToStep",
					"text": "Prak de bananen met een vork of pureer ze in een blender. Voeg toe aan je beslag en roer dit weer om tot het goed verdeeld is. Als laatste voeg je de chocola en de noten toe.",
					"name": "Prak de bananen met een vork of pureer ze in een blender. Voeg toe aan je beslag en roer dit weer om tot het goed verdeeld is. Als laatste voeg je de chocola en de noten toe.",
					"url": "https://miljuschka.nl/het-beste-bananenbrood-recept-ooit/#wprm-recipe-39443-step-0-3"
				},
				{
					"@type": "HowToStep",
					"text": "Stort het beslag in de klaarstaande vorm en bak dit in de voorverwarmde oven voor ongeveer 50-65 minuten of tot een houten prikker er droog uitkomt.",
					"name": "Stort het beslag in de klaarstaande vorm en bak dit in de voorverwarmde oven voor ongeveer 50-65 minuten of tot een houten prikker er droog uitkomt.",
					"url": "https://miljuschka.nl/het-beste-bananenbrood-recept-ooit/#wprm-recipe-39443-step-0-4"
				},
				{
					"@type": "HowToStep",
					"text": "(Wil je de extra banaan gebruiken, snij deze dan in de lengte doormidden en druk met de bolle kanten in het beslag.)",
					"name": "(Wil je de extra banaan gebruiken, snij deze dan in de lengte doormidden en druk met de bolle kanten in het beslag.)",
					"url": "https://miljuschka.nl/het-beste-bananenbrood-recept-ooit/#wprm-recipe-39443-step-0-5"
				},
				{
					"@type": "HowToStep",
					"text": "Laat het 10 minuten afkoelen in het blik voordat je het uitstort en verder af laat koelen. Serveer eventueel met een lekkere bol vanille ijs! Eet smakelijk!",
					"name": "Laat het 10 minuten afkoelen in het blik voordat je het uitstort en verder af laat koelen. Serveer eventueel met een lekkere bol vanille ijs! Eet smakelijk!",
					"url": "https://miljuschka.nl/het-beste-bananenbrood-recept-ooit/#wprm-recipe-39443-step-0-6"
				}
			],
			"aggregateRating": {
				"@type": "AggregateRating",
				"ratingValue": "4.38",
				"ratingCount": "2310",
				"reviewCount": "8"
			},
			"review": [
				{
					"@type": "Review",
					"reviewRating": { "@type": "Rating", "ratingValue": "5" },
					"reviewBody": "Al zo vaak dit recept opgezocht om een heerlijk bananenbrood te gaan bakken. Echt super lekker!",
					"author": { "@type": "Person", "name": "Inge" },
					"datePublished": "2025-05-20"
				},
				{
					"@type": "Review",
					"reviewRating": { "@type": "Rating", "ratingValue": "5" },
					"reviewBody": "Ook heerlijk als muffin!\nOmdat ik mn cakeblik niet kon vinden maar wel muffinblik, heb ik er 12 muffins van gemaakt. Uiteindelijk is dit ook prettig als je minder geduld hebt, want je kan hierdoor de  baktijd halveren. \nHet is heel lekker luchtig met een klein krokantje in de korst, maar totaal niet droog! En de zoetheid zit precies goed naar mijn mening. Al met al, makkelijk en leuk recept!",
					"author": { "@type": "Person", "name": "Aaron" },
					"datePublished": "2025-03-23"
				},
				{
					"@type": "Review",
					"reviewRating": { "@type": "Rating", "ratingValue": "5" },
					"reviewBody": "Heerlijk!!",
					"author": { "@type": "Person", "name": "Ronald" },
					"datePublished": "2025-02-09"
				},
				{
					"@type": "Review",
					"reviewRating": { "@type": "Rating", "ratingValue": "5" },
					"reviewBody": "De titel zegt het al en dit is echt niet overdreven: het beste ooit. Al is het meer een cake dan een brood, maar een kniesoor die daarop let ;) ik heb hem al meerdere malen voor diverse mensen gemaakt, en nog nooit is er ook maar een kruimel van blijven liggen. Wel een beetje zoet, dus je zou ook prima net wat minder suiker kunnen doen als je dat lekkerder vindt. Maar ik vind het echt het beste bananenbroodrecept ooit!",
					"author": { "@type": "Person", "name": "Sabrina" },
					"datePublished": "2025-01-17"
				},
				{
					"@type": "Review",
					"reviewRating": { "@type": "Rating", "ratingValue": "5" },
					"reviewBody": "Ik heb er wat geraspte kokos, stukjes pure chocolade en gehakte pecannoten doorheen gedaan en het was heerlijk!",
					"author": { "@type": "Person", "name": "Jane" },
					"datePublished": "2025-01-10"
				},
				{
					"@type": "Review",
					"reviewRating": { "@type": "Rating", "ratingValue": "5" },
					"reviewBody": "Absoluut het beste bananenbrood. Al zo vaak gemaakt. En je kunt er eindeloos mee vari\u00ebren. Nu een brood met een caramel swirl er door. Ben benieuwd.",
					"author": { "@type": "Person", "name": "Maggy" },
					"datePublished": "2024-08-20"
				},
				{
					"@type": "Review",
					"reviewRating": { "@type": "Rating", "ratingValue": "5" },
					"reviewBody": "Erg lekkere cake! Ik heb er alleen veel minder suiker in gedaan. 70 gram.\nEn nog zoet genoeg!",
					"author": { "@type": "Person", "name": "Joana" },
					"datePublished": "2024-07-18"
				},
				{
					"@type": "Review",
					"reviewRating": { "@type": "Rating", "ratingValue": "5" },
					"reviewBody": "Heerlijk!!!",
					"author": { "@type": "Person", "name": "JvA" },
					"datePublished": "2024-05-31"
				}
			],
			"recipeCategory": ["Koek &amp; Gebak"],
			"recipeCuisine": ["Hollandse keuken"],
			"keywords": "banaan",
			"nutrition": {
				"@type": "NutritionInformation",
				"calories": "325 kcal",
				"carbohydrateContent": "47 g",
				"proteinContent": "5 g",
				"fatContent": "13 g",
				"saturatedFatContent": "8 g",
				"transFatContent": "1 g",
				"cholesterolContent": "64 mg",
				"sodiumContent": "312 mg",
				"fiberContent": "2 g",
				"sugarContent": "24 g",
				"servingSize": "1 portie"
			},
			"@id": "https://miljuschka.nl/het-beste-bananenbrood-recept-ooit/#recipe",
			"mainEntityOfPage": "https://miljuschka.nl/het-beste-bananenbrood-recept-ooit/#webpage",
			"isPartOf": {
				"@id": "https://miljuschka.nl/het-beste-bananenbrood-recept-ooit/#richSnippet"
			}
		},
		{
			"@type": "WebSite",
			"@id": "https://miljuschka.nl/#website",
			"url": "https://miljuschka.nl",
			"name": "miljuschka.nl",
			"alternateName": "Miljuschka",
			"publisher": { "@id": "https://miljuschka.nl/#organization" },
			"inLanguage": "nl-NL"
		}]}`

	foundRecipe, err := getRecipeJSON(mockJSON)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(foundRecipe)
}
