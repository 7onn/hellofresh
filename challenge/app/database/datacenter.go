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
func GetDataCenterByName(n string) (DataCenter, error) {
	var dc DataCenter
	err := db.Collection(dataCenterCollection).FindOne(bson.M{"name": n}, &dc)

	if err != nil {
		return DataCenter{}, err
	}

	return dc, nil
}

//AddDataCenter is for signing up a new data center
func AddDataCenter(dc *DataCenter) (DataCenter, error) {
	existent, _ := GetDataCenterByName(dc.Name)
	if existent.Name != "" {
		return existent, errors.New(dc.Name + " already exists")
	}

	err := db.Collection(dataCenterCollection).Save(dc)

	if err != nil {
		return *dc, err
	}

	return GetDataCenterByName(dc.Name)
}

//UpdateDataCenter is for patching an existing data center
func UpdateDataCenter(dc *DataCenter) (*DataCenter, error) {
	err := db.Collection(dataCenterCollection).Save(dc)

	if err != nil {
		return dc, err
	}

	return dc, nil
}

//DeleteDataCenter is for removing an existing data center
func DeleteDataCenter(n string) error {
	dc, _ := GetDataCenterByName(n)

	if dc.Name == "" {
		return errors.New(n + " does not exists")
	}

	return db.Collection(dataCenterCollection).DeleteDocument(&dc)
}
