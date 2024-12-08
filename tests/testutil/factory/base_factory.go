package factory

import (
	"fmt"

	"github.com/jambo0624/blog/tests/testutil/fixtures"
)

// BaseFactory provides common functionality for all factories.
type BaseFactory struct {
	sequence uint
}

func NewBaseFactory() BaseFactory {
	return BaseFactory{
		sequence: fixtures.InitFixturesLength,
	}
}

// ApplyOptions applies a list of option functions to an entity.
func ApplyOptions[T any](entity T, opts []func(T)) T {
	for _, opt := range opts {
		opt(entity)
	}
	return entity
}

// NextSequence increments and returns the sequence.
func (f *BaseFactory) NextSequence() uint {
	f.sequence++
	return f.sequence
}

// FormatTestName formats a test name with sequence.
func (f *BaseFactory) FormatTestName(prefix string) string {
	return fmt.Sprintf("Test %s %d", prefix, f.sequence)
}

// FormatUpdatedName formats an updated name with sequence.
func (f *BaseFactory) FormatUpdatedName(prefix string) string {
	return fmt.Sprintf("Updated %s %d", prefix, f.sequence)
}

// FormatTestSlug formats a test slug with sequence.
func (f *BaseFactory) FormatTestSlug(prefix string) string {
	return fmt.Sprintf("test-%s-%d", prefix, f.sequence)
}

// FormatUpdatedSlug formats an updated slug with sequence.
func (f *BaseFactory) FormatUpdatedSlug(prefix string) string {
	return fmt.Sprintf("updated-%s-%d", prefix, f.sequence)
}

// FormatHexColor formats a hex color with sequence.
func (f *BaseFactory) FormatHexColor() string {
	return fmt.Sprintf("#%06X", f.sequence)
}

// BuildRequest is a generic function to build requests.
func BuildRequest[T any](isUpdate bool, buildFn func(bool) interface{}) T {
	result := buildFn(isUpdate)
	req, ok := result.(T)
	if !ok {
		panic(fmt.Sprintf("invalid type assertion for %T", result))
	}
	return req
}
