package fixtures

type TestRecordInt struct {
	Input  []int32 `json:"input"`
	Result int32   `json:"result"`
}

type TestRecordsInt []TestRecordInt

type TestRecordIntV2 struct {
	Input  []int32 `json:"input"`
	Result []int32 `json:"result"`
}

type TestRecordsIntV2 []TestRecordIntV2
//
type InputV3 struct {
	Budget int32 `json:"budget"`
	Keyboard  []int32 `json:"keyboard"`
	Usb  []int32 `json:"usb"`
}

type TestRecordIntV3 struct {
	Input  InputV3 `json:"input"`
	Result int32 `json:"result"`
}

type TestRecordsIntV3 []TestRecordIntV3