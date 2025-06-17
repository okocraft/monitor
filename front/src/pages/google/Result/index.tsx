import { useAuth } from "../../../hooks/useAuth.ts";
import { getSearchParam } from "../../../utils/searchParams.ts";
import { Failure } from "./Failure.tsx";
import type { PageType } from "./pageTypes.ts";
import { RedirectingSuccess, Success } from "./Success.tsx";

export const Component = ({ type }: { type?: PageType }) => {
    const auth = useAuth();
    switch (type) {
        case "success": {
            const nickname = auth.me.current?.nickname ?? "unknown";
            const redirectTo = getSearchParam("redirectTo");
            return redirectTo ? (
                <RedirectingSuccess nickname={nickname} url={redirectTo} />
            ) : (
                <Success nickname={nickname} />
            );
        }
        case "not_enabled":
            return (
                <Failure
                    title="Google login is not enabled by the server"
                    description="If this error is unintentional, contact the administrator to change the server configuration."
                />
            );
        case "try_again":
            return (
                <Failure
                    title="Login failed due to expiration or other reasons"
                    description="Please try logging in again."
                />
            );
        case "user_not_found":
            return (
                <Failure
                    title="Your Google account is not linked to any user"
                    description="For new accounts, please request the administrator to provide a login URL."
                />
            );
        case "login_key_not_found":
            return (
                <Failure
                    title="The login key provided for account linking could not be found"
                    description="Request the administrator to provide the login URL again."
                />
            );
        case "invalid_token":
            return (
                <Failure
                    title="Invalid token received while logged in"
                    description="If this error persists even retrying to log in again, please contact the administrator."
                />
            );
        case "internal_error":
            return (
                <Failure
                    title="Internal server error occurred while logging in"
                    description="Please report this error to the administrator."
                />
            );
        default:
            return (
                <Failure
                    title="Unknown result received"
                    description="Please report this error to the administrator."
                />
            );
    }
};
