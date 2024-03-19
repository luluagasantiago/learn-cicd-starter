package auth

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {

	tests := map[string]struct {
		input      http.Header
		wantString string
		wantErr    error
	}{
		"No Header": {input: http.Header{}, wantString: "", wantErr: ErrNoAuthHeaderIncluded},
		"Malforded": {input: http.Header{"Authorization": {"NotApiKey"}}, wantString: "", wantErr: MalformedAuthHeader},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotString, gotErr := GetAPIKey(tc.input)
			diffString := cmp.Diff(tc.wantString, gotString)

			// check the error first
			// if we've got an error, Expeced out string doesn't matter
			if !errors.Is(tc.wantErr, gotErr) {
				fmt.Printf("%#v, %#v", tc.wantErr, gotErr)

				t.Fatalf("%v != %v", tc.wantErr, gotErr)
			}

			if diffString != "" {
				t.Fatalf(diffString)
			}

		})

	}

}
