package hive

type Message struct {
	Sequence  uint        `json:"seq,omitempty"`
	Operation Operation   `json:"op,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type MessageWithStatus struct {
	Message
	Status bool `json:"status"`
}
