// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package gui

import "fmt"

type meter struct {
	data
}

func (e *meter) update(id string) []jsCall {
	return []jsCall{{F: "setAttr", Args: []interface{}{id, "value", e.value()}}}
}

func (d *Page) Meter(id string, min, max, value int, extra ...string) string {
	e := &meter{data: data{value}}
	d.addElem(id, e)
	return fmt.Sprintf(`<meter id=%v min=%v max=%v></meter>`, id, min, max)
}
