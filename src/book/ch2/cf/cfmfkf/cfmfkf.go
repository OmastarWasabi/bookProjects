package cfmfkf

import "fmt"

type Celsius float64
type Fahrenheit float64
type Meter float64
type Foot float64
type Kilogram float64
type Lb float64

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtToM(ft Foot) Meter       { return Meter(ft / 3.28084) }
func MToFt(m Meter) Foot        { return Foot(m * 3.28084) }
func KgToLb(kg Kilogram) Lb     { return Lb(kg / 0.453592) }
func LbToKg(lb Lb) Kilogram     { return Kilogram(lb * 0.453592) }

func (c Celsius) String() string    { return fmt.Sprintf("%.2f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.2f°F", f) }
func (m Meter) String() string      { return fmt.Sprintf("%.2fM", m) }
func (ft Foot) String() string      { return fmt.Sprintf("%.2fFt", ft) }
func (kg Kilogram) String() string  { return fmt.Sprintf("%gKg", kg) }
func (lb Lb) String() string        { return fmt.Sprintf("%gLb", lb) }
