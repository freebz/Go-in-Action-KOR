// 예제 4.14  함수에 커다란 배열을 값으로 전달하는 예제

// 8메가바이트 크기의 배열을 선언한다.
var array [1e6]int

// 배열을 foo 함수에 전달한다.
foo(array)

// 백만 개의 정수 배열을 매개변수로 전달받는 함수.
func foo(array [1e6]int) {
	//...
}
