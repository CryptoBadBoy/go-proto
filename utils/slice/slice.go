package slice

func ConvertTo32(src []byte) [32]byte {
	var array [32]byte
	copy(array[:], src[:32])
	return array
}
