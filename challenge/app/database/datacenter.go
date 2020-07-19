package database

import (
	"errors"

	"github.com/globalsign/mgo/bson"
)

const dataCenterCollection = "datacenters"

//GetAllDataCenters is for filterlessly return available data centers
func GetAllDataCenters() []DataCenter {
	find := db.Collection(dataCenterCollection).Find(bson.M{})

	var dcs []DataCenter
	dc := &DataCenter{}
	for find.Next(dc) {
		dcs = append(dcs, *dc)
	}

	return dcs
}

//GetDataCenterByName is for filtering one data center by its name field
func GetDataCenterByName(n string) DataCenter {
	var dc DataCenter
	db.Collection(dataCenterCollection).FindOne(bson.M{"name": n}, &dc)
	return dc
}

//AddDataCenter is for signing up a new data center
func AddDataCenter(dc *DataCenter) (DataCenter, error) {
	existent := GetDataCenterByName(dc.Name)
	if existent.Name != "" {
		return existent, errors.New(dc.Name + " already exists")
	}

	db.Collection(dataCenterCollection).Save(dc)
	return GetDataCenterByName(dc.Name), nil
}

//UpdateDataCenter is for patching an existing data center
func UpdateDataCenter(dc *DataCenter) *DataCenter {
	db.Collection(dataCenterCollection).Save(dc)
	return dc
}

//DeleteDataCenter is for removing an existing data center
func DeleteDataCenter(dc *DataCenter) bool {
	db.Collection(dataCenterCollection).DeleteDocument(dc)
	return true
}
