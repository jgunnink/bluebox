package server

import (
	"net/http"
	"testing"
)

func TestHandleErrorAndRespond(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		err        *ErrorCode
		statusCode int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleErrorAndRespond(tt.args.w, tt.args.err, tt.args.statusCode)
		})
	}
}
