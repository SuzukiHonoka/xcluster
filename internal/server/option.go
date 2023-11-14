package server

type Option uint

const (
	OptionForceSSL Option = iota
	OptionKeepAlive
	OptionMaxRetry
	OptionTimeout
)
