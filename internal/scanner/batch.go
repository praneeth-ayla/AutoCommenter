package scanner

func BatchByLines(files []Info, maxLines int) [][]Info {
	var result [][]Info
	var group []Info
	used := 0 // Tracks the total lines in the current group

	for _, f := range files {
		// If adding the current file exceeds maxLines and the group is not empty,
		// finalize the current group and start a new one.
		if used+f.Lines > maxLines && len(group) > 0 {
			result = append(result, group)
			group = nil
			used = 0
		}
		group = append(group, f)
		used += f.Lines
	}

	// Add the last group if it's not empty.
	if len(group) > 0 {
		result = append(result, group)
	}

	return result
}