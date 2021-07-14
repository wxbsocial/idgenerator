package models

import (
	"testing"
)

type FixedTimeGetter struct {
}

var BASE_TIMESTAMP int64 = 1622476800000

func (getter *FixedTimeGetter) Now() int64 {
	return 1626187855000 - BASE_TIMESTAMP
}

func TestSignMask(t *testing.T) {

	t.Logf("%02X", POSITIVE_SIGN_MASK)
}

func TestGenerator_genIdWithoutSeq(t *testing.T) {
	timeGetter := &FixedTimeGetter{}
	type fields struct {
		Config     *Config
		TimeGetter TimeGetter
	}
	type args struct {
		nowTs int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				Config: NewConfig(
					1, 1, 5, 5, 12, BASE_TIMESTAMP, 10,
				),
			},
			args: args{timeGetter.Now()},
			want: ((timeGetter.Now()-int64(BASE_TIMESTAMP))<<22 | 1<<17 | 1<<12) & POSITIVE_SIGN_MASK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gen := NewGenerator(tt.fields.Config, timeGetter)
			if got := gen.genIdWithoutSeq(tt.args.nowTs); got != tt.want {
				t.Errorf("Generator.genIdWithoutSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}
