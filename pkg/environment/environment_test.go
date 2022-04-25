package environment_test

import (
	"testing"

	environment "github.com/Sensory-Cloud/go-sdk/pkg/environment"
	test_util "github.com/Sensory-Cloud/go-sdk/pkg/util/test"
	"github.com/google/uuid"
)

func TestGetSet(t *testing.T) {
	testKey := environment.CloudHost
	testValue := uuid.NewString()
	fallback := uuid.NewString()

	oldConfig := environment.Get(testKey, "")
	environment.Set(testKey, "")
	defer environment.Set(testKey, oldConfig)

	config := environment.Get(testKey, fallback)
	test_util.AssertEquals(t, config, fallback)

	err := environment.Set(testKey, testValue)
	test_util.AssertOk(t, err)

	config = environment.Get(testKey, fallback)
	test_util.AssertEquals(t, config, testValue)
}

func TestGetEnvironment(t *testing.T) {
	environment.Set(environment.GoEnv, "local")
	env := environment.GetEnvironment()
	test_util.Assert(t, env == environment.LocalEnv, "current environment should be set to LOCAL")
}

func TestGetEmptyValue(t *testing.T) {
	testKey := environment.CloudHost

	oldConfig := environment.Get(testKey, "")
	environment.Set(testKey, "")
	defer environment.Set(testKey, oldConfig)

	defaultValue := "default!"
	value := environment.Get(testKey, defaultValue)
	test_util.AssertEquals(t, value, defaultValue)
}

func TestIsTruthy(t *testing.T) {
	type args struct {
		key      environment.EnvironmentVariable
		fallback bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{environment.EnvironmentVariable("1"), false}, true},
		{"t", args{environment.EnvironmentVariable("t"), false}, true},
		{"T", args{environment.EnvironmentVariable("T"), false}, true},
		{"TRUE", args{environment.EnvironmentVariable("TRUE"), false}, true},
		{"True", args{environment.EnvironmentVariable("True"), false}, true},
		{"true", args{environment.EnvironmentVariable("true"), false}, true},
		{"0", args{environment.EnvironmentVariable("0"), true}, false},
		{"f", args{environment.EnvironmentVariable("f"), true}, false},
		{"F", args{environment.EnvironmentVariable("F"), true}, false},
		{"FALSE", args{environment.EnvironmentVariable("FALSE"), true}, false},
		{"False", args{environment.EnvironmentVariable("False"), true}, false},
		{"false", args{environment.EnvironmentVariable("false"), true}, false},
		{"test", args{environment.EnvironmentVariable(""), false}, false},
		{"test", args{environment.EnvironmentVariable(""), true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			environment.Set(tt.args.key, string(tt.args.key))
			if got := environment.IsTruthy(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("IsTruthy() = %v, want %v", got, tt.want)
			}
		})
	}
}
