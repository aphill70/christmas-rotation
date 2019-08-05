package gifts

import "strings"

// NormalizeName handles typos and toLowers
func NormalizeName(name string) string {
	normalized := strings.Trim(strings.ToLower(name), " ")

	if strings.HasPrefix(normalized, "chris") {
		return "christopher"
	}

	if strings.HasPrefix(normalized, "mich") {
		return "michael"
	}

	return normalized
}
