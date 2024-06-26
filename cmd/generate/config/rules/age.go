package rules

import (
	"regexp"

	"github.com/zricethezav/gitleaks/v8/config"
)

func AgeSecretKey() *config.Rule {
	// define rule
	r := config.Rule{
		Description: "Discovered a potential Age encryption tool secret key, risking data decryption and unauthorized access to sensitive information.",
		RuleID:      "age-secret-key",
		Regex:       regexp.MustCompile(`AGE-SECRET-KEY-1[QPZRY9X8GF2TVDW0S3JN54KHCE6MUA7L]{58}`),
		Keywords:    []string{"AGE-SECRET-KEY-1"},
	}

	// validate
	tps := []string{
	         generateSampleSecret("generic", "CLOJARS_34bf0e88955ff5a1c328d6a7491acc4f48e865a7b8dd4d70a70749037443"),
		`apiKey := "AGE-SECRET-KEY-1QQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQ`, // gitleaks:allow
	}
	return validate(r, tps, nil)
}
