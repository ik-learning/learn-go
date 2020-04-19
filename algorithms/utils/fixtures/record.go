package fixtures

type TestRecordInt struct {
	Input  []int32 `json:"input"`
	Result int32   `json:"result"`
}

type TestRecordsInt struct {
	Test []TestRecordInt `json:"test"`
}
