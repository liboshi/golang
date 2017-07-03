package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	FILENAME    = "filename"
	PARAMETER   = "parameter"
	CONCURRENCY = "concurrency"
	METRICS     = "metrics"
	POOLSIZE    = "poolsize"
	INITPATTER  = "TenPool"
)

type myError struct {
	errMsg string
}

type ConfigStruct struct {
	Parameters    Parameter     `json:"parameters"`
	Concurrencies []Concurrency `json:"concurrency"`
	Metrics       []Metrics     `json:"metrics"`
	PartitionKey  PartitionKey  `json:"partitionKey"`
}

type Parameter struct {
	PU  string `json:"u"`
	PR  string `json:"r"`
	PC  string `json:"c"`
	PI  string `json:"i"`
	PB  string `json:"b"`
	PM  string `json:"m"`
	PK  string `json:"k"`
	PP  string `json:"p"`
	PS  string `json:"s"`
	PD  string `json:"d"`
	PBT string `json:"bt"`
}

type Concurrency struct {
	Path      string  `json:"path"`
	Type      string  `json:"type,omitempty"`
	Generator string  `json:"generator"`
	Formula   Formula `json:"formula"`
}

type Metrics struct {
	Path      string  `json:"path"`
	Type      string  `json:"type"`
	Generator string  `json:"generator"`
	Formula   Formula `json:"formula"`
}

type PartitionKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Item `Except` can be omitted.
type Formula struct {
	Start       string   `json:"start,omitempty"`
	Initial     string   `json:"initial,omitempty"`
	LeastLength string   `json:"leastLength,omitempty"`
	Step        string   `json:"step,omitempty"`
	Type        string   `json:"type,omitempty"`
	Max         string   `json:"max,omitempty"`
	Min         string   `json:"min,omitempty"`
	Value       string   `json:"value,omitempty"`
	Count       string   `json:"count,omitempty"`
	Except      []string `json:"except,omitempty"`
	List        []string `json:"list,omitempty"`
}

// This function is used to extend concurrency list
func (configStruct *ConfigStruct) AddConcurrency(item Concurrency) []Concurrency {
	configStruct.Concurrencies = append(configStruct.Concurrencies, item)
	return configStruct.Concurrencies
}

// This function is used to extend metrics list
func (configStruct *ConfigStruct) AddMetrics(item Metrics) []Metrics {
	configStruct.Metrics = append(configStruct.Metrics, item)
	return configStruct.Metrics
}

// This function is used to generate Paramemters
func GenerateParameter(paras []string) Parameter {
	if len(paras) != 11 {
		return Parameter{}
	}
	retParameters := Parameter{
		PU:  paras[0],
		PR:  paras[1],
		PC:  paras[2],
		PI:  paras[3],
		PB:  paras[4],
		PM:  paras[5],
		PK:  paras[6],
		PP:  paras[7],
		PS:  paras[8],
		PD:  paras[9],
		PBT: paras[10],
	}
	return retParameters
}

// This function is used to generate Concurrency
func GenerateConcurrency(cPath, cType, cGenerator string, cFValues []string) (Concurrency, error) {
	cFormula, err := GenerateConcurrencyFormula(cGenerator, cFValues)
	retConcurrency := Concurrency{
		Path:      cPath,
		Type:      cType,
		Generator: cGenerator,
		Formula:   cFormula,
	}
	return retConcurrency, err
}

// This function is used to generate Metrics
func GenerateMetrics(mPath, mType, mGenerator string, mFValues []string) (Metrics, error) {
	mFormula, err := GenerateMetricsFormula(mGenerator, mFValues)
	retMetrics := Metrics{
		Path:      mPath,
		Type:      mType,
		Generator: mGenerator,
		Formula:   mFormula,
	}
	return retMetrics, err
}

// This function is raise error for function GenerateFormula
func (e *myError) Error() string {
	return fmt.Sprintf("Error %s", e.errMsg)
}

// This function is used to generate Formula
func GenerateConcurrencyFormula(formulaType string, fValues []string) (Formula, error) {
	switch formulaType {
	// option_data
	case DataGeneratorType["Concurrent"][0]:
		platformMap := fValues
		return GenerateOptionData(platformMap), nil
	// range_data
	case DataGeneratorType["Concurrent"][1]:
		if len(fValues) != 3 {
			return Formula{}, &myError{"The number of parameters is not correct"}
		}
		return GenerateRangeData(fValues[0], fValues[1], fValues[2]), nil
	// fixed_data
	case DataGeneratorType["Concurrent"][2]:
		if len(fValues) != 1 {
			return Formula{}, &myError{"The number of parameters is not correct"}
		}
		return GenerateFixedData(fValues[0]), nil
	// sequential_string_data
	case DataGeneratorType["Concurrent"][3]:
		if len(fValues) < 3 {
			return Formula{}, &myError{"The number of parameters is not correct"}
		}
		exceptList := fValues[3:]
		return GenerateSequentialStringData(fValues[0], fValues[1], fValues[2], exceptList), nil
	// timestamp_data
	case DataGeneratorType["Concurrent"][4]:
		return GenerateTimestampData(), nil
	// iterative_data
	case DataGeneratorType["Concurrent"][5]:
		return GenerateIterativeData(fValues), nil
	// array_data
	case DataGeneratorType["Concurrent"][6]:
		if len(fValues) <= 4 {
			return Formula{}, &myError{"The number of parameters is not correct"}
		}
		exceptList := fValues[3:]
		return GenerateArrayData(fValues[0], fValues[1], fValues[2], fValues[3], exceptList), nil
	default:
		fmt.Println("It is UnsupportedType...")
		return Formula{}, &myError{"It is UnsupportedType..."}
	}
}

// This function is used to generate Formula
func GenerateMetricsFormula(formulaType string, fValues []string) (Formula, error) {
	switch formulaType {
	// option_data
	case DataGeneratorType["Metrics"][0]:
		platformMap := fValues
		return GenerateOptionData(platformMap), nil
	// range_data
	case DataGeneratorType["Metrics"][1]:
		if len(fValues) != 3 {
			return Formula{}, &myError{"The number of parameters is not correct"}
		}
		return GenerateRangeData(fValues[0], fValues[1], fValues[2]), nil
	// fixed_data
	case DataGeneratorType["Metrics"][2]:
		if len(fValues) != 1 {
			return Formula{}, &myError{"The number of parameters is not correct"}
		}
		return GenerateFixedData(fValues[0]), nil
	// sequential_data
	case DataGeneratorType["Metrics"][3]:
		if len(fValues) != 3 {
			return Formula{}, &myError{"The number of parameters is not correct"}
		}
		return GenerateSequentialData(fValues[0], fValues[1], fValues[2]), nil
	// timestamp_data
	case DataGeneratorType["Metrics"][4]:
		return GenerateTimestampData(), nil
	// random_switch_data
	case DataGeneratorType["Metrics"][5]:
		return GenerateRandomSwitchData(), nil
	default:
		fmt.Println("It is UnsupportedType...")
		return Formula{}, &myError{"It is UnsupportedType..."}
	}
}

func CalcStart(sPattern, poolSize string, currentStart int) string {
	var start string
	if strings.HasPrefix(sPattern, POOLSIZE) {
		rate, _ := strconv.ParseFloat(strings.Split(sPattern, "*")[1], 64)
		poolsize, _ := strconv.Atoi(poolSize)
		start = fmt.Sprintf("%d", (int(float64(poolsize)*rate) + currentStart))
	}
	return start
}

func CalcConcurrent(sPattern, poolSize string) string {
	var concurrent string
	if strings.HasPrefix(sPattern, POOLSIZE) {
		rate, _ := strconv.ParseFloat(strings.Split(sPattern, "*")[1], 64)
		poolsize, _ := strconv.Atoi(poolSize)
		concurrent = fmt.Sprintf("%d", (int(float64(poolsize) * rate)))
	}
	return concurrent
}

func GenerateConfigJson(v interface{}, tid string, pSize []string) {
	var (
		iConcurrency   Concurrency
		iMetrics       Metrics
		fContent       []string
		configFilePath string
		config         ConfigStruct
		currentStart   int
		currentStep    int
		tenPrefix      string
		poolPrefix     string
		poolDir        string
	)

	// Initialize config strcuture
	configInit := ConfigStruct{
		Parameters:    Parameter{},
		Concurrencies: []Concurrency{},
		Metrics:       []Metrics{},
		PartitionKey: PartitionKey{
			Type:  "path",
			Value: "tenantid",
		},
	}

	config = configInit
	currentStart = 0
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	for i, psize := range pSize {
		tenPrefix = fmt.Sprintf("Ten%s", tid)
		poolPrefix = fmt.Sprintf("Pool0%d", i)
		// Create pool directory
		poolDir = path.Join(currentDir, "config", tenPrefix, poolPrefix)
		if isExists, _ := PathExists(poolDir); isExists == false {
			os.MkdirAll(poolDir, 0777)
		}
		switch sessionCfgContent := v.(type) {
		case []SessionCfgCSVStruct:
			for _, r := range sessionCfgContent {
				if r.Section == PARAMETER {
					fContent = strings.Split(r.Formula, ",")
					if strings.HasPrefix(fContent[2], POOLSIZE) {
						fContent[2] = CalcConcurrent(fContent[2], psize)
					}
					config.Parameters = GenerateParameter(fContent)
				}
				if r.Section == CONCURRENCY {
					fContent = strings.Split(r.Formula, ",")
					if len(fContent) > 1 {
						if strings.HasPrefix(fContent[0], POOLSIZE) {
							currentStep, _ = strconv.Atoi(CalcStart(fContent[0], psize, currentStart))
							fContent[0] = fmt.Sprintf("%d", currentStart)
						}
						if strings.HasPrefix(fContent[1], INITPATTER) {
							fContent[1] = strings.Replace(fContent[1], INITPATTER, tenPrefix+poolPrefix, 1)
						}
					}
					iConcurrency, _ = GenerateConcurrency(r.Path, r.Type, r.Generator, fContent)
					config.AddConcurrency(iConcurrency)
				} else if r.Section == METRICS {
					fContent = strings.Split(r.Formula, ",")
					iMetrics, _ = GenerateMetrics(r.Path, r.Type, r.Generator, fContent)
					config.AddMetrics(iMetrics)
				}
				if r.Section == FILENAME {
					// Write to file
					configFilePath = fmt.Sprintf("%s.json", r.Path)
					configFilePath = path.Join(tenPrefix, poolPrefix, configFilePath)
					GenerateJSONFile(&config, configFilePath, true)
					config = configInit
					currentStart += currentStep
				}
			}
		default:
			fmt.Println("UnsupportedType...")
		}
		currentStart = 0
	}

}
