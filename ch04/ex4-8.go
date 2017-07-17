// 예제 4.8  다른 타입의 배열에 대입하려는 경우에는 컴파일 에러가 발생한다

// 네개의 원소를 가지는 문자열 배열을 선언한다.
var array1 [4]string

// 다섯 개의 원소를 가지는 두 번째 배열을 선언한다.
// 이 배열을 색상을 표현하는 문자열로 초기화한다.
array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

// array2의 값을 array1로 복사한다.
array1 = array2


Compile Error:
cannot use array2 (type [5]string) as type [4]string in assignment
