package main

import ()

func UpdateStruct(iStruct struct{}) struct{} {
	return struct{}{}
}

func GenerateOptionData(platformMap []string) Formula {
	optionData := Formula{
		List: platformMap,
	}
	return optionData
}

func GenerateRangeData(rType, rMax, rMin string) Formula {
	rangeData := Formula{
		Type: rType,
		Max:  rMax,
		Min:  rMin,
	}
	return rangeData
}

func GenerateFixedData(fixedValue string) Formula {
	fixedData := Formula{
		Value: fixedValue,
	}
	return fixedData
}

func GenerateSequentialData(initialValue, sStep, sType string) Formula {
	sequentialData := Formula{
		Initial: initialValue,
		Step:    sStep,
		Type:    sType,
	}
	return sequentialData
}

func GenerateSequentialStringData(sStart, initialValue, leastLength string, exceptList []string) Formula {
	sequentialStringData := Formula{
		Start:       sStart,
		Initial:     initialValue,
		LeastLength: leastLength,
		Except:      exceptList,
	}
	return sequentialStringData
}

func GenerateTimestampData() Formula {
	return Formula{}
}

func GenerateIterativeData(iterativeList []string) Formula {
	iterativeData := Formula{
		List: iterativeList,
	}
	return iterativeData
}

func GenerateArrayData(aStart, initialValue, leastLength, aCount string, exceptList []string) Formula {
	arrayData := Formula{
		Start:       aStart,
		Initial:     initialValue,
		LeastLength: leastLength,
		Count:       aCount,
		Except:      exceptList,
	}
	return arrayData
}

func GenerateRandomSwitchData() Formula {
	return Formula{}
}
