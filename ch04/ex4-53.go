// 예제 4.53  맵에서 아이템 제거하기

// 키 "Coral"에 해당하는 키와 값의 쌍을 제거한다.
delete(colors, "Coral")

// 맵에 저장된 모든 색상을 출력한다.
for key, value := range colors {
	fmt.Printf("키: %s  값: %s\n", key, value)
}
