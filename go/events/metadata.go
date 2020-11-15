// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    metadata, err := UnmarshalMetadata(bytes)
//    bytes, err = metadata.Marshal()

package events

import "encoding/json"

func UnmarshalMetadata(data []byte) (Metadata, error) {
	var r Metadata
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Metadata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Metadata struct {
	TraceID string `json:"traceId"`
	UserID  string `json:"userId"` 
}
