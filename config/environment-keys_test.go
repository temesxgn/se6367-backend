package config_test

import "testing"

import "os"

import "strconv"

func TestEnvironment(t *testing.T) {
	t.Parallel()
	if err := os.Setenv("foo", "bar"); err != nil {
		t.Errorf("Error setting env variable: %v", err)
	}

	if err := os.Setenv("num", strconv.Itoa(10)); err != nil {
		t.Errorf("Error setting env variable: %v", err)
	}

	tables := []struct {
		name       string
		key        string
		defaultVal string
		expected   string
		isString   bool
	}{
		{"Should find os value", "foo", "dog", "bar", true},
		{"Should NOT find os value and return default", "bazz", "dog", "dog", true},
		{"Should find os int value", "num", "3", "10", false},
	}

	for _, tt := range tables {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)
			if tt.isString {
				val := GetValue(tt.key, tt.defaultVal)
				if val != tt.expected {
					t.Errorf("Got %v, want %v", val, tt.expected)
				}
			} else {
				dVal, _ := strconv.ParseInt(tt.defaultVal, 10, 64)
				expected, _ := strconv.ParseInt(tt.expected, 10, 64)
				val := GetIntValue(tt.key, dVal)
				if val != expected {
					t.Errorf("Got %v, want %v", val, tt.expected)
				}
			}

		})
	}
}
