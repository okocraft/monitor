const PageTypes = {
    // see app/domain/auth/google.go
    success: "success",
    notEnabled: "not_enabled",
    tryAgain: "try_again",
    userNotFound: "user_not_found",
    loginKeyNotFound: "login_key_not_found",
    invalidToken: "invalid_token",
    internalError: "internal_error",
} as const;

export type PageType = (typeof PageTypes)[keyof typeof PageTypes];

export function toPageType(value: string | null): PageType | undefined {
    if (Object.values(PageTypes).includes(value as PageType)) {
        return value as PageType;
    }
    return undefined;
}
