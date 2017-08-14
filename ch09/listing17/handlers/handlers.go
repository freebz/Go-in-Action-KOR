// 웹 서비스의 종단점을 제공하는 패키지
package handlers

import (
	"encoding/json"
	"net/http"
)

// 웹 서비스의 라우트를 설정한다.
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON 함수는 간단한 JSON 문서를 리턴한다.
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "장현희",
		Email: "webgenie@email.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
