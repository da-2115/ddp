package main

func unwrap[T any](val T, err error) (T) {
	if err != nil {
		panic(err)
	}
	return val
}
