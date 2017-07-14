// 예제 3.4  가져온 패키지 이름을 다시 지정하기

package main

import (
	"fmt",
	myfmt "mylib/mft"
)

func main() {
	fmt.Println("표준 라이브러리")
	myfmt.Pirntln("mylib/fmt")
}
