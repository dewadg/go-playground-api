package store

import (
	"testing"
)

func Test_generateIDs(t *testing.T) {
	type args struct {
		numOfIDs int
		length   int
	}
	tests := []struct {
		name       string
		args       args
		assertWant func([]string)
	}{
		{
			name: "success",
			args: args{
				numOfIDs: 5,
				length:   5,
			},
			assertWant: func(ids []string) {
				tracker := make(map[string]int)

				for _, id := range ids {
					count, exists := tracker[id]
					if !exists {
						tracker[id] = 0
					} else {
						tracker[id] = count + 1
						if count > 1 {
							t.Error("duplicate id detected", id)
						}
					}
				}

				t.Log(ids)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateIDs(tt.args.numOfIDs, tt.args.length)
			tt.assertWant(got)
		})
	}
}
