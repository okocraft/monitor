package queries

import "strings"

var replacerForLikeParameter = strings.NewReplacer("%", "\\%", "_", "\\_", "\\", "\\\\")

func escapeLikeParameter(param string) string {
	return replacerForLikeParameter.Replace(param)
}

func likePartialMatch(param string) string {
	return "%" + escapeLikeParameter(param) + "%"
}
