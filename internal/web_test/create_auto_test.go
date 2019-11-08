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

var _ = Describe("Check create auto handler", func() {

	var r *mux.Router
	var w *httptest.ResponseRecorder
	serviceName := "autoteka"

	ctx := context.Background()
	db, mock, _ := sqlmock.New()

	xdb := sqlx.NewDb(db, "fake")

	ctx = metadata.SetContextValues(ctx, serviceName, xdb)

	handlerCreate := createAutoHandler(ctx)

	pathPost := "/autos"

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO auto (id, brand, model, engine_volume) VALUES(?, ?, ?, ?)`)).WithArgs(
		1, "skoda", "yeti", 1.2).WillReturnResult(sqlmock.NewResult(1, 0))

	BeforeEach(func() {
		r = mux.NewRouter()
	})
	Describe("Create autos", func() {

		Context("Provide a valid object", func() {
			It("should send via POST valid JSON and get HTTP Status: 200", func() {
				commandJson := `{"id":1,"brand":"skoda","model":"yeti","engine_volume":1.2}`
				reqUrl := "/autos"
				testCommand(r, w, handlerCreate, "POST", pathPost, reqUrl, commandJson, 200, true)
			})

		})

		Context("Provide invalid objects", func() {
			It("should send via POST empty JSON and get HTTP Status: 400", func() {
				commandJson := `{}`
				reqUrl := "/autos"
				testCommand(r, w, handlerCreate, "POST", pathPost, reqUrl, commandJson, 400, false)
			})

			It("should send via POST empty body and get HTTP Status: 400", func() {
				commandJson := ``
				reqUrl := "/autos"
				testCommand(r, w, handlerCreate, "POST", pathPost, reqUrl, commandJson, 400, false)
			})

			It("emulating conn.Exec error and get HTTP Status: 400", func() {
				commandJson := `{"id":1,"brand":"skoda","model":"yeti","engine_volume":1.8}`
				reqUrl := "/autos"
				testCommand(r, w, handlerCreate, "POST", pathPost, reqUrl, commandJson, 400, false)
			})
		})

	})
})

func createAutoHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(handlers.CreateAuto(ctx))
}
