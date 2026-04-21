// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package data

// Holds meta data to be saved together with a slice.
// Typically winds up in OVF or DUMP header
type Meta struct {
	Name, Unit     string
	Time, TimeStep float64
	CellSize       [3]float64
	MeshUnit       string
}
