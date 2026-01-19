package template

import (
	"testing"

	"github.com/moq77111113/kite/internal/domain/models"
)

func TestEngine_ExtractVariables(t *testing.T) {
	engine := NewEngine()
	vars, err := engine.ExtractVariables("Hello [[name]], your ID is [[id]]. Welcome [[name]]!")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(vars) != 2 || vars[0] != "name" || vars[1] != "id" {
		t.Errorf("expected [name, id], got %v", vars)
	}
}

func TestEngine_Interpolate_Simple(t *testing.T) {
	engine := NewEngine()
	result, err := engine.Interpolate("Hello [[name]]!", map[string]string{"name": "World"}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "Hello World!" {
		t.Errorf("expected 'Hello World!', got %q", result)
	}
}

func TestEngine_Interpolate_WithDefaults(t *testing.T) {
	engine := NewEngine()
	meta := []models.Variable{{Name: "port", Default: "5432"}}
	result, err := engine.Interpolate("Port: [[port]]", nil, meta)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "Port: 5432" {
		t.Errorf("expected 'Port: 5432', got %q", result)
	}
}

func TestEngine_Interpolate_ValueOverridesDefault(t *testing.T) {
	engine := NewEngine()
	meta := []models.Variable{{Name: "port", Default: "5432"}}
	result, err := engine.Interpolate("Port: [[port]]", map[string]string{"port": "3306"}, meta)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "Port: 3306" {
		t.Errorf("expected 'Port: 3306', got %q", result)
	}
}

func TestEngine_Interpolate_UndefinedOptional(t *testing.T) {
	engine := NewEngine()
	result, err := engine.Interpolate("Value: [[optional]]", nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "Value: " {
		t.Errorf("expected 'Value: ', got %q", result)
	}
}

func TestEngine_Interpolate_MissingRequired(t *testing.T) {
	engine := NewEngine()
	meta := []models.Variable{{Name: "password", Required: true}}
	_, err := engine.Interpolate("Password: [[password]]", nil, meta)
	if err == nil {
		t.Fatal("expected error for missing required variable")
	}
}

func TestEngine_ExtractFromFiles(t *testing.T) {
	engine := NewEngine()
	files := []models.File{
		{Path: "config.yaml", Content: "db: [[db_name]]\nport: [[db_port]]"},
		{Path: "docker.yml", Content: "password: [[db_password]]"},
	}
	vars := engine.ExtractFromFiles(files)
	if len(vars) != 3 {
		t.Fatalf("expected 3 variables, got %d: %v", len(vars), vars)
	}
}

func TestEngine_ExtractFromFiles_SkipsBinary(t *testing.T) {
	engine := NewEngine()
	files := []models.File{
		{Path: "binary.bin", Content: "has\x00null\x00bytes[[ignored]]"},
		{Path: "config.txt", Content: "[[included]]"},
	}
	vars := engine.ExtractFromFiles(files)
	if len(vars) != 1 || vars[0] != "included" {
		t.Errorf("expected only 'included', got %v", vars)
	}
}
