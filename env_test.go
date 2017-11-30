package gotools

import "testing"
import "os"

func TestMustGetEnv(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantEnv string
		wantErr bool
	}{
		{
			name: "get a env",
			args: args{
				name: "ENV_NAME",
			},
			wantEnv: "env_value",
			wantErr: false,
		},
		{
			name: "get null env",
			args: args{
				name: "TEST_NULL",
			},
			wantEnv: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == false {
				os.Setenv(tt.args.name, tt.wantEnv)
			}
			gotEnv, err := MustGetEnv(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("MustGetEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEnv != tt.wantEnv {
				t.Errorf("MustGetEnv() = %v, want %v", gotEnv, tt.wantEnv)
			}
		})
	}
}
