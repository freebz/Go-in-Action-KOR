// 예제 3.5  init 함수의 사용 예

package postgres

import (
	"database/sql"
)

func init() {
	sql.Register("postgres", new (PostgresDriver))
}
