package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
)

type conStatus struct {
	status  int
	minutes int
}

//path like a.b.c or a.b[i].c
func getNode(json *simplejson.Json, path string) *simplejson.Json {

	pathArr := strings.Split(path, ".")
	data := json
	length := 0
	index := 0
	for _, node := range pathArr {
		length = len(node)
		if length > 3 && strings.EqualFold(string(node[length-1]), "]") && strings.EqualFold(string(node[length-3]), "[") {
			index, _ = strconv.Atoi(string(node[length-2]))
			arrName := node[:length-3]
			data = data.Get(arrName).GetIndex(index)
		} else {
			data = data.Get(node)
		}
	}
	return data

}

//path like a.b.c or a.b[i].c
func setValue(json *simplejson.Json, path string, value interface{}) {

	pathArr := strings.Split(path, ".")

	nodeName := path
	parentNode := json

	if len(pathArr) != 1 {
		parentPath := path[:strings.LastIndex(path, ".")]
		nodeName = path[strings.LastIndex(path, ".")+1:]
		parentNode = getNode(json, parentPath)
	}

	parentNode.Set(nodeName, value)
}

func getValue(json *simplejson.Json, path string) *simplejson.Json {

	pathArr := strings.Split(path, ".")

	nodeName := path
	parentNode := json

	if len(pathArr) != 1 {
		parentPath := path[:strings.LastIndex(path, ".")]
		nodeName = path[strings.LastIndex(path, ".")+1:]
		parentNode = getNode(json, parentPath)
	}
	return parentNode.Get(nodeName)
}
func optionalDataGenerator(origin *simplejson.Json, path string, formula *simplejson.Json) {
	var value string
	optionList, err := formula.Get("list").StringArray()
	if err != nil {
		//to-do : handle error
	}
	n := rand.Int() % len(optionList)
	value = optionList[n]
	setValue(origin, path, value)
}

func iterativeDataGenerator(origin *simplejson.Json, path string, formula *simplejson.Json, count int) {
	var value string
	optionList, err := formula.Get("list").StringArray()
	if err != nil {
		//to-do : handle error
	}
	n := count % len(optionList)
	value = optionList[n]
	setValue(origin, path, value)
}

func rangeDataGenerator(origin *simplejson.Json, path string, formula *simplejson.Json) {
	dataType, _ := formula.Get("type").String()
	item1, _ := formula.Get("max").String()
	item2, _ := formula.Get("min").String()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch dataType {
	case "int", "integer":
		maxValue, _ := strconv.Atoi(item1)
		minValue, _ := strconv.Atoi(item2)
		list := r.Intn(maxValue-minValue) + minValue
		value := strconv.Itoa(list)
		setValue(origin, path, value)

	case "double", "float":
		maxValue, _ := strconv.ParseFloat(item1, 64)
		minValue, _ := strconv.ParseFloat(item2, 64)
		value := strconv.FormatFloat((maxValue-minValue)*r.Float64()+minValue, 'f', 10, 64)
		setValue(origin, path, value)

	default:
		maxValue, _ := strconv.ParseFloat(item1, 64)
		minValue, _ := strconv.ParseFloat(item2, 64)
		value := strconv.FormatFloat((maxValue-minValue)*r.Float64()+minValue, 'f', 10, 64)
		setValue(origin, path, value)
	}
}

func sequentialDataGenerator(origin *simplejson.Json, path string, formula *simplejson.Json, counter int) {
	dataType, _ := formula.Get("type").String()
	btStr, _ := formula.Get("initial").String()
	stepStr, _ := formula.Get("step").String()

	switch dataType {
	case "int", "integer":
		bt, _ := strconv.Atoi(btStr)
		step, _ := strconv.Atoi(stepStr)

		value := strconv.Itoa(bt + step*counter)
		setValue(origin, path, value)

	case "double", "float":
		bt, _ := strconv.ParseFloat(btStr, 64)
		step, _ := strconv.ParseFloat(stepStr, 64)

		value := strconv.FormatFloat(bt+step*float64(counter), 'f', 10, 64)
		setValue(origin, path, value)

	default:
		bt, _ := strconv.ParseFloat(btStr, 64)
		step, _ := strconv.ParseFloat(stepStr, 64)

		value := strconv.FormatFloat(bt+step*float64(counter), 'f', 10, 64)
		setValue(origin, path, value)
	}

}

func sequentialStringDataGenerator(origin *simplejson.Json, path string, formula *simplejson.Json, value int) int {
	except, _ := formula.Get("except").Array()

	for _, s := range except {
		source, _ := strconv.Atoi(s.(string))
		if source == value {
			value++
		}
	}
	sequentialString := ""
	leastLengthString, _ := formula.Get("leastLength").String()
	leastLength, _ := strconv.Atoi(leastLengthString)

	for i := len(strconv.Itoa(value)); i < leastLength; i++ {
		sequentialString += "0"
	}
	sequentialString += strconv.Itoa(value)

	pattern, _ := formula.Get("initial").String()
	// replace all * in pattern with index
	srp := strings.NewReplacer("*", sequentialString)
	targetString := srp.Replace(pattern)

	setValue(origin, path, targetString)

	value++

	return value
}

func initSequentialStringDataGenerator(origin *simplejson.Json, path string, formula *simplejson.Json) int {
	startStr, _ := formula.Get("start").String()
	start, _ := strconv.Atoi(startStr)
	return sequentialStringDataGenerator(origin, path, formula, start)
}

func arrayDataGenerator(origin *simplejson.Json, path string, formula *simplejson.Json) {
	startStr, _ := formula.Get("start").String()
	countStr, _ := formula.Get("count").String()
	start, _ := strconv.Atoi(startStr)
	count, _ := strconv.Atoi(countStr)
	pattern, _ := formula.Get("initial").String()
	leastLengthString, _ := formula.Get("leastLength").String()
	leastLength, _ := strconv.Atoi(leastLengthString)
	except, _ := formula.Get("except").Array()

	values := make([]string, 0, count)

	for c, v := 0, start; c < count; c, v = c+1, v+1 {
		for _, s := range except {
			source, _ := strconv.Atoi(s.(string))
			if v == source {
				v++
			}
		}

		sequentialString := ""

		for j := len(strconv.Itoa(v)); j < leastLength; j++ {
			sequentialString += "0"
		}
		sequentialString += strconv.Itoa(v)

		// replace all * in pattern with index
		srp := strings.NewReplacer("*", sequentialString)
		vStr := srp.Replace(pattern)
		values = append(values, vStr)
	}
	setValue(origin, path, values)
}

func randomSwitchDataGenerator(origin *simplejson.Json, formula *simplejson.Json, status conStatus, ct string, minute int) conStatus {
	change := false
	status.minutes++

	if status.status == -1 {
		change = true
	} else {
		condition := formula.Get("condition").GetIndex(status.status)
		conditionType := condition.Get("type").MustString()
		threshholdString := condition.Get("threshhold").MustString()
		switch conditionType {
		case "random":
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			threshhold, _ := strconv.ParseFloat(threshholdString, 64)
			temp := r.Float64()
			if temp < threshhold {
				change = true
			}
		case "count":
			threshhold, _ := strconv.Atoi(threshholdString)
			if (status.minutes > threshhold) && (threshhold != 0) {
				change = true
			}
		}
	}
	var subFormulaList []interface{}
	if change {
		totalConditionString := formula.Get("totalCondition").MustString()
		totalCondition, _ := strconv.Atoi(totalConditionString)
		status.status = (status.status + 1) % totalCondition
		status.minutes = 1
		subFormulaList = formula.Get("subFormula").GetIndex(status.status).Get("init").MustArray()

	} else {
		formula.Get("subFormula").GetIndex(status.status)
		subFormulaList = formula.Get("subFormula").GetIndex(status.status).Get("normal").MustArray()

	}

	for _, f := range subFormulaList {
		subFormula, _ := f.(map[string]interface{})
		path := subFormula["path"].(string)
		value := subFormula["value"].(string)
		switch subFormula["generator"] {
		case "string":
			setValue(origin, path, value)
		case "step":
			valueType := subFormula["type"].(string)

			switch valueType {
			case "int", "integer":
				base, _ := strconv.Atoi(getValue(origin, path).MustString())
				step, _ := strconv.Atoi(value)

				value := strconv.Itoa(base + step)
				setValue(origin, path, value)

			case "double", "float":
				base, _ := strconv.ParseFloat(getValue(origin, path).MustString(), 64)
				step, _ := strconv.ParseFloat(value, 64)

				value := strconv.FormatFloat(base+step, 'f', 10, 64)
				setValue(origin, path, value)
			}
		case "timestamp":
			setValue(origin, path, ct)
		}

	}

	return status
}

func timestampDataGenerator(origin *simplejson.Json, path string, formula *simplejson.Json, timestamp string) {
	setValue(origin, path, timestamp)
}

func fixedDataGenerator(origin *simplejson.Json, path string, formula *simplejson.Json) {
	value, _ := formula.Get("value").String()
	setValue(origin, path, value)
}
