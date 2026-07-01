// "Copyright (c) 2026 MetaX Integrated Circuits (Shanghai) Co., Ltd. All rights reserved."
package lib

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <stdlib.h>
import "C"

const mxsmlLibName = "libmxsml.so"

var g_mxsmlLib = newMxsmlLib()

type library struct {
	handle unsafe.Pointer
	mu     sync.Mutex
	loaded bool
	flag   int
}

func newMxsmlLib() *library {
	mxsmlLib := &library{
		flag: C.RTLD_LAZY | C.RTLD_GLOBAL,
	}

	return mxsmlLib
}

func Load() error {
	g_mxsmlLib.mu.Lock()
	defer g_mxsmlLib.mu.Unlock()

	if g_mxsmlLib.loaded {
		return nil
	}

	libName := C.CString(mxsmlLibName)
	defer C.free(unsafe.Pointer(libName))

	runtime.LockOSThread()
	handle := C.dlopen(libName, C.int(g_mxsmlLib.flag))
	runtime.UnlockOSThread()

	if handle != nil {
		g_mxsmlLib.handle = handle
		g_mxsmlLib.loaded = true
		return nil
	}

	var installPath string
	// check lib path
	libPathList := []string{
		"/opt/mxdriver/lib/" + mxsmlLibName,
		"/opt/maca/lib/" + mxsmlLibName,
		"/opt/mxn100/lib/" + mxsmlLibName,
	}
	for _, path := range libPathList {
		if _, err := os.Stat(path); err == nil {
			installPath = path
			break
		}
	}

	if len(installPath) == 0 {
		return fmt.Errorf("invalid mxsml lib path")
	}

	libPath := C.CString(installPath)
	defer C.free(unsafe.Pointer(libPath))

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	handle = C.dlopen(libPath, C.int(g_mxsmlLib.flag))
	if handle == nil {
		return getDlError()
	}

	g_mxsmlLib.handle = handle
	g_mxsmlLib.loaded = true
	return nil
}

func Unload() error {
	g_mxsmlLib.mu.Lock()
	defer g_mxsmlLib.mu.Unlock()

	if !g_mxsmlLib.loaded {
		return nil
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if C.dlclose(g_mxsmlLib.handle) != 0 {
		return getDlError()
	}

	g_mxsmlLib.handle = nil
	g_mxsmlLib.loaded = false
	return nil
}

func getDlError() error {
	lastErr := C.dlerror()
	if lastErr != nil {
		return errors.New(C.GoString(lastErr))
	}

	return nil
}
