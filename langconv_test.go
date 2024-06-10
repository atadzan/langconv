package langconv

import "testing"

func TestGetISO6391Code(t *testing.T) {
	testCases := []struct {
		Title        string
		Language     string
		ExpectedCode string
	}{
		{
			Title:        "Turkmen language",
			Language:     "Turkmen",
			ExpectedCode: "tk",
		},
		{
			Title:        "Russian language",
			Language:     "Russian",
			ExpectedCode: "ru",
		},
		{
			Title:        "English language",
			Language:     "English",
			ExpectedCode: "en",
		},
		{
			Title:        "Invalid language Plany",
			Language:     "Plany",
			ExpectedCode: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			lang, _ := GetISO6391Code(testCase.Language)
			if lang.Code != testCase.ExpectedCode {
				t.Errorf("Expected language ISO 639-1 is %s. But actual lang code is %s", testCase.ExpectedCode, lang.Code)
			}
		})
	}
}
