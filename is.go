package abolish

// IsNumber string. should >= 0
func IsNumber(v any) bool {
	// get type with err check
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	}
	return false
}
