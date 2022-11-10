package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type K float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	ZeroK         Celsius = -273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k K) String() string          { return fmt.Sprintf("%g°K", k) }

//e2.1 新增开尔文为单位(K)的值
