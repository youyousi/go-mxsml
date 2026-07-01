// "Copyright (c) 2026 MetaX Integrated Circuits (Shanghai) Co., Ltd. All rights reserved."
package main

import (
	"fmt"
	"log"

	"github.com/MetaX-MACA/go-mxsml/pkg/mxsml"
)

func getProcessInfo(deviceCount uint32) {
	processNum, ret := mxsml.MxSmlGetNumberOfProcess()
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get number of process failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("number of process: %d\n", processNum)
	}

	if processNum <= 0 {
		return
	}

	processInfo := make([]mxsml.MxSmlProcessInfo, processNum)
	ret = mxsml.MxSmlGetProcessInfo(processNum, &processInfo[0])
	if ret == mxsml.MXSML_Success {
		for _, info := range processInfo {
			fmt.Printf("process number: %d, name: %s, gpu number: %d\n",
				info.ProcessId, uint8SliceToString(info.ProcessName[:]), info.GpuNumber)
		}
	}

	processInfo_v2 := make([]mxsml.MxSmlProcessInfo_v2, processNum)
	ret = mxsml.MxSmlGetProcessInfo_v2(processNum, &processInfo_v2[0])
	if ret == mxsml.MXSML_Success {
		for _, info := range processInfo_v2 {
			fmt.Printf("process number: %d, name: %s, gpu number: %d\n",
				info.ProcessId, uint8SliceToString(info.ProcessName[:]), info.GpuNumber)
		}
	}

	processInfo_v3 := make([]mxsml.MxSmlProcessInfo_v3, processNum)
	ret = mxsml.MxSmlGetProcessInfo_v3(processNum, &processInfo_v3[0])
	if ret == mxsml.MXSML_Success {
		for _, info := range processInfo_v3 {
			fmt.Printf("process number: %d, name: %s, gpu number: %d\n",
				info.ProcessId, uint8SliceToString(info.ProcessName[:]), info.GpuNumber)
		}
	}

	var id uint32
	for ; id < deviceCount; id++ {
		processInfo, ret := mxsml.MxSmlGetSingleGpuProcess(id)
		if ret == mxsml.MXSML_Success {
			for _, info := range processInfo {
				fmt.Printf("device: %d, process number: %d, name: %s, gpu number: %d\n",
					id, info.ProcessId, uint8SliceToString(info.ProcessName[:]), info.GpuNumber)
			}
		}

		processInfo_v2, ret := mxsml.MxSmlGetSingleGpuProcess_v2(id)
		if ret == mxsml.MXSML_Success {
			for _, info := range processInfo_v2 {
				fmt.Printf("device: %d, process number: %d, name: %s, gpu number: %d\n",
					id, info.ProcessId, uint8SliceToString(info.ProcessName[:]), info.GpuNumber)
			}
		}

		processInfo_v3, ret := mxsml.MxSmlGetSingleGpuProcess_v3(id)
		if ret == mxsml.MXSML_Success {
			for _, info := range processInfo_v3 {
				fmt.Printf("device: %d, process number: %d, name: %s, gpu number: %d\n",
					id, info.ProcessId, uint8SliceToString(info.ProcessName[:]), info.GpuNumber)
			}
		}
	}
}

func getDeviceStats(id uint32) {
	state, ret := mxsml.MxSmlGetDeviceState(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get device state failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("device state: %d\n", state)
	}

	memoryInfo, ret := mxsml.MxSmlGetMemoryInfo(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get memory info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("memory info: %+v\n", memoryInfo)
	}

	temperature, ret := mxsml.MxSmlGetTemperatureInfo(id, mxsml.MXSML_Temperature_Hotspot)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get hotspot temperature failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("hotspot temperature: %.2f C\n", float32(temperature)/100.0)
	}

	boardPowerInfo, ret := mxsml.MxSmlGetBoardPowerInfo(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get board power info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("board power info: %+v\n", boardPowerInfo)
	}

	pmBusInfo, ret := mxsml.MxSmlGetPmbusInfo(id, mxsml.MXSML_Pmbus_Core)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get core pmbus info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("core pmbus info: %+v\n", pmBusInfo)
	}

	xcoreClock, ret := mxsml.MxSmlGetClocks(id, mxsml.MXSML_Clock_Xcore)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get xcore clocks failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("xcore clocks: %+v\n", xcoreClock)
	}

	dpmClock, ret := mxsml.MxSmlGetDpmIpClockInfo(id, mxsml.MXSML_Dpm_Xcore)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get xcore dpm clock info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("xcore dpm clock info: %+v\n", dpmClock)
	}

	vddInfo, ret := mxsml.MxSmlGetDpmIpVddInfo(id, mxsml.MXSML_Dpm_Xcore)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get xcore vdd info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("xcore vdd info: %+v\n", vddInfo)
	}

	dpmPerfLevel, ret := mxsml.MxSmlGetCurrentDpmIpPerfLevel(id, mxsml.MXSML_Dpm_Xcore)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get xcore dpm perf level: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("xcore dpm perf level: %+v\n", dpmPerfLevel)
	}

	ret = mxsml.MxSmlSetDpmIpMaxPerfLevel(id, mxsml.MXSML_Dpm_Xcore, 7)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("set dpm max perf level failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	}

	dpmMaxPerfLevel, ret := mxsml.MxSmlGetDpmIpMaxPerfLevel(id, mxsml.MXSML_Dpm_Xcore)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get xcore dpm max perf level failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("xcore dpm max perf level: %+v\n", dpmMaxPerfLevel)
	}

	serial, ret := mxsml.MxSmlGetChipSerial(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get chip serial failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("chip serial: %s\n", serial)
	}

	pptableVer, ret := mxsml.MxSmlGetPptableVersion(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get pptable version failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("pptable version: %s\n", pptableVer)
	}

	mxlkStateCode, mxlkState, ret := mxsml.MxSmlGetMetaXLinkState(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get metaxlink state failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("metaxlink code: %d, state: %s\n", mxlkStateCode, mxlkState)
	}

	devRealPath, ret := mxsml.MxSmlGetDeviceRealPath(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get device real path failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("device real path: %s\n", devRealPath)
	}

	xcoreUsage, ret := mxsml.MxSmlGetDeviceIpUsage(id, mxsml.MXSML_Usage_Xcore)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get device xcore usage failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("device xcore usage: %+v\n", xcoreUsage)
	}

	eepromInfo, ret := mxsml.MxSmlGetEepromInfo(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get eeprom info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("eeprom info: %s\n", uint8SliceToString(eepromInfo.Content[:]))
	}

	dirverVer, ret := mxsml.MxSmlGetDeviceVersion(id, mxsml.MXSML_Version_Driver)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get driver version failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("driver version: %s\n", dirverVer)
	}

	rasErrorData, ret := mxsml.MxSmlGetRasErrorData(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get ras error data failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		printRasErrorData(&rasErrorData)
	}

	rasErrorDataV2, ret := mxsml.MxSmlGetRasErrorData_v2(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get ras error data v2 failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		printRasErrorData_v2(&rasErrorDataV2)
	}

	rasStatusData, ret := mxsml.MxSmlGetRasStatusData(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get ras status data failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		printRasStatusData(&rasStatusData)
	}

	pcieInfo, ret := mxsml.MxSmlGetPcieInfo(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get pcie info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("pcie info: %+v\n", pcieInfo)
	}

	mxlkPcieInfo, ret := mxsml.MxSmlGetPcieMaxLinkInfo(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get pcie metaxlink info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("pcie metaxlink info: %+v\n", mxlkPcieInfo)
	}

	powerStateInfo, ret := mxsml.MxSmlGetPowerStateInfo(id, mxsml.MXSML_Dpm_Xcore)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get xcore power state info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("xcore power state info: %+v\n", powerStateInfo)
	}

	pciPowerState, ret := mxsml.MxSmlGetPciPowerState(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get pci power state info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("pci power state info: %+v\n", pciPowerState)
	}

	codeStatus, ret := mxsml.MxSmlGetCodecStatus(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get code status failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("code status info: %+v\n", codeStatus)
	}

	hbmBandwidth, ret := mxsml.MxSmlGetHbmBandWidth(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get hbm bandwidth failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("hbm bandwidth: %+v\n", hbmBandwidth)
	}

	pcieThoughout, ret := mxsml.MxSmlGetPcieThroughput(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get pcie thoughout failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("pcie thoughout: %+v\n", pcieThoughout)
	}

	dmBandwidth, ret := mxsml.MxSmlGetDmaBandwidth(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get dma bandwidth failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("dma bandwidth: %+v\n", dmBandwidth)
	}

	loglevel, ret := mxsml.MxSmlGetFwIpLoglevel(id, mxsml.MXSML_Fw_IpName_SMP0)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get smp0 log level failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("smp0 log level: %d\n", loglevel)
	}

	ret = mxsml.MxSmlSetFwLoglevel(id, mxsml.MXSML_Fw_IpName_SMP0, mxsml.MXSML_Loglevel_WARN)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("set smp0 log level to warn failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		loglevel, ret = mxsml.MxSmlGetFwIpLoglevel(id, mxsml.MXSML_Fw_IpName_SMP0)
		if ret == mxsml.MXSML_Success {
			fmt.Printf("get smp0 log level: %d\n", loglevel)
		}
	}

	delay, ret := mxsml.MxSmlGetPciDelay(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get pci delay failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("pci delay: %d\n", delay)
	}

	if mxsml.MxSmlSetPciDelay(id, 200) == mxsml.MXSML_Success {
		fmt.Println("set pci delay with 200 succeeded")
		if delay, ret = mxsml.MxSmlGetPciDelay(id); ret == mxsml.MXSML_Success {
			fmt.Printf("pci delay: %d\n", delay)
		}
	}

	mxlkBandwidth, ret := mxsml.MxSmlGetMetaXLinkBandwidth(id, mxsml.MXSML_MetaXLink_Input)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get metaxlink input bandwidth failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("metaxlink input bandwidth: %+v\n", mxlkBandwidth)
	}

	boardSerial, ret := mxsml.MxSmlGetBoardSerial(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get board serial failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("board serial: %s\n", boardSerial)
	}

	maxPerfLevel, ret := mxsml.MxSmlGetDpmMaxPerfLevel(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get dpm max perf level failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("dpm max perf level: %d\n", maxPerfLevel)
	}

	apUsageToggle, ret := mxsml.MxSmlGetApUsageToggle(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get apusage toggle failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("apusage toggle: %d\n", apUsageToggle)
	}

	if ret := mxsml.MxSmlSetApUsageToggle(id, 1); ret == mxsml.MXSML_Success {
		fmt.Println("set ap usage toggle to 1 succeeded")
		if apUsageToggle, ret = mxsml.MxSmlGetApUsageToggle(id); ret == mxsml.MXSML_Success {
			fmt.Printf("apusage toggle: %d\n", apUsageToggle)
		}
	} else {
		fmt.Printf("set apusage toggle to 1 failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	}

	logLevel, ret := mxsml.MxSmlGetFwLoglevel(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get fw log lovel failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("fw log lovel: %+v\n", logLevel)
	}

	xcoreApUsage, ret := mxsml.MxSmlGetXcoreApUsage(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get xcore ap usage failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("xcore ap usage: %+v\n", xcoreApUsage)
	}

	vpuStatus, ret := mxsml.MxSmlGetVpuStatus(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get vpu status failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("vpu status: %+v\n", vpuStatus)
	}

	cpuSets, ret := mxsml.MxSmlGetCpuAffinity(id, 5)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get cpu affinity failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("cpu affinity: %+v\n", cpuSets)
	}

	nodeSets, ret := mxsml.MxSmlGetNodeAffinity(id, 5)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get node affinity failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("node affinity: %+v\n", nodeSets)
	}

	eccState, ret := mxsml.MxSmlGetEccState(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get ecc state failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("ecc state: %+v\n", eccState)
	}

	mxlkLinkInfo, ret := mxsml.MxSmlGetMetaXLinkInfo(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get metaxlink info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("metaxlink info: %+v\n", mxlkLinkInfo)
	}

	mxlkLinkInfoV2, ret := mxsml.MxSmlGetMetaXLinkInfo_v2(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get metaxlink info v2 failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("metaxlink info v2: %+v\n", mxlkLinkInfoV2)
	}

	// get metaxlink aer count
	mxlkAer, ret := mxsml.MxSmlGetMetaXLinkAer(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get metaxlink aer failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Println("metaxlink aer count:")
		for i, info := range mxlkAer {
			fmt.Printf("  port: %d, ue: %d, ce: %d\n", i+1, info.UeAer, info.CeAer)
		}
	}

	sriovState, ret := mxsml.MxSmlGetSriovState(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get sriov state failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("sriov state: %d\n", sriovState)
	}

	slot, ret := mxsml.MxSmlGetDeviceSlot(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get device slot failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("device slot: %d\n", slot)
	}

	omStatus, ret := mxsml.MxSmlGetOpticalModuleStatus(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get optical module status failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("optical module status: %+v\n", omStatus)
	}

	omInfo, ret := mxsml.MxSmlGetOpticalModuleInfo(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get optical module info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("optical module info: %+v\n", omInfo)
	}

	throttleReason, ret := mxsml.MxSmlGetCurrentClocksThrottleReason(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get current clocks throttle reason failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("current clocks throttle reason: %d\n", throttleReason)
	}

	boardPowerLimit, ret := mxsml.MxSmlGetBoardPowerLimit(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get board power limit failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("board power limit: %d\n", boardPowerLimit)
	}

	portState, ret := mxsml.MxSmlGetMetaXLinkPortState(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get metaxlink power state failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("metaxlink power state: %+v\n", portState)
	}

	eventInfo, ret := mxsml.MxSmlGetPciEventInfo(id, mxsml.MXSML_Pci_Event_AER_UE)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get pcie event aer ue info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Println("pcie event aer ue info:")
		printEventInfo(eventInfo)
	}

	isaVer, ret := mxsml.MxSmlGetDeviceIsaVersion(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get isa version failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("isa version: %d\n", isaVer)
	}

	reason, ret := mxsml.MxSmlGetDeviceUnavailableReason(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get device unavailable reason failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("device unavailable reason: %+v\n", reason)
	}

	eccErrors, ret := mxsml.MxSmlGetTotalEccErrors(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get total ecc errors failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("total ecc errors: %+v\n", eccErrors)
	}

	fmt.Println("MetaxLink remote info:")
	for i := range mxsml.MXSML_METAX_LINK_NUM + 1 {
		if remoteInfo, ret := mxsml.MxSmlGetMetaXLinkRemoteInfo(id, uint32(i)); ret == mxsml.MXSML_Success {
			fmt.Printf("PORT#%d,device id: %d, bdf id: %s\n", i, remoteInfo.DeviceId, uint8SliceToString(remoteInfo.BdfId[:]))
		}
	}

	remoteInfo, ret := mxsml.MxSmlGetMetaXLinkRemoteInfo(id, 4)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get remoteInfo failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("remoteInfo of device %d, link 4: %+v\n", id, remoteInfo)
	}

	linkTopo, ret := mxsml.MxSmlGetMetaXLinkTopo(id)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get metaxlink topo failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("metaxlink topo of device %d: %+v\n", id, linkTopo)
	}
}

func main() {
	ret := mxsml.MxSmlInit()
	if ret != mxsml.MXSML_Success {
		log.Fatal(mxsml.MxSmlGetErrorString(ret))
	}
	defer mxsml.MxSmlShutDown()

	macaVersion, ret := mxsml.MxSmlGetMacaVersion()
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get maca version failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("maca version: %s\n", macaVersion)
	}

	pfCount := mxsml.MxSmlGetPfDeviceCount()
	if pfCount > 0 {
		fmt.Printf("get pf device count: %d\n", pfCount)
		var pfId uint32 = 100
		for ; pfId < 100+pfCount; pfId++ {
			virtaulIds, ret := mxsml.MxSmlGetVirtualDevicesByPhysicalId(pfId)
			if ret == mxsml.MXSML_Success {
				fmt.Printf("get virtual devices Id: %+v\n", virtaulIds)
			}
		}
	}

	limitedDevices, ret := mxsml.MxSmlGetAllLimitedDevices()
	if ret == mxsml.MXSML_Success && limitedDevices.Number > 0 {
		fmt.Printf("get limited devices: %+v\n", limitedDevices)
		for number := range limitedDevices.Number {
			limitedDevice, ret := mxsml.MxSmlGetLimitedDeviceInfo(uint32(limitedDevices.DeviceId[number]))
			if ret != mxsml.MXSML_Success {
				fmt.Printf("get limited device info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
				continue
			}
			fmt.Printf("limited device Id: %d, type: %d, bdfId: %s, gpuId: %d, nodeId: %d, "+
				"uuid: %s, band: %d, mode: %d, device name: %s\n",
				limitedDevice.DeviceId, limitedDevice.Type, uint8SliceToString(limitedDevice.BdfId[:]),
				limitedDevice.GpuId, limitedDevice.NodeId, uint8SliceToString(limitedDevice.Uuid[:]),
				limitedDevice.Brand, limitedDevice.Mode, uint8SliceToString(limitedDevice.DeviceName[:]))
		}
	}

	count := mxsml.MxSmlGetDeviceCount()
	fmt.Printf("device count: %d\n", count)

	getProcessInfo(count)

	var i uint32 = 0
	for ; i < count; i++ {
		deviceInfo, ret := mxsml.MxSmlGetDeviceInfo(i)
		if ret != mxsml.MXSML_Success {
			log.Fatalf("%v\n", mxsml.MxSmlGetErrorString(ret))
		}
		fmt.Printf(" ------ print device %d stats start ------\n", deviceInfo.DeviceId)
		fmt.Println("devices info:")
		printDeviceInfo(&deviceInfo)

		getDeviceStats(deviceInfo.DeviceId)

		fmt.Println("Topo info: ")
		for dst := range count {
			topLevel, ret := mxsml.MxSmlGetDeviceTopology(deviceInfo.DeviceId, dst)
			if ret == mxsml.MXSML_Success {
				fmt.Printf("  ->GPU#%d: %d\n", dst, topLevel)
			}
		}

		fmt.Println("Distance: ")
		for dst := range count {
			distance, ret := mxsml.MxSmlGetDeviceDistance(deviceInfo.DeviceId, dst)
			if ret == mxsml.MXSML_Success {
				fmt.Printf("  ->GPU#%d: %d\n", dst, distance)
			}
		}

		dieCount, ret := mxsml.MxSmlGetDeviceDieCount(deviceInfo.DeviceId)
		if ret != mxsml.MXSML_Success {
			fmt.Printf("get die count failed: %s\n", mxsml.MxSmlGetErrorString(ret))
		} else {
			fmt.Printf("device die count: %d\n", dieCount)

			for dieId := range dieCount {
				getDeviceDieStats(deviceInfo.DeviceId, dieId)
			}
		}

		fmt.Printf(" ------ print device %d stats end ------\n", deviceInfo.DeviceId)
		fmt.Println()
	}

	opMode, ret := mxsml.MxSmlGetOpMode()
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get operation mode failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("operation mode: %d\n", opMode)
	}

	if opMode == 0 {
		if ret := mxsml.MxSmlSetOpMode(1); ret != mxsml.MXSML_Success {
			fmt.Printf("set operation mode to 1 failed: %s\n", mxsml.MxSmlGetErrorString(ret))
		} else {
			opMode, ret = mxsml.MxSmlGetOpMode()
			if ret == mxsml.MXSML_Success {
				fmt.Printf("operation mode: %d\n", opMode)
				mxsml.MxSmlSetOpMode(0)
			}
		}
	}
}

func getDeviceDieStats(deviceId, dieId uint32) {
	fmt.Printf("--- get stats for device %d die %d start ---\n", deviceId, dieId)
	driverVer, ret := mxsml.MxSmlGetDeviceDieVersion(deviceId, dieId, mxsml.MXSML_Version_Driver)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get driver vbios failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("driver vbios version: %s\n", driverVer)
	}

	pptableVer, ret := mxsml.MxSmlGetDiePptableVersion(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get pptable version failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("pptable version: %s\n", pptableVer)
	}

	serial, ret := mxsml.MxSmlGetDieChipSerial(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get die chip serial failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("die chip serial: %s\n", serial)
	}

	memoryInfo, ret := mxsml.MxSmlGetDieMemoryInfo(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get memory info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("memory info: %+v", memoryInfo)
	}

	temperature, ret := mxsml.MxSmlGetDieTemperatureInfo(deviceId, dieId, mxsml.MXSML_Temperature_Core)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get core temperature failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("core temperature: %d\n", temperature)
	}

	pmBusInfo, ret := mxsml.MxSmlGetDiePmbusInfo(deviceId, dieId, mxsml.MXSML_Pmbus_Core)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get core pmbus info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("core pmbus info: %+v\n", pmBusInfo)
	}

	xcoreClock, ret := mxsml.MxSmlGetDieClocks(deviceId, dieId, mxsml.MXSML_Clock_Xcore)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get xcore clocks failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("xcore clocks: %+v\n", xcoreClock)
	}

	eepromInfo, ret := mxsml.MxSmlGetDieEepromInfo(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get eeprom info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("eeprom info: %s\n", uint8SliceToString(eepromInfo.Content[:]))
	}

	powerStateInfo, ret := mxsml.MxSmlGetDiePowerStateInfo(deviceId, dieId, mxsml.MXSML_Dpm_Xcore)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get xcore power state info failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("xcore power state info: %+v\n", powerStateInfo)
	}

	hbmBandwidth, ret := mxsml.MxSmlGetDieHbmBandWidth(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get hbm bandwidth failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("hbm bandwidth: %+v\n", hbmBandwidth)
	}

	eccState, ret := mxsml.MxSmlGetDieEccState(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get ecc state failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("ecc state: %d\n", eccState)
	}

	apUsageToggle, ret := mxsml.MxSmlGetDieApUsageToggle(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get apusage toggle failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("apusage toggle: %d\n", apUsageToggle)
	}

	xcoreUsage, ret := mxsml.MxSmlGetDieIpUsage(deviceId, dieId, mxsml.MXSML_Usage_Xcore)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get die xcore usage failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("die xcore usage: %+v\n", xcoreUsage)
	}

	xcoreApUsage, ret := mxsml.MxSmlGetDieXcoreApUsage(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get die xcore ap usage failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("die xcore ap usage: %+v\n", xcoreApUsage)
	}

	dpmPerfLevel, ret := mxsml.MxSmlGetCurrentDieDpmIpPerfLevel(deviceId, dieId, mxsml.MXSML_Dpm_Dla)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get dpm perf level: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("dla dpm perf level: %+v\n", dpmPerfLevel)
	}

	dpmMaxPerfLevel, ret := mxsml.MxSmlGetDieDpmIpMaxPerfLevel(deviceId, dieId, mxsml.MXSML_Dpm_Xcore)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get xcore dpm max perf level failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("xcore dpm max perf level: %+v\n", dpmMaxPerfLevel)
	}

	codeStatus, ret := mxsml.MxSmlGetDieCodecStatus(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get code status failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("code status info: %+v\n", codeStatus)
	}

	rasStatusData, ret := mxsml.MxSmlGetDieRasStatusData(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get die ras status data failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		printRasStatusData(&rasStatusData)
	}

	rasErrorData, ret := mxsml.MxSmlGetDieRasErrorData(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get die ras error data failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		printRasErrorData(&rasErrorData)
	}

	remoteInfo, ret := mxsml.MxSmlGetDieMetaXLinkRemoteInfo(deviceId, dieId, 4)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get remoteInfo failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("remoteInfo of device %d, die %d, link 4: %+v\n", deviceId, dieId, remoteInfo)
	}

	reason, ret := mxsml.MxSmlGetDieUnavailableReason(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get die unavailable reason failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("die unavailable reason: %+v\n", reason)
	}

	eccErrors, ret := mxsml.MxSmlGetDieTotalEccErrors(deviceId, dieId)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("get die total ecc errors failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("die total ecc errors: %+v\n", eccErrors)
	}

	fmt.Printf("--- get stats for device %d die %d end ---\n", deviceId, dieId)
}

func uint8SliceToString(slice []int8) string {
	end := 0
	var tmpArr []uint8
	for ; end < len(slice); end++ {
		if slice[end] == 0 {
			break
		}

		tmpArr = append(tmpArr, uint8(slice[end]))
	}
	return string(tmpArr)
}

func printDeviceInfo(devInfo *mxsml.MxSmlDeviceInfo) {
	fmt.Printf("deviceId: %d, type: %d, bdfId: %s,"+
		"gpuId: %d, nodeId: %d, Uuid: %s, "+
		"brand: %d, mode: %d, DeviceName: %s\n",
		devInfo.DeviceId, devInfo.Type, uint8SliceToString(devInfo.BdfId[:]),
		devInfo.GpuId, devInfo.NodeId, uint8SliceToString(devInfo.Uuid[:]),
		devInfo.Brand, devInfo.Mode, uint8SliceToString(devInfo.DeviceName[:]))
}

func printRasErrorData(rasErrorData *mxsml.MxSmlRasErrorData) {
	fmt.Println("ras error data:")
	for i := range int(rasErrorData.ShowRasErrorSize) {
		fmt.Printf("ras ip: %d, ue: %d, ce: %d\n",
			rasErrorData.RasErrorRegister[i].RasIp,
			rasErrorData.RasErrorRegister[i].RasErrorUe,
			rasErrorData.RasErrorRegister[i].RasErrorCe,
		)
	}
}

func printRasErrorData_v2(rasStatusData *mxsml.MxSmlRasErrorData_v2) {
	fmt.Println("ras error data v2:")
	for i := range int(rasStatusData.ShowRasErrorSize) {
		fmt.Printf("ras ip: %d, ue: %d, ce: %d\n",
			rasStatusData.RasErrorRegister[i].RasIp,
			rasStatusData.RasErrorRegister[i].RasErrorUe,
			rasStatusData.RasErrorRegister[i].RasErrorCe,
		)
	}
}

func printRasStatusData(rasStatusData *mxsml.MxSmlRasStatusData) {
	fmt.Println("ras status data:")
	for i := range int(rasStatusData.ShowRasStatusSize) {
		fmt.Printf("ras ip: %d, index: %d, register data: 0x%x\n",
			rasStatusData.RasStatusRegister[i].RasIp,
			rasStatusData.RasStatusRegister[i].RegisterIndex,
			rasStatusData.RasStatusRegister[i].RegisterData,
		)
	}
}

func printEventInfo(eventInfo []mxsml.MxSmlPciEventInfo) {
	if len(eventInfo) == 0 {
		fmt.Println("  no event")
		return
	}

	for _, eventInfo := range eventInfo {
		fmt.Printf("  bit: %d, count: %d, first time: %s, name: %s\n",
			eventInfo.BitNumber, eventInfo.Count,
			uint8SliceToString(eventInfo.FirstTime[:]), uint8SliceToString(eventInfo.Name[:]))
	}
}
