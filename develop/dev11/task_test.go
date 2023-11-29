package main

import "testing"

func Test_myTelnet(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "опять тестить",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := myTelnet(); (err != nil) != tt.wantErr {
				t.Errorf("myTelnet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
