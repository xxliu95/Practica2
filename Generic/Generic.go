package Generic

func AnyOf[T any](slice []T, f func(T) bool) bool {
	for _, item := range slice {
		if f(item) {
			return true
		}
	}
	return false
}

func Equal[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

// func RemoveIf ([]T, func (T) bool) int {
// 	return
// }
