package database

import (
	"errors"
	"strings"

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
func GetMealByName(n string) (Meal, error) {
	var m Meal
	err := db.Collection(mealCollection).FindOne(bson.M{"name": n}, &m)

	if err != nil && !strings.Contains(err.Error(), "not found") {
		return m, err
	}
	return m, nil
}

//AddMeal is for signing up a new meal
func AddMeal(m *Meal) (Meal, error) {
	existent, _ := GetMealByName(m.Name)

	if existent.Name != "" {
		return existent, errors.New(m.Name + " already exists")
	}

	err := db.Collection(mealCollection).Save(m)
	if err != nil {
		return *m, err
	}

	return GetMealByName(m.Name)
}

//UpdateMeal is for patching an existing meal
func UpdateMeal(m *Meal) error {
	return db.Collection(mealCollection).Save(m)
}

//DeleteMeal is for removing an existing meal
func DeleteMeal(n string) error {
	m, _ := GetMealByName(n)

	if m.Name == "" {
		return errors.New(n + " does not exists")
	}

	return db.Collection(mealCollection).DeleteDocument(&m)
}
