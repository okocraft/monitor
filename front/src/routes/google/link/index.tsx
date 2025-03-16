import { createFileRoute } from "@tanstack/react-router";
import { Header } from "../../../components/ui/Header";
import * as Link from "../../../pages/google/Link";
import { getSearchParam } from "../../../utils/searchParams.ts";

export const Route = createFileRoute("/google/link/")({
    component: Component,
});

function Component() {
    const loginKey = getSearchParam("loginKey");
    return (
        <>
            <Header />
            <Link.Component initialLoginKey={loginKey ? loginKey : ""} />
        </>
    );
}
