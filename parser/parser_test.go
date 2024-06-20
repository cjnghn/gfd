package parser

import (
	"testing"
)

func TestParseGitHubURL(t *testing.T) {
	tests := []struct {
		input       string
		wantOwner   string
		wantRepo    string
		wantPath    string
		expectError bool
	}{
		{
			input:       "https://github.com/id-Software/DOOM/tree/master/linuxdoom-1.10",
			wantOwner:   "id-Software",
			wantRepo:    "DOOM",
			wantPath:    "linuxdoom-1.10",
			expectError: false,
		},
		{
			input:       "invalid-url",
			expectError: true,
		},
		{
			input:       "https://github.com/id-Software/DOOM",
			expectError: true,
		},
	}

	for _, test := range tests {
		gotOwner, gotRepo, gotPath, err := ParseGitHubURL(test.input)
		if (err != nil) != test.expectError {
			t.Errorf("ParseGitHubURL(%q) error = %v, expectError %v", test.input, err, test.expectError)
			continue
		}
		if gotOwner != test.wantOwner {
			t.Errorf("ParseGitHubURL(%q) gotOwner = %q, want %q", test.input, gotOwner, test.wantOwner)
		}
		if gotRepo != test.wantRepo {
			t.Errorf("ParseGitHubURL(%q) gotRepo = %q, want %q", test.input, gotRepo, test.wantRepo)
		}
		if gotPath != test.wantPath {
			t.Errorf("ParseGitHubURL(%q) gotPath = %q, want %q", test.input, gotPath, test.wantPath)
		}
	}
}
