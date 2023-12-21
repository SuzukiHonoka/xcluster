package hive

var Central *Hive

func init() {
	Central = NewHive()
	Central.Open()
}
