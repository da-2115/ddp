package main
// easy func to replace 
// err := function()
// if err != nil {
//    panic(err)
// }
// with unwrap(function())
// like rust's .unwrap()

func unwrap[T any](val T, err error) (T) {
	if err != nil {
		panic(err)
	}
	return val
}
