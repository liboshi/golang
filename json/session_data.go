package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type DataStruct struct {
	Product  string    `json:"product"`
	TenantId string    `json:"tenantid"`
	DaasVMId string    `json:"daas_vm_id"`
	DataTop  []DataTop `json:"data"`
}

type DataTop struct {
	Type      string    `json:"type"`
	DataInner DataInner `json:"data"`
}

type DataInner struct {
	DataType string `json:"dataType"`
	Client   Client `json:"client"`
	Sample   Sample `json:"sample"`
}

type Client struct {
	Number     string `json:"n"`
	Version    string `json:"version"`
	DeviceType string `json:"deviceType"`
}

type Sample struct {
	Time    string        `json:"time"`
	Metrics []DataMetrics `json:"metrics"`
	S       []S           `json:"s"`
}

type DataMetrics struct {
	N           string `json:"n"`
	DV          string `json:"dv,omitempty"`
	LV          string `json:"lv,omitempty"`
	StringValue string `json:"stringValue,omitempty"`
}

// Section for session
type S struct {
	ID           string `json:"id"`
	SID          string `json:"sid"`
	UN           string `json:"un"`
	UD           string `json:"ud"`
	IsConnnected string `json:"isConnnected"`
	P            []P    `json:"p"`
	Metrics      string `json:"metrics"`
	RemoteApps   string `json:"remoteApps"`
}

type P struct {
	N string `json:"n"`
	V string `json:"v"`
}

func GenerateDataJson(isSession bool, daasVMId, tenantId, tid string) {
	// Read template JSON file
	var templateFile string
	var outputFile string
	if isSession {
		templateFile = "session_data_template.json"
		outputFile = "session_data.json"
	} else {
		templateFile = "session_data_no_session_template.json"
		outputFile = "session_data_nosession.json"
	}

	// Create tenant directory
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	tenPrefix := fmt.Sprintf("Ten%s", tid)
	tenDir := path.Join(currentDir, "data", tenPrefix)
	if isExists, _ := PathExists(tenDir); isExists == false {
		os.MkdirAll(tenDir, 0777)
	}

	fp, e := ioutil.ReadFile(templateFile)
	if e != nil {
		fmt.Println("File error: ", e)
		return
	}
	// Update fields and generate JSON object
	var data DataStruct
	json.Unmarshal(fp, &data)
	data.DaasVMId = daasVMId
	data.TenantId = tenantId

	// Write to file
	outputFile = path.Join(tenPrefix, outputFile)
	GenerateJSONFile(&data, outputFile, false)
}
