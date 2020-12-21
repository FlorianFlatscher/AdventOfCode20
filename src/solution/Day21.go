package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"sort"
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

	return fmt.Sprintf("1: %v\n2: %v\n", d.countIngredientsWithNoAllergy(foods), d.generateCanonicalDangerousIngredientList(foods))
}

//192 too low

func (d Day21) countIngredientsWithNoAllergy(foods []Food) int {
	ingredients := make([]Ingredient, 0)
	allergens := make([]Allergen, 0)

	for _, f := range foods {
		for _, a := range f.allergens {
			allergens = append(allergens, a)
		}
		for _, i := range f.ingredients {
			ingredients = append(ingredients, i)
		}
	}

	hasAllergen := make(map[Ingredient]map[Allergen]bool)
	for _, i := range ingredients {
		hasAllergen[i] = make(map[Allergen]bool)
		for _, a := range allergens {
			hasAllergen[i][a] = true
		}
	}

	for _, f := range foods {
		for _, a := range f.allergens {
			for _, i := range ingredients {
				found := false
				for _, foodI := range f.ingredients {
					if foodI == i {
						found = true
						break
					}
				}
				if !found {
					delete(hasAllergen[i], a)
				}
			}
		}
	}

	count := 0
	for i, v := range hasAllergen {
		if len(v) == 0 {
			for _, f := range foods {
				for _, fi := range f.ingredients {
					if i == fi {
						count++
						break
					}
				}
			}
		}
	}

	return count
}

func (d Day21) generateCanonicalDangerousIngredientList(foods []Food) string {
	ingredients := make([]Ingredient, 0)
	allergens := make([]Allergen, 0)

	for _, f := range foods {
		for _, a := range f.allergens {
			allergens = append(allergens, a)
		}
		for _, i := range f.ingredients {
			ingredients = append(ingredients, i)
		}
	}

	hasAllergen := make(map[Ingredient]map[Allergen]bool)

	for _, i := range ingredients {
		hasAllergen[i] = make(map[Allergen]bool)
		for _, a := range allergens {
			hasAllergen[i][a] = true
		}
	}

	for _, f := range foods {
		for _, a := range f.allergens {
			for _, i := range ingredients {
				found := false
				for _, foodI := range f.ingredients {
					if foodI == i {
						found = true
						break
					}
				}
				if !found {
					delete(hasAllergen[i], a)
				}
			}
		}
	}

	count := 0
	for i, v := range hasAllergen {
		if len(v) == 0 {
			for _, f := range foods {
				for _, fi := range f.ingredients {
					if i == fi {
						count++
						break
					}
				}
			}
		}
	}

	finalList := make(map[Ingredient]Allergen)

	for true {
		somethingChanged := false

		for _, f := range foods {
			for _, ing := range f.ingredients {
				if len(hasAllergen[ing]) == 0 {
					continue
				}
				if len(hasAllergen[ing]) == 1 {
					somethingChanged = true
					for a, _ := range hasAllergen[ing] {
						finalList[ing] = a
					}
					for _, allIng := range ingredients {
						for a, _ := range hasAllergen[allIng] {
							if a == finalList[ing] {
								delete(hasAllergen[allIng], a)
							}
						}
					}
				}
			}
		}

		if !somethingChanged {
			break
		}
	}

	fmt.Println(finalList)

	keys := make([]string, 0, len(finalList))
	for key := range finalList {
		keys = append(keys, string(key))
	}
	sort.Slice(keys, func(i1, i2 int) bool {
		return 0 > strings.Compare(string(finalList[Ingredient(keys[i1])]), string(finalList[Ingredient(keys[i2])]))
	})
	//values := make([]string, 0, len(finalList))
	//for _, k := range keys {
	//	values = append(values, string(finalList[Ingredient(k)]))
	//}

	return strings.Join(keys, ",")
}
