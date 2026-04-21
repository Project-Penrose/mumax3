// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package script

func Contains(tree, search Expr) bool {
	if tree == search {
		return true
	} else {
		children := tree.Child()
		for _, e := range children {
			if Contains(e, search) {
				return true
			}
		}
	}
	return false
}
