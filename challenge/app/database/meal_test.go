package database

import (
	"testing"
)

var mealFats = func() map[string]string {
	f := make(map[string]string)
	f["saturated-fat"] = "0g"
	f["trans-fat"] = "1g"
	return f
}()

var mealCarbohydrates = func() map[string]string {
	c := make(map[string]string)
	c["dietary-fiber"] = "4g"
	c["sugars"] = "1g"
	return c
}()

var mealAllergens = func() map[string]bool {
	a := make(map[string]bool)
	a["nuts"] = false
	a["seafood"] = false
	a["eggs"] = true
	return a
}()

const mealName = "test meal"

var mealMock = &Meal{
	Name: mealName,
	Metadata: MealMetadata{
		Calories:      230,
		Fats:          mealFats,
		Carbohydrates: mealCarbohydrates,
		Allergens:     mealAllergens,
	},
}

func TestAddMeal(t *testing.T) {
	m, err := AddMeal(mealMock)
	if err != nil {
		t.Errorf("Expected to sign up meal but got error instead\n%v", err)
	}
	if m.GetId() == "" {
		t.Errorf("Expected to sign up meal but got no registered name")
	}
}
func TestGetAllMeals(t *testing.T) {
	ms := GetAllMeals()
	if len(ms) == 0 {
		t.Errorf("Expected some meal but got none")
	}
}
func TestGetMealByName(t *testing.T) {
	m, err := GetMealByName(mealName)

	if err != nil {
		t.Errorf(err.Error())
	}

	if m.GetId() == "" {
		t.Errorf("Expected 'test meal' to return but got empty record")
	}
}
func TestUpdateMeal(t *testing.T) {
	mealMock.Name = "tested meal"
	err := UpdateMeal(mealMock)

	if err != nil {
		t.Errorf(err.Error())
	}

	if mealMock.Name != "tested meal" {
		t.Errorf("Expected 'tested meal' to return but got out-of-date record")
	}
}
func TestDeleteMeal(t *testing.T) {
	err := DeleteMeal(mealMock.Name)
	if err != nil {
		t.Errorf("Expected truthy deletion but got false")
	}
}
