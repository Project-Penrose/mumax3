// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package cuda

import (
	"github.com/mumax/3/data"
	"github.com/mumax/3/util"
)

// Normalize vec to unit length, unless length or vol are zero.
func Normalize(vec, vol *data.Slice) {
	util.Argument(vol == nil || vol.NComp() == 1)
	N := vec.Len()
	cfg := make1DConf(N)
	k_normalize_async(vec.DevPtr(X), vec.DevPtr(Y), vec.DevPtr(Z), vol.DevPtr(0), N, cfg)
}
