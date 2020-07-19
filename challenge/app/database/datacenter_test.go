package database

import (
	"testing"
)

const dataCenterName = "test center"

var dcMock = &DataCenter{
	Name: dataCenterName,
	Metadata: DataCenterMetadata{
		Monitoring: Monitoring{
			Enabled: true,
		},
		Limits: Limits{
			CPU: CPU{
				Enabled: true,
				Value:   "200m",
			},
		},
	},
}

func TestAddDataCenter(t *testing.T) {
	dc, err := AddDataCenter(dcMock)
	if err != nil {
		t.Errorf("Expected to sign up data center but got error instead\n%v", err)
	}
	if dc.GetId() == "" {
		t.Errorf("Expected to sign up data center but got no registered name")
	}
}
func TestGetAllDataCenters(t *testing.T) {
	dcs := GetAllDataCenters()
	if len(dcs) == 0 {
		t.Errorf("Expected some data center but got none")
	}
}
func TestGetDataCenterByName(t *testing.T) {
	dc := GetDataCenterByName(dataCenterName)
	if dc.GetId() == "" {
		t.Errorf("Expected 'test center' to return but got empty record")
	}
}
func TestUpdateDataCenter(t *testing.T) {
	dcMock.Name = "tested center"
	dc := UpdateDataCenter(dcMock)
	if dc.Name != "tested center" {
		t.Errorf("Expected 'tested center' to return but got out-of-date record")
	}
}
func TestDeleteDataCenter(t *testing.T) {
	d := DeleteDataCenter(dcMock)
	if !d {
		t.Errorf("Expected truthy deletion but got false")
	}
}
