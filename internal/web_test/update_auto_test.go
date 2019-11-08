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

var _ = Describe("Check update auto handler", func() {

	var r *mux.Router
	var w *httptest.ResponseRecorder
	serviceName := "autoteka"

	ctx := context.Background()
	db, mock, _ := sqlmock.New()

	xdb := sqlx.NewDb(db, "fake")

	ctx = metadata.SetContextValues(ctx, serviceName, xdb)

	handlerUpdate := updateAutoHandler(ctx)

	pathPut := "/autos"

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE auto set brand=?, model=?, engine_volume=? where id=?`)).WithArgs(
		"skoda", "yeti", 1.2, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	BeforeEach(func() {
		r = mux.NewRouter()
	})
	Describe("Update autos", func() {

		Context("Provide a valid object", func() {
			It("should send via PUT valid JSON and get HTTP Status: 200", func() {
				commandJson := `{"id":1,"brand":"skoda","model":"yeti","engine_volume":1.2}`
				reqUrl := "/autos"
				testCommand(r, w, handlerUpdate, "PUT", pathPut, reqUrl, commandJson, 200, true)
			})

		})

		Context("Provide invalid objects", func() {
			It("should send via PUT empty JSON and get HTTP Status: 400", func() {
				commandJson := `{}`
				reqUrl := "/autos"
				testCommand(r, w, handlerUpdate, "PUT", pathPut, reqUrl, commandJson, 400, false)
			})

			It("should send via PUT empty body and get HTTP Status: 400", func() {
				commandJson := ``
				reqUrl := "/autos"
				testCommand(r, w, handlerUpdate, "POST", pathPut, reqUrl, commandJson, 400, false)
			})

			It("emulating conn.Exec error and get HTTP Status: 400", func() {
				commandJson := `{"id":1,"brand":"skoda","model":"yeti","engine_volume":1.8}`
				reqUrl := "/autos"
				testCommand(r, w, handlerUpdate, "PUT", pathPut, reqUrl, commandJson, 400, false)
			})
		})

	})
})

func updateAutoHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(handlers.UpdateAuto(ctx))
}
