// "Copyright (c) 2026 MetaX Integrated Circuits (Shanghai) Co., Ltd. All rights reserved."
package main

import (
	"fmt"
	"sync"
	"log"

	"github.com/MetaX-MACA/go-mxsml/pkg/mxsml"
	"github.com/MetaX-MACA/go-mxsml/pkg/utils"
)

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

	count := mxsml.MxSmlGetDeviceCount()
	fmt.Printf("device count: %d\n", count)

	// --- getDeviceEid Start ---
	var wg sync.WaitGroup
	wg.Add(int(count))
	for deviceId := range count {
		go func() {
			defer wg.Done()
			getDeviceEid(deviceId)
		}()
	}

	fmt.Println("event set waiting...")
	wg.Wait()
	// --- getDeviceEid End ---
}

func getDeviceEid(id uint32) {
	eventSet, ret := mxsml.MxSmlEventSetCreate()
	if ret != mxsml.MXSML_Success {
		fmt.Printf("create event set failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	}

	ret = mxsml.MxSmlDeviceRegisterEvents(id, mxsml.MXSML_mxSmlEventTypeEid, eventSet)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("register event set failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	}

	data, ret := mxsml.MxSmlEventSetWait(eventSet)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("wait event set failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	} else {
		fmt.Printf("device id: %d, eid: 0x%x\nevent data: %s",
					data.DeviceId, data.Type, utils.ExtractInt8ArrayValidString(data.EventData[:]))
	}

	ret = mxsml.MxSmlEventSetFree(eventSet)
	if ret != mxsml.MXSML_Success {
		fmt.Printf("free event set failed: %s\n", mxsml.MxSmlGetErrorString(ret))
	}
}