// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package gui

// concatenate elements
func cat(s []string) string {
	str := ""
	for _, s := range s {
		str += s + " "
	}
	return str
}
