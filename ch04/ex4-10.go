// 예제 4.10  이차원 배열 선언하기

// 두 개의 요소를 가지는 네 개의 이차원 정수 배열을 선언한다.
var array [4][2]int

// 배열 리터럴을 이용하여 이차원 정수 배열을 초기화한다.
array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

// 바깥 배열의 인덱스 1과 3을 초기화한다.
array := [4][2]int{1: {20, 21}, 3: {40, 41}}

// 바깥 배열과 내부 배열의 요소들을 선언과 동시에 초기화한다.
array := [4][2]int{1: {0: 20}, 3: {1: 41}}
