package open

import "testing"

func TestOpen(t *testing.T) {
	type args struct {
		things []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Open URLs",
			args: args{
				things: []string{"https://google.com", "https://github.com"},
			},
		},
		{
			name: "Open files",
			args: args{
				things: []string{".editorconfig", "README.md"},
			},
		},
		{
			name: "Open mixed",
			args: args{
				things: []string{".editorconfig", "https://github.com"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Open(tt.args.things)
		})
	}
}
