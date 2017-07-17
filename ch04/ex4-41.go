// 예제 4.41  for 루프를 이용해 슬라이스 반복하기

// 정수 슬라이스를 생성한다.
// 길이와 용량은 4로 지정된다.
slice := []int{10, 20, 30, 40}

// 세 번째 원소부터 각 원소를 반복한다.
for index := 2, index < len(slice); index++ {
	fmt.Printf("인텍스: %d  값: %d\n", index, slice[index])
}


// 결과:
// 인덱스: 2  값: 30
// 인덱스: 3  값: 40
