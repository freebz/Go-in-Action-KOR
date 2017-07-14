// 예제 3.6  빈 식별자를 이용한 가져오기

package main

import (
	"database/sql"

	_ "githib.com/webgenie/go-in-action/chapter3/dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "mydb")
}
