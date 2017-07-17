// 예제 4.39  값의 복사본을 제공하는 range 키워드

// 정수 슬라이스를 생성한다.
// 길이와 용량은 4로 지정된다.
slice := []int{10, 20, 30, 40}

// 각 원소를 반복하면서 그 값과 주소를 출력한다.
for index, value := range slice {
	fmt.Printf("값: %d  값의 주소: %X  원소의 주소: %X\n", value, &value, &slice[index])
}


// 결과
// 값: 10  값의 주소: 10500168  원소의 주소: 1052E100
// 값: 20  값의 주소: 10500168  원소의 주소: 1052E104
// 값: 30  값의 주소: 10500168  원소의 주소: 1052E108
// 값: 40  값의 주소: 10500168  원소의 주소: 1052E10C
