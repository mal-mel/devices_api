package entity

import "encoding/json"

const (
	unicodeQuoteByte = 34
)

type CustomBool struct {
	Bool bool
}

func (cb *CustomBool) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"true"`, `true`, `"1"`, `1`:
		cb.Bool = true
	case `"false"`, `false`, `"0"`, `0`:
		cb.Bool = false
	}
	return nil
}

type CustomFloat32 struct {
	Float32 float32
}

func (cf *CustomFloat32) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if data[0] != unicodeQuoteByte {
		return json.Unmarshal(data, &cf.Float32)
	}

	return json.Unmarshal(data[1:len(data)-1], &cf.Float32)
}
