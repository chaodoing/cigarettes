package models

import "math"

var OptionsValue = map[string]float64{
	"orthostate": 3,
	"rectangle":  math.Sqrt(3),
	"triangle":   math.Sqrt(6),
	"arcsine":    math.Sqrt(2),
	"trapezoid":  2,
	"points":     1,
}
