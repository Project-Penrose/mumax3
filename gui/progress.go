// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package gui

import "fmt"

type progress struct {
	data
}

func (e *progress) update(id string) []jsCall {
	return []jsCall{{F: "setAttr", Args: []interface{}{id, "value", e.value()}}}
}

func (d *Page) Progress(id string, max, value int, extra ...string) string {
	e := &progress{data: data{value}}
	d.addElem(id, e)
	return fmt.Sprintf(`<progress id=%v max=%v></progress>`, id, max)
}
