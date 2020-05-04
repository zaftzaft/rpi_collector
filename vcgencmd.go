package main

import (
	"os/exec"
	"strconv"
	"strings"
)

var (
	VcgencmdClockList = []string{
		"arm",
		"core",
		"H264",
		"isp",
		"v3d",
		"uart",
		"pwm",
		"emmc",
		"pixel",
		"vec",
		"hdmi",
		"dpi",
	}

	VcgencmdVoltsList = []string{
		"core",
		"sdram_c",
		"sdram_i",
		"sdram_p",
	}

	VcgencmdThrottledBits = map[string]int{
		"under_voltage_detected":              0,
		"arm_frequency_capped":                1,
		"currently_throttled":                 2,
		"soft_temperature_limit_active":       3,
		"under_voltage_has_occurred":          16,
		"arm_frequency_capping_has_occurred":  17,
		"throttling_has_occurred":             18,
		"soft_temperature_limit_has_occurred": 19,
	}
)

type Vcgencmd struct {
	Bin string
}

func (v *Vcgencmd) Exec(args string) (string, error) {
	out, err := exec.Command(v.Bin, strings.Split(args, " ")...).Output()

	return string(out), err
}

func (v *Vcgencmd) MeasureTemp() (float64, error) {
	s, err := v.Exec("measure_temp")
	if err != nil {
		return 0, nil
	}

	temp := strings.Split(s, "=")[1]
	temp = strings.TrimSpace(temp)
	temp = strings.Replace(temp, "C", "", 1)
	temp = strings.Replace(temp, "'", "", 1)

	return strconv.ParseFloat(temp, 64)
}

func (v *Vcgencmd) MeasureClocks() (map[string]float64, error) {
	clocks := map[string]float64{}

	for _, c := range VcgencmdClockList {
		s, err := v.Exec("measure_clock " + c)
		if err != nil {
			continue
		}

		clk := strings.Split(s, "=")[1]
		clk = strings.TrimSpace(clk)

		if n, err := strconv.ParseFloat(clk, 64); err == nil {
			clocks[c] = n
		}
	}

	return clocks, nil
}

func (v *Vcgencmd) MeasureVolts() (map[string]float64, error) {
	volts := map[string]float64{}

	for _, c := range VcgencmdVoltsList {
		s, err := v.Exec("measure_volts " + c)
		if err != nil {
			continue
		}

		volt := strings.Split(s, "=")[1]
		volt = strings.TrimSpace(volt)
		volt = strings.Replace(volt, "V", "", 1)

		if n, err := strconv.ParseFloat(volt, 64); err == nil {
			volts[c] = n
		}
	}

	return volts, nil
}

func (v *Vcgencmd) GetThrottled() (map[string]float64, error) {
	throttled := map[string]float64{}

	s, err := v.Exec("get_throttled")
	if err != nil {
		return throttled, err
	}

	s = strings.Split(s, "=")[1]
	s = strings.TrimSpace(s)
	thr, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return throttled, err
	}

	for c, bit := range VcgencmdThrottledBits {
		throttled[c] = float64((thr & (1 << bit)) >> bit)
	}

	return throttled, nil
}
