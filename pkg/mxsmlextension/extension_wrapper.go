// "Copyright (c) 2026 MetaX Integrated Circuits (Shanghai) Co., Ltd. All rights reserved."
package mxsmlextension

// #include <dlfcn.h>
// #include <stdint.h>
import "C"
import (
	"github.com/MetaX-MACA/go-mxsml/pkg/lib"
	"github.com/MetaX-MACA/go-mxsml/pkg/utils"
)

const (
	SYSTEM_NVML_VERSION_BUFFER_SIZE   = MXSMLEX_SYSTEM_MXSMLEX_VERSION_BUFFER_SIZE
	SYSTEM_DRIVER_VERSION_BUFFER_SIZE = MXSMLEX_DRIVER_VERSION_BUFFER_SIZE
	NVML_FI_DEV_NVLINK_LINK_COUNT     = MXSMLEX_FI_DEV_METAXLINK_LINK_COUNT
)

type Return = MxSmlExReturn

const (
	SUCCESS                         Return = MXSMLEX_SUCCESS
	ERROR_UNINITIALIZED             Return = MXSMLEX_ERROR_UNINITIALIZED
	ERROR_INVALID_ARGUMENT          Return = MXSMLEX_ERROR_INVALID_ARGUMENT
	ERROR_NOT_SUPPORTED             Return = MXSMLEX_ERROR_NOT_SUPPORTED
	ERROR_NO_PERMISSION             Return = MXSMLEX_ERROR_NO_PERMISSION
	ERROR_ALREADY_INITIALIZED       Return = MXSMLEX_ERROR_ALREADY_INITIALIZED
	ERROR_NOT_FOUND                 Return = MXSMLEX_ERROR_NOT_FOUND
	ERROR_INSUFFICIENT_SIZE         Return = MXSMLEX_ERROR_INSUFFICIENT_SIZE
	ERROR_INSUFFICIENT_POWER        Return = MXSMLEX_ERROR_INSUFFICIENT_POWER
	ERROR_DRIVER_NOT_LOADED         Return = MXSMLEX_ERROR_DRIVER_NOT_LOADED
	ERROR_TIMEOUT                   Return = MXSMLEX_ERROR_TIMEOUT
	ERROR_IRQ_ISSUE                 Return = MXSMLEX_ERROR_IRQ_ISSUE
	ERROR_LIBRARY_NOT_FOUND         Return = MXSMLEX_ERROR_LIBRARY_NOT_FOUND
	ERROR_FUNCTION_NOT_FOUND        Return = MXSMLEX_ERROR_FUNCTION_NOT_FOUND
	ERROR_CORRUPTED_INFOROM         Return = MXSMLEX_ERROR_CORRUPTED_INFOROM
	ERROR_GPU_IS_LOST               Return = MXSMLEX_ERROR_GPU_IS_LOST
	ERROR_RESET_REQUIRED            Return = MXSMLEX_ERROR_RESET_REQUIRED
	ERROR_OPERATING_SYSTEM          Return = MXSMLEX_ERROR_OPERATING_SYSTEM
	ERROR_LIB_RM_VERSION_MISMATCH   Return = MXSMLEX_ERROR_LIB_RM_VERSION_MISMATCH
	ERROR_IN_USE                    Return = MXSMLEX_ERROR_IN_USE
	ERROR_MEMORY                    Return = MXSMLEX_ERROR_MEMORY
	ERROR_NO_DATA                   Return = MXSMLEX_ERROR_NO_DATA
	ERROR_VGPU_ECC_NOT_SUPPORTED    Return = MXSMLEX_ERROR_VGPU_ECC_NOT_SUPPORTED
	ERROR_INSUFFICIENT_RESOURCES    Return = MXSMLEX_ERROR_INSUFFICIENT_RESOURCES
	ERROR_FREQ_NOT_SUPPORTED        Return = MXSMLEX_ERROR_FREQ_NOT_SUPPORTED
	ERROR_ARGUMENT_VERSION_MISMATCH Return = MXSMLEX_ERROR_ARGUMENT_VERSION_MISMATCH
	ERROR_DEPRECATED                Return = MXSMLEX_ERROR_DEPRECATED
	ERROR_UNKNOWN                   Return = MXSMLEX_ERROR_UNKNOWN
)

type EnableState = MxSmlExEnableState

const (
	FEATURE_DISABLED EnableState = MXSMLEX_FEATURE_DISABLED
	FEATURE_ENABLED  EnableState = MXSMLEX_FEATURE_ENABLED
)

type TemperatureSensors = MxSmlExTemperatureSensors

const (
	TEMPERATURE_GPU   TemperatureSensors = MXSMLEX_TEMPERATURE_GPU
	TEMPERATURE_COUNT TemperatureSensors = MXSMLEX_TEMPERATURE_COUNT
)

type PcieUtilCounter = MxSmlExPcieUtilCounter

const (
	PCIE_UTIL_TX_BYTES PcieUtilCounter = MXSMLEX_PCIE_UTIL_TX_BYTES
	PCIE_UTIL_RX_BYTES PcieUtilCounter = MXSMLEX_PCIE_UTIL_RX_BYTES
	PCIE_UTIL_COUNT    PcieUtilCounter = MXSMLEX_PCIE_UTIL_COUNT
)

type TemperatureThresholds = MxSmlExTemperatureThresholds

const (
	TEMPERATURE_THRESHOLD_SHUTDOWN      TemperatureThresholds = MXSMLEX_TEMPERATURE_THRESHOLD_SHUTDOWN
	TEMPERATURE_THRESHOLD_SLOWDOWN      TemperatureThresholds = MXSMLEX_TEMPERATURE_THRESHOLD_SLOWDOWN
	TEMPERATURE_THRESHOLD_MEM_MAX       TemperatureThresholds = MXSMLEX_TEMPERATURE_THRESHOLD_MEM_MAX
	TEMPERATURE_THRESHOLD_GPU_MAX       TemperatureThresholds = MXSMLEX_TEMPERATURE_THRESHOLD_GPU_MAX
	TEMPERATURE_THRESHOLD_ACOUSTIC_MIN  TemperatureThresholds = MXSMLEX_TEMPERATURE_THRESHOLD_ACOUSTIC_MIN
	TEMPERATURE_THRESHOLD_ACOUSTIC_CURR TemperatureThresholds = MXSMLEX_TEMPERATURE_THRESHOLD_ACOUSTIC_CURR
	TEMPERATURE_THRESHOLD_ACOUSTIC_MAX  TemperatureThresholds = MXSMLEX_TEMPERATURE_THRESHOLD_ACOUSTIC_MAX

	// this threshold count is 8 in nvml, but 7 in mxsml
	TEMPERATURE_THRESHOLD_COUNT TemperatureThresholds = MXSMLEX_TEMPERATURE_THRESHOLD_COUNT
)

type ClockType = MxSmlExClockType

const (
	CLOCK_GRAPHICS ClockType = MXSMLEX_CLOCK_GRAPHICS
	CLOCK_SM       ClockType = MXSMLEX_CLOCK_SM
	CLOCK_MEM      ClockType = MXSMLEX_CLOCK_MEM
	CLOCK_VIDEO    ClockType = MXSMLEX_CLOCK_VIDEO
	CLOCK_COUNT    ClockType = MXSMLEX_CLOCK_COUNT
)

type GpuP2PCapsIndex = MxSmlExGpuP2PCapsIndex

const (
	P2P_CAPS_INDEX_READ    GpuP2PCapsIndex = MXSMLEX_P2P_CAPS_INDEX_READ
	P2P_CAPS_INDEX_WRITE   GpuP2PCapsIndex = MXSMLEX_P2P_CAPS_INDEX_WRITE
	P2P_CAPS_INDEX_NVLINK  GpuP2PCapsIndex = MXSMLEX_P2P_CAPS_INDEX_MXLINK
	P2P_CAPS_INDEX_ATOMICS GpuP2PCapsIndex = MXSMLEX_P2P_CAPS_INDEX_ATOMICS
	P2P_CAPS_INDEX_PCI     GpuP2PCapsIndex = MXSMLEX_P2P_CAPS_INDEX_PROP
	P2P_CAPS_INDEX_PROP    GpuP2PCapsIndex = MXSMLEX_P2P_CAPS_INDEX_PROP
	P2P_CAPS_INDEX_UNKNOWN GpuP2PCapsIndex = MXSMLEX_P2P_CAPS_INDEX_UNKNOWN
)

type GpuP2PStatus = MxSmlExGpuP2PStatus

const (
	P2P_STATUS_OK                         GpuP2PStatus = MXSMLEX_P2P_STATUS_OK
	P2P_STATUS_CHIPSET_NOT_SUPPORED       GpuP2PStatus = MXSMLEX_P2P_STATUS_CHIPSET_NOT_SUPPORTED
	P2P_STATUS_CHIPSET_NOT_SUPPORTED      GpuP2PStatus = MXSMLEX_P2P_STATUS_CHIPSET_NOT_SUPPORTED
	P2P_STATUS_GPU_NOT_SUPPORTED          GpuP2PStatus = MXSMLEX_P2P_STATUS_GPU_NOT_SUPPORTED
	P2P_STATUS_IOH_TOPOLOGY_NOT_SUPPORTED GpuP2PStatus = MXSMLEX_P2P_STATUS_IOH_TOPOLOGY_NOT_SUPPORTED
	P2P_STATUS_DISABLED_BY_REGKEY         GpuP2PStatus = MXSMLEX_P2P_STATUS_DISABLED_BY_REGKEY
	P2P_STATUS_NOT_SUPPORTED              GpuP2PStatus = MXSMLEX_P2P_STATUS_NOT_SUPPORTED
	P2P_STATUS_UNKNOWN                    GpuP2PStatus = MXSMLEX_P2P_STATUS_UNKNOWN
)

type Pstates = MxSmlExPstates

const (
	PSTATE_0       Pstates = MXSMLEX_PSTATE_0
	PSTATE_1       Pstates = MXSMLEX_PSTATE_1
	PSTATE_2       Pstates = MXSMLEX_PSTATE_2
	PSTATE_3       Pstates = MXSMLEX_PSTATE_3
	PSTATE_4       Pstates = MXSMLEX_PSTATE_4
	PSTATE_5       Pstates = MXSMLEX_PSTATE_5
	PSTATE_6       Pstates = MXSMLEX_PSTATE_6
	PSTATE_7       Pstates = MXSMLEX_PSTATE_7
	PSTATE_8       Pstates = MXSMLEX_PSTATE_8
	PSTATE_9       Pstates = MXSMLEX_PSTATE_9
	PSTATE_10      Pstates = MXSMLEX_PSTATE_10
	PSTATE_11      Pstates = MXSMLEX_PSTATE_11
	PSTATE_12      Pstates = MXSMLEX_PSTATE_12
	PSTATE_13      Pstates = MXSMLEX_PSTATE_13
	PSTATE_14      Pstates = MXSMLEX_PSTATE_14
	PSTATE_15      Pstates = MXSMLEX_PSTATE_15
	PSTATE_UNKNOWN Pstates = MXSMLEX_PSTATE_UNKNOWN
)

type GpuTopologyLevel = MxSmlExGpuTopologyLevel

const (
	TOPOLOGY_INTERNAL   GpuTopologyLevel = MXSMLEX_TOPOLOGY_INTERNAL
	TOPOLOGY_SINGLE     GpuTopologyLevel = MXSMLEX_TOPOLOGY_SINGLE
	TOPOLOGY_MULTIPLE   GpuTopologyLevel = MXSMLEX_TOPOLOGY_MULTIPLE
	TOPOLOGY_HOSTBRIDGE GpuTopologyLevel = MXSMLEX_TOPOLOGY_HOSTBRIDGE
	TOPOLOGY_NODE       GpuTopologyLevel = MXSMLEX_TOPOLOGY_NODE
	TOPOLOGY_SYSTEM     GpuTopologyLevel = MXSMLEX_TOPOLOGY_SYSTEM
)

type IntNvLinkDeviceType = MxSmlExMetaXLinkDeviceType

const (
	NVLINK_DEVICE_TYPE_GPU     IntNvLinkDeviceType = MXSMLEX_METAXLINK_DEVICE_TYPE_GPU
	NVLINK_DEVICE_TYPE_IBMNPU  IntNvLinkDeviceType = MXSMLEX_METAXLINK_DEVICE_TYPE_NPU
	NVLINK_DEVICE_TYPE_SWITCH  IntNvLinkDeviceType = MXSMLEX_METAXLINK_DEVICE_TYPE_SWITCH
	NVLINK_DEVICE_TYPE_UNKNOWN IntNvLinkDeviceType = MXSMLEX_METAXLINK_DEVICE_TYPE_UNKNOWN
)

type ProcessInfo = MxSmlExProcessInfo
type FieldValue = MxSmlExFieldValue
type PciInfo = MxSmlExPciInfo
type Memory = MxSmlExMemory
type Utilization = MxSmlExUtilization

type ExtensionWrapper struct {
}

var extensionWrapper = createExtensionWrapper()

func createExtensionWrapper() *ExtensionWrapper {
	return &ExtensionWrapper{}
}

func getErrorString(r Return) string {
	return utils.BytePtrToStr(mxSmlExErrorString(r))
}

func (r Return) String() string {
	return r.Error()
}

func (r Return) Error() string {
	return getErrorString(r)
}

func (n *ExtensionWrapper) ErrorString(r Return) string {
	return r.Error()
}

func (n *ExtensionWrapper) Init() Return {
	err := lib.Load()
	if err != nil {
		return ERROR_LIBRARY_NOT_FOUND
	}

	return mxSmlExInit()
}

func (n *ExtensionWrapper) DeviceGetCount() (int, Return) {
	var devCount uint32
	ret := mxSmlExDeviceGetCount(&devCount)
	return int(devCount), ret
}

func (n *ExtensionWrapper) ShutDown() Return {
	mxSmlExShutdown()
	if lib.Unload() == nil {
		return SUCCESS
	}

	return ERROR_UNKNOWN
}

func (n *ExtensionWrapper) SystemGetDriverVersion() (string, Return) {
	var version [SYSTEM_DRIVER_VERSION_BUFFER_SIZE]byte
	var size uint32 = SYSTEM_DRIVER_VERSION_BUFFER_SIZE

	ret := mxSmlExSystemGetDriverVersion(&version[0], size)
	if ret != SUCCESS {
		return "", ret
	}

	return string(version[:size]), SUCCESS
}

func (n *ExtensionWrapper) SystemGetNVMLVersion() (string, Return) {
	var version [SYSTEM_NVML_VERSION_BUFFER_SIZE]byte
	var size uint32 = SYSTEM_NVML_VERSION_BUFFER_SIZE

	ret := mxSmlExSystemGetMxsmlVersion(&version[0], size)
	if ret != SUCCESS {
		return "", ret
	}

	return string(version[:size]), SUCCESS
}

func (n *ExtensionWrapper) DeviceGetHandleByIndex(index int) (Device, Return) {
	var device MxSmlExDevice
	ret := mxSmlExGetDeviceHandleByIndex(uint32(index), &device)
	return device, ret
}

func (n *ExtensionWrapper) DeviceGetHandleByUUID(uuid string) (Device, Return) {
	if len(uuid) == 0 {
		return nil, ERROR_INVALID_ARGUMENT
	}

	var device MxSmlExDevice
	uuidBytes := []byte(uuid)
	ret := mxSmlExDeviceGetHandleByUUID(&uuidBytes[0], &device)
	return device, ret
}

func (n *ExtensionWrapper) DeviceGetHandleByPciBusId(pciBusId string) (Device, Return) {
	if len(pciBusId) == 0 {
		return nil, ERROR_INVALID_ARGUMENT
	}

	var device MxSmlExDevice
	ret := mxSmlExGetDeviceHandleByPciBusId(pciBusId, &device)
	return device, ret
}

func (n *ExtensionWrapper) DeviceGetName(device Device) (string, Return) {
	return device.GetName()
}

func (n *ExtensionWrapper) DeviceGetUUID(device Device) (string, Return) {
	return device.GetUUID()
}

func (n *ExtensionWrapper) DeviceGetPciInfo(device Device) (PciInfo, Return) {
	return device.GetPciInfo()
}

func (n *ExtensionWrapper) DeviceGetMemoryInfo(device Device) (Memory, Return) {
	return device.GetMemoryInfo()
}

func (n *ExtensionWrapper) DeviceGetFanSpeed(device Device) (uint32, Return) {
	return device.GetFanSpeed()
}

func (n *ExtensionWrapper) DeviceGetFanSpeed_v2(device Device, fan int) (uint32, Return) {
	return device.GetFanSpeed_v2(fan)
}

func (n *ExtensionWrapper) DeviceGetUtilizationRates(device Device) (Utilization, Return) {
	return device.GetUtilizationRates()
}

func (n *ExtensionWrapper) DeviceGetMinorNumber(device Device) (int, Return) {
	return device.GetMinorNumber()
}

func (n *ExtensionWrapper) DeviceGetNvLinkState(device Device, link int) (EnableState, Return) {
	return device.GetNvLinkState(link)
}

func (n *ExtensionWrapper) DeviceGetNvLinkRemotePciInfo(device Device, link int) (PciInfo, Return) {
	return device.GetNvLinkRemotePciInfo(link)
}

func (n *ExtensionWrapper) DeviceGetTemperature(device Device, sensorType TemperatureSensors) (uint32, Return) {
	return device.GetTemperature(sensorType)
}

func (n *ExtensionWrapper) DeviceGetCurrPcieLinkWidth(device Device) (int, Return) {
	return device.GetCurrPcieLinkWidth()
}

func (n *ExtensionWrapper) DeviceGetComputeRunningProcesses(device Device) ([]ProcessInfo, Return) {
	return device.GetComputeRunningProcesses()
}

func (n *ExtensionWrapper) DeviceGetCudaComputeCapability(device Device) (int, int, Return) {
	return device.GetCudaComputeCapability()
}

func (n *ExtensionWrapper) DeviceGetPowerUsage(device Device) (uint32, Return) {
	return device.GetPowerUsage()
}

func (n *ExtensionWrapper) DeviceGetPowerManagementLimit(device Device) (uint32, Return) {
	return device.GetPowerManagementLimit()
}

func (n *ExtensionWrapper) DeviceGetMaxPcieLinkWidth(device Device) (int, Return) {
	return device.GetMaxPcieLinkWidth()
}

func (n *ExtensionWrapper) DeviceGetPcieThroughput(device Device, counter PcieUtilCounter) (uint32, Return) {
	return device.GetPcieThroughput(counter)
}

func (n *ExtensionWrapper) DeviceGetTemperatureThreshold(device Device, thresholdType TemperatureThresholds) (uint32, Return) {
	return device.GetTemperatureThreshold(thresholdType)
}

func (n *ExtensionWrapper) DeviceGetClockInfo(device Device, clockType ClockType) (uint32, Return) {
	return device.GetClockInfo(clockType)
}

func (n *ExtensionWrapper) DeviceGetCurrentClocksThrottleReasons(device Device) (uint64, Return) {
	return device.GetCurrentClocksThrottleReasons()
}

func (n *ExtensionWrapper) DeviceGetSupportedClocksThrottleReasons(device Device) (uint64, Return) {
	return device.GetSupportedClocksThrottleReasons()
}

func (n *ExtensionWrapper) DeviceGetCpuAffinity(device Device, numCPUs int) ([]uint, Return) {
	return device.GetCpuAffinity(numCPUs)
}

func (n *ExtensionWrapper) DeviceGetSupportedMemoryClocks(device Device) (int, []uint32, Return) {
	return device.GetSupportedMemoryClocks()
}

func (n *ExtensionWrapper) DeviceGetSupportedGraphicsClocks(device Device, memoryClockMHz int) (int, []uint32, Return) {
	return device.GetSupportedGraphicsClocks(memoryClockMHz)
}

func (n *ExtensionWrapper) DeviceSetApplicationsClocks(device Device, memClockMHz, graphicsClockMHz uint32) Return {
	return device.SetApplicationsClocks(memClockMHz, graphicsClockMHz)
}

func (n *ExtensionWrapper) DeviceResetApplicationsClocks(device Device) Return {
	return device.ResetApplicationsClocks()
}

func (n *ExtensionWrapper) DeviceGetApplicationsClock(device Device, clockType ClockType) (uint32, Return) {
	return device.GetApplicationsClock(clockType)
}

func (n *ExtensionWrapper) DeviceGetP2PStatus(device1 Device, device2 Device, p2pIndex GpuP2PCapsIndex) (GpuP2PStatus, Return) {
	return device1.GetP2PStatus(device2, p2pIndex)
}

func (n *ExtensionWrapper) DeviceGetPerformanceState(device Device) (Pstates, Return) {
	return device.GetPerformanceState()
}

func (n *ExtensionWrapper) DeviceSetAutoBoostedClocksEnabled(device Device, enabled EnableState) Return {
	return device.SetAutoBoostedClocksEnabled(enabled)
}

func (n *ExtensionWrapper) DeviceGetAutoBoostedClocksEnabled(device Device) (EnableState, EnableState, Return) {
	return device.GetAutoBoostedClocksEnabled()
}

func (n *ExtensionWrapper) DeviceGetTopologyCommonAncestor(device1 Device, device2 Device) (GpuTopologyLevel, Return) {
	return device1.GetTopologyCommonAncestor(device2)
}

func (n *ExtensionWrapper) DeviceGetNvLinkRemoteDeviceType(device Device, link int) (IntNvLinkDeviceType, Return) {
	return device.GetNvLinkRemoteDeviceType(link)
}

func (n *ExtensionWrapper) DeviceGetFieldValues(device Device, values []FieldValue) Return {
	return device.GetFieldValues(values)
}
