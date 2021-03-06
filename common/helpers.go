package common

func Unique(intSlice []uint) []uint {
	keys := make(map[uint]bool)
	var list []uint
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func GetIdFromCtx(id interface{}) uint {
	idFloat, _ := id.(float64)
	return uint(idFloat)
}