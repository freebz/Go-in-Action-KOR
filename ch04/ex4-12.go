// 예제 4.12  동일한 타입의 다차원 배열을 대입하기

// 두 개의 정수 배열을 가지는 이차원 배열을 두 개 선언한다.
var array1 [2][2]int
var array2 [2][2]int

// 각 원소에 정수 값을 대입한다.
array2[0][0] = 10
array2[0][1] = 20
array2[1][0] = 30
array2[1][1] = 40

// array2의 값들을 array1으로 복사한다.
array1 = array2
