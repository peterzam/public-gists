package main

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

func main() {
	sensors, _ := host.SensorsTemperatures()
	for _, sensor := range sensors {
		if sensor.SensorKey == "k10temp_tctl_input" {
			fmt.Printf("%.2f", sensor.Temperature)
		}
	}
}
