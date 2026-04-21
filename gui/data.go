// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package gui

type data struct {
	val interface{}
}

func (d *data) set(v interface{})  { d.val = v }
func (d *data) value() interface{} { return d.val }
