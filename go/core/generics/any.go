package generics

func ToAnySlice[T any](ts []T) []interface{} {
	s := make([]interface{}, 0, len(ts))
	for _, e := range ts {
		s = append(s, e)
	}
	return s
}
