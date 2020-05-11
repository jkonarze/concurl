package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name         string

		givenUrl string
		givenConc int
		givenPayload string
		givenMethod string
		givenHeaders []string

		wantNil bool
	}{
		{
			name: "should return not nil",

			givenUrl: "www.google.com",
			givenConc: 5,
			givenMethod: "get",

			wantNil: false,
		},
	}

	required := require.New(t)

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			svc := NewSvc(tc.givenConc,tc.givenUrl, tc.givenMethod, tc.givenPayload, tc.givenHeaders)

			required.NotNil(svc)
			required.Equal( "", svc.payload)
			required.Equal( 5, svc.concurrent)
			required.Equal( "get", svc.method)
		})
	}
}