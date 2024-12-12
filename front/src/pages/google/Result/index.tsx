import type { PageType } from "../../../types/google/pageTypes.ts";
import { NotFound } from "../../NotFound.tsx";
import { Success } from "./Success.tsx";

export const Component = ({ type }: { type: PageType }) => {
    switch (type) {
        case "success":
            return <Success />;
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
