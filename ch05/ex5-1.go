// 시스템에 user 타입을 선언한다.
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

// user 타입의 변수를 선언한다.
var bill user

// user 타입의 변수를 선언하고 모든 필드를 초기화한다.
lisa := user {
	name:       "Lisa"
	email:      "lisa@email.com"
	ext:        123,
	privileged: true,
}

// 권한이 할당된 관리자 계정을 포표현하 admin 타입
type admin {
	person user
	level  string
}

// admin 타입의 변수를 선언한다.
fred := admin {
	person: user {
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123
		privileged: true,
	},
	level: "super",
}
