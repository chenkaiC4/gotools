package gotools

import (
	"os"
	"reflect"
	"testing"
)

const zipExt = ".zip"

func TestUnZipFile(t *testing.T) {
	type args struct {
		src  string
		dest string
	}
	var emptyArrayString []string
	filesToBeZipped := []string{"zip.go", "zip_test.go"}
	zipFile := "test-unzip"
	err := ZipFiles(zipFile+zipExt, filesToBeZipped)
	if err != nil {
		t.Errorf("ZipFiles() error = %v", err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "unzip a zip file",
			args: args{
				src:  zipFile + zipExt,
				dest: ".",
			},
			want:    filesToBeZipped,
			wantErr: false,
		},
		{
			name: "unzip file, but zip not exist",
			args: args{
				src:  "xxx.zip",
				dest: ".",
			},
			want:    emptyArrayString,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnZipFile(tt.args.src, tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unzip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unzip() = %v, want %v", got, tt.want)
			}
		})
	}
	os.Remove(zipFile)
	os.Remove(zipFile + zipExt)
}

func TestZipFiles(t *testing.T) {
	type args struct {
		filename string
		files    []string
	}
	zipName := "test"
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "zip files",
			args: args{
				filename: zipName + zipExt,
				files:    []string{"file.go", "file_test.go", "env.go"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ZipFiles(tt.args.filename, tt.args.files); (err != nil) != tt.wantErr {
				t.Errorf("ZipFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	os.Remove(zipName + zipExt)
}
