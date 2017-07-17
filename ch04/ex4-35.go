// 예제 4.35  너무 큰 용량을 지정했을 때 발생하는 런타임 에러

// 이 예제에서는 용량을 4로 지정하려고 한다.
// 이 값은 현재 남아있는 여분의 용량보다 큰 값이다.
slice := source[2:3:6]


Runtime Error:
panic: runtime error: slice bounds out of range
