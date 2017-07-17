// 예제 4.49  nil 맵에 값을 대입할 때 발생하는 런타임 에러

// 맵을 선언하기만 하면 `nil` 맵이 생성된다.
var colors map[string]string

// Red 색상의 색상코드를 맵에 추가한다.
colors["Red"] = "#da1337"


// runtime error:
// panic: runtime error: assignment to entry in nil map
