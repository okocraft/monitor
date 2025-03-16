import { createFileRoute } from "@tanstack/react-router";
import { Header } from "../../../components/ui/Header";
import * as Result from "../../../pages/google/Result";
import { toPageType } from "../../../pages/google/Result/pageTypes.ts";
import { getSearchParam } from "../../../utils/searchParams.ts";

export const Route = createFileRoute("/google/result/")({
    component: Component,
});

function Component() {
    const type = getSearchParam("type");
    return (
        <>
            <Header />
            <Result.Component type={type ? toPageType(type) : undefined} />
        </>
    );
}
