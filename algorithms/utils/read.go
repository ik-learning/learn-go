package utils

import (
	"encoding/json"
	"io/ioutil"

	. "./fixtures"
)

func ReadFixturesInt(filename string) TestRecordsInt {
	file, _ := ioutil.ReadFile(filename)
	data := TestRecordsInt{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func ReadFixturesIntV2(filename string) TestRecordsIntV2 {
	file, _ := ioutil.ReadFile(filename)
	data := TestRecordsIntV2{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func ReadFixturesIntV3(filename string) TestRecordsIntV3 {
	file, _ := ioutil.ReadFile(filename)
	data := TestRecordsIntV3{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}