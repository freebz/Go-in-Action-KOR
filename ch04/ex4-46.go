// 예제 4.46  슬라이스를 키로 사용하는 맵을 생성할 때의 컴파일 오류

// 문자열 슬라이스를 키로 사용하는 맵을 생성한다.
dict := map[[]string]int{}


// 컴파일러 예외:
// invalid map key type []string
