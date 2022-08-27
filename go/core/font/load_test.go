package font

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFace(t *testing.T) {
	type args struct {
		t    Type
		size int
	}
	tests := []struct {
		args    args
		wantErr bool
	}{
		{args{MonoSpaced, 20}, false},
		{args{MonoSpaced, 100}, false},
		{args{Propotional, 20}, false},
		{args{Propotional, 100}, false},
		{args{Type(10), 20}, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%d", tt.args.t, tt.args.size), func(t *testing.T) {
			assert := assert.New(t)
			got, err := LoadFace(tt.args.t, tt.args.size)
			if tt.wantErr {
				assert.Error(err)
				t.Log(err)
				return
			}
			assert.NoError(err)
			assert.NotNil(got)
		})
	}
}
