// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package main

import (
	"encoding/json"
	"io"

	"github.com/mumax/3/data"
)

func dumpJSON(f *data.Slice, info data.Meta, out io.Writer) {
	w := json.NewEncoder(out)
	w.Encode(f.Tensors())
}
