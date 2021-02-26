package models

import (
	"math"
)

// Calculate 称量引入不确定度
func (w Weigh) Calculate() (data float64, value float64) {
	w.Selection = OptionsValue[w.Options]
	data = math.Sqrt(math.Pow((w.Data/2)/w.Selection, 2) + math.Pow(w.Value/w.Selection, 2))
	value = data / 100
	return
}

// Calculate 加入基体改进剂引入不确定度
func (m Matrix) Calculate() (data float64, value float64) {
	if m.Category {
		value = m.Data / m.Selection
		data = m.Value / m.Selection
	} else {
		m.Selection = OptionsValue[m.Options]
		data = m.Data / m.Selection
		value = m.Value / m.Selection
	}
	return
}

// Figure 表二
func (m Matrix) Figure() (data float64, value float64, option float64) {
	data, value = m.Calculate()
	if m.Category {
		value = value * 0.1
	} else {
		value = value / 0.1
		option = math.Sqrt(data + value)
	}
	return
}

// Calculate 标准溶液配制过程中引入的体积相对标准不确定度
func (s Solution) Calculate(m Matrix) (data float64, value float64) {
	s.Selection = OptionsValue[s.Options]
	data = s.Data / s.Selection
	value = data / 10
	ua, ub, _ := m.Figure()
	for key, profile := range s.Profile {
		pop := (profile.N1 * ua) + ub + profile.N1*value
		s.Profile[key].Value = math.Sqrt(pop)
	}
	return
}

// Calculate 测量重复性引入不确定度
func (g Gauging) Calculate() Gauging {
	for key, value := range g {
		average := value.Data.Avg()
		g[key].Values.UC = math.Sqrt(value.Data.Sigma(average) / (float64(len(value.Data)) - 1))
		g[key].Values.UR = g[key].Values.UC / average
		g[key].Values.DL = 3 * g[key].Values.UC
		g[key].Values.QL = 10 * g[key].Values.UC
	}
	return g
}

// Calculate 标准曲线拟合引入不确定度
//func (s Standard) Calculate(g Gauging) {
//
//	// step 1. 得到Yi 以及 Xi 值
//	var (
//		Yo utils.ArrayFloat64
//		Xo utils.ArrayFloat64
//	)
//	for _, Yi := range s.Graticule.YI {
//		for _, Si := range s.Graticule.SI {
//			Yo = append(Yo, Yi/Si) // 标准物质峰面积
//		}
//	}
//	for _, Xi := range s.Graticule.XI {
//		Xo = append(Xo, Xi/s.Thickness) // 标准物质浓度
//	}
//	s.Graticule.YI = Yo
//	s.Graticule.XI = Xo
//	// step 2.计算SR值
//	var (
//		loop  bool    = true
//		index int     = 0
//		Pow   float64 = 0
//		SR    float64 = 0
//		Xa    float64 = 0 // 标准曲线标注浓度 平均值
//
//	)
//	for loop {
//		Pow = Pow + math.Pow(Yo[index] - (s.Bight.ResultA + s.Bight.ResultB*Xo[index]), 2)
//		index = index + 1
//	}
//	SR = math.Sqrt(Pow / (s.Number - 2))
//
//	// step 3. 标准曲线标注浓度 平均值
//	Xa = s.Graticule.SI.Avg()
//	var B utils.ArrayFloat64
//	B = append(B, g[0].Data[0], g[0].Data[1])
//	// step 4. 线性拟合标准不确定度
//	SR/s.Bight.ResultB * math.Sqrt(1/s.Parallel + 1/s.Number + )
//}
