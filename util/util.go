package util

import "strconv"

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
func CalcularDigitoRut(rut int) rune {
	m := 0
	s := 1
	for ; rut != 0; rut /= 10 {
		s = (s + rut%10*(9-(m%6))) % 11
		m++
	}
	if s != 0 {
		return rune(s + 47)
	}
	return rune(75)

}
