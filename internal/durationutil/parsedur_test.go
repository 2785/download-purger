package durationutil

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name: "Success Case: minute",
			args: args{
				s: "5min",
			},
			want: time.Duration(5 * time.Minute),
		},
		{
			name: "Success Case: day",
			args: args{
				s: "1day",
			},
			want: time.Duration(24 * time.Hour),
		},
		{
			name: "Success Case: week",
			args: args{
				s: "2week",
			},
			want: time.Duration(2 * 24 * 7 * time.Hour),
		},
		{
			name: "Success Case: day",
			args: args{
				s: "3month",
			},
			want: time.Duration(3 * 30 * 24 * time.Hour),
		},
		{
			name: "Success Case: day",
			args: args{
				s: "4year",
			},
			want: time.Duration(4 * 365 * 24 * time.Hour),
		},
		{
			name: "Failure Case: garbage input",
			args: args{
				s: "not valid string",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTime(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
