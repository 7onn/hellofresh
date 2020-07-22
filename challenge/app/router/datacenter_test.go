package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"echotom.dev/hellofresh/database"
)

const dataCenterName = "testCenter"

var dcMock = &database.DataCenter{
	Name: dataCenterName,
	Metadata: database.DataCenterMetadata{
		Monitoring: database.Monitoring{
			Enabled: true,
		},
		Limits: database.Limits{
			CPU: database.CPU{
				Enabled: true,
				Value:   "200m",
			},
		},
	},
}

func TestAddDataCenterHandler(t *testing.T) {
	payload, _ := json.Marshal(dcMock)
	req, err := http.NewRequest("POST", "/configs", strings.NewReader(string(payload)))
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

func TestGetAllDataCentersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/configs", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	mockMuxRouter.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Expected status code 200 but got %v instead", status)
	}

	database.DeleteDataCenter(dcMock.Name)
}
