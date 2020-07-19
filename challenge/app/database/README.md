## package database
This directory contains the responsible files for database interactions

### testing
It is necessary to have running some local mongodb instance (with no authentication) at `localhost:27017/`, the database `testing` will be automatically created once you run:
```bash
go test
```

### current schemas
- DataCenter
```
{
    "name": "datacenter-1",
    "metadata": {
        "monitoring": {
            "enabled": "true"
        },
        "limits": {
            "cpu": {
                "enabled": "false",
                "value": "300m"
            }
        }
    }
}
```

- Meal
```
{
    "name": "burger-nutrition",
    "metadata": {
        "calories": 230,
        "fats": {
            "saturated-fat": "0g",
            "trans-fat": "1g"
        },
        "carbohydrates": {
            "dietary-fiber": "4g",
            "sugars": "1g"
        },
        "allergens": {
            "nuts": "false",
            "seafood": "false",
            "eggs": "true"
        }
    }
}

```
