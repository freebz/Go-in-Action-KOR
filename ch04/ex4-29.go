// 예제 4.29  인덱스가 범위를 벗어난 경우에 발생하는 런타임 예외

// 정수 슬라이스를 생성한다.
// 길이 및 용량이 5로 지정된다.
slice := []int{10, 20, 30, 40, 50}

// 새로운 슬라이스를 생성한다.
// 길이는 2, 용량은 4로 지정된다.
newSlice := slice[1:3]

// `newSlice` 슬라이스의 인덱스 3의 값을 변경한다.
// 이 원소는 `newSlice`슬라이스에 존재하지 않는다.
newSlice[3] = 45


Runtime Exception:
panic: runtime error: index out of range
