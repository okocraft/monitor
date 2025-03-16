import { useAuth } from "../../../hooks/useAuth.ts";
import type { PageType } from "../../../types/google/pageTypes.ts";
import { getSearchParam } from "../../../utils/searchParams.ts";
import { NotFound } from "../../NotFound.tsx";
import { RedirectingSuccess, Success } from "./Success.tsx";

export const Component = ({ type }: { type: PageType }) => {
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
            return <TempResultPage type={type} />;
        case "try_again":
            return <TempResultPage type={type} />;
        case "user_not_found":
            return <TempResultPage type={type} />;
        case "login_key_not_found":
            return <TempResultPage type={type} />;
        case "invalid_token":
            return <TempResultPage type={type} />;
        case "internal_error":
            return <TempResultPage type={type} />;
        default:
            return <NotFound />;
    }
};

const TempResultPage = ({ type }: { type: PageType }) => {
    return (
        <>
            <p>{type}</p>
        </>
    );
};
