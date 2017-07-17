// 예제 4.54  맵을 함수에 전달하기

func main() {
	// 색상과 색상 코드를 저장하는 맵을 생성한다.
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	// 맵에 저장된 모든 색상을 출력한다.
	for key, value := range colors {
		fmt.Printf("키: %s  값: %s\n", key, value)
	}

	// 특정 키를 삭제하는 함수를 호출한다.
	removeColor(colors, "Coral")

	// 맵에 저장된 모든 색상을 출력한다.
	for key, value := range colors {
		fmt.Printf("키: %s  값: %s\n", key, value)
	}
}

// removeColor 함수는 맵에서 저장된 키를 제거한다.
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
