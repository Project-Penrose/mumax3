// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package cu

// This file implements CUDA driver version management

//#include <cuda.h>
import "C"

const CUDA_VERSION = C.CUDA_VERSION

// Returns the CUDA driver version.
func Version() int {
	var version C.int
	err := Result(C.cuDriverGetVersion(&version))
	if err != SUCCESS {
		panic(err)
	}
	return int(version)
}
