package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CtoF(c Celsius) {return Fahrenheit(c * 9/5 + 32)}
func FtoC(f Fahrenheit) {return Celsius((f - 32) * 5/9)}
