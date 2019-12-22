package main

import (
	"os"
	"testing"
)

func Test_isDirExist(t *testing.T) {
	type args struct {
		dirName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Dir exists",
			args: args{dirName: ".vscode"},
			want: true,
		},
		{
			name: "Dir not exists",
			args: args{dirName: "not-exist"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDirExist(tt.args.dirName); got != tt.want {
				t.Errorf("isDirExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseArgsAndExecute(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	//defer func() { os.RemoveAll("git-clone-open-code") }()
	os.Args = []string{"cmd", "https://github.com/rsrini7/git-clone-open-code"}
	// append(os.Args,"https://github.com/rsrini7/git-clone-open-code")

	tests := []struct {
		name string
		want bool
	}{
		{
			name: "e2e main function",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseArgsAndExecute(); got != tt.want {
				t.Errorf("parseArgsAndExecute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openInVSCode(t *testing.T) {
	type args struct {
		folder string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "open in code",
			args: args{folder: ".vscode"},
			want: true,
		},
		{
			name: "open in code",
			args: args{folder: "no folder"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openInVSCode(tt.args.folder); got != tt.want {
				t.Errorf("openInVSCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDir(t *testing.T) {
	type args struct {
		dirName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "remove dir",
			args: args{dirName: "./git-clone-open-code"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removeDir(tt.args.dirName)
		})
	}
}
