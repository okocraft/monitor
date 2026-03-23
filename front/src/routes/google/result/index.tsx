import { createFileRoute } from "@tanstack/react-router";
import { Header } from "../../../components/ui/Header";
import * as Result from "../../../pages/google/Result";
import { toPageType } from "../../../pages/google/Result/pageTypes.ts";
import { useSearchParam } from "../../../utils/searchParams.ts";

export const Route = createFileRoute("/google/result/")({
    component: Component,
});

function Component() {
    const type = useSearchParam("type");
    return (
        <>
            <Header />
            <Result.Component type={type ? toPageType(type) : undefined} />
        </>
    );
}
