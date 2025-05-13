package util

// easy func to replace
// err := function()
// if err != nil {
//    panic(err)
// }
// with unwrap(function())
// like rust's .unwrap()

func Unwrap[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
