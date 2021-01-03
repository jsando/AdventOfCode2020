package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	foods := readInput(os.Args[1])
	fmt.Printf("Part 1: %d\n", foods.part1()) // 2302
	fmt.Printf("Part 2: %s\n", foods.part2()) // smfz,vhkj,qzlmr,tvdvzd,lcb,lrqqqsg,dfzqlk,shp
}

func readInput(path string) *Foods {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	foods := newFoods()
	for _, line := range lines {
		i := strings.Index(line, "(")
		contains := line[i+len("contains ") : len(line)-1]
		line = line[:i]
		foods.addRecipe(strings.Split(strings.TrimSpace(line), " "), strings.Split(strings.TrimSpace(contains), ","))
	}
	for changed := true; changed; {
		changed = false
		for k := range foods.allergens {
			if len(foods.allergens[k]) == 1 {
				// must be the ingredient containing the allergen, remove ingrent from other allergen lists
				ingredient := ""
				for ingredient = range foods.allergens[k] {
					break
				}
				for k2 := range foods.allergens {
					if k != k2 && foods.allergens[k2][ingredient] {
						delete(foods.allergens[k2], ingredient)
						changed = true
					}
				}
			}
		}
	}
	foods.ingredientsWithAllergens = map[string]bool{}
	for _, allergenIngredients := range foods.allergens {
		for badIngredient := range allergenIngredients {
			foods.ingredientsWithAllergens[badIngredient] = true
		}
	}
	return foods
}

// Foods is a list of all ingredients and recipes.
type Foods struct {
	allergens                map[string]map[string]bool // allergen name -> set of possible ingredients
	recipes                  []*Recipe
	ingredientsWithAllergens map[string]bool
}

func newFoods() *Foods {
	return &Foods{
		recipes:   []*Recipe{},
		allergens: map[string]map[string]bool{},
	}
}

func (f *Foods) addAllergen(name string, ingredients map[string]bool) {
	allergen := f.allergens[name]
	if allergen == nil {
		f.allergens[name] = map[string]bool{}
		for k := range ingredients {
			f.allergens[name][k] = true
		}
	} else {
		// retain common subset
		retained := map[string]bool{}
		for k := range allergen {
			if ingredients[k] {
				retained[k] = true
			}
		}
		f.allergens[name] = retained
	}
}

func (f *Foods) addRecipe(ingredients []string, allergens []string) {
	recipe := &Recipe{
		ingredients: map[string]bool{},
		allergens:   map[string]bool{},
	}
	for _, ingredient := range ingredients {
		ingredient = strings.TrimSpace(ingredient)
		recipe.ingredients[ingredient] = true
	}
	for _, allergen := range allergens {
		allergen = strings.TrimSpace(allergen)
		recipe.allergens[allergen] = true
		f.addAllergen(allergen, recipe.ingredients)
	}
	f.recipes = append(f.recipes, recipe)
}

// Ingredient is a single ingredient and its allergens.
type Ingredient struct {
	name string
}

// Recipe is a list of ingredients and all known allergens.
type Recipe struct {
	ingredients map[string]bool
	allergens   map[string]bool
}

func (f *Foods) part1() int {
	count := 0
	for _, recipe := range f.recipes {
		for ingredient := range recipe.ingredients {
			if !f.ingredientsWithAllergens[ingredient] {
				count++
			}
		}
	}
	return count
}

func (f *Foods) part2() string {
	alist := []string{}
	for k := range f.allergens {
		alist = append(alist, k)
	}
	sort.Strings(alist)
	result := ""
	for i, k := range alist {
		if i != 0 {
			result += ","
		}
		aIngredient := ""
		for aIngredient = range f.allergens[k] {
			break
		}
		result += aIngredient
	}
	return result
}
