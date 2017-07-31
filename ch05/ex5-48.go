// 인터페이스의 다형성을 보여주기 위한 예제
package main

import (
	"fmt"
)

// 알림 동작을 정의하는
// notifier 인터페이스를 선언한다.
type notifier interface {
	notify()
}

// 사용자를 표현하는 user 타입을 선언한다.
type user struct {
	name  string
	email string
}

// 포인터 수신자를 이용하여 notify 메서드를 구현한다.
func (u *user) notify() {
	fmt.Printf("사용자에게 메일을 전송합니다: %s<%s>\n",
		u.name,
		u.email)
	
