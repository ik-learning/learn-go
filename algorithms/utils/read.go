package utils

import (
	"encoding/json"
	"io/ioutil"
)

type TestRecordsInt struct {
}

func ReadFixturesInt(filename string) TestRecordsInt {
	file, _ := ioutil.ReadFile(filename)
	data := TestRecordsInt{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

type TestRecordsIntV2 struct {
}

func ReadFixturesIntV2(filename string) TestRecordsIntV2 {
	file, _ := ioutil.ReadFile(filename)
	data := TestRecordsIntV2{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

type TestRecordsIntV3 struct {
}

func ReadFixturesIntV3(filename string) TestRecordsIntV3 {
	file, _ := ioutil.ReadFile(filename)
	data := TestRecordsIntV3{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
