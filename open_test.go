package open

import "testing"

func TestOpen(t *testing.T) {
	type args struct {
		obj string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Open URL",
			args: args{
				obj: "https://google.com",
			},
		},
		{
			name: "Open file",
			args: args{
				obj: "README.md",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Open(tt.args.obj)
		})
	}
}
