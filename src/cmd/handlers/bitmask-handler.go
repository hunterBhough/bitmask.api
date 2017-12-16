package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/hunterBhough/bitmask.api/src/restapi/operations"
	"github.com/alejandroq/locator.hiv.gov/src/services"
)


func PostAnswersHandlerFunc(params operations.PostAnswers) middleware.Responder {
	mutParams := services.Params{
		Geohash: params.Geohash,
	}
	return handle(mutParams)
}

func handle (params services.Params) middleware.Responder {

}