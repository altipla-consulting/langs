package langs

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Content represents a translatable string that can store a different value in
// each language.
//
// It can be serialized to JSON to send it to a client application. It can also
// be used with libs.altipla.consulting/rdb.
type Content struct {
	v map[string]string
}

// NewContent returns an empty content without any values.
func NewContent() Content {
	return Content{}
}

// NewContentValue builds a new content with a single translated value.
func NewContentValue(lang Lang, value string) Content {
	content := NewContent()
	content.Set(lang.Code, value)
	return content
}

// NewContentFromMap builds a new content from a map containing one or multiple values.
func NewContentFromMap(values map[string]string) Content {
	return Content{
		v: values,
	}
}

// ParseContent reads a new content map parsing every language name to check if
// it's correct. It only returns an error for invalid languages.
func ParseContent(values map[string]string) (Content, error) {
	content := NewContent()
	for lang, value := range values {
		if !IsValid(lang) {
			return Content{}, fmt.Errorf("unknown lang %q", lang)
		}
		content.Set(lang, value)
	}
	return content, nil
}

func (content *Content) init() {
	if content.v == nil {
		content.v = make(map[string]string)
	}
}

// Set changes the translated value of a language. If empty that language will be
// discarded from the content.
func (content *Content) Set(lang string, value string) {
	content.init()
	if value == "" {
		delete(content.v, lang)
	} else {
		content.v[lang] = value
	}
}

// Get returns the translated value for a language or empty if not present.
func (content Content) Get(lang string) string {
	if content.v == nil {
		return ""
	}
	return content.v[lang]
}

// IsEmpty returns if the content does not contains any translated value in any language.
func (content Content) IsEmpty() bool {
	return len(content.v) == 0
}

// Clear removes a specific language translated value.
func (content *Content) Clear(lang string) {
	content.init()
	delete(content.v, lang)
}

// ClearAll removes all translated values in all languages.
func (content *Content) ClearAll() {
	content.v = nil
}

// Chain helps configuring the chain of fallbacks for a project.
type Chain struct {
	fallbacks []Lang
}

// NewChain initializes a new chain.
func NewChain(fallbacks ...Lang) Chain {
	return Chain{
		fallbacks: fallbacks,
	}
}

// GetChain does the following steps:
//
// 1. Return the content in the requested lang if available.
// 2. Use the fallback languages if one of them is available. Order is important here.
// 3. Return any lang available randomly to have something.
//
// If the content is empty it returns an empty string.
func (content Content) GetChain(chain Chain, lang string) string {
	if content.IsEmpty() {
		return ""
	}

	value, ok := content.v[lang]
	if ok {
		return value
	}

	for _, l := range chain.fallbacks {
		value, ok := content.v[l.Code]
		if ok {
			return value
		}
	}

	for k := range content.v {
		return content.v[k]
	}

	panic("should not reach here")
}

// MarshalJSON implements the JSON interface.
func (content Content) MarshalJSON() ([]byte, error) {
	v := content.v
	if v == nil {
		v = make(map[string]string)
	}
	return json.Marshal(v)
}

// UnmarshalJSON implements the JSON interface.
func (content *Content) UnmarshalJSON(b []byte) error {
	v := make(map[string]string)
	if err := json.Unmarshal(b, &v); err != nil {
		return fmt.Errorf("cannot unmarshal langs: %w", err)
	}

	content.v = v
	return nil
}

// Map returns a map of every language of the content and its value.
func (content Content) Map() map[string]string {
	c := make(map[string]string)
	for k, v := range content.v {
		c[k] = v
	}
	return c
}

// PlainMap returns a map of every language of the content and its value in a
// format that can be serialized to protobufs.
func (content Content) PlainMap() map[string]string {
	c := make(map[string]string)
	for k, v := range content.v {
		c[string(k)] = v
	}
	return c
}

func (content *Content) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch value.(type) {
	case []byte:
		return json.Unmarshal(value.([]byte), &content.v)
	case string:
		return json.Unmarshal([]byte(value.(string)), &content.v)
	default:
		return fmt.Errorf("sqltypes: unknown content type %T", value)
	}
}

func (content Content) Value() (driver.Value, error) {
	if content.v == nil {
		return "{}", nil
	}
	return json.Marshal(content)
}
