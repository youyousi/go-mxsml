// "Copyright (c) 2026 MetaX Integrated Circuits (Shanghai) Co., Ltd. All rights reserved."
package mxsml

// #include <dlfcn.h>
// #include <stdint.h>
import "C"
import (
	"github.com/MetaX-MACA/go-mxsml/pkg/lib"
	"github.com/MetaX-MACA/go-mxsml/pkg/utils"
)

type MxSmlWrapper struct {
}

var mxsmlWrapper = createMxsmlWrapper()

func createMxsmlWrapper() *MxSmlWrapper {
	return &MxSmlWrapper{}
}


func (m *MxSmlWrapper) mxSmlInit() MxSmlReturn {
	err := lib.Load()
	if err != nil {
		return MXSML_LoadDllFailure
	}

	return mxSmlInit()
}

func (m *MxSmlWrapper) mxSmlInitWithFlags(flags uint32) MxSmlReturn {
	err := lib.Load()
	if err != nil {
		return MXSML_LoadDllFailure
	}

	return mxSmlInitWithFlags(flags)
}

func (m *MxSmlWrapper) mxSmlShutDown() error {
	return lib.Unload()
}

func (m *MxSmlWrapper) mxSmlGetErrorString(result MxSmlReturn) string {
	return utils.BytePtrToStr(mxSmlGetErrorString(result))
}

func (m *MxSmlWrapper) mxSmlGetDeviceCount() uint32 {
	return mxSmlGetDeviceCount()
}

func (m *MxSmlWrapper) mxSmlGetPfDeviceCount() uint32 {
	return mxSmlGetPfDeviceCount()
}

func (m *MxSmlWrapper) mxSmlGetMacaVersion() (string, MxSmlReturn) {
	var version [MXSML_VERSION_INFO_SIZE]byte
	var size uint32 = MXSML_VERSION_INFO_SIZE

	ret := mxSmlGetMacaVersion(&version[0], &size)
	if ret != MXSML_Success {
		return "", ret
	}

	return utils.ExtractValidString(version[:size]), MXSML_Success
}

func (m *MxSmlWrapper) mxSmlGetDeviceVersion(deviceId uint32, versionUnit MxSmlVersionUnit) (string, MxSmlReturn) {
	var version [MXSML_VERSION_INFO_SIZE]byte
	var size uint32 = MXSML_VERSION_INFO_SIZE

	ret := mxSmlGetDeviceVersion(deviceId, versionUnit, &version[0], &size)
	if ret != MXSML_Success {
		return "", ret
	}

	return utils.ExtractValidString(version[:size]), MXSML_Success
}

func (m *MxSmlWrapper) mxSmlGetChipSerial(deviceId uint32) (string, MxSmlReturn) {
	var serial [MXSML_CHIP_SERIAL_SIZE]byte
	var size uint32 = MXSML_CHIP_SERIAL_SIZE

	ret := mxSmlGetChipSerial(deviceId, &serial[0], &size)
	if ret != MXSML_Success {
		return "", ret
	}

	return utils.ExtractValidString(serial[:size]), MXSML_Success
}

func (m *MxSmlWrapper) mxSmlGetBoardSerial(deviceId uint32) (string, MxSmlReturn) {
	var serial [MXSML_BOARD_SERIAL_SIZE]byte
	var size uint32 = MXSML_BOARD_SERIAL_SIZE
	ret := mxSmlGetBoardSerial(deviceId, &serial[0], &size)
	if ret != MXSML_Success {
		return "", ret
	}

	return utils.ExtractValidString(serial[:size]), MXSML_Success
}

func (m *MxSmlWrapper) mxSmlGetPptableVersion(deviceId uint32) (string, MxSmlReturn) {
	var version [MXSML_VERSION_INFO_SIZE]byte
	var size uint32 = MXSML_VERSION_INFO_SIZE
	ret := mxSmlGetPptableVersion(deviceId, size, &version[0])
	if ret != MXSML_Success {
		return "", ret
	}

	return utils.ExtractValidString(version[:size]), MXSML_Success
}

func (m *MxSmlWrapper) mxSmlGetDeviceRealPath(deviceId uint32) (string, MxSmlReturn) {
	var relPath [4096]byte
	var size uint32 = 4096
	ret := mxSmlGetDeviceRealPath(deviceId, &relPath[0], &size)
	if ret != MXSML_Success {
		return "", ret
	}

	return utils.ExtractValidString(relPath[:size]), MXSML_Success
}

func (m *MxSmlWrapper) mxSmlGetMetaXLinkState(deviceId uint32) (MxSmlMetaXLinkState, string, MxSmlReturn) {
	var state [128]byte
	var size uint32 = 128
	var stateCode MxSmlMetaXLinkState
	ret := mxSmlGetMetaXLinkState(deviceId, &stateCode, &state[0], &size)
	if ret != MXSML_Success {
		return 0, "", ret
	}

	return stateCode, utils.ExtractValidString(state[:size]), MXSML_Success
}

func (m *MxSmlWrapper) mxSmlGetDeviceInfo(deviceId uint32) (MxSmlDeviceInfo, MxSmlReturn) {
	var devInfo MxSmlDeviceInfo
	ret := mxSmlGetDeviceInfo(deviceId, &devInfo)
	return devInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetVirtualDevicesByPhysicalId(phyDeviceId uint32) (MxSmlVirtualDeviceIds, MxSmlReturn) {
	var deviceIds MxSmlVirtualDeviceIds
	ret := mxSmlGetVirtualDevicesByPhysicalId(phyDeviceId, &deviceIds)
	return deviceIds, ret
}

func (m *MxSmlWrapper) mxSmlGetAllLimitedDevices() (MxSmlLimitedDeviceIds, MxSmlReturn) {
	var deviceIds MxSmlLimitedDeviceIds
	ret := mxSmlGetAllLimitedDevices(&deviceIds)
	return deviceIds, ret
}

func (m *MxSmlWrapper) mxSmlGetLimitedDeviceInfo(deviceId uint32) (MxSmlDeviceInfo, MxSmlReturn) {
	var deviceInfo MxSmlDeviceInfo
	ret := mxSmlGetLimitedDeviceInfo(deviceId, &deviceInfo)
	return deviceInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetMemoryInfo(deviceId uint32) (MxSmlMemoryInfo, MxSmlReturn) {
	var memoryInfo MxSmlMemoryInfo
	ret := mxSmlGetMemoryInfo(deviceId, &memoryInfo)
	return memoryInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetPmbusInfo(deviceId uint32, pmbusUnit MxSmlPmbusUnit) (MxSmlPmbusInfo, MxSmlReturn) {
	var pmbusInfo MxSmlPmbusInfo
	ret := mxSmlGetPmbusInfo(deviceId, pmbusUnit, &pmbusInfo)
	return pmbusInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetTemperatureInfo(deviceId uint32, tempType MxSmlTemperatureSensors) (int32, MxSmlReturn) {
	var temp int32
	ret := mxSmlGetTemperatureInfo(deviceId, tempType, &temp)
	return temp, ret
}

func (m *MxSmlWrapper) mxSmlGetBoardPowerInfo(deviceId uint32) ([]MxSmlBoardWayElectricInfo, MxSmlReturn) {
	var maxSize uint32 = 36
	boardInfo := make([]MxSmlBoardWayElectricInfo, maxSize)
	ret := mxSmlGetBoardPowerInfo(deviceId, &maxSize, &boardInfo[0])
	return boardInfo[:maxSize], ret
}

func (m *MxSmlWrapper) mxSmlGetDpmMaxPerfLevel(deviceId uint32) (MxSmlMxcDpmPerfLevel, MxSmlReturn) {
	var dpmMaxLevel MxSmlMxcDpmPerfLevel
	ret := mxSmlGetDpmMaxPerfLevel(deviceId, &dpmMaxLevel)
	return dpmMaxLevel, ret
}

func (m *MxSmlWrapper) mxSmlGetDpmIpMaxPerfLevel(deviceId uint32, dpmIp MxSmlDpmIp) (uint32, MxSmlReturn) {
	var dpmMaxLevel uint32
	ret := mxSmlGetDpmIpMaxPerfLevel(deviceId, dpmIp, &dpmMaxLevel)
	return dpmMaxLevel, ret
}

func (m *MxSmlWrapper) mxSmlGetEepromInfo(deviceId uint32) (MxSmlEepromInfo, MxSmlReturn) {
	var eepromInfo MxSmlEepromInfo
	ret := mxSmlGetEepromInfo(deviceId, &eepromInfo)
	return eepromInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetRasErrorData(deviceId uint32) (MxSmlRasErrorData, MxSmlReturn) {
	var rasErrorData MxSmlRasErrorData
	ret := mxSmlGetRasErrorData(deviceId, &rasErrorData)
	return rasErrorData, ret
}

func (m *MxSmlWrapper) mxSmlGetRasErrorData_v2(deviceId uint32) (MxSmlRasErrorData_v2, MxSmlReturn) {
	var rasErrorData MxSmlRasErrorData_v2
	ret := mxSmlGetRasErrorData_v2(deviceId, &rasErrorData)
	return rasErrorData, ret
}

func (m *MxSmlWrapper) mxSmlGetRasStatusData(deviceId uint32) (MxSmlRasStatusData, MxSmlReturn) {
	var rasStatusData MxSmlRasStatusData
	ret := mxSmlGetRasStatusData(deviceId, &rasStatusData)
	return rasStatusData, ret
}

func (m *MxSmlWrapper) mxSmlGetPcieInfo(deviceId uint32) (MxSmlPcieInfo, MxSmlReturn) {
	var pcieInfo MxSmlPcieInfo
	ret := mxSmlGetPcieInfo(deviceId, &pcieInfo)
	return pcieInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetPcieMaxLinkInfo(deviceId uint32) (MxSmlPcieInfo, MxSmlReturn) {
	var pcieInfo MxSmlPcieInfo
	ret := mxSmlGetPcieMaxLinkInfo(deviceId, &pcieInfo)
	return pcieInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetPowerStateInfo(deviceId uint32, dpmIp MxSmlDpmIp) ([]int32, MxSmlReturn) {
	var size uint32 = 32
	powerState := make([]int32, size)
	ret := mxSmlGetPowerStateInfo(deviceId, dpmIp, &powerState[0], &size)
	return powerState[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetPciPowerState(deviceId uint32) (MxSmlPciPowerState, MxSmlReturn) {
	var powerState MxSmlPciPowerState
	ret := mxSmlGetPciPowerState(deviceId, &powerState)
	return powerState, ret
}

func (m *MxSmlWrapper) mxSmlGetVpuStatus(deviceId uint32) (MxSmlVpuStatus, MxSmlReturn) {
	var vpuStatus MxSmlVpuStatus
	ret := mxSmlGetVpuStatus(deviceId, &vpuStatus)
	return vpuStatus, ret
}

func (m *MxSmlWrapper) mxSmlGetCodecStatus(deviceId uint32) (MxSmlCodecStatus, MxSmlReturn) {
	var codecStatus MxSmlCodecStatus
	ret := mxSmlGetCodecStatus(deviceId, &codecStatus)
	return codecStatus, ret
}

func (m *MxSmlWrapper) mxSmlGetClocks(deviceId uint32, clockIp MxSmlClockIp) ([]uint32, MxSmlReturn) {
	var size uint32 = 32
	clocks := make([]uint32, size)
	ret := mxSmlGetClocks(deviceId, clockIp, &size, &clocks[0])
	if ret == MXSML_InsufficientSize {
		clocks = make([]uint32, size)
		ret = mxSmlGetClocks(deviceId, clockIp, &size, &clocks[0])
	}
	if ret != MXSML_Success {
		return nil, ret
	}
	return clocks[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetHbmBandWidth(deviceId uint32) (MxSmlHbmBandWidth, MxSmlReturn) {
	var hbmBandwidth MxSmlHbmBandWidth
	ret := mxSmlGetHbmBandWidth(deviceId, &hbmBandwidth)
	return hbmBandwidth, ret
}

func (m *MxSmlWrapper) mxSmlGetPcieThroughput(deviceId uint32) (MxSmlPcieThroughput, MxSmlReturn) {
	var pcieThroughput MxSmlPcieThroughput
	ret := mxSmlGetPcieThroughput(deviceId, &pcieThroughput)
	return pcieThroughput, ret
}

func (m *MxSmlWrapper) mxSmlGetDmaBandwidth(deviceId uint32) ([]MxSmlDmaEngineBandwidth, MxSmlReturn) {
	var size uint32 = 32
	for {
		dmaBandwidth := make([]MxSmlDmaEngineBandwidth, size)
		ret := mxSmlGetDmaBandwidth(deviceId, &dmaBandwidth[0], &size)
		if ret == MXSML_Success {
			return dmaBandwidth[:size], ret
		}

		if ret == MXSML_InsufficientSize {
			size *= 2
		} else {
			return nil, ret
		}
	}
}

func (m *MxSmlWrapper) mxSmlGetMetaXLinkBandwidth(deviceId uint32,
	mxlkType MxSmlMetaXLinkType) ([]MxSmlMetaXLinkBandwidth, MxSmlReturn) {

	var size uint32 = MXSML_METAX_LINK_NUM
	mxlkBandwidth := make([]MxSmlMetaXLinkBandwidth, size)
	ret := mxSmlGetMetaXLinkBandwidth(deviceId, mxlkType, &size, &mxlkBandwidth[0])
	return mxlkBandwidth[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetMetaXLinkRemoteInfo(deviceId, linkId uint32) (MxSmlMetaXLinkRemoteInfo, MxSmlReturn) {
	var mxlkRemoteInfo MxSmlMetaXLinkRemoteInfo
	ret := mxSmlGetMetaXLinkRemoteInfo(deviceId, linkId, &mxlkRemoteInfo)
	return mxlkRemoteInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetNumberOfProcess() (uint32, MxSmlReturn) {
	var num uint32
	ret := mxSmlGetNumberOfProcess(&num)
	return num, ret
}

func (m *MxSmlWrapper) mxSmlGetProcessInfo(processNumber uint32, processInfo *MxSmlProcessInfo) MxSmlReturn {
	return mxSmlGetProcessInfo(processNumber, processInfo)
}

func (m *MxSmlWrapper) mxSmlGetProcessInfo_v2(processNumber uint32, processInfo *MxSmlProcessInfo_v2) MxSmlReturn {
	return mxSmlGetProcessInfo_v2(processNumber, processInfo)
}

func (m *MxSmlWrapper) mxSmlGetProcessInfo_v3(processNumber uint32, processInfo *MxSmlProcessInfo_v3) MxSmlReturn {
	return mxSmlGetProcessInfo_v3(processNumber, processInfo)
}

func (m *MxSmlWrapper) mxSmlGetSingleGpuProcess(deviceId uint32) ([]MxSmlProcessInfo, MxSmlReturn) {
	var count uint32 = 0
	infos := make([]MxSmlProcessInfo, 1)
	mxSmlGetSingleGpuProcess(deviceId, &count, &infos[0])

	if count == 0 {
		return nil, MXSML_Success
	}

	infos = make([]MxSmlProcessInfo, count)
	ret := mxSmlGetSingleGpuProcess(deviceId, &count, &infos[0])
	return infos[:count], ret
}

func (m *MxSmlWrapper) mxSmlGetSingleGpuProcess_v2(deviceId uint32) ([]MxSmlProcessInfo_v2, MxSmlReturn) {
	var count uint32 = 0
	infos := make([]MxSmlProcessInfo_v2, 1)
	mxSmlGetSingleGpuProcess_v2(deviceId, &count, &infos[0])

	if count == 0 {
		return nil, MXSML_Success
	}

	infos = make([]MxSmlProcessInfo_v2, count)
	ret := mxSmlGetSingleGpuProcess_v2(deviceId, &count, &infos[0])
	return infos[:count], ret
}

func (m *MxSmlWrapper) mxSmlGetSingleGpuProcess_v3(deviceId uint32) ([]MxSmlProcessInfo_v3, MxSmlReturn) {
	var count uint32 = 0
	infos := make([]MxSmlProcessInfo_v3, 1)
	mxSmlGetSingleGpuProcess_v3(deviceId, &count, &infos[0])

	if count == 0 {
		return nil, MXSML_Success
	}

	infos = make([]MxSmlProcessInfo_v3, count)
	ret := mxSmlGetSingleGpuProcess_v3(deviceId, &count, &infos[0])
	return infos[:count], ret
}

func (m *MxSmlWrapper) mxSmlGetDeviceTopology(deviceId1, deviceId2 uint32) (MxSmlGpuTopologyLevel, MxSmlReturn) {
	var gpuTopoLevel MxSmlGpuTopologyLevel
	ret := mxSmlGetDeviceTopology(deviceId1, deviceId2, &gpuTopoLevel)
	return gpuTopoLevel, ret
}

func (m *MxSmlWrapper) mxSmlGetDeviceDistance(deviceId1, deviceId2 uint32) (uint32, MxSmlReturn) {
	var distance uint32
	ret := mxSmlGetDeviceDistance(deviceId1, deviceId2, &distance)
	return distance, ret
}

func (m *MxSmlWrapper) mxSmlGetCpuAffinity(deviceId uint32, cpuSetSize uint32) ([]uint32, MxSmlReturn) {
	cpuSet := make([]uint32, cpuSetSize)
	ret := mxSmlGetCpuAffinity(deviceId, cpuSetSize,  &cpuSet[0])
	return cpuSet, ret
}

func (m *MxSmlWrapper) mxSmlGetNodeAffinity(deviceId uint32, nodeSetSize uint32) ([]uint32, MxSmlReturn) {
	nodeSet := make([]uint32, nodeSetSize)
	ret := mxSmlGetNodeAffinity(deviceId, nodeSetSize,  &nodeSet[0])
	return nodeSet, ret
}

func (m *MxSmlWrapper) mxSmlGetPciDelay(deviceId uint32) (uint32, MxSmlReturn) {
	var delay uint32
	ret := mxSmlGetPciDelay(deviceId, &delay)
	return delay, ret
}

func (m *MxSmlWrapper) mxSmlGetDpmIpClockInfo(deviceId uint32, dpmIp MxSmlDpmIp) ([]uint32, MxSmlReturn) {
	var size uint32 = 8
	clockInfo := make([]uint32, size)
	ret := mxSmlGetDpmIpClockInfo(deviceId, dpmIp, &clockInfo[0], &size)
	if ret == MXSML_InsufficientSize {
		clockInfo = make([]uint32, size)
		ret = mxSmlGetDpmIpClockInfo(deviceId, dpmIp, &clockInfo[0], &size)
	}

	if ret != MXSML_Success {
		return nil, ret
	}
	return clockInfo[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetDpmIpVddInfo(deviceId uint32, dpmIp MxSmlDpmIp) ([]uint32, MxSmlReturn) {
	var size uint32 = 8
	voltageInfo := make([]uint32, size)
	ret := mxSmlGetDpmIpVddInfo(deviceId, dpmIp, &voltageInfo[0], &size)
	if ret == MXSML_InsufficientSize {
		voltageInfo = make([]uint32, size)
		ret = mxSmlGetDpmIpClockInfo(deviceId, dpmIp, &voltageInfo[0], &size)
	}
	if ret != MXSML_Success {
		return nil, ret
	}

	return voltageInfo[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetCurrentDpmIpPerfLevel(deviceId uint32, dpmIp MxSmlDpmIp) (uint32, MxSmlReturn) {
	var dpmIpPerLevel uint32
	ret := mxSmlGetCurrentDpmIpPerfLevel(deviceId, dpmIp, &dpmIpPerLevel)
	return dpmIpPerLevel, ret
}

func (m *MxSmlWrapper) mxSmlGetDeviceIpUsage(deviceId uint32, ip MxSmlUsageIp) (int32, MxSmlReturn) {
	var usage int32
	ret := mxSmlGetDeviceIpUsage(deviceId, ip, &usage)
	return usage, ret
}

func (m *MxSmlWrapper) mxSmlGetXcoreApUsage(deviceId uint32) ([]uint32, MxSmlReturn) {
	var size uint32 = 1
	var dpmNum uint32 = 1
	apusage := make([]uint32, 1)
	ret := mxSmlGetXcoreApUsage(deviceId, &apusage[0], &size, &dpmNum)
	if ret == MXSML_InsufficientSize {
		apusage = make([]uint32, size)
		ret =  mxSmlGetXcoreApUsage(deviceId, &apusage[0], &size, &dpmNum)
	}

	if ret != MXSML_Success {
		return nil, ret
	}

	return apusage[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetApUsageToggle(deviceId uint32) (uint32, MxSmlReturn) {
	var toggle uint32
	ret := mxSmlGetApUsageToggle(deviceId, &toggle)
	return toggle, ret
}

func (m *MxSmlWrapper) mxSmlGetFwIpLoglevel(deviceId uint32, fwIpName MxSmlFwIpName) (MxSmlLoglevel, MxSmlReturn) {
	var loglevel MxSmlLoglevel
	ret := mxSmlGetFwIpLoglevel(deviceId, fwIpName, &loglevel)
	return loglevel, ret
}

func (m *MxSmlWrapper) mxSmlGetFwLoglevel(deviceId uint32) (MxSmlFwLoglevel, MxSmlReturn) {
	var logLevel MxSmlFwLoglevel
	ret := mxSmlGetFwLoglevel(deviceId, &logLevel)
	return logLevel, ret
}

func (m *MxSmlWrapper) mxSmlGetEccState(deviceId uint32) (uint32, MxSmlReturn) {
	var eccState uint32
	ret := mxSmlGetEccState(deviceId, &eccState)
	return eccState, ret
}

func (m *MxSmlWrapper) mxSmlGetMetaXLinkInfo(deviceId uint32) (MxSmlMetaXLinkInfo, MxSmlReturn) {
	var metaxLinkInfo MxSmlMetaXLinkInfo
	ret := mxSmlGetMetaXLinkInfo(deviceId, &metaxLinkInfo)
	return metaxLinkInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetMetaXLinkInfo_v2(deviceId uint32) ([]MxSmlSingleMxlkInfo, MxSmlReturn) {
	var size uint32 = 1
	mxlkLinkInfo := make([]MxSmlSingleMxlkInfo, size)
	ret := mxSmlGetMetaXLinkInfo_v2(deviceId, &size, &mxlkLinkInfo[0])
	if ret == MXSML_InsufficientSize {
		mxlkLinkInfo = make([]MxSmlSingleMxlkInfo, size)
		ret = mxSmlGetMetaXLinkInfo_v2(deviceId, &size, &mxlkLinkInfo[0])
	}

	if ret != MXSML_Success {
		return nil, ret
	}

	return mxlkLinkInfo[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetSriovState(deviceId uint32) (uint32, MxSmlReturn) {
	var sriovState uint32
	ret := mxSmlGetSriovState(deviceId, &sriovState)
	return sriovState, ret
}

func (m *MxSmlWrapper) mxSmlGetDeviceSlot(deviceId uint32) (uint32, MxSmlReturn) {
	var slotId uint32
	ret := mxSmlGetDeviceSlot(deviceId, &slotId)
	return slotId, ret
}

func (m *MxSmlWrapper) mxSmlGetOpticalModuleStatus(deviceId uint32) ([]MxSmlOpticalModuleStatus, MxSmlReturn) {
	var size uint32 = 1
	opModuleStatus := make([]MxSmlOpticalModuleStatus, size)
	ret := mxSmlGetOpticalModuleStatus(deviceId, &opModuleStatus[0], &size)
	if ret == MXSML_InsufficientSize {
		opModuleStatus = make([]MxSmlOpticalModuleStatus, size)
		ret = mxSmlGetOpticalModuleStatus(deviceId, &opModuleStatus[0],  &size)
	}

	if ret != MXSML_Success {
		return nil, ret
	}

	return opModuleStatus[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetOpticalModuleInfo(deviceId uint32) ([]MxSmlOpticalModuleInfo, MxSmlReturn) {
	var size uint32 = 1
	opModuleInfo := make([]MxSmlOpticalModuleInfo, size)
	ret := mxSmlGetOpticalModuleInfo(deviceId, &opModuleInfo[0], &size)
	if ret == MXSML_InsufficientSize {
		opModuleInfo = make([]MxSmlOpticalModuleInfo, size)
		ret = mxSmlGetOpticalModuleInfo(deviceId, &opModuleInfo[0],  &size)
	}

	if ret != MXSML_Success {
		return nil, ret
	}

	return opModuleInfo[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetCurrentClocksThrottleReason(deviceId uint32) (uint64, MxSmlReturn) {
	var reason uint64
	ret := mxSmlGetCurrentClocksThrottleReason(deviceId, &reason)
	return reason, ret
}

func (m *MxSmlWrapper) mxSmlGetBoardPowerLimit(deviceId uint32) (uint32, MxSmlReturn) {
	var limit uint32
	ret := mxSmlGetBoardPowerLimit(deviceId, &limit)
	return limit, ret
}

func (m *MxSmlWrapper) mxSmlSetDpmIpMaxPerfLevel(deviceId uint32, dpmIp MxSmlDpmIp, maxPerfLevel uint32) MxSmlReturn {
	return mxSmlSetDpmIpMaxPerfLevel(deviceId, dpmIp, maxPerfLevel)
}

func (m *MxSmlWrapper) mxSmlDumpVbios(deviceId uint32, timeLimit uint32, binPath string) (uint32, MxSmlReturn) {
	var dupmRet uint32
	ret := mxSmlDumpVbios(deviceId, timeLimit, binPath, &dupmRet)
	return dupmRet, ret
}

func (m *MxSmlWrapper) mxSmlVbiosUpgrade(deviceId uint32, vbiosUpgradeArg *MxSmlVbiosUpgradeArg) MxSmlReturn {
	return mxSmlVbiosUpgrade(deviceId, vbiosUpgradeArg)
}

func (m *MxSmlWrapper) mxSmlFunctionLevelReset(deviceId uint32) MxSmlReturn {
	return mxSmlFunctionLevelReset(deviceId)
}

func (m *MxSmlWrapper) mxSmlFunctionLevelResetVfFromPf(deviceId uint32) MxSmlReturn {
	return mxSmlFunctionLevelResetVfFromPf(deviceId)
}

func  (m *MxSmlWrapper) mxSmlSetUnlockKey(deviceId uint32, unlockKey string) MxSmlReturn {
	return mxSmlSetUnlockKey(deviceId, unlockKey)
}

func (m *MxSmlWrapper) mxSmlSetCpuAffinity(deviceId uint32) MxSmlReturn {
	return mxSmlSetCpuAffinity(deviceId)
}

func (m *MxSmlWrapper) mxSmlClearCpuAffinity(deviceId uint32) MxSmlReturn {
	return mxSmlClearCpuAffinity(deviceId)
}

func (m *MxSmlWrapper) mxSmlReset(deviceId uint32) MxSmlReturn {
	return mxSmlReset(deviceId)
}

func (m *MxSmlWrapper) mxSmlSetPciSpeed(deviceId uint32, pciGen MxSmlPciGen) MxSmlReturn {
	return mxSmlSetPciSpeed(deviceId, pciGen)
}

func (m *MxSmlWrapper) mxSmlSetPciDelay(deviceId uint32, pciDelay uint32) MxSmlReturn {
	return mxSmlSetPciDelay(deviceId, pciDelay)
}

func (m *MxSmlWrapper) mxSmlSetFwLoglevel(deviceId uint32, fwIpName MxSmlFwIpName,
	loglevel MxSmlLoglevel) MxSmlReturn {
	return mxSmlSetFwLoglevel(deviceId, fwIpName, loglevel)
}

func (m *MxSmlWrapper) mxSmlSetApUsageToggle(deviceId uint32, toggle int32) MxSmlReturn {
	return mxSmlSetApUsageToggle(deviceId, toggle)
}

func (m *MxSmlWrapper) mxSmlSetEccState(deviceId uint32, eccState uint32) MxSmlReturn {
	return mxSmlSetEccState(deviceId, eccState)
}

func (m *MxSmlWrapper) mxSmlGetMetaXLinkPortState(deviceId uint32) ([]MxSmlMxlkPortState, MxSmlReturn) {
	var size uint32 = 16
	mxlkPortState := make([]MxSmlMxlkPortState, size)
	ret := mxSmlGetMetaXLinkPortState(deviceId, &mxlkPortState[0], &size)
	if ret == MXSML_InsufficientSize {
		mxlkPortState = make([]MxSmlMxlkPortState, size)
		ret = mxSmlGetMetaXLinkPortState(deviceId, &mxlkPortState[0], &size)
	}

	if ret != MXSML_Success {
		return nil, ret
	}

	return mxlkPortState[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetPciMmioState(deviceId uint32) (uint32, MxSmlReturn) {
	var state uint32
	ret := mxSmlGetPciMmioState(deviceId, &state)
	return state, ret
}

func (m *MxSmlWrapper) mxSmlGetPciEventInfo(deviceId uint32, eventType MxSmlPciEventType) ([]MxSmlPciEventInfo, MxSmlReturn) {
	var size uint32 = 8
	eventInfo := make([]MxSmlPciEventInfo, size)
	ret := mxSmlGetPciEventInfo(deviceId, eventType, &eventInfo[0], &size)
	if ret == MXSML_InsufficientSize {
		eventInfo = make([]MxSmlPciEventInfo, size)
		ret = mxSmlGetPciEventInfo(deviceId, eventType, &eventInfo[0], &size)
	}

	if ret != MXSML_Success {
		return nil, ret
	}

	return eventInfo[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetDeviceIsaVersion(deviceId uint32) (int32, MxSmlReturn) {
	var ver int32
	ret := mxSmlGetDeviceIsaVersion(deviceId, &ver)
	return ver, ret
}

func (m *MxSmlWrapper) mxSmlGetDeviceUnavailableReason(deviceId uint32) (MxSmlDeviceUnavailableReasonInfo, MxSmlReturn) {
	var reasonInfo MxSmlDeviceUnavailableReasonInfo
	ret := mxSmlGetDeviceUnavailableReason(deviceId, &reasonInfo)
	return reasonInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetTotalEccErrors(deviceId uint32) (MxSmlEccErrorCount, MxSmlReturn) {
	var errCount MxSmlEccErrorCount
	ret := mxSmlGetTotalEccErrors(deviceId, &errCount)
	return errCount, ret
}

func (m *MxSmlWrapper) mxSmlGetDeviceDieVersion(deviceId, dieId uint32, versionUnit MxSmlVersionUnit) (string, MxSmlReturn) {
	var version [MXSML_VERSION_INFO_SIZE]byte
	var size uint32 = MXSML_VERSION_INFO_SIZE
	ret := mxSmlGetDeviceDieVersion(deviceId, dieId, versionUnit, &version[0], &size)
	if ret != MXSML_Success {
		return "", ret
	}

	return utils.ExtractValidString(version[:size]), MXSML_Success
}

func (m *MxSmlWrapper) mxSmlGetDieChipSerial(deviceId, dieId uint32) (string, MxSmlReturn) {
	var serial [MXSML_CHIP_SERIAL_SIZE]byte
	var size uint32 = MXSML_CHIP_SERIAL_SIZE

	ret := mxSmlGetDieChipSerial(deviceId, dieId, &serial[0], &size)
	if ret != MXSML_Success {
		return "", ret
	}

	return utils.ExtractValidString(serial[:size]), MXSML_Success
}

func (m *MxSmlWrapper) mxSmlGetDiePptableVersion(deviceId, dieId uint32) (string, MxSmlReturn) {
	var version [MXSML_VERSION_INFO_SIZE]byte
	var size uint32 = MXSML_VERSION_INFO_SIZE
	ret := mxSmlGetDiePptableVersion(deviceId, dieId, size, &version[0])
	if ret != MXSML_Success {
		return "", ret
	}

	return utils.ExtractValidString(version[:size]), MXSML_Success
}

func (m *MxSmlWrapper) mxSmlGetDeviceDieCount(deviceId uint32) (uint32, MxSmlReturn) {
	var dieCount uint32
	ret := mxSmlGetDeviceDieCount(deviceId, &dieCount)
	return dieCount, ret
}

func (m *MxSmlWrapper) mxSmlGetDieMemoryInfo(deviceId, dieId uint32) (MxSmlMemoryInfo, MxSmlReturn) {
	var memoryInfo MxSmlMemoryInfo
	ret := mxSmlGetDieMemoryInfo(deviceId, dieId, &memoryInfo)
	return memoryInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetDieTemperatureInfo(deviceId, dieId uint32,
	temperatureType MxSmlTemperatureSensors) (int32, MxSmlReturn) {

	var temp int32
	ret := mxSmlGetDieTemperatureInfo(deviceId, dieId, temperatureType, &temp)
	return temp, ret
}

func (m *MxSmlWrapper) mxSmlGetDiePmbusInfo(deviceId, dieId uint32,
	pmbusUnit MxSmlPmbusUnit) (MxSmlPmbusInfo, MxSmlReturn) {

	var pmbusInfo MxSmlPmbusInfo
	ret := mxSmlGetDiePmbusInfo(deviceId, dieId, pmbusUnit, &pmbusInfo)
	return pmbusInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetDieClocks(deviceId, dieId uint32, clockIp MxSmlClockIp) ([]uint32, MxSmlReturn) {
	var size uint32 = 32
	clocks := make([]uint32, size)
	ret := mxSmlGetDieClocks(deviceId, dieId, clockIp, &size, &clocks[0])
	if ret == MXSML_InsufficientSize {
		clocks = make([]uint32, size)
		ret = mxSmlGetDieClocks(deviceId, dieId, clockIp, &size, &clocks[0])
	}
	if ret != MXSML_Success {
		return nil, ret
	}
	return clocks[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetDieApUsageToggle(deviceId, dieId uint32) (uint32, MxSmlReturn) {
	var toggle uint32
	ret := mxSmlGetDieApUsageToggle(deviceId, dieId, &toggle)
	return toggle, ret
}

func (m *MxSmlWrapper) mxSmlGetCurrentDieDpmIpPerfLevel(deviceId, dieId uint32,
	dpmIp MxSmlDpmIp) (uint32, MxSmlReturn) {

	var dpmLevel uint32
	ret := mxSmlGetCurrentDieDpmIpPerfLevel(deviceId, dieId, dpmIp, &dpmLevel)
	return dpmLevel, ret
}

func (m *MxSmlWrapper) mxSmlGetDieIpUsage(deviceId, dieId uint32, ip MxSmlUsageIp) (int32, MxSmlReturn) {
	var usage int32
	ret := mxSmlGetDieIpUsage(deviceId, dieId, ip, &usage)
	return usage, ret
}

func (m *MxSmlWrapper) mxSmlGetDieDpmIpMaxPerfLevel(deviceId, dieId uint32, dpmIp MxSmlDpmIp) (uint32, MxSmlReturn) {
	var dpmMaxLevel uint32
	ret := mxSmlGetDieDpmIpMaxPerfLevel(deviceId, dieId, dpmIp, &dpmMaxLevel)
	return dpmMaxLevel, ret
}

func (m *MxSmlWrapper) mxSmlGetDieXcoreApUsage(deviceId, dieId uint32) ([]uint32, MxSmlReturn) {
	var size uint32 = 1
	var dpmNum uint32 = 1
	apusage := make([]uint32, 1)
	ret := mxSmlGetDieXcoreApUsage(deviceId, dieId, &apusage[0], &size, &dpmNum)
	if ret == MXSML_InsufficientSize {
		apusage = make([]uint32, size)
		ret =  mxSmlGetDieXcoreApUsage(deviceId, dieId, &apusage[0], &size, &dpmNum)
	}

	if ret != MXSML_Success {
		return nil, ret
	}

	return apusage[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetDieCodecStatus(deviceId, dieId uint32) (MxSmlCodecStatus, MxSmlReturn) {
	var codecStatus MxSmlCodecStatus
	ret := mxSmlGetDieCodecStatus(deviceId, dieId, &codecStatus)
	return codecStatus, ret
}

func (m *MxSmlWrapper) mxSmlGetDieEepromInfo(deviceId, dieId uint32) (MxSmlEepromInfo, MxSmlReturn) {
	var eepromInfo MxSmlEepromInfo
	ret := mxSmlGetDieEepromInfo(deviceId, dieId, &eepromInfo)
	return eepromInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetDieEccState(deviceId, dieId uint32) (uint32, MxSmlReturn) {
	var eccState uint32
	ret := mxSmlGetDieEccState(deviceId, dieId, &eccState)
	return eccState, ret
}

func (m *MxSmlWrapper) mxSmlGetDiePowerStateInfo(deviceId, dieId uint32, dpmIp MxSmlDpmIp) ([]int32, MxSmlReturn) {
	var size uint32 = 32
	powerState := make([]int32, size)
	ret := mxSmlGetDiePowerStateInfo(deviceId, dieId, dpmIp, &powerState[0], &size)
	return powerState[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetDieHbmBandWidth(deviceId, dieId uint32) (MxSmlHbmBandWidth, MxSmlReturn) {
	var hbmBandwidth MxSmlHbmBandWidth
	ret := mxSmlGetDieHbmBandWidth(deviceId, dieId, &hbmBandwidth)
	return hbmBandwidth, ret
}

func (m *MxSmlWrapper) mxSmlGetDieRasStatusData(deviceId, dieId uint32) (MxSmlRasStatusData, MxSmlReturn) {
	var rasStatusData MxSmlRasStatusData
	ret := mxSmlGetDieRasStatusData(deviceId, dieId, &rasStatusData)
	return rasStatusData, ret
}

func (m *MxSmlWrapper) mxSmlGetDieRasErrorData(deviceId, dieId uint32) (MxSmlRasErrorData, MxSmlReturn) {
	var rasErrorData MxSmlRasErrorData
	ret := mxSmlGetDieRasErrorData(deviceId, dieId, &rasErrorData)
	return rasErrorData, ret
}

func (m *MxSmlWrapper) mxSmlGetDieMetaXLinkRemoteInfo(deviceId, dieId, linkId uint32) (
	MxSmlMcmMetaXLinkRemoteInfo, MxSmlReturn) {

	var remoteLinkInfo MxSmlMcmMetaXLinkRemoteInfo
	ret := mxSmlGetDieMetaXLinkRemoteInfo(deviceId, dieId, linkId, &remoteLinkInfo)
	return remoteLinkInfo, ret
}

func (m *MxSmlWrapper) mxSmlGetDieUnavailableReason(deviceId, dieId uint32) (
	MxSmlDeviceUnavailableReasonInfo, MxSmlReturn) {

	var reason MxSmlDeviceUnavailableReasonInfo
	ret := mxSmlGetDieUnavailableReason(deviceId, dieId, &reason)
	return reason, ret
}

func (m *MxSmlWrapper) mxSmlGetDieTotalEccErrors(deviceId, dieId uint32) (MxSmlEccErrorCount, MxSmlReturn) {
	var eccCount MxSmlEccErrorCount
	ret := mxSmlGetDieTotalEccErrors(deviceId, dieId, &eccCount)
	return eccCount, ret
}

func (m *MxSmlWrapper) mxSmlGetMetaXLinkTopo(deviceId uint32) (MxSmlMetaXLinkTopo, MxSmlReturn) {
	var linkTopo MxSmlMetaXLinkTopo
	ret := mxSmlGetMetaXLinkTopo(deviceId, &linkTopo)
	return linkTopo, ret
}

func (m *MxSmlWrapper) mxSmlGetMetaXLinkAer(deviceId uint32) ([]MxSmlMetaXLinkAer, MxSmlReturn) {
	var size uint32 = 7
	mxlkAer := make([]MxSmlMetaXLinkAer, size)
	ret := mxSmlGetMetaXLinkAer(deviceId, &size, &mxlkAer[0])
	if ret == MXSML_InsufficientSize {
		mxlkAer = make([]MxSmlMetaXLinkAer, size)
		ret = mxSmlGetMetaXLinkAer(deviceId, &size, &mxlkAer[0])
	}

	if ret != MXSML_Success {
		return nil, ret
	}
	return mxlkAer[:size], ret
}

func (m *MxSmlWrapper) mxSmlGetDeviceState(deviceId uint32) (int32, MxSmlReturn) {
	var state int32
	ret := mxSmlGetDeviceState(deviceId, &state)
	return state, ret
}

func (m *MxSmlWrapper) mxSmlSetOpMode(mode uint32) MxSmlReturn {
	return mxSmlSetOpMode(mode)
}

func (m *MxSmlWrapper) mxSmlGetOpMode() (uint32, MxSmlReturn) {
	var mode uint32
	ret := mxSmlGetOpMode(&mode)
	return mode, ret
}

func (m *MxSmlWrapper) mxSmlEventSetCreate() (MxSmlEventSet, MxSmlReturn) {
	var eventSet MxSmlEventSet
	ret := mxSmlEventSetCreate(&eventSet)
	return eventSet, ret
}

func (m *MxSmlWrapper) mxSmlDeviceRegisterEvents(deviceId uint32, eventTypes uint64, eventSet MxSmlEventSet) (MxSmlReturn) {
	ret := mxSmlDeviceRegisterEvents(deviceId, eventTypes, eventSet)
	return ret
}

func (m *MxSmlWrapper) mxSmlEventSetWait(eventSet MxSmlEventSet) (MxSmlEvent, MxSmlReturn) {
	var eventData MxSmlEvent
	ret := mxSmlEventSetWait(eventSet, &eventData)
	return eventData, ret
}

func (m *MxSmlWrapper) mxSmlEventSetFree(eventSet MxSmlEventSet) (MxSmlReturn) {
	ret := mxSmlEventSetFree(eventSet)
	return ret
}