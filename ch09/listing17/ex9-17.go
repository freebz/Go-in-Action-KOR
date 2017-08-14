// 간단한 웹 서비스 예제
package main

import (
	"log"
	"net/http"

	"github.com/webgenie/go-in-action/chapter9/listing17/handlers"
)

// 애플리케이션 진입점
func main() {
	handlers.Routes()

	log.Println("웹 서비스 실행 중: 포트: 4000")
	http.ListenAndServe(":4000", nil)
}
