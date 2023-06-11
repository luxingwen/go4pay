package model

import (
	xdb "go4pay/pkg/db"
	"strconv"
)

// SysConfig 结构体
type SysConfig struct {
	Name    string `json:"name"`
	KeyName string `json:"key"`
	Value   string `json:"value"`
	Comment string `json:"comment"`
}

func GetExchangeRate() (r float64, err error) {
	var sysConfig SysConfig
	db := xdb.GetDB().Table("sys_config")
	has, err := db.Where("key_name = ?", "ExchangeRate").Get(&sysConfig)
	if err != nil {
		return
	}

	if !has {
		r = 1.0
		return
	}

	if sysConfig.Value == "" {
		r = 1.0
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	if err != nil {
		return
	}

	var floatRateConf SysConfig

	db1 := xdb.GetDB().Table("sys_config")
	_, err = db1.Where("key_name = ?", "FloatExchangeRate").Get(&sysConfig)
	if err != nil {
		return
	}

	floatRate, _ := strconv.ParseFloat(floatRateConf.Value, 64)

	r += floatRate
	return
}

// 交易金额大于该数值，按百分比收取手续费
func GetPercentageNum() (r float64, err error) {
	var sysConfig SysConfig
	db := xdb.GetDB().Table("sys_config")
	has, err := db.Where("key_name = ?", "PercentageNum").Get(&sysConfig)
	if err != nil {
		return
	}

	if !has {
		r = 50
		return
	}
	if sysConfig.Value == "" {
		r = 50
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	return
}

// 交易金额大于收取百分比值
func GetPercentageFee() (r float64, err error) {
	var sysConfig SysConfig
	db := xdb.GetDB().Table("sys_config")
	has, err := db.Where("key_name = ?", "PercentageFee").Get(&sysConfig)
	if err != nil {
		return
	}

	if !has {
		r = 1.0
		return
	}

	if sysConfig.Value == "" {
		r = 1.0
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	return
}

// 交易金额小于多少值
func GetFixedValueNum() (r float64, err error) {
	var sysConfig SysConfig
	db := xdb.GetDB().Table("sys_config")
	has, err := db.Where("key_name = ?", "FixedValueNum").Get(&sysConfig)
	if err != nil {
		return
	}

	if !has {
		r = 50
		return
	}

	if sysConfig.Value == "" {
		r = 50
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	return
}

// 交易金额小于固定值手续费用
func GetFixedValueFee() (r float64, err error) {
	var sysConfig SysConfig
	db := xdb.GetDB().Table("sys_config")
	has, err := db.Where("key_name = ?", "FixedValueFee").Get(&sysConfig)
	if err != nil {
		return
	}

	if !has {
		r = 0.5
		return
	}

	if sysConfig.Value == "" {
		r = 0.5
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	return
}

// 获取平台佣金比例
func GetCommissionRate() (r float64, err error) {
	var sysConfig SysConfig
	db := xdb.GetDB().Table("sys_config")
	has, err := db.Where("key_name = ?", "CommissionRate").Get(&sysConfig)
	if err != nil {
		return
	}

	if !has {
		r = 0.5
		return
	}

	if sysConfig.Value == "" {
		r = 0.5
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	return
}

// 计算手续费
func CalculationFee(val float64) (r float64, err error) {
	percentageNum, err := GetPercentageNum()
	if err != nil {
		return
	}
	if val > percentageNum {
		percentageFee, err := GetPercentageFee()
		if err != nil {
			return 0, err
		}
		r = val * percentageFee / 100
		return r, nil
	}

	fixedValueNum, err := GetFixedValueNum()

	if err != nil {
		return
	}

	if val <= fixedValueNum {
		fixedValueFee, err := GetFixedValueFee()
		if err != nil {
			return 0, err
		}
		return fixedValueFee, nil
	}
	r = 0.5
	return
}
