package handlers

import (
	"context"
	"net/http"

	"github.com/yetialex/autoteka/internal/metadata"
	"github.com/yetialex/autoteka/internal/web/kit"
	"github.com/yetialex/autoteka/internal/web/model"
)

func UpdateAuto(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
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

		_, err = conn.Exec(`UPDATE auto set brand=?, model=?, engine_volume=? where id=?`,
			autotekaRequest.Brand, autotekaRequest.Model, autotekaRequest.EngineVolume, autotekaRequest.ID)

		if err != nil {
			kit.WriteErrorResponse(w, http.StatusBadRequest, []kit.ErrorItem{{
				Code:          "query",
				Message:       "update query exec error: %v",
				MessageParams: []interface{}{err},
			}}, "update auto error")
			return
		}

		kit.WriteSuccessResponse(w, nil, "updated auto with id: %d", autotekaRequest.ID)
	}
}
