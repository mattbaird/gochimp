package mandrill

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAPIErrorConversion(t *testing.T) {
	testCases := map[string]struct {
		JSON     string
		Expected interface{}
	}{
		"invalid_key_error": {
			JSON:     fmt.Sprintf(`{"name":"Invalid_Key", "message":"%s","status":"error","code":24}`, t.Name()),
			Expected: &InvalidKeyError{},
		},
		"template_error": {
			JSON:     fmt.Sprintf(`{"name":"Unknown_Template", "message":"%s","status":"error","code":24}`, t.Name()),
			Expected: &TemplateError{},
		},
		"payment_error": {
			JSON:     fmt.Sprintf(`{"name":"PaymentRequired", "message":"%s","status":"error","code":24}`, t.Name()),
			Expected: &PaymentRequiredError{},
		},
		"subaccount_error": {
			JSON:     fmt.Sprintf(`{"name":"Unknown_Subaccount", "message":"%s","status":"error","code":24}`, t.Name()),
			Expected: &SubAccountError{},
		},
		"validation_error": {
			JSON:     fmt.Sprintf(`{"name":"ValidationError", "message":"%s","status":"error","code":24}`, t.Name()),
			Expected: &ValidationError{},
		},
		"general_error": {
			JSON:     fmt.Sprintf(`{"name":"GeneralError", "message":"%s","status":"error","code":24}`, t.Name()),
			Expected: &GeneralError{},
		},
		"service_error": {
			JSON:     fmt.Sprintf(`{"name":"ServiceUnavailable", "message":"%s","status":"error","code":24}`, t.Name()),
			Expected: &ServiceUnavailableError{},
		},
		"unknown_error": {
			JSON:     fmt.Sprintf(`{"name":"A_New_Error_Type", "message":"%s","status":"error","code":24}`, t.Name()),
			Expected: &UnknownError{},
		},
	}
	testName := t.Name()
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := parseErrorJSON([]byte(tc.JSON))
			require.Error(t, err)
			require.IsType(t, tc.Expected, err)
			require.Contains(t, err.Error(), testName)
		})
	}
}

func TestNonMandrillErrorJSON(t *testing.T) {
	testJSON := `{"key":"value"}`
	err := parseErrorJSON([]byte(testJSON))
	require.NoError(t, err)
}
