// "Copyright (c) 2026 MetaX Integrated Circuits (Shanghai) Co., Ltd. All rights reserved."
package main

import (
	"fmt"
	"log"

	nvml "go-mxsml/pkg/mxsmlextension"
)

func main() {
	ret := nvml.Init()
	if ret != nvml.SUCCESS {
		log.Fatal("init failed:", nvml.ErrorString(ret))
	}
	defer nvml.ShutDown()

	driverVer, ret := nvml.SystemGetDriverVersion()
	if ret != nvml.SUCCESS {
		log.Fatal("get driver version failed:", nvml.ErrorString(ret))
	}
	fmt.Printf("driver version: %s\n", driverVer)

	mxsmlVersion, ret := nvml.SystemGetNVMLVersion()
	if ret != nvml.SUCCESS {
		log.Fatal("get mxsml version failed:", nvml.ErrorString(ret))
	}
	fmt.Printf("mxsml version: %s\n", mxsmlVersion)

	count, ret := nvml.DeviceGetCount()
	if ret != nvml.SUCCESS {
		log.Fatal("get device count failed:", nvml.ErrorString(ret))
	}
	fmt.Printf("device count %d\n", count)

	for i := range count {
		fmt.Printf("----- get device stats start -----\n")
		getDeviceInfo(i)
		fmt.Printf("----- get device stats end -----\n")
	}

	if count > 1 {
		// test between 2 devices
		device1, ret := nvml.DeviceGetHandleByIndex(0)
		if ret != nvml.SUCCESS {
			log.Fatalf("get handle of device 0 failed: %s\n", nvml.ErrorString(ret))
		}

		device2, ret := nvml.DeviceGetHandleByIndex(1)
		if ret != nvml.SUCCESS {
			log.Fatalf("get handle of device 1 failed: %s\n", nvml.ErrorString(ret))
		}

		p2pStatus, ret := nvml.DeviceGetP2PStatus(device1, device2, nvml.P2P_CAPS_INDEX_NVLINK)
		if ret != nvml.SUCCESS {
			fmt.Printf("get p2p status between device 0 and device 1 failed: %s\n", nvml.ErrorString(ret))
		} else {
			fmt.Printf("p2p status between device 0 and device 1: %+v\n", p2pStatus)
		}
	}
}

func getDeviceInfo(index int) {
	device, ret := nvml.DeviceGetHandleByIndex(index)
	if ret != nvml.SUCCESS {
		log.Fatalf("get handle of devicefailed: %s\n", nvml.ErrorString(ret))
	}

	name, ret := device.GetName()
	if ret != nvml.SUCCESS {
		log.Fatalf("get name of device failed: %s\n", nvml.ErrorString(ret))
	}
	fmt.Printf("device name: %s\n", name)

	uuid, ret := device.GetUUID()
	if ret != nvml.SUCCESS {
		log.Fatalf("get uuid of device failed: %s\n", nvml.ErrorString(ret))
	}
	fmt.Printf("device uuid: %s\n", uuid)

	memory, ret := device.GetMemoryInfo()
	if ret != nvml.SUCCESS {
		fmt.Printf("get device memory failed: %s\n", nvml.ErrorString(ret))
	}
	fmt.Printf("device memory: %+v\n", memory)

	fanspeed, ret := device.GetFanSpeed()
	if ret != nvml.SUCCESS {
		fmt.Printf("get fanspeed of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device fanspeed: %d\n", fanspeed)
	}

	utilization, ret := device.GetUtilizationRates()
	if ret != nvml.SUCCESS {
		fmt.Printf("get utilization of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device utilization: %+v\n", utilization)
	}

	minorNum, ret := device.GetMinorNumber()
	if ret != nvml.SUCCESS {
		fmt.Printf("get monior number of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device monior number: %d\n", minorNum)
	}

	linkState, ret := device.GetNvLinkState(0)
	if ret != nvml.SUCCESS {
		fmt.Printf("get link state of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device link state: %d\n", linkState)
	}

	pciInfo, ret := device.GetNvLinkRemotePciInfo(4)
	if ret != nvml.SUCCESS {
		fmt.Printf("get remote pci info of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("remote pci info, busId legacy: %s, domain: 0x%x, bus: 0x%x, device: 0x%x, deviceId: 0x%x, " +
			"subsystemId: 0x%x, busId: %s\n",
			uint8SliceToString(pciInfo.BusIdLegacy[:]), pciInfo.Domain, pciInfo.Bus, pciInfo.Device,
			pciInfo.PciDeviceId, pciInfo.PciSubSystemId, uint8SliceToString(pciInfo.BusId[:]))
	}

	enableState, isEnable, ret := device.GetAutoBoostedClocksEnabled()
	if ret != nvml.SUCCESS {
		fmt.Printf("get auto boosted clocks of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("auto boosted clocks enable state: %d, isEnable: %d\n", enableState, isEnable)
	}

	temp, ret := device.GetTemperature(nvml.TEMPERATURE_GPU)
	if ret != nvml.SUCCESS {
		fmt.Printf("get temperature of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device temperature : %d\n", temp)
	}

	width, ret := device.GetCurrPcieLinkWidth()
	if ret != nvml.SUCCESS {
		fmt.Printf("get link width of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device link width : %d\n", width)
	}

	processInfo, ret := device.GetComputeRunningProcesses()
	if ret != nvml.SUCCESS {
		fmt.Printf("get process info of device failed: %s\n", nvml.ErrorString(ret))
	}
	if len(processInfo) > 0 {
		fmt.Printf("device process info: %d\n", processInfo)
	} else {
		fmt.Println("no process")
	}

	major, minor, ret := device.GetCudaComputeCapability()
	if ret != nvml.SUCCESS {
		fmt.Printf("get compute capability of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device compute capability major: %d, minor: %d\n", major, minor)
	}

	powerUsage, ret := device.GetPowerUsage()
	if ret != nvml.SUCCESS {
		fmt.Printf("get power usage of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device power usage: %+v\n", powerUsage)
	}

	limit, ret := device.GetPowerManagementLimit()
	if ret != nvml.SUCCESS {
		fmt.Printf("get power management limit of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device power management limit: %d\n", limit)
	}

	maxWidth, ret := device.GetMaxPcieLinkWidth()
	if ret != nvml.SUCCESS {
		fmt.Printf("get max link width of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device max link width: %d\n", maxWidth)
	}

	throughputTx, ret := device.GetPcieThroughput(nvml.PCIE_UTIL_TX_BYTES)
	if ret != nvml.SUCCESS {
		fmt.Printf("get tx throughput of device failed: %s\n", nvml.ErrorString(ret))
	}

	throughputRx, ret := device.GetPcieThroughput(nvml.PCIE_UTIL_RX_BYTES)
	if ret != nvml.SUCCESS {
		fmt.Printf("get rx throughput of device failed: %s\n", nvml.ErrorString(ret))
	}

	fmt.Printf("Pcie throughput: tx %d, rx %d\n", throughputTx, throughputRx)

	threshold, ret := device.GetTemperatureThreshold(nvml.TEMPERATURE_THRESHOLD_SHUTDOWN)
	if ret != nvml.SUCCESS {
		fmt.Printf("get temperature threshold of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device temperature threshold: %d\n", threshold)
	}

	clock, ret := device.GetClockInfo(nvml.CLOCK_MEM)
	if ret != nvml.SUCCESS {
		fmt.Printf("get mc0 clock failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("mc0 clock: %d\n", clock)
	}

	clock, ret = device.GetClockInfo(nvml.MXSMLEX_CLOCK_SM)
	if ret != nvml.SUCCESS {
		fmt.Printf("get xcore clock failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("xcore clock: %d\n", clock)
	}

	clock, ret = device.GetClockInfo(nvml.MXSMLEX_CLOCK_VIDEO)
	if ret != nvml.SUCCESS {
		fmt.Printf("get vpue0 clock failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("vpue0 clock: %d\n", clock)
	}

	clock, ret = device.GetApplicationsClock(nvml.MXSMLEX_CLOCK_VIDEO)
	if ret != nvml.SUCCESS {
		fmt.Printf("get applications vpue0 clock of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("applications vpue0 clock: %d\n", clock)
	}

	reason, ret := device.GetCurrentClocksThrottleReasons()
	if ret != nvml.SUCCESS {
		fmt.Printf("get current clocks throttle reasons of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device current clocks throttle reasons: %d\n", reason)
	}

	sReasons, ret := device.GetSupportedClocksThrottleReasons()
	if ret != nvml.SUCCESS {
		fmt.Printf("get supported clocks throttle reasons of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device supported clocks throttle reasons: %d\n", sReasons)
	}

	affinity, ret := device.GetCpuAffinity(0)
	if ret != nvml.SUCCESS {
		fmt.Printf("get cpu affinity of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device cpu affinity: %+v\n", affinity)
	}

	count, clocksMHz, ret := device.GetSupportedMemoryClocks()
	if ret != nvml.SUCCESS {
		fmt.Printf("get supported memory clocks of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device supported memory clocks count: %d, MHz: %+v\n", count, clocksMHz)
	}

	count, clocksMHz, ret = device.GetSupportedGraphicsClocks(0)
	if ret != nvml.SUCCESS {
		fmt.Printf("get supported graphics clocks of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device supported graphics clocks count: %d, MHz: %+v\n", count, clocksMHz)
	}

	pciInfo, ret = device.GetPciInfo()
	if ret != nvml.SUCCESS {
		fmt.Printf("get pci info of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("pci info: busId legacy: %s, domain: 0x%x, bus: 0x%x, device: 0x%x, deviceId: 0x%x, " +
			"subsystemId: 0x%x, busId: %s\n",
			uint8SliceToString(pciInfo.BusIdLegacy[:]), pciInfo.Domain, pciInfo.Bus, pciInfo.Device,
			pciInfo.PciDeviceId, pciInfo.PciSubSystemId, uint8SliceToString(pciInfo.BusId[:]))
	}

	device2, ret := nvml.DeviceGetHandleByUUID(uuid)
	if ret != nvml.SUCCESS {
		log.Fatalf("get device by uuid failed: %s", nvml.ErrorString(ret))
	}

	speed2, ret := nvml.DeviceGetFanSpeed_v2(device2, 0)
	if ret != nvml.SUCCESS {
		fmt.Printf("get speed v2 of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device fanspeed v2: %d\n", speed2)
	}

	perfState, ret := nvml.DeviceGetPerformanceState(device2)
	if ret != nvml.SUCCESS {
		fmt.Printf("get performance state of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device performance state: %d\n", perfState)
	}

	deviceType, ret := nvml.DeviceGetNvLinkRemoteDeviceType(device2, 4)
	if ret != nvml.SUCCESS {
		fmt.Printf("get remote device type of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device's remote device type state: %d for link 4\n", deviceType)
	}

	var values []nvml.FieldValue
	values = append(values, nvml.FieldValue{
		FieldId: nvml.NVML_FI_DEV_NVLINK_LINK_COUNT,
		ScopeId: 5,
	})
	values = append(values, nvml.FieldValue{
		FieldId: 1,
	})
	ret = nvml.DeviceGetFieldValues(device2, values)
	if ret != nvml.SUCCESS {
		fmt.Printf("get field values of device failed: %s\n", nvml.ErrorString(ret))
	} else {
		fmt.Printf("device's field values: %+v\n", values)
	}
}

func uint8SliceToString(data []int8) string {
	end := 0
	var tmpArr []uint8
	for ; end < len(data); end++ {
		if data[end] == 0 {
			break
		}

		tmpArr = append(tmpArr, uint8(data[end]))
	}
	return string(tmpArr)
}
