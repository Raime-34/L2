package main

import "testing"

func Test_ewfnwe(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Чек на дедлок",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ewfnwe(); (err != nil) != tt.wantErr {
				t.Errorf("ewfnwe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
