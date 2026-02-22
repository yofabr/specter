package scanner

import (
	"testing"
)

func TestNewScanner(t *testing.T) {
	cfg := Config{
		Target:    "127.0.0.1",
		StartPort: 1,
		EndPort:   1024,
		Workers:   100,
	}

	s := NewScanner(cfg)
	if s.config.Target != cfg.Target {
		t.Errorf("expected target %s, got %s", cfg.Target, s.config.Target)
	}
	if s.config.StartPort != cfg.StartPort {
		t.Errorf("expected start port %d, got %d", cfg.StartPort, s.config.StartPort)
	}
	if s.config.EndPort != cfg.EndPort {
		t.Errorf("expected end port %d, got %d", cfg.EndPort, s.config.EndPort)
	}
	if s.config.Workers != cfg.Workers {
		t.Errorf("expected workers %d, got %d", cfg.Workers, s.config.Workers)
	}
}

func TestScan(t *testing.T) {
	cfg := Config{
		Target:    "127.0.0.1",
		StartPort: 80,
		EndPort:   80,
		Workers:   1,
	}

	s := NewScanner(cfg)
	results := s.Scan()

	for _, r := range results {
		if r.State != "open" {
			t.Errorf("expected state 'open', got '%s'", r.State)
		}
	}
}
