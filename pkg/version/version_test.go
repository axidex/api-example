package version_test

import (
	"github.com/axidex/api-example/pkg/version"
	"testing"
)

func TestVersionMethods(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		version     version.Version
		wantVersion string
		wantDate    string
		wantCommit  string
	}{
		{
			name:        "Default BuildVersion",
			version:     version.NewVersion(),
			wantVersion: "N/A",
			wantDate:    "N/A",
			wantCommit:  "N/A",
		},
		{
			name: "Custom BuildVersion",
			version: &version.BuildVersion{
				BuildVersion: "v1.2.3",
				BuildDate:    "02.03.2025",
				BuildCommit:  "abcd1234",
			},
			wantVersion: "v1.2.3",
			wantDate:    "02.03.2025",
			wantCommit:  "abcd1234",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.version.Version(); got != tt.wantVersion {
				t.Errorf("BuildVersion() = %q, want %q", got, tt.wantVersion)
			}

			if got := tt.version.Date(); got != tt.wantDate {
				t.Errorf("Date() = %q, want %q", got, tt.wantDate)
			}

			if got := tt.version.Commit(); got != tt.wantCommit {
				t.Errorf("Commit() = %q, want %q", got, tt.wantCommit)
			}
		})
	}
}
