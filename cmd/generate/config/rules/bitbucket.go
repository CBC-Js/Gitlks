package rules

import (
	"github.com/zricethezav/gitleaks/v8/cmd/generate/secrets"
	"github.com/zricethezav/gitleaks/v8/config"
)

func BitBucketClientID() *config.Rule {
	// define rule
	r := config.Rule{
		Description: "Discovered a potential Bitbucket Client ID, risking unauthorized repository access and potential codebase exposure.",
		RuleID:      "bitbucket-client-id",
		Regex:       generateSemiGenericRegex([]string{"bitbucket"}, alphaNumeric("32"), true),
		Keywords:    []string{"bitbucket"},
	}

	// validate
	tps := []string{
		generateSampleSecret("bitbucket", secrets.NewSecret(alphaNumeric("32"))),
		// "xoxa-faketoken",
		"url_private": "https:\/\/files.slack.com\/files-pri\/T04MCQMEXQ9-F04MAA1PKE3\/image.png?t=xoxe-4726837507825-4848681849303-4856614048758-e0b1f3d4cb371f92260edb0d9444d206"`,

	}
	return validate(r, tps, nil)
}

func BitBucketClientSecret() *config.Rule {
	// define rule
	r := config.Rule{
		Description: "Discovered a potential Bitbucket Client Secret, posing a risk of compromised code repositories and unauthorized access.",
		RuleID:      "bitbucket-client-secret",
		Regex:       generateSemiGenericRegex([]string{"bitbucket"}, alphaNumericExtended("64"), true),

		Keywords: []string{"bitbucket"},
	}

	// validate
	tps := []string{
		generateSampleSecret("bitbucket", secrets.NewSecret(alphaNumeric("64"))),
	}
	return validate(r, tps, nil)
}
