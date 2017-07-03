package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

var DataGeneratorType = map[string][]string{
	"Concurrent": {
		"option_data",
		"range_data",
		"fixed_data",
		"sequential_string_data",
		"timestamp_data",
		"iterative_data",
		"array_data",
	},
	"Metrics": {
		"option_data",
		"range_data",
		"fixed_data",
		"sequential_data",
		"timestamp_data",
		"random_switch_data",
	},
}

var Platform = map[string]string{
	"linux":      "Linux",
	"mac":        "Mac",
	"win":        "Windows",
	"iOS":        "iOS",
	"htmlaccess": "Htmlaccess",
	"android":    "Android",
	"uwp":        "UWP",
}

type FieldMismatch struct {
	expected, found int
}

type UnsupportedType struct {
	utype string
}

type TopoCSVStruct struct {
	Id        string
	Pnum      string
	Psize     string
	Namespace string
	Object    string
	Type      string
}

type SessionCfgCSVStruct struct {
	Section   string
	Path      string
	Type      string
	Generator string
	Formula   string
}

// exists returns whether the given file or directory exists or not
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func GenerateJSONFile(v interface{}, filePath string, isConfig bool) {
	var genFilePath string
	// Create directories to store JSON files.
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	cfgDir := path.Join(currentDir, "config")
	dataDir := path.Join(currentDir, "data")
	// Create config folder
	if isExists, _ := PathExists(cfgDir); isExists == false {
		os.MkdirAll(cfgDir, 0777)
	}
	// Create data folder
	if isExists, _ := PathExists(dataDir); isExists == false {
		os.MkdirAll(dataDir, 0777)
	}

	if isConfig {
		genFilePath = path.Join(cfgDir, filePath)
	} else {
		genFilePath = path.Join(dataDir, filePath)
	}

	output, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	// Write to JSON file.
	err = ioutil.WriteFile(genFilePath, output, 0664)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}

func CSVUnmarshal(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			// Trick: remove leading and tailing "{" or "}"
			f.SetString(strings.Trim(strings.Replace(record[i], " ", "", -1), "{|}"))
		case "int":
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			return &UnsupportedType{f.Type().String()}
		}
	}
	return nil
}

func (e *FieldMismatch) Error() string {
	return fmt.Sprintf(`CSV line fields mismatch.
Expected %s found %s`, strconv.Itoa(e.expected), strconv.Itoa(e.found))
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.utype
}

func TopoCSVFileToMap(filename string) []TopoCSVStruct {
	var csvstruct TopoCSVStruct
	topoConfig := []TopoCSVStruct{}
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error when read file in TopoCSVFileToMap: ", err)
		return make([]TopoCSVStruct, 0, 0)
	}
	defer fp.Close()

	r := csv.NewReader(fp)
	for {
		err = CSVUnmarshal(r, &csvstruct)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		topoConfig = append(topoConfig, csvstruct)
	}
	return topoConfig
}

func SessionCfgCSVFileToMap(filename string) []SessionCfgCSVStruct {
	var csvstruct SessionCfgCSVStruct
	sessionConfig := []SessionCfgCSVStruct{}
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error when read file in SessionCfgCSVFileToMap: ", err)
		return make([]SessionCfgCSVStruct, 0, 0)
	}
	defer fp.Close()

	r := csv.NewReader(fp)
	for {
		err = CSVUnmarshal(r, &csvstruct)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		sessionConfig = append(sessionConfig, csvstruct)
	}
	return sessionConfig
}
