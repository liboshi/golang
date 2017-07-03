package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	iPoolCapacity  int
	PoolCapacities = []string{
		"4", "10", "200", "100", "300", "400",
		"500", "600", "700", "800", "900", "1000",
	}
)

type TopoStruct struct {
	Product     string `json:"product"`
	SP          string `json:"sp"`
	TenantId    string `json:"tenantid"`
	Timestamp   int64  `json:"timestamp"`
	CMSTenantID string `json:"cmstenantid"`
	T           string `json:"t"`
	Pools       []Pool `json:"pools"`
}

type Pool struct {
	PoolID       string   `json:"poolid"`
	PoolName     string   `json:"poolname"`
	PoolCapacity string   `json:"poolcapacity"`
	VM           []string `json:"vms"`
}

func (topoStruct *TopoStruct) AddPool(item Pool) []Pool {
	topoStruct.Pools = append(topoStruct.Pools, item)
	return topoStruct.Pools
}

func GenerateTopoPool(poolId, poolName, poolCapacity string, vms []string) Pool {
	topoPool := Pool{
		PoolID:       poolId,
		PoolName:     poolName,
		PoolCapacity: poolCapacity,
		VM:           vms,
	}
	return topoPool
}

func generateTopoDataJson(spid, tenantid, cmstenantid, product string, pStart, poolSize int) {
	var poolName string
	var iPool Pool
	Pools := []Pool{}
	tempVMs := []string{
		"9811eba4-5ceb-4573-9495-dece36184995",
		"9811eba4-5ceb-4573-9495-dece36184996",
		"9811eba4-5ceb-4573-9495-dece36184997",
	}
	topo := TopoStruct{
		Product:     "horizonair",
		SP:          spid,
		TenantId:    tenantid,
		Timestamp:   time.Now().Unix(),
		CMSTenantID: cmstenantid,
		T:           "f",
		Pools:       Pools,
	}
	for i := pStart; i < (pStart + poolSize); i++ {
		iPoolCapacity = i - pStart
		if iPoolCapacity > poolSize {
			iPoolCapacity = iPoolCapacity - poolSize
		}
		poolName = fmt.Sprintf("Ten%sPool%02d", tenantid, (i - pStart))
		iPool = GenerateTopoPool(strconv.Itoa(i), poolName,
			PoolCapacities[iPoolCapacity], tempVMs)
		topo.AddPool(iPool)
	}

	// Write to file
	topoDataFilePath := fmt.Sprintf("full_topo_data_ten%02s.json", tenantid)
	GenerateJSONFile(&topo, topoDataFilePath, false)
}

func generateTopoConfigJson(tenantid string, poolSize int) {
	var tempConcurrency Concurrency
	var tempPath string
	var tempFormula Formula
	var tempInit string

	topoConcurrencies := []Concurrency{}
	topoMetrics := []Metrics{}
	topoConfig := ConfigStruct{
		Concurrencies: topoConcurrencies,
		Metrics:       topoMetrics,
		PartitionKey: PartitionKey{
			Type:  "path",
			Value: "tenantid",
		},
	}
	for i := 0; i < poolSize; i++ {
		iPoolCapacity = i
		if i > iPoolCapacity {
			iPoolCapacity = i - poolSize
		}
		tempPath = fmt.Sprintf("pools[%d].vms", i)
		tempInit = fmt.Sprintf("Ten%sPool%02dVM*", tenantid, i)
		tempFormula = GenerateArrayData("0", tempInit, "4",
			PoolCapacities[iPoolCapacity], make([]string, 0))
		tempConcurrency = Concurrency{
			Path:      tempPath,
			Generator: "array_data",
			Formula:   tempFormula,
		}
		topoConfig.AddConcurrency(tempConcurrency)
	}
	// Write to file
	topoConfigFilePath := fmt.Sprintf("topo-config-ten%02s.json", tenantid)
	GenerateJSONFile(&topoConfig, topoConfigFilePath, true)
}

func GenerateTopoJson(spid, tenantid, cmstenantid, product string, pStart, poolSize int) {
	generateTopoDataJson(spid, tenantid, cmstenantid, product, pStart, poolSize)
	generateTopoConfigJson(tenantid, poolSize)
}
