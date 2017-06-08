package main

func main() {
	admin := &Admin{}
	str :=
		`{
			"id": 1,
			"Username": "Root",
			"Password": "123456"
		 }`
	admin.Jtoo(str)
	admin.Otoj()
}
