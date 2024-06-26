package configloader

import (
	"testing"
)

type Config struct {
	Field1 string `yaml:"field1"`
	Field2 int    `yaml:"field2"`
	Nested struct {
		Field3 bool `yaml:"field3"`
	}
}

func TestMockLoader_Load(t *testing.T) {
	mockData := map[string]interface{}{
		"Field1":        "value1",
		"Field2":        2,
		"Nested.Field3": true,
	}

	var config Config

	loader := NewMockLoader(mockData)
	if err := loader.Load(&config); err != nil {
		t.Fatalf("Failed to load configuration: %s", err)
	}

	if config.Field1 != "value1" {
		t.Fatalf("Field1 is not set correctly; expected: %s, got: %s", "value1", config.Field1)
	}
	if config.Field2 != 2 {
		t.Fatalf("Field2 is not set correctly; expected: %d, got: %d", 2, config.Field2)
	}
	if config.Nested.Field3 != true {
		t.Fatalf("Nested field is not set correctly; expected: %t, got: %t", true, config.Nested.Field3)
	}
}

func TestMockLoader_Override(t *testing.T) {
	mockData := map[string]interface{}{
		"Field1":        "value1",
		"Field2":        2,
		"Nested.Field3": true,
	}

	var config Config

	loader := NewMockLoader(mockData)

	loader.Override("Field1", "newvalue1")
	loader.Override("Nested.Field3", false)

	if err := loader.Load(&config); err != nil {
		t.Fatalf("Failed to load configuration: %s", err)
	}

	if config.Field1 != "newvalue1" {
		t.Fatalf("Field1 is not set correctly; expected: %s, got: %s", "newvalue1", config.Field1)
	}
	if config.Field2 != 2 {
		t.Fatalf("Field2 is not set correctly; expected: %d, got: %d", 2, config.Field2)
	}
	if config.Nested.Field3 != false {
		t.Fatalf("Nested field is not set correctly; expected: %t, got: %t", false, config.Nested.Field3)
	}
}

func TestMockLoader_MockNonExistentField(t *testing.T) {
	mockData := map[string]interface{}{
		"Field1": "value1",
		"Field2": 2,
		"FieldX": "valueX",
	}

	var config Config

	loader := NewMockLoader(mockData)

	if err := loader.Load(&config); err == nil {
		t.Fatalf("Expected error when trying to override non-existent field")
	}
}

func TestMockLoader_OvertideNonExistentField(t *testing.T) {
	mockData := map[string]interface{}{
		"Field1": "value1",
		"Field2": 2,
	}

	var config Config

	loader := NewMockLoader(mockData)

	loader.Override("FieldX", "valueX")

	if err := loader.Load(&config); err == nil {
		t.Fatalf("Expected error when trying to override non-existent field")
	}
}

type ConfigWithEmbeds struct {
	Config
	Field4 float64 `yaml:"field4"`
}

func TestMockLoader_MockEmbeddedDirectly(t *testing.T) {
	mockData := map[string]interface{}{
		"Field1":        "value1",
		"Field2":        2,
		"Nested.Field3": true,
		"Field4":        3.14,
	}

	var config ConfigWithEmbeds

	loader := NewMockLoader(mockData)

	if err := loader.Load(&config); err != nil {
		t.Fatalf("Failed to load configuration: %s", err)
	}

	if config.Field1 != "value1" {
		t.Fatalf("Field1 is not set correctly; expected: %s, got: %s", "value1", config.Field1)
	}
	if config.Field2 != 2 {
		t.Fatalf("Field2 is not set correctly; expected: %d, got: %d", 2, config.Field2)
	}
	if config.Nested.Field3 != true {
		t.Fatalf("Nested field is not set correctly; expected: %t, got: %t", true, config.Nested.Field3)
	}
	if config.Field4 != 3.14 {
		t.Fatalf("Field4 is not set correctly; expected: %f, got: %f", 3.14, config.Field4)
	}
}

func TestMockLoader_MockEmbeddedIndirectly(t *testing.T) {
	mockData := map[string]interface{}{
		"Config.Field1":        "value1",
		"Config.Field2":        2,
		"Config.Nested.Field3": true,
		"Field4":               3.14,
	}

	var config ConfigWithEmbeds

	loader := NewMockLoader(mockData)

	if err := loader.Load(&config); err != nil {
		t.Fatalf("Failed to load configuration: %s", err)
	}

	if config.Field1 != "value1" {
		t.Fatalf("Field1 is not set correctly; expected: %s, got: %s", "value1", config.Field1)
	}
	if config.Field2 != 2 {
		t.Fatalf("Field2 is not set correctly; expected: %d, got: %d", 2, config.Field2)
	}
	if config.Nested.Field3 != true {
		t.Fatalf("Nested field is not set correctly; expected: %t, got: %t", true, config.Nested.Field3)
	}
	if config.Field4 != 3.14 {
		t.Fatalf("Field4 is not set correctly; expected: %f, got: %f", 3.14, config.Field4)
	}
}

func TestMockLoader_OverrideEmbeddedDirectly(t *testing.T) {
	mockData := map[string]interface{}{
		"Field1":        "value1",
		"Field2":        2,
		"Nested.Field3": true,
		"Field4":        3.14,
	}

	var config ConfigWithEmbeds

	loader := NewMockLoader(mockData)

	loader.Override("Field1", "newvalue1")

	if err := loader.Load(&config); err != nil {
		t.Fatalf("Failed to load configuration: %s", err)
	}

	if config.Field1 != "newvalue1" {
		t.Fatalf("Field1 is not set correctly; expected: %s, got: %s", "newvalue1", config.Field1)
	}
	if config.Field2 != 2 {
		t.Fatalf("Field2 is not set correctly; expected: %d, got: %d", 2, config.Field2)
	}
	if config.Nested.Field3 != true {
		t.Fatalf("Nested field is not set correctly; expected: %t, got: %t", true, config.Nested.Field3)
	}
	if config.Field4 != 3.14 {
		t.Fatalf("Field4 is not set correctly; expected: %f, got: %f", 3.14, config.Field4)
	}
}

func TestMockLoader_OverrideEmbeddedIndirectly(t *testing.T) {
	mockData := map[string]interface{}{
		"Config.Field1":        "value1",
		"Config.Field2":        2,
		"Config.Nested.Field3": true,
		"Field4":               3.14,
	}

	var config ConfigWithEmbeds

	loader := NewMockLoader(mockData)

	loader.Override("Config.Field1", "newvalue1")

	if err := loader.Load(&config); err != nil {
		t.Fatalf("Failed to load configuration: %s", err)
	}

	if config.Field1 != "newvalue1" {
		t.Fatalf("Field1 is not set correctly; expected: %s, got: %s", "newvalue1", config.Field1)
	}
	if config.Field2 != 2 {
		t.Fatalf("Field2 is not set correctly; expected: %d, got: %d", 2, config.Field2)
	}
	if config.Nested.Field3 != true {
		t.Fatalf("Nested field is not set correctly; expected: %t, got: %t", true, config.Nested.Field3)
	}
	if config.Field4 != 3.14 {
		t.Fatalf("Field4 is not set correctly; expected: %f, got: %f", 3.14, config.Field4)
	}
}
