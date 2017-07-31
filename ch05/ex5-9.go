// 예제 5.9  값을 다른 타입에 대입할 때 발생하는 커마일 오류

package main

type Duration int64

func main() {
	var dur Duration
	dur = int64(1000)
}


// prog.go:7 cannot use int64(1000) (type int64) as type Duration in assignment
