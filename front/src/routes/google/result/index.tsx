import { createFileRoute } from "@tanstack/react-router";
import * as Result from "../../../pages/google/Result";
import { toPageType } from "../../../types/google/pageTypes.ts";
import { NotFound } from "../../../pages/NotFound.tsx";
import { getSearchParam } from "../../../utils/searchParams.ts";

export const Route = createFileRoute("/google/result/")({
    component: Component,
});

function Component() {
    const type = getSearchParam("type");
    if (!type) {
        return <NotFound />;
    }

    const pageType = toPageType(type);
    if (!pageType) {
        return <NotFound />;
    }

    return <Result.Component type={pageType} />;
}
