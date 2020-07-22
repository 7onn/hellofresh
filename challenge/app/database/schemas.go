package database

import (
	"github.com/go-bongo/bongo"
)

//CPU for Limits type
type CPU struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

//Limits for DataCenterMetadata type
type Limits struct {
	CPU CPU `json:"cpu"`
}

//Monitoring for DataCenterMetadata type
type Monitoring struct {
	Enabled bool `json:"enabled"`
}

//DataCenterMetadata for DataCenter type
type DataCenterMetadata struct {
	Monitoring Monitoring `json:"monitoring"`
	Limits     Limits     `json:"limits"`
}

//DataCenter ... what you think it is, linter?
type DataCenter struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string             `json:"name"`
	Metadata           DataCenterMetadata `json:"metadata"`
}

// {
//     "name": "datacenter-1",
//     "metadata": {
//       "monitoring": {
//         "enabled": "true"
//       },
//       "limits": {
//         "cpu": {
//           "enabled": "false",
//           "value": "300m"
//         }
//       }
//     }
//   }

//MealMetadata for Meal type
type MealMetadata struct {
	Calories      int               `json:"calories"`
	Fats          map[string]string `json:"fats"`
	Carbohydrates map[string]string `json:"carbohydrates"`
	Allergens     map[string]bool   `json:"allergens"`
}

//Meal ... what you think it is, linter?
type Meal struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string       `json:"name"`
	Metadata           MealMetadata `json:"metadata"`
}

// {
//     "name": "burger-nutrition",
//     "metadata": {
//       "calories": 230,
//       "fats": {
//         "saturated-fat": "0g",
//         "trans-fat": "1g"
//       },
//       "carbohydrates": {
//           "dietary-fiber": "4g",
//           "sugars": "1g"
//       },
//       "allergens": {
//         "nuts": "false",
//         "seafood": "false",
//         "eggs": "true"
//       }
//     }
//   }
