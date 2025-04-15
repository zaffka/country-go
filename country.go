package country

import (
	_ "embed"
	"errors"
	"strings"

	"github.com/goccy/go-json"
)

//go:embed dictionary/countries.json
var dictionary []byte

var (
	// List contains all countries as a slice of Ident structs.
	List []Ident

	// ListedByName provides quick lookup by country name (uppercase).
	ListedByName map[string]*Ident

	// ListedByAlpha2 provides quick lookup by 2-letter country code (uppercase).
	ListedByAlpha2 map[string]*Ident

	// ListedByAlpha3 provides quick lookup by 3-letter country code (uppercase).
	ListedByAlpha3 map[string]*Ident

	// ListedByIsoNum provides quick lookup by ISO numeric country code.
	ListedByIsoNum map[uint16]*Ident

	// ErrNotFound is returned when a country lookup fails.
	ErrNotFound = errors.New("country not found")
)

// Ident represents a country with its standard identifiers.
// All fields use official ISO 3166-1 naming and codes.
type Ident struct {
	// Name - is an official country name in uppercase (e.g., "CANADA").
	Name string `json:"name"`

	// Alpha2 - 2-letter country code (e.g., "CA").
	Alpha2 string `json:"alpha2"`

	// Alpha3 - 3-letter country code (e.g., "CAN").
	Alpha3 string `json:"alpha3"`

	// ISONum - ISO numeric country code (e.g., 124).
	ISONum uint16 `json:"iso_num"`
}

func init() {
	if err := json.Unmarshal(dictionary, &List); err != nil {
		panic("failed to unmarshal countries dictionary: " + err.Error())
	}

	l := len(List)
	ListedByName = make(map[string]*Ident, l)
	ListedByAlpha2 = make(map[string]*Ident, l)
	ListedByAlpha3 = make(map[string]*Ident, l)
	ListedByIsoNum = make(map[uint16]*Ident, l)

	for i, c := range List {
		nameKey := strings.ToUpper(strings.TrimSpace(c.Name))
		alpha2Key := strings.ToUpper(strings.TrimSpace(c.Alpha2))
		alpha3Key := strings.ToUpper(strings.TrimSpace(c.Alpha3))

		ListedByName[nameKey] = &List[i]
		ListedByAlpha2[alpha2Key] = &List[i]
		ListedByAlpha3[alpha3Key] = &List[i]
		ListedByIsoNum[c.ISONum] = &List[i]
	}
}

// ListLen returns the total number of countries in the dataset.
func ListLen() int {
	return len(List)
}

// ByName finds a country by its name (case-insensitive).
// Returns ErrNotFound if no matching country exists.
func ByName(name string) (*Ident, error) {
	if c, exists := ListedByName[strings.ToUpper(strings.TrimSpace(name))]; exists {
		return c, nil
	}

	return nil, ErrNotFound
}

// ByAlpha2Code finds a country by its 2-letter ISO code (case-insensitive).
// Returns ErrNotFound if no matching country exists.
func ByAlpha2Code(code string) (*Ident, error) {
	if c, exists := ListedByAlpha2[strings.ToUpper(strings.TrimSpace(code))]; exists {
		return c, nil
	}

	return nil, ErrNotFound
}

// ByAlpha3Code finds a country by its 3-letter ISO code (case-insensitive).
// Returns ErrNotFound if no matching country exists.
func ByAlpha3Code(code string) (*Ident, error) {
	if c, exists := ListedByAlpha3[strings.ToUpper(strings.TrimSpace(code))]; exists {
		return c, nil
	}

	return nil, ErrNotFound
}

// ByISONum finds a country by its ISO numeric code.
// Returns ErrNotFound if no matching country exists.
func ByISONum(isoNum uint16) (*Ident, error) {
	if c, exists := ListedByIsoNum[isoNum]; exists {
		return c, nil
	}

	return nil, ErrNotFound
}

// Exists checks if a country with the given name exists (case-insensitive).
func Exists(name string) bool {
	_, exists := ListedByName[strings.ToUpper(strings.TrimSpace(name))]

	return exists
}

// ExistsAlpha2 checks if a country with the given Alpha2 code exists (case-insensitive).
func ExistsAlpha2(code string) bool {
	_, exists := ListedByAlpha2[strings.ToUpper(strings.TrimSpace(code))]

	return exists
}

// ExistsAlpha3 checks if a country with the given Alpha3 code exists (case-insensitive).
func ExistsAlpha3(code string) bool {
	_, exists := ListedByAlpha3[strings.ToUpper(strings.TrimSpace(code))]

	return exists
}

// ExistsISONum checks if a country with the given ISO numeric code exists.
func ExistsISONum(isoNum uint16) bool {
	_, exists := ListedByIsoNum[isoNum]

	return exists
}
