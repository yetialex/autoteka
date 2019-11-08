package handlers

import (
	"context"
	"net/http"

	"github.com/yetialex/autoteka/internal/metadata"
	"github.com/yetialex/autoteka/internal/web/kit"
	"github.com/yetialex/autoteka/internal/web/model"
)

func CreateAuto(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		autotekaRequest, err := model.AutoFromRequest(r)
		if err != nil {
			kit.WriteErrorResponse(w, http.StatusBadRequest, []kit.ErrorItem{{
				Code:          "number",
				Message:       "parse request parameter error",
				MessageParams: nil,
			}}, err.Error())
			return
		}
		if autotekaRequest.ID == 0 {
			kit.WriteErrorResponse(w, http.StatusBadRequest, []kit.ErrorItem{{
				Code:          "id",
				Message:       "field is empty",
				MessageParams: nil,
			}}, kit.ErrInvalidRequestParams.Error())
			return
		}

		conn := metadata.ContextDbConnection(ctx)

		_, err = conn.Exec("INSERT INTO auto (id, brand, model, engine_volume) VALUES(?, ?, ?, ?)",
			autotekaRequest.ID, autotekaRequest.Brand, autotekaRequest.Model, autotekaRequest.EngineVolume)

		if err != nil {
			kit.WriteErrorResponse(w, http.StatusBadRequest, []kit.ErrorItem{{
				Code:          "query",
				Message:       "insert query exec error: %v",
				MessageParams: []interface{}{err},
			}}, "create auto error")
			return
		}

		kit.WriteSuccessResponse(w, nil, "created auto with id: %d", autotekaRequest.ID)
	}
}
