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
