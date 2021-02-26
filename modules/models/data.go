package models

import "github.com/chaodoing/cigarettes/providers/utils"

type (
	// 称量引入不确定度
	Weigh struct {
		Data      float64 `json:"data" xml:"data"`
		Value     float64 `json:"value" xml:"value"`
		Options   string  `json:"options" xml:"options"`
		Selection float64 `json:"selection" xml:"selection"`
	}
	// 加入基体改进剂引入不确定度
	Matrix struct {
		Category  bool    `json:"category" xml:"category"`
		Data      float64 `json:"data" xml:"data"`
		Value     float64 `json:"value" xml:"value"`
		Selection float64 `json:"selection" xml:"selection"`
		Options   string  `json:"options" xml:"options"`
	}
	// 标准溶液配体积引入不确定度
	Solution struct {
		Data      float64 `json:"data" xml:"data"`
		Selection float64 `json:"value" xml:"value"`
		Options   string  `json:"options" xml:"options"`
		Profile   []struct {
			ID    int     `json:"id" xml:"id"`
			Name  string  `json:"name" xml:"name"`
			N1    float64 `json:"n1" xml:"n1"`
			Value float64 `json:"value" xml:"value"`
		} `json:"profile" xml:"profile"`
	}
	values struct {
		UC float64 `json:"uc" xml:"uc"`
		UR float64 `json:"ur" xml:"ur"`
		DL float64 `json:"dl" xml:"dl"`
		QL float64 `json:"ql" xml:"ql"`
	}
	// 测量重复性引入不确定度
	Gauging []struct {
		ID     int                `json:"id" xml:"id"`
		Name   string             `json:"name" xml:"name"`
		Data   utils.ArrayFloat64 `json:"data" xml:"data"`
		Values values             `json:"values" xml:"values"`
	}
	// 标准曲线拟合引入不确定度
	Standard struct {
		Graticule struct {
			ID   int                `json:"id" xml:"id"`
			Name string             `json:"name" xml:"name"`
			YI   utils.ArrayFloat64 `json:"yi" xml:"yi"`
			BI   utils.ArrayFloat64 `json:"bi" xml:"bi"`
			XI   utils.ArrayFloat64 `json:"xi" xml:"xi"`
			SI   utils.ArrayFloat64 `json:"si" xml:"si"`
		} `json:"standard" xml:"standard"` // Graticule 标线拟合
		Bight struct {
			ID      int     `json:"id" xml:"id"`
			Name    string  `json:"name" xml:"name"`
			ResultA float64 `json:"result_a" xml:"result_a"`
			ResultB float64 `json:"result_b" xml:"result_b"`
		} `json:"bight" xml:"bight"`                         // 标准曲线
		Thickness float64 `json:"thickness" xml:"thickness"` // 浓度
		Parallel  float64 `json:"parallel" xml:"parallel"`   // 平行次数
		Number    float64 `json:"number" xml:"number"` // 标线拟合文件文件夹中的文件数目
	}
	// 整体应用数据
	Application struct {
		King     float64  `json:"king" xml:"king"`         // 包含因子K3
		Weigh    Weigh    `json:"weigh" xml:"weigh"`       // 称量引入不确定度
		Matrix   Matrix   `json:"matrix" xml:"matrix"`     // 加入基体改进剂引入不确定度
		Solution Solution `json:"solution" xml:"solution"` // 标准溶液配体积引入不确定度
		Gauging  Gauging  `json:"gauging" xml:"gauging"`   // 测量重复性引入不确定度
		Standard Standard `json:"standard" xml:"standard"` // 标准曲线拟合引入不确定度
	}
)
