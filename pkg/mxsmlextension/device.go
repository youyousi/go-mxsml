// "Copyright (c) 2026 MetaX Integrated Circuits (Shanghai) Co., Ltd. All rights reserved."
package mxsmlextension

import (
	"unsafe"

	"go-mxsml/pkg/utils"
)

const MXSMLEX_SYSTEM_MXSMLEX_VERSION_BUFFER_SIZE = 80

func (m MxSmlExDevice) GetUUID() (string, MxSmlExReturn) {
	var uuid [MXSMLEX_DRIVER_VERSION_BUFFER_SIZE]byte
	var length uint32 = MXSMLEX_DRIVER_VERSION_BUFFER_SIZE
	ret := mxSmlExDeviceGetUUID(m, &uuid[0], length)
	if ret != MXSMLEX_SUCCESS {
		return "", ret
	}

	return utils.ExtractValidString(uuid[:length]), MXSMLEX_SUCCESS
}

func (m MxSmlExDevice) GetName() (string, MxSmlExReturn) {
	var name [MXSMLEX_DRIVER_VERSION_BUFFER_SIZE]byte
	var size uint32 = MXSMLEX_DRIVER_VERSION_BUFFER_SIZE
	ret := mxSmlExDeviceGetName(m, &name[0], size)
	if ret != MXSMLEX_SUCCESS {
		return "", ret
	}

	return utils.ExtractValidString(name[:size]), MXSMLEX_SUCCESS
}

func (m MxSmlExDevice) GetPciInfo() (MxSmlExPciInfo, MxSmlExReturn) {
	var pciInfo MxSmlExPciInfo
	ret := mxSmlExDeviceGetPciInfo(m, &pciInfo)
	return pciInfo, ret
}

func (m MxSmlExDevice) GetMemoryInfo() (MxSmlExMemory, MxSmlExReturn) {
	var memory MxSmlExMemory
	ret := mxSmlExDeviceGetMemoryInfo(m, &memory)
	return memory, ret
}

func (m MxSmlExDevice) GetFanSpeed() (uint32, MxSmlExReturn) {
	var speed uint32
	ret := mxSmlExDeviceGetFanSpeed(m, &speed)
	return speed, ret
}

func (m MxSmlExDevice) GetFanSpeed_v2(fan int) (uint32, MxSmlExReturn) {
	var speed uint32
	ret := mxSmlExDeviceGetFanSpeed_v2(m, uint32(fan), &speed)
	return speed, ret
}

func (m MxSmlExDevice) GetUtilizationRates() (MxSmlExUtilization, MxSmlExReturn) {
	var utilization MxSmlExUtilization
	ret := mxSmlExDeviceGetUtilization(m, &utilization)
	return utilization, ret
}

func (m MxSmlExDevice) GetMinorNumber() (int, MxSmlExReturn) {
	var minorNumber uint32
	ret := mxSmlExDeviceGetMinorNumber(m, &minorNumber)
	return int(minorNumber), ret
}

func (m MxSmlExDevice) GetNvLinkState(link int) (MxSmlExEnableState, MxSmlExReturn) {
	var enable MxSmlExEnableState
	ret := mxSmlExDeviceGetMetaXLinkState(m, uint32(link), &enable)
	return enable, ret
}

func (m MxSmlExDevice) GetNvLinkRemotePciInfo(link int) (PciInfo, MxSmlExReturn) {
	var pciInfo MxSmlExPciInfo
	ret := mxSmlExDeviceGetMetaXLinkRemotePciInfo_v2(m, uint32(link), &pciInfo)
	return pciInfo, ret
}

func (m MxSmlExDevice) GetTemperature(sensor TemperatureSensors) (uint32, MxSmlExReturn) {
	var temperature uint32
	ret := mxSmlExDeviceGetTemperature(m, sensor, &temperature)
	return temperature, ret
}

func (m MxSmlExDevice) GetCurrPcieLinkWidth() (int, MxSmlExReturn) {
	var linkWidth uint32
	ret := mxSmlExDeviceGetCurrPcieLinkWidth(m, &linkWidth)
	return int(linkWidth), ret
}

func (m MxSmlExDevice) GetComputeRunningProcesses() ([]MxSmlExProcessInfo, MxSmlExReturn) {
	var infoCount uint32 = 1
	infos := make([]MxSmlExProcessInfo, infoCount)
	ret := mxSmlExDeviceGetComputeRunningProcesses(m, &infoCount, &infos[0])
	if ret == SUCCESS {
		return infos[:infoCount], ret
	}
	if ret != MXSMLEX_ERROR_INSUFFICIENT_SIZE {
		return nil, ret
	}

	infos = make([]MxSmlExProcessInfo, infoCount)
	ret = mxSmlExDeviceGetComputeRunningProcesses(m, &infoCount, &infos[0])
	return infos[:infoCount], ret
}

func (m MxSmlExDevice) GetCudaComputeCapability() (int, int, MxSmlExReturn) {
	var major, minor int32
	ret := mxSmlExDeviceGetComputeCapability(m, &major, &minor)
	return int(major), int(minor), ret
}

func (m MxSmlExDevice) GetPowerUsage() (uint32, MxSmlExReturn) {
	var powerUsage uint32
	ret := mxSmlExGetPowerUsage(m, &powerUsage)
	return powerUsage, ret
}

func (m MxSmlExDevice) GetPowerManagementLimit() (uint32, MxSmlExReturn) {
	var limit uint32
	ret := mxSmlExGetPowerManagementLimit(m, &limit)
	return limit, ret
}

func (m MxSmlExDevice) GetMaxPcieLinkWidth() (int, MxSmlExReturn) {
	var maxLinkWidth uint32
	ret := mxSmlExGetMaxPcieLinkWidth(m, &maxLinkWidth)
	return int(maxLinkWidth), ret
}

func (m MxSmlExDevice) GetPcieThroughput(counter MxSmlExPcieUtilCounter) (uint32, MxSmlExReturn) {
	var val uint32
	ret := mxSmlExGetPcieThroughput(m, counter, &val)
	return val, ret
}

func (m MxSmlExDevice) GetTemperatureThreshold(thresholdType MxSmlExTemperatureThresholds) (uint32, MxSmlExReturn) {
	var temp uint32
	ret := mxSmlExDeviceGetTemperatureThreshold(m, thresholdType, &temp)
	return temp, ret
}

func (m MxSmlExDevice) GetClockInfo(clockType MxSmlExClockType) (uint32, MxSmlExReturn) {
	var clock uint32
	ret := mxSmlExDeviceGetClockInfo(m, clockType, &clock)
	return clock, ret
}

func (m MxSmlExDevice) GetCurrentClocksThrottleReasons() (uint64, MxSmlExReturn) {
	var clocksThrottleReasons uint64
	ret := mxSmlExDeviceGetCurrentClocksThrottleReasons(m, &clocksThrottleReasons)
	return clocksThrottleReasons, ret
}

func (m MxSmlExDevice) GetSupportedClocksThrottleReasons() (uint64, MxSmlExReturn) {
	var supportedClocksThrottleReasons uint64
	ret := mxSmlExDeviceGetSupportedClocksThrottleReasons(m, &supportedClocksThrottleReasons)
	return supportedClocksThrottleReasons, ret
}

func (m MxSmlExDevice) GetCpuAffinity(numCPUs int) ([]uint, MxSmlExReturn) {
	cpuSetSize := uint32((numCPUs-1)/int(unsafe.Sizeof(uint64(0))) + 1)
	cpuSet := make([]uint64, cpuSetSize)
	ret := mxSmlExDeviceGetCpuAffinity(m, cpuSetSize, &cpuSet[0])
	if ret != MXSMLEX_SUCCESS {
		return []uint{}, ret
	}

	var resSet []uint
	for _, set := range cpuSet {
		resSet = append(resSet, uint(set))
	}
	return resSet, MXSMLEX_SUCCESS
}

func (m MxSmlExDevice) GetSupportedMemoryClocks() (int, []uint32, MxSmlExReturn) {
	var count uint32 = 16
	clocksMHz := make([]uint32, count)
	ret := mxSmlExDeviceGetSupportedMemoryClocks(m, &count, &clocksMHz[0])
	if ret == MXSMLEX_ERROR_INSUFFICIENT_SIZE {
		clocksMHz = make([]uint32, count)
		ret = mxSmlExDeviceGetSupportedMemoryClocks(m, &count, &clocksMHz[0])
	}
	if ret != MXSMLEX_SUCCESS {
		return int(count), nil, ret
	}
	return int(count), clocksMHz, ret
}

func (m MxSmlExDevice) GetSupportedGraphicsClocks(memoryClockMHz int) (int, []uint32, MxSmlExReturn) {
	var count uint32 = 16
	clocksMHz := make([]uint32, count)
	ret := mxSmlExDeviceGetSupportedGraphicsClocks_v2(m, uint32(memoryClockMHz), &count, &clocksMHz[0])
	if ret == MXSMLEX_ERROR_INSUFFICIENT_SIZE {
		clocksMHz = make([]uint32, count)
		ret = mxSmlExDeviceGetSupportedGraphicsClocks_v2(m, uint32(memoryClockMHz), &count, &clocksMHz[0])
	}
	if ret != MXSMLEX_SUCCESS {
		return int(count), nil, ret
	}
	return int(count), clocksMHz, ret
}

func (m MxSmlExDevice) SetApplicationsClocks(memClockMHz, graphicsClockMHz uint32) MxSmlExReturn {
	return mxSmlExDeviceSetApplicationsClocks(m, memClockMHz, graphicsClockMHz)
}

func (m MxSmlExDevice) ResetApplicationsClocks() MxSmlExReturn {
	return mxSmlExDeviceResetApplicationsClocks(m)
}

func (m MxSmlExDevice) GetApplicationsClock(clockType MxSmlExClockType) (uint32, MxSmlExReturn) {
	var clockMHz uint32
	ret := mxSmlExDeviceGetApplicationsClock(m, clockType, &clockMHz)
	return clockMHz, ret
}

func (m MxSmlExDevice) GetP2PStatus(device Device, index MxSmlExGpuP2PCapsIndex) (MxSmlExGpuP2PStatus, MxSmlExReturn) {
	var p2pStatus MxSmlExGpuP2PStatus
	dev := device.(MxSmlExDevice)
	ret := mxSmlExDeviceGetP2PStatus(m, dev, index, &p2pStatus)
	return p2pStatus, ret
}

func (m MxSmlExDevice) GetPerformanceState() (MxSmlExPstates, MxSmlExReturn) {
	var pStates MxSmlExPstates
	ret := mxSmlExDeviceGetPerformanceState(m, &pStates)
	return pStates, ret
}

func (m MxSmlExDevice) SetAutoBoostedClocksEnabled(enable MxSmlExEnableState) MxSmlExReturn {
	return mxSmlExDeviceSetAutoBoostedClocksEnabled(m, enable)
}

func (m MxSmlExDevice) GetAutoBoostedClocksEnabled() (MxSmlExEnableState, MxSmlExEnableState, MxSmlExReturn) {
	var isEnable, defaultIsEnable MxSmlExEnableState
	ret := mxSmlExDeviceGetAutoBoostedClocksEnabled(m, &isEnable, &defaultIsEnable)
	return isEnable, defaultIsEnable, ret
}

func (m MxSmlExDevice) GetTopologyCommonAncestor(device Device) (MxSmlExGpuTopologyLevel, Return) {
	var pathInfo MxSmlExGpuTopologyLevel
	dev := device.(MxSmlExDevice)
	ret := mxSmlExDeviceGetTopologyCommonAncestor(m, dev, &pathInfo)
	return pathInfo, ret
}

func (m MxSmlExDevice) GetNvLinkRemoteDeviceType(link int) (MxSmlExMetaXLinkDeviceType, Return) {
	var linkType MxSmlExMetaXLinkDeviceType
	ret := mxSmlExDeviceGetMetaXLinkRemoteDeviceType(m, uint32(link), &linkType)
	return linkType, ret
}

func (m MxSmlExDevice) GetFieldValues(values []MxSmlExFieldValue) Return {
	valuesCount := len(values)
	return mxSmlExDeviceGetFieldValues(m, int32(valuesCount), &values[0])
}
