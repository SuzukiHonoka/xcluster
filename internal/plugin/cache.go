package plugin

var Registers []Register

func RegisterPlugin(register Register) {
	Registers = append(Registers, register)
}
