// 예제 4.36  길이와 용량을 동일하게 지정할 때의 장점

// 문자열 슬라이스를 생성한다.
// 길이 및 용량이 5로 지정된다.
source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

// 세 번째 원소를 잘라낸다. 이때 용량을 설정한다.
// 길이 및 용량이 1로 지정된다.
slice := source[2:3:3]

// 슬라이스에 새로운 문자열을 추가한다.
slice = append(slice, "Kiwi")
