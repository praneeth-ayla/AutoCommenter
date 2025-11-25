package contextstore

// MapToSlice converts a map of string keys to FileDetails values into a slice of FileDetails.
func MapToSlice(m map[string]FileDetails) []FileDetails {
	// Initialize a slice with a capacity equal to the map's length for efficiency.
	out := make([]FileDetails, 0, len(m))
	// Iterate over the values of the map.
	for _, v := range m {
		// Append each FileDetails value to the slice.
		out = append(out, v)
	}
	// Return the resulting slice.
	return out
}