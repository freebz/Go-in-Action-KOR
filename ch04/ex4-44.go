// 예제 4.44  함수에 슬라이스 전달하기

// 백만 개의 정수를 갖는 슬라이스를 생성한다.
slice := make([]int, 1e6)

// 슬라이스를 함수 foo에 전달한다.
slice = foo(slice)

// foo 함수는 정수 슬라이스를 전달받아 다시 리턴한다.
func foo(slice []int) []int {
	//...
	return slice
}
