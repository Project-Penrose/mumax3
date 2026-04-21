// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package cu

import (
	"fmt"
	"testing"
)

func TestVersion(t *testing.T) {
	fmt.Println("CUDA driver version: ", Version())
}
