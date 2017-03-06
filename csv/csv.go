package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type FieldMismatch struct {
	expected, found int
}

type UnsupportedType struct {
	utype string
}

type CSVStruct struct {
	Name string
	Age  int
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
			f.SetString(record[i])
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

func CSVFileToMap(filename string) []CSVStruct {
	var csvstruct CSVStruct
	topoConfig := []CSVStruct{}
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error when read file CSVFileToMap: ", err)
		return make([]CSVStruct, 0, 0)
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
