// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package cu

import (
	"fmt"
)

// needed for all other tests.
func init() {
	Init(0)
	ctx := CtxCreate(CTX_SCHED_AUTO, 0)
	CtxSetCurrent(ctx)
	fmt.Println("Created CUDA context")
}
