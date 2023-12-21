package hive

type Operation string

const (
	OperationEngage  Operation = "engage"
	OperationRefrain Operation = "refrain"
	OperationCall    Operation = "call"
	OperationPing    Operation = "ping"
)
