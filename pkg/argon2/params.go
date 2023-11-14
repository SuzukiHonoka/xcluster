package argon2

type Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var ParamDefault = &Params{
	Memory:      64 * 1024, // 64M
	Iterations:  2,         // 1 ~= 34ms, 2 ~= 57ms, 3 ~= 84ms
	Parallelism: 2,         // Dual CPU Threads
	SaltLength:  16,        // 128 bits
	KeyLength:   32,        // 256 bits
}
