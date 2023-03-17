package archive

import "testing"

func Test_extensionsFromFile(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Format
	}{
		{"Invalid", "", ""},
		{"None", "binary", ""},
		{"Single", "single.tar", Tar},
		{"Single-Zip", "single.zip", Zip},
		{"Double", "double.tar.gz", TarGzip},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extensionsFromFile(tt.input); got != tt.want {
				t.Errorf("extensionsFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
