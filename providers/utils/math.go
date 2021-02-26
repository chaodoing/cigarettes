package utils

import "math"

type ArrayFloat64 []float64

// Sum 求和
func (this ArrayFloat64) Sum() (data float64) {
	for _, value := range this {
		data = data + value
	}
	return data
}

// Avg 求平均值
func (this ArrayFloat64) Avg() (data float64) {
	data = this.Sum()
	data = data / float64(len(this))
	return data
}
// Sigma 平均求和
func (this ArrayFloat64) Sigma(average float64) (sigma float64) {
	for _, value := range this {
		sigma = sigma + math.Pow(value - average, 2)
	}
	return
}