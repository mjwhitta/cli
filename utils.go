package cli

import "strings"

func less(i int, j int) bool {
	var left string = flags[i].long
	var right string = flags[j].long

	// Sort by long flag, unless it is not defined
	if left == "" {
		left = flags[i].short
	}

	if right == "" {
		right = flags[j].short
	}

	// Fallback to short flag if comparing same long flag (should
	// not happen)
	if strings.EqualFold(left, right) {
		if flags[i].short != "" {
			left = flags[i].short
		}

		if flags[j].short != "" {
			right = flags[j].short
		}
	}

	// Check again b/c values may have changed
	if strings.EqualFold(left, right) {
		return left < right
	}

	return strings.ToLower(left) < strings.ToLower(right)
}
