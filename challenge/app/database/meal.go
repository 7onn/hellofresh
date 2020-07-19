package database

import (
	"errors"

	"github.com/globalsign/mgo/bson"
)

const mealCollection = "meals"

//GetAllMeals is for filterlessly return available meals
func GetAllMeals() []Meal {
	find := db.Collection(mealCollection).Find(bson.M{})

	var ms []Meal
	m := &Meal{}
	for find.Next(m) {
		ms = append(ms, *m)
	}

	return ms
}

//GetMealByName is for filtering one meal by its name field
func GetMealByName(n string) Meal {
	var m Meal
	db.Collection(mealCollection).FindOne(bson.M{"name": n}, &m)
	return m
}

//AddMeal is for signing up a new meal
func AddMeal(m *Meal) (Meal, error) {
	existent := GetMealByName(m.Name)
	if existent.Name != "" {
		return existent, errors.New(m.Name + " already exists")
	}

	db.Collection(mealCollection).Save(m)
	return GetMealByName(m.Name), nil
}

//UpdateMeal is for patching an existing meal
func UpdateMeal(m *Meal) *Meal {
	db.Collection(mealCollection).Save(m)
	return m
}

//DeleteMeal is for removing an existing meal
func DeleteMeal(n string) bool {
	m := GetMealByName(n)
	if m.Name == "" {
		return false
	}
	db.Collection(mealCollection).DeleteDocument(&m)
	return true
}
