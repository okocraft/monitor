package auth

type GoogleResultPageType string

const (
	GoogleResultPageSuccess          GoogleResultPageType = "success"
	GoogleResultPageNotEnabled       GoogleResultPageType = "not_enabled"
	GoogleResultPageTryAgain         GoogleResultPageType = "try_again"
	GoogleResultPageUserNotFound     GoogleResultPageType = "user_not_found"
	GoogleResutlPageLoginKeyNotFound GoogleResultPageType = "login_key_not_found"
	GoogleResultPageInvalidToken     GoogleResultPageType = "invalid_token"
	GoogleResultPageInternalError    GoogleResultPageType = "internal_error"
)
