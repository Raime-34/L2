package main

import "testing"

func Test_getPage(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "тупо тест",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getPage(); (err != nil) != tt.wantErr {
				t.Errorf("getPage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
