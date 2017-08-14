// HTTP GET의 모의 호출을 사용하는 예제
// 책에서 사용한 예제와는 다소 다른 부분이 있다.
package listing12

import (
//	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// feed 변수에는 우리가 기대하는 모의 응답 데이터를 대입한다.
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Goint Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    <item>
</channel>
</rss>`

// mockServer 함수는 GET 요청을 처리할 서버에 대한 포인터를 리턴한다.
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-TYpe", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload 함수는 HTTP GET 요청을 이용해 콘텐츠를 다운로드한 후
// 해당 콘텐츠를 언마샬링할 수 있는지 확인한다.
func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close();

	t.Log("콘텐츠 다운로드 기능 테스트 시작.")
	{
		t.Logf("\tURL \"%s\" 호출 시 상태 코드가 \"%d\"인지 확인.",
			server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\thTTP GET 요청을 보냈는지 확인.",
					ballotX, err)
			}
			t.Log("\t\tHTTP GET 요청을 보냈는지 확인",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\t상태코드가 \"%d\" 인지 확인. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\t상태 코드가 \"%d\" 인지 확인. %v",
				statusCode, checkMark)
		}
	}
}
