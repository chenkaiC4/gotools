package gotools

import "testing"

func TestIsFileExist(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "file exists",
			args: args{
				filename: "file.go",
			},
			want: true,
		},
		{
			name: "file exists",
			args: args{
				filename: "file_test.go",
			},
			want: true,
		},
		{
			name: "file doesn't exist",
			args: args{
				filename: "foo.go",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFileExist(tt.args.filename); got != tt.want {
				t.Errorf("IsFileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
