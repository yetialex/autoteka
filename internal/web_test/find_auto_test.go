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

var _ = Describe("Check find auto handler", func() {

	var r *mux.Router
	var w *httptest.ResponseRecorder
	serviceName := "autoteka"

	ctx := context.Background()
	db, mock, _ := sqlmock.New()

	xdb := sqlx.NewDb(db, "fake")

	ctx = metadata.SetContextValues(ctx, serviceName, xdb)

	handlerFind := findAutoHandler(ctx)

	pathGet := "/autos/{id}"

	columns := []string{"id", "brand", "model", "engine_volume"}

	mock.ExpectQuery(regexp.QuoteMeta(`select * from auto where id=?`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "skoda", "yeti", 1.2))

	BeforeEach(func() {
		r = mux.NewRouter()
	})
	Describe("Find autos", func() {

		Context("Provide a valid object", func() {
			It("should send via GET valid object ID and get HTTP Status: 200", func() {
				commandJson := ``
				reqUrl := "/autos/1"
				testCommand(r, w, handlerFind, "GET", pathGet, reqUrl, commandJson, 200, false)
			})

		})

		Context("Provide invalid objects", func() {
			It("should send via GET wrong url and get HTTP Status: 400", func() {
				commandJson := ``
				reqUrl := "/autos/qwe"
				testCommand(r, w, handlerFind, "GET", pathGet, reqUrl, commandJson, 400, false)
			})

			It("emulating conn.Exec error via GET unmocked id and get HTTP Status: 400", func() {
				commandJson := ``
				reqUrl := "/autos/2"
				testCommand(r, w, handlerFind, "GET", pathGet, reqUrl, commandJson, 400, false)
			})

		})

	})
})

func findAutoHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(handlers.FindAuto(ctx))
}
