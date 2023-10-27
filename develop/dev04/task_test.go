package main

import (
	"reflect"
	"testing"
)

func Test_getAnagrams(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[string]map[string]struct{}
	}{
		{
			name: "Проверка работает ли",
			args: args{words: []string{
				"мама", "амам", "лист", "слит",
			}},
			wantResult: map[string]map[string]struct{}{
				"аамм": {"мама": {}, "амам": {}},
				"илст": {"лист": {}, "слит": {}},
			},
		},

		{
			name: "Проверка на удаление одинарного элемента",
			args: args{words: []string{
				"яблоко", "груша", "апельсин",
			}},
			wantResult: map[string]map[string]struct{}{},
		},

		{
			name: "Проверка на повторения",
			args: args{words: []string{
				"мама", "амам", "лист", "слит", "сплит",
			}},
			wantResult: map[string]map[string]struct{}{
				"аамм": {"мама": {}, "амам": {}},
				"илст": {"лист": {}, "слит": {}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := getAnagrams(tt.args.words); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("getAnagrams() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
