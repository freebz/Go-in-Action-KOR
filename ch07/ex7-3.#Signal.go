// 운영체제의 신호를 표현하기 위한 타입
// 실제 구현은 운영체제에 따라 다르다.
// UNIX 환경의 신호는 syscall.Signal에 구현되어 있다.
type Signal interface {
	String() string
	Signal() // 다른 문자열과 구분하기 위한 함수
}
