package web_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/yetialex/autoteka/internal/metadata"
	"github.com/yetialex/autoteka/internal/web/handlers"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Check delete auto handler", func() {

	var r *mux.Router
	var w *httptest.ResponseRecorder
	serviceName := "autoteka"

	ctx := context.Background()
	db, mock, _ := sqlmock.New()

	xdb := sqlx.NewDb(db, "fake")

	ctx = metadata.SetContextValues(ctx, serviceName, xdb)

	handlerDelete := deleteAutoHandler(ctx)

	pathDelete := "/autos/{id}"

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM auto where id=?`)).WithArgs(
		1).WillReturnResult(sqlmock.NewResult(0, 1))

	BeforeEach(func() {
		r = mux.NewRouter()
	})
	Describe("Delete autos", func() {

		Context("Provide a valid object", func() {
			It("should send via DELETE valid object ID and get HTTP Status: 200", func() {
				commandJson := ``
				reqUrl := "/autos/1"
				testCommand(r, w, handlerDelete, "DELETE", pathDelete, reqUrl, commandJson, 200, true)
			})

		})

		Context("Provide invalid objects", func() {
			It("should send via DELETE wrong url and get HTTP Status: 400", func() {
				commandJson := ``
				reqUrl := "/autos/qwe"
				testCommand(r, w, handlerDelete, "DELETE", pathDelete, reqUrl, commandJson, 400, false)
			})

			It("emulating conn.Exec error via DELETE unmocked id and get HTTP Status: 400", func() {
				commandJson := ``
				reqUrl := "/autos/2"
				testCommand(r, w, handlerDelete, "DELETE", pathDelete, reqUrl, commandJson, 400, false)
			})

		})

	})
})

func deleteAutoHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(handlers.DeleteAuto(ctx))
}
