// 예제 4.50  맵에서 값과 키의 존재 여부를 동시에 확인하는 방법

// 키 "Blue"의 값을 조회한다.
value, exists := colors["Blue"]

// 키가 존재하는지를 확인한다.
if exists {
	fmt.Println(value)
}
