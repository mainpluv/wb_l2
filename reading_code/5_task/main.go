package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

// выведется error, так как test() возвращает указатель на nil, следовательно,
// err снова не будет равна nil, так как будет содержать указатель на него.
