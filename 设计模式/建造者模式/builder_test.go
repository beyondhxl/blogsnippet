package builder

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResourcePoolCfgBuilder_Build(t *testing.T) {
	tests := []struct {
		name    string
		builder *ResourcePoolCfgBuilder
		want    *ResourcePoolCfg
		wantErr bool
	}{
		{
			name: "name empty",
			builder: &ResourcePoolCfgBuilder{
				name:     "",
				maxTotal: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "maxIdle < minIdle",
			builder: &ResourcePoolCfgBuilder{
				name:     "test",
				maxTotal: 0,
				maxIdle:  10,
				minIdle:  20,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			builder: &ResourcePoolCfgBuilder{
				name: "test",
			},
			want: &ResourcePoolCfg{
				name:     "test",
				maxTotal: MaxToTal,
				maxIdle:  MaxIdle,
				minIdle:  MinIdle,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get, err := tt.builder.Build()
			require.Equalf(t, tt.wantErr, err != nil, "Build() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, get)
		})
	}
}
