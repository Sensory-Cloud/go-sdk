package crypto_service_test

import (
	"testing"

	crypto_service "github.com/Sensory-Cloud/go-sdk/pkg/service/crypto"
)

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"short string", args{2}, 2, false},
		{"long string", args{64}, 64, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := crypto_service.GenerateRandomString(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRandomString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("GenerateRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}
