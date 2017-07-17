// 예제 4.9  포인터 배열을 다른 배열에 복사하기

// 세개의 원소를 가지는 문자열 포인터 배열을 선언한다.
var array1 [3]*string

// 세개의 원소를 가지는 두 번째 문자열 포인터 배열을 선언한다.
// 배열을 문자열 포인터로 초기화한다.
array2 := [3]*string{new(string), new(string), new(string)}

// 각 원소에 색상을 표현하는 문자열을 대입한다.
*array[0] = "Red"
*array[1] = "Blue"
*array[2] = "Green"

// array2의 값들을 array1으로 복사한다.
array1 = array2
