package handlers

import (
	"context"
	"net/http"

	"github.com/yetialex/autoteka/internal/metadata"
	"github.com/yetialex/autoteka/internal/web/kit"
	"github.com/yetialex/autoteka/internal/web/model"
)

func DeleteAuto(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
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

		_, err = conn.Exec("DELETE FROM auto where id=?", autotekaRequest.ID)

		if err != nil {
			kit.WriteErrorResponse(w, http.StatusBadRequest, []kit.ErrorItem{{
				Code:          "query",
				Message:       "delete query exec error: %v",
				MessageParams: []interface{}{err},
			}}, "delete auto error")
			return
		}

		kit.WriteSuccessResponse(w, nil, "deleted auto with id: %d", autotekaRequest.ID)

	}
}
