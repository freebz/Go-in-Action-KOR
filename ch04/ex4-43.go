// 예제 4.43  슬라이스의 슬라이스 다루기

// 정수의 슬라이스의 슬라이스를 생성한다.
slice := [][]int{{10}, {100, 200}}

// 값 20을 첫 번째 정수 슬라이스에 추가한다.
slice[0] = append(slice[0], 20)
