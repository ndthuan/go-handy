package env_test

import (
	"github.com/ndthuan/go-handy/env"
	"os"
	"testing"
)

type envVar struct {
	name  string
	value string
}

func TestGet(t *testing.T) {
	type args struct {
		key      string
		fallback string
	}

	tests := []struct {
		name string
		env  envVar
		args args
		want string
	}{
		{
			name: "Undefined env key without fallback",
			env:  envVar{name: "YYY", value: "ZZZ"},
			args: args{key: "XXX"},
			want: "",
		},
		{
			name: "Undefined env key with fallback",
			env:  envVar{name: "YYY", value: "ZZZ"},
			args: args{key: "XXX", fallback: "yyy"},
			want: "yyy",
		},
		{
			name: "Defined env key",
			env:  envVar{name: "YYY", value: "ZZZ"},
			args: args{key: "YYY", fallback: "yyy"},
			want: "ZZZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.Setenv(tt.env.name, tt.env.value)
			if err != nil {
				t.Error("Error setting env for testing: " + err.Error())
			}
			if got := env.Get(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		key      string
		fallback int
	}
	tests := []struct {
		name string
		env  envVar
		args args
		want int
	}{
		{
			name: "Undefined env key without fallback",
			env:  envVar{name: "YYY", value: "ZZZ"},
			args: args{key: "XXX"},
			want: 0,
		},
		{
			name: "Undefined env key with fallback",
			env:  envVar{name: "YYY", value: "ZZZ"},
			args: args{key: "XXX", fallback: 1},
			want: 1,
		},
		{
			name: "Defined env key",
			env:  envVar{name: "YYY", value: "ZZZ"},
			args: args{key: "YYY", fallback: 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.Setenv(tt.env.name, tt.env.value)
			if err != nil {
				t.Error("Error setting env for testing: " + err.Error())
			}
			if got := env.ToInt(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
