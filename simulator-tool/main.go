package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/bitly/go-simplejson"
)

type testRecord struct {
	//Related seconds for a sending interval
	sendingTime int
	//Sending content
	content      string
	clientID     string
	clientName   string
	requestTime  string
	partitionKey string
}

var stream string
var region string
var credentialsFile string
var credentialsProvider string

var connNum int
var minute int
var sampleFile string
var configFile string
var dump bool
var interval int
var burst int
var beginTimestamp string
var endTimestamp string

var dumpPath string

const goroutineNum int = 300 //base on linux threads limit
const unit int = 60          //s

type result struct {
	success int
	fail    int
}

func main() {

	//init logger
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	//get parameter
	flag.StringVar(&stream, "u", "", "Post Stream Name")
	flag.StringVar(&region, "r", "us-east-1", "Post AWS Region")
	flag.IntVar(&connNum, "c", 1, "Concurrent number")
	flag.IntVar(&interval, "i", 1, "Interval")
	flag.IntVar(&burst, "b", -1, "Send data in first few seconds of interval")
	flag.IntVar(&minute, "m", 0, "Total simulation time in Minutes")
	flag.StringVar(&credentialsFile, "k", "credentials", "Credentials file path")
	flag.StringVar(&credentialsProvider, "p", "default", "Credentials provider in file")
	flag.StringVar(&sampleFile, "s", "data/data.json", "Sample file path")
	flag.StringVar(&configFile, "f", "config/config.json", "Config file path")
	flag.BoolVar(&dump, "d", false, "Enable local disk dumper. If enabled, post data will be saved under directory dumpfile")
	flag.StringVar(&beginTimestamp, "bt", "0", "Begin timestamp")

	flag.Parse()

	if dump {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0])) //get the directory of the currently running file
		dumpPath = path.Join(dir, "dumpfile")
		os.MkdirAll(dumpPath, 0777) //create dump file directory
	}

	if (burst == -1) || (burst > interval) {
		burst = interval
	}

	var beginTime int64
	if beginTimestamp != "0" {
		beginTime, _ = strconv.ParseInt(beginTimestamp, 10, 64) //convert to int64
	} else {
		beginTime = time.Now().Unix()
	}

	kinesisStream := generateKinesisStream(region, credentialsFile, credentialsProvider)
	sample := readJsonFile(sampleFile)
	sampleConfig := readJsonFile(configFile)
	templateSets := generateTemplateSets(connNum, beginTime, interval, burst, sample, sampleConfig.GetPath("concurrency"))

	var totalRecords int
	totalRecords = minute * connNum
	Info.Printf("Supposed total records:%d", totalRecords)

	var totalTime int
	totalTime = minute * interval
	Info.Printf("Supposed total time:%d", totalTime)

	cRecord := make(chan testRecord, connNum)
	cResult := make(chan result, totalRecords)

	for i := 0; i < goroutineNum; i++ {
		go createGoroutine(kinesisStream, cRecord, cResult)
	}

	startTime := time.Now()
	Info.Printf("Starting Time: %v\n", startTime)

	generateAndSendRecord(connNum, interval, minute, templateSets, sampleConfig.GetPath("metrics"), sampleConfig.GetPath("partitionKey"), cRecord, beginTime)

	success := 0
	fail := 0
	for i := 0; i < totalRecords; i++ {
		result := <-cResult
		success = success + result.success
		fail = fail + result.fail
	}

	Info.Printf("Total Success Records:%d\n", success)
	Info.Printf("Total failed Records:%d\n", fail)
	Info.Printf("Ending Time: %v\n", time.Now())
	Info.Printf("Total sending time: %s\n", time.Since(startTime))

}

func createGoroutine(kinesisStream *kinesis.Kinesis, cRecord chan testRecord, cResult chan result) {
	for {
		result := result{0, 0}
		record, ok := <-cRecord
		//check if processing is ended, chanel is closed
		if ok == false {
			return
		}

		//body, status, err := (&record).post()
		//Info.Printf("status:%s, body:%s", status, body)

		_, _, err := (&record).post(kinesisStream)

		if err != nil {
			Error.Println(err)
			result.fail = result.fail + 1
		} else {
			result.success = result.success + 1
			//local disk dumper
			if dump {
				clientName := record.clientName
				//os.MkdirAll(path.Join(dumpPath, clientName), 0777)
				path := path.Join(dumpPath, clientName+"-"+record.requestTime+".json")
				f, err1 := os.Create(path)
				if err1 != nil {
					Error.Println(err1)
				}
				defer f.Close()
				err2 := ioutil.WriteFile(path, []byte(record.content), 0777)
				if err2 != nil {

					Error.Println(err2)
				}
			}
		}
		cResult <- result
	}
}

func (tc *testRecord) post(kinesisStream *kinesis.Kinesis) (string, string, error) {
	if kinesisStream == nil {
		return "", "", errors.New("Kinesis Stream is not initialized")
	}
	content := tc.content
	params := &kinesis.PutRecordInput{
		Data:         []byte(content),
		PartitionKey: aws.String(tc.partitionKey),
		StreamName:   aws.String(stream),
	}
	resp, err := kinesisStream.PutRecord(params)

	if err != nil {
		Error.Println(err)
		return "", "", err
	}

	return *(resp.SequenceNumber), *(resp.ShardId), nil
}

func generatePartitionKey(content *simplejson.Json, partitionKeyConfig *simplejson.Json) string {
	value, _ := partitionKeyConfig.Get("value").String()
	keyType, _ := partitionKeyConfig.Get("type").String()
	switch keyType {
	case "string":
		return value
	case "path":
		pathValue, _ := getNode(content, value).String()
		return pathValue
	}
	return value
}

func generateAndSendRecord(connNum int, interval int, minute int, templateSets map[int][]testRecord, metricConfigs *simplejson.Json, partitionKeyConfigs *simplejson.Json, cRecord chan testRecord, beginTime int64) {
	var totalTime int = minute * interval

	counter := 0
	ticker := time.NewTicker(time.Second)
	metricConfigsArray, _ := metricConfigs.Array()
	var ct string
	switchStatus := make([][]conStatus, 0, 5)
	c := -1
	for x := 0; x < totalTime; x++ {
		<-ticker.C
		Info.Printf("Ticked:%d\n", x)
		m := x / interval
		c = (c + 1) % connNum
		set := templateSets[x%interval]
		ct = strconv.FormatInt((beginTime+int64(unit*counter+x%interval))*1000, 10)

		for j := 0; j < len(set); j++ {
			set[j].requestTime = strconv.Itoa(x) //ms
			newRecord, _ := simplejson.NewJson([]byte(set[j].content))
			switchIndex := -1
			for p := 0; p < len(metricConfigsArray); p++ {
				path, _ := metricConfigs.GetIndex(p).GetPath("path").String()
				//valuetype, _ := metricConfigs.GetIndex(p).GetPath("type").String()
				generator, _ := metricConfigs.GetIndex(p).GetPath("generator").String()
				formula := metricConfigs.GetIndex(p).GetPath("formula")
				switch generator {
				case "optional_data":
					optionalDataGenerator(newRecord, path, formula)
				case "range_data":
					rangeDataGenerator(newRecord, path, formula)
				case "timestamp_data":
					//requestTime := strconv.FormatInt(time.Now().Unix()*1000, 10)
					timestampDataGenerator(newRecord, path, formula, ct)
				case "sequential_data":
					sequentialDataGenerator(newRecord, path, formula, counter)
				case "fixed_data":
					fixedDataGenerator(newRecord, path, formula)
				case "random_switch_data":
					switchIndex++
					if (m == 0) && (c == 0) {
						status := make([]conStatus, connNum, connNum)
						for v := 0; v < len(status); v++ {
							status[v].status = -1
						}
						switchStatus = append(switchStatus, status)
					}
					switchStatus[switchIndex][c] = randomSwitchDataGenerator(newRecord, formula, switchStatus[switchIndex][c], ct, m)
				}
			}
			set[j].partitionKey = generatePartitionKey(newRecord, partitionKeyConfigs)
			newContent, _ := newRecord.Encode()
			set[j].content = string(newContent)
			cRecord <- set[j]
		}
		if (x+1)%interval == 0 {
			counter++
		}
	}
	ticker.Stop()
}

func generateTemplateSets(connNum int, beginTime int64, interval int, burst int, samples *simplejson.Json, configs *simplejson.Json) map[int][]testRecord {
	m := make(map[int][]testRecord)
	for n := 0; n < interval; n++ {
		clients := []testRecord{}
		m[n] = clients
	}

	connConfigs, _ := configs.Array()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	counters := make([]int, 0, 5)
	for q := 0; q < connNum; q++ {
		counterIndex := -1
		for p := 0; p < len(connConfigs); p++ {
			path, _ := configs.GetIndex(p).GetPath("path").String()
			//valuetype, _ := metricConfigs.GetIndex(p).GetPath("type").String()
			generator, _ := configs.GetIndex(p).GetPath("generator").String()
			formula := configs.GetIndex(p).GetPath("formula")
			switch generator {
			case "optional_data":
				optionalDataGenerator(samples, path, formula)
			case "range_data":
				rangeDataGenerator(samples, path, formula)
			case "timestamp_data":
				timestampDataGenerator(samples, path, formula, strconv.FormatInt(beginTime*1000, 10))
			case "iterative_data":
				iterativeDataGenerator(samples, path, formula, q)
			case "array_data":
				arrayDataGenerator(samples, path, formula)
			case "sequential_string_data":
				counterIndex++
				if q == 0 {
					sequentialCounter := initSequentialStringDataGenerator(samples, path, formula)
					counters = append(counters, sequentialCounter)
				} else {
					sequentialCounter := sequentialStringDataGenerator(samples, path, formula, counters[counterIndex])
					counters[counterIndex] = sequentialCounter
				}
			}
		}

		id := strconv.Itoa(q + 1)
		clientName, _ := samples.GetPath("daas_vm_id").String()
		sendingTime := r.Intn(burst)
		content, _ := samples.Encode()
		record := testRecord{clientID: id, content: string(content), sendingTime: sendingTime, clientName: clientName}
		m[sendingTime] = append(m[sendingTime], record)

	}

	for i := 0; i < interval; i++ {
		Info.Printf("=====%d seconds\n", i)
		sets := m[i]
		for j := 0; j < len(sets); j++ {
			Info.Println(sets[j].clientID)
		}
	}

	return m
}

func generateKinesisStream(region string, file string, provider string) *kinesis.Kinesis {
	kinesisSession := session.New(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewSharedCredentials(file, provider),
	})

	return kinesis.New(kinesisSession)
}

func readJsonFile(file string) *simplejson.Json {
	fileBytes, err := ioutil.ReadFile(file)
	check(err, "can not open config file : "+file)
	fileJson, _ := simplejson.NewJson(fileBytes)
	return fileJson
}
