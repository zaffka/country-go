package country_test

import (
	"testing"

	"github.com/zaffka/country-go"
)

func TestListLen(t *testing.T) {
	if count := country.ListLen(); count == 0 {
		t.Error("Expected non-zero country count, got 0")
	}
}

func TestByName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantName string
		wantErr  bool
	}{
		{"Valid uppercase", "CANADA", "CANADA", false},
		{"Valid lowercase", "canada", "CANADA", false},
		{"Valid mixed case", "CaNaDa", "CANADA", false},
		{"Valid with spaces", "  canada  ", "CANADA", false},
		{"Invalid country", "WAKANDA", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := country.ByName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.Name != tt.wantName {
				t.Errorf("ByName() = %v, want %v", got.Name, tt.wantName)
			}
		})
	}
}

func TestByAlpha2Code(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantAlpha2 string
		wantErr    bool
	}{
		{"Valid uppercase", "CA", "CA", false},
		{"Valid lowercase", "ca", "CA", false},
		{"Valid with spaces", " ca ", "CA", false},
		{"Invalid code", "XX", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := country.ByAlpha2Code(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByAlpha2Code() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.Alpha2 != tt.wantAlpha2 {
				t.Errorf("ByAlpha2Code() = %v, want %v", got.Alpha2, tt.wantAlpha2)
			}
		})
	}
}

func TestByAlpha3Code(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantAlpha3 string
		wantErr    bool
	}{
		{"Valid uppercase", "CAN", "CAN", false},
		{"Valid lowercase", "can", "CAN", false},
		{"Valid with spaces", " can ", "CAN", false},
		{"Invalid code", "XXX", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := country.ByAlpha3Code(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByAlpha3Code() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.Alpha3 != tt.wantAlpha3 {
				t.Errorf("ByAlpha3Code() = %v, want %v", got.Alpha3, tt.wantAlpha3)
			}
		})
	}
}

func TestByISONum(t *testing.T) {
	tests := []struct {
		name    string
		input   uint16
		wantNum uint16
		wantErr bool
	}{
		{"Valid code", 124, 124, false},
		{"Invalid code", 999, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := country.ByISONum(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByISONum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.ISONum != tt.wantNum {
				t.Errorf("ByISONum() = %v, want %v", got.ISONum, tt.wantNum)
			}
		})
	}
}

func TestExistsFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func() bool
		want bool
	}{
		{"Exists - valid", func() bool { return country.Exists("canada") }, true},
		{"Exists - invalid", func() bool { return country.Exists("wakanda") }, false},
		{"ExistsAlpha2 - valid", func() bool { return country.ExistsAlpha2("ca") }, true},
		{"ExistsAlpha2 - invalid", func() bool { return country.ExistsAlpha2("xx") }, false},
		{"ExistsAlpha3 - valid", func() bool { return country.ExistsAlpha3("can") }, true},
		{"ExistsAlpha3 - invalid", func() bool { return country.ExistsAlpha3("xxx") }, false},
		{"ExistsISONum - valid", func() bool { return country.ExistsISONum(124) }, true},
		{"ExistsISONum - invalid", func() bool { return country.ExistsISONum(999) }, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn(); got != tt.want {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
