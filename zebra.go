package ciris

import "fmt"

func Simple99() string {
	return Gen("scg")
}

func ReadableSerial(id uint8, n int8) string {
	if id == 0 {
		id = 1
	}
	var sign string
	if n >= 0 {
		sign = "+"
	}
	return Gen(fmt.Sprintf("zit:%d%s%d", id, sign, n))
}

func RandomToken() string {
	return Gen("efls")
}

func RandomPassword() string {
	return Gen("rpsw")
}

func TimeBasedToken(mode uint8) string {
	return Gen(fmt.Sprintf("ptsl:%d", mode))
}
