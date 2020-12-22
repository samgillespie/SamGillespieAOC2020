package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(21)
	SolveRecipes(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

type Recipe struct {
	Ingredients []string
	Allergens   []string
}

type PotentialAllergens struct {
	ingredient string
}

func (r Recipe) ContainsIngredient(ingredient string) bool {
	for _, i := range r.Ingredients {
		if i == ingredient {
			return true
		}
	}
	return false
}

func (r Recipe) ContainsAllergen(allergen string) bool {
	for _, i := range r.Allergens {
		if i == allergen {
			return true
		}
	}
	return false
}

func ConvertEntryToRecipe(entry string) Recipe {
	split := strings.Split(entry, " ")
	ingredients := []string{}
	allergens := []string{}
	isIngredients := true
	for _, ingredient := range split {
		if strings.Contains(ingredient, "(contains") {
			isIngredients = false
			continue
		} else if isIngredients == false {
			allergen := strings.Replace(ingredient, ")", "", -1)
			allergen = strings.Replace(allergen, ",", "", -1)
			allergens = append(allergens, allergen)
		} else {
			ingredients = append(ingredients, ingredient)
		}
	}
	return Recipe{
		Ingredients: ingredients,
		Allergens:   allergens,
	}
}

// Return all elements common in a and b
func Intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

// Return only elements in the left
func LeftOnly(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			delete(m, item)
		}
	}
	for i := range m {
		c = append(c, i)
	}
	return
}

func ParseInput(data []string) ([]Recipe, []string, []string) {
	recipes := []Recipe{}
	allergens := make(map[string]bool)
	ingredients := make(map[string]bool)
	for _, row := range data {
		recipes = append(recipes, ConvertEntryToRecipe(row))
		for _, allergen := range recipes[len(recipes)-1].Allergens {
			allergens[allergen] = true
		}
		for _, ingredient := range recipes[len(recipes)-1].Ingredients {
			ingredients[ingredient] = true
		}
	}
	allergensList := []string{}
	for key := range allergens {
		allergensList = append(allergensList, key)
	}

	ingredientsList := []string{}
	for key := range ingredients {
		ingredientsList = append(ingredientsList, key)
	}

	return recipes, allergensList, ingredientsList
}

func FindNoAllergens(ingredients []string, allergens []string, recipeList []Recipe) []string {
	allergensToIngredients := make(map[string][]string)
	for _, allergen := range allergens {
		var validIngredients []string = nil
		for _, recipe := range recipeList {
			if recipe.ContainsAllergen(allergen) == false {
				continue
			} else if validIngredients == nil {
				validIngredients = recipe.Ingredients
			} else {
				validIngredients = Intersection(validIngredients, recipe.Ingredients)
			}
		}
		allergensToIngredients[allergen] = validIngredients
	}

	for _, list := range allergensToIngredients {
		ingredients = LeftOnly(ingredients, list)
	}
	return ingredients
}

func RemoveFromRecipes(recipeList []Recipe, ingredients []string) (r []Recipe) {
	for _, recipe := range recipeList {
		recipe.Ingredients = LeftOnly(recipe.Ingredients, ingredients)
		r = append(r, recipe)
	}
	return
}

func CalculateSolution(recipes []Recipe, allergens []string) map[string]string {
	solution := make(map[string]string)

	for len(solution) < len(allergens) {
		candidates := make(map[string][]string)
		for _, recipe := range recipes {
			bailout := false
			for _, allergen := range recipe.Allergens {
				if _, ok := solution[allergen]; ok {
					continue
				}

				if len(recipe.Ingredients) == 1 {
					solution[allergen] = recipe.Ingredients[0]
					recipes = RemoveFromRecipes(recipes, []string{solution[allergen]})
					bailout = true
					break
				}

				_, present := candidates[allergen]
				if present == false {
					candidates[allergen] = recipe.Ingredients
				} else {
					candidates[allergen] = Intersection(candidates[allergen], recipe.Ingredients)
					if len(candidates[allergen]) == 1 {
						solution[allergen] = candidates[allergen][0]
						recipes = RemoveFromRecipes(recipes, []string{solution[allergen]})
						bailout = true
						break
					}
				}
			}
			if bailout == true {
				break
			}
		}
	}
	return solution
}

func SolveRecipes(data []string) {
	recipes, allergens, ingredients := ParseInput(data)
	noAllergens := FindNoAllergens(ingredients, allergens, recipes)
	totalScore := 0
	for _, ingredient := range noAllergens {
		for _, recipe := range recipes {
			if recipe.ContainsIngredient(ingredient) {
				totalScore++
			}
		}
	}
	fmt.Println("Day 21 Part 1 Solution")
	fmt.Println(totalScore)
	recipes = RemoveFromRecipes(recipes, noAllergens)
	solution := CalculateSolution(recipes, allergens)

	//Alphabetical
	keys := make([]string, 0, len(solution))
	for key := range solution {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	submission := ""
	for _, key := range keys {
		submission += solution[key] + ","
	}
	fmt.Println("Day 21 Part 2 Solution")
	fmt.Println(submission[0 : len(submission)-1]) // Remove trailing comma
}
