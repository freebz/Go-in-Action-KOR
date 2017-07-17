// 예제 4.52  for range 키워드를 이용해 맵을 반복하기

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
