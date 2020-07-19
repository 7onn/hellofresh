package database

import (
	"github.com/go-bongo/bongo"
)

//CPU for Limits type
type CPU struct {
	Enabled bool
	Value   string
}

//Limits for DataCenterMetadata type
type Limits struct {
	CPU CPU
}

//Monitoring for DataCenterMetadata type
type Monitoring struct {
	Enabled bool
}

//DataCenterMetadata for DataCenter type
type DataCenterMetadata struct {
	Monitoring Monitoring
	Limits     Limits
}

//DataCenter ... what you think it is, linter?
type DataCenter struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string
	Metadata           DataCenterMetadata
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
	Calories      int
	Fats          map[string]string
	Carbohydrates map[string]string
	Allergens     map[string]bool
}

//Meal ... what you think it is, linter?
type Meal struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string
	Metadata           MealMetadata
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
