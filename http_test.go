package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if len(name) == 0 {
		fmt.Fprintf(writer, "Hello World")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=kojek", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println("Result : ", bodyString)
}
