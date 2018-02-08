package main_test

import (
	"testing"

	main "github.com/hunterBhough/go-doge/src/cmd/go-doge-server"

	"github.com/stretchr/testify/assert"
	"github.com/hunterBhough/go-doge/src/models"
)

//var foodWallet = models.TransmissionType{
//	Name: "NNDSS Food borne",
//	//Records:
//}

func TestHandler(t *testing.T) {

	tests := []struct {
		expect  []models.TransmissionType
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			expect:  []models.TransmissionType{

			},
			err:     nil,
		},
		//{
		//	// Test that the handler responds ErrNameNotProvided
		//	// when no name is provided in the HTTP body
		//	expect:  "",
		//	err:     main.ErrNameNotProvided,
		//},
	}

	for _, test := range tests {
		response, err := main.Handler()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response)
	}

}
