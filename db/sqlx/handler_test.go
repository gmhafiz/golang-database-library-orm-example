package sqlx

import (
	"net/http"
	"testing"
)

func Test_handler_Create(t *testing.T) {
	type fields struct {
		db *database
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				db: tt.fields.db,
			}
			h.Create(tt.args.w, tt.args.r)
		})
	}
}
