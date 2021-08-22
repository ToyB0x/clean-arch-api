package store_test

import (
	"reflect"
	"testing"

	"github.com/toaru/clean-arch-api/cmd"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store"
)

func init() {
	cmd.GetConfigs("../../../../../config")
}

func TestNewSqlHandler(t *testing.T) {
	type args struct {
		projectID string
	}
	tests := []struct {
		name string
		args args
		want *store.SqlHandler
	}{
		{
			"success",
			args{projectID: "test"},
			&store.SqlHandler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := store.NewSqlHandler(tt.args.projectID); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewSqlHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
