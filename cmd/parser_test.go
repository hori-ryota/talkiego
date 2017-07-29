package cmd

import (
	"reflect"
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_parseProps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "multiline",
			args: args{
				s: ` invert
          backface=https://farm8.staticflickr.com/7469/16209884502_211d01ac0d_z_d.jpg
          backfaceFilter="blur(3px) brightness(.9)"
					`,
			},
			want: map[string]string{
				"invert":         "",
				"backface":       "https://farm8.staticflickr.com/7469/16209884502_211d01ac0d_z_d.jpg",
				"backfaceFilter": "blur(3px) brightness(.9)",
			},
		},
		{
			name: "escaped double quote",
			args: args{
				s: `backfaceFilter="blu\"r(3px) brightness(.9)"`,
			},
			want: map[string]string{
				"backfaceFilter": `blu\"r(3px) brightness(.9)`,
			},
		},
		{
			name: "contain double quote",
			args: args{
				s: `backfaceFilter=blu"r`,
			},
			want: map[string]string{
				"backfaceFilter": `blu"r`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseProps(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseProps() = %v, want %v", got, tt.want)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
