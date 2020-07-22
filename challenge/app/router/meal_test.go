package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"echotom.dev/hellofresh/database"
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

const mealName = "testMeal"
const updtMealName = "testMeal2"

var mealMock = &database.Meal{
	Name: mealName,
	Metadata: database.MealMetadata{
		Calories:      230,
		Fats:          mealFats,
		Carbohydrates: mealCarbohydrates,
		Allergens:     mealAllergens,
	},
}

func TestAddMealHandler(t *testing.T) {
	payload, _ := json.Marshal(mealMock)

	req, err := http.NewRequest("PUT", "/configs/"+mealName, strings.NewReader(string(payload)))
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	mockMuxRouter.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Expected status code 200 but got %v instead", status)
	}
}

func TestGetMealByNameHandlerHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/configs/"+mealName, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	mockMuxRouter.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Expected status code 200 but got %v instead", status)
	}
}

func TestUpdateMealHandler(t *testing.T) {
	mealMock.Name = updtMealName
	payload, _ := json.Marshal(mealMock)
	req, err := http.NewRequest("PATCH", "/configs/"+mealName, strings.NewReader(string(payload)))
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	mockMuxRouter.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Expected status code 200 but got %v instead", status)
	}
}

func TestDeleteMealHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/configs/"+updtMealName, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	mockMuxRouter.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Expected status code 200 but got %v instead", status)
	}
}
