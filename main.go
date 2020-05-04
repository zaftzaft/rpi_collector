package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	filename = kingpin.Flag("output", "output filename").Short('o').Default("/tmp/node_exporter/rpi.prom").String()
	mkdirp   = kingpin.Flag("mkdirp", "mkdirp").Short('p').Bool()

	cmd = kingpin.Flag("command", "vcgencmd command path").Short('c').Default("/opt/vc/bin/vcgencmd").String()
)

func Run() int {
	v := &Vcgencmd{}
	v.Bin = *cmd

	registry := prometheus.NewRegistry()

	rpiTemp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "rpi_temp",
		Help: "temperature of the SoC as measured by the on-board temperature sensor",
	})

	rpiClock := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "rpi_clock",
		Help: "current frequency of the specified clock",
	}, []string{"clock"})

	rpiVolts := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "rpi_volts",
		Help: "current voltages used by the specific block.",
	}, []string{"block"})

	rpiThrottled := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "rpi_throttled",
		Help: " throttled state of the system",
	}, []string{"meaning"})

	registry.MustRegister(rpiTemp)
	registry.MustRegister(rpiClock)
	registry.MustRegister(rpiVolts)
	registry.MustRegister(rpiThrottled)

	t, err := v.MeasureTemp()
	if err != nil {
		fmt.Println(err)
	}
	rpiTemp.Set(t)

	clocks, err := v.MeasureClocks()
	if err != nil {
		fmt.Println(err)
	}
	for c, f := range clocks {
		rpiClock.WithLabelValues([]string{c}...).Add(f)
	}

	volts, err := v.MeasureVolts()
	if err != nil {
		fmt.Println(err)
	}
	for b, f := range volts {
		rpiVolts.WithLabelValues([]string{b}...).Add(f)
	}

	thr, err := v.GetThrottled()
	if err != nil {
		fmt.Println(err)
	}

	for m, f := range thr {
		rpiThrottled.WithLabelValues([]string{m}...).Add(f)
	}

	if *mkdirp {
		if err := os.MkdirAll(filepath.Dir(*filename), 0755); err != nil {
			fmt.Println(err)
			return 1
		}
	}

	if err := prometheus.WriteToTextfile(*filename, registry); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func main() {
	kingpin.Version("0.0.1")
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	os.Exit(Run())
}
