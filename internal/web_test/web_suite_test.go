package web_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/yetialex/autoteka/internal/web/kit"

	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWeb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Web Suite")
}

func testCommand(r *mux.Router, w *httptest.ResponseRecorder, handler http.Handler,
	method, path, reqUrl, commandJson string, code int, success bool) []byte {
	r.Handle(path, handler).Methods(method)
	req, err := http.NewRequest(
		method,
		reqUrl,
		strings.NewReader(commandJson),
	)
	Expect(err).NotTo(HaveOccurred())
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	Expect(w.Code).To(Equal(code))
	body := w.Body.Bytes()
	message := kit.SuccessMessage{}
	_ = json.Unmarshal(body, &message)
	Expect(message.Success).To(Equal(success))
	return body
}
