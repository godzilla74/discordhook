package discordhook

import (
	"testing"
)

func Test_sendMessage(t *testing.T) {
	type args struct {
		url     string
		message Message
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMessage(tt.args.url, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("sendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
