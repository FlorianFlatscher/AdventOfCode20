package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"strings"
)

type Day21 struct{}

type Allergen string

type Ingredient string

type Food struct {
	ingredients []Ingredient
	allergens   []Allergen
}

func newFood(line string) *Food {
	inputLine := strings.Split(line, " (contains ")

	var allergens []Allergen
	var ingredients []Ingredient

	ingredientsInput := strings.Split(inputLine[0], " ")
	for _, i := range ingredientsInput {
		ingredients = append(ingredients, Ingredient(i))
	}
	if len(inputLine[1]) > 0 {
		inputLine[1] = strings.TrimSuffix(inputLine[1], ")")
		allergensInput := strings.Split(inputLine[1], ", ")
		for _, a := range allergensInput {
			allergens = append(allergens, Allergen(a))
		}
	}

	return &Food{
		ingredients: ingredients,
		allergens:   allergens,
	}
}

func (d Day21) Calc() string {
	lines := strings.Split(input.ReadInputFile(21), constants.LineSeparator)
	var foods []Food

	for _, line := range lines {
		foods = append(foods, *newFood(line))
	}

	d.countIngredientsWithNoAllergy(foods)
	return fmt.Sprintf("1: %v\n2: %v\n", nil, nil)
}

func (d Day21) countIngredientsWithNoAllergy(foods []Food) {
	allIngredients := make(map[Ingredient]bool)
	allAllergens := make(map[Allergen]bool)

	doesNotHaveAllergen := make(map[Ingredient]map[Allergen]bool)

	for _, f := range foods {
		for _, a := range f.allergens {
			allAllergens[a] = true
		}
		for _, i := range f.ingredients {
			allIngredients[i] = true
		}
	}

	for i := range allIngredients {
		doesNotHaveAllergen[i] = make(map[Allergen]bool)
	}

	//for _, f := range foods {
	//	for _, i := range f.ingredients {
	//		for a := range allAllergens {
	//			doesNotHaveAllergen[i][a] = true
	//		}
	//	}
	//}
	//
	//for _, f := range foods {
	//	for _, i := range f.ingredients {
	//		for _, a := range f.allergens {
	//			delete(doesNotHaveAllergen[i], a)
	//		}
	//	}
	//}

	for true {
		somethingChanged := false

		for _, f := range foods {
			for i := 0; i < len(f.ingredients); i++ {
				ing := f.ingredients[i]
				if len(doesNotHaveAllergen[ing]) == len(allAllergens)-1 {
					continue
				}

				for _, a := range f.allergens {
					unique := true

					for _, checkI := range f.ingredients {
						if checkI == ing {
							continue
						}

						if _, ok := doesNotHaveAllergen[checkI][a]; ok {
							unique = false
							break
						}
					}

					if unique {
						somethingChanged = true

						for _, checkI := range f.ingredients {
							if checkI == ing {
								continue
							}
							doesNotHaveAllergen[checkI][a] = true
						}
						for all := range allAllergens {
							doesNotHaveAllergen[ing][all] = true
						}
						delete(doesNotHaveAllergen[ing], a)
						break
					}
				}
			}
			fmt.Println(doesNotHaveAllergen)
		}

		if !somethingChanged {
			break
		}
	}

	for ingredient, allergens := range doesNotHaveAllergen {
		if 0 == len(allergens) {
			fmt.Println(ingredient)
		}
	}
}
