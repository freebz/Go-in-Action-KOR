// 예제 4.18  길이보다 작은 크기의 용량을 지정할 경우에 발생하는 컴파일러 에러

// 정수 슬라이스를 생성한다.
// 용량보다 길이의 값을 크게 지정한다.
slice := make([]int, 5, 3)


Compile Error:
len larger than cap in make([]int)
