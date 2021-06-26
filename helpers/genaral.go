package helpers

import "github.com/jasonbronson/ldd/models"

func ConvertMatchestoMatchString(matches []*models.Matches) []string {

	var match_strings []string

	for _, el := range matches {
		match_strings = append(match_strings, el.Matching_string)
	}

	return match_strings
}

func ConvertLogMatchestoMatchString(logmatches []*models.Logs) []string {

	var match_strings []string

	for _, el := range logmatches {
		match_strings = append(match_strings, el.Matching_string)
	}

	return match_strings
}
