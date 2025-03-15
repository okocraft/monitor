import { createFileRoute } from "@tanstack/react-router";
import * as Link from "../../../pages/google/Link";
import { getSearchParam } from "../../../utils/searchParams.ts";

export const Route = createFileRoute("/google/link/")({
    component: Component,
});

function Component() {
    const loginKey = getSearchParam("loginKey");
    if (!loginKey) {
        return <p>No login key provided</p>;
    }
    return <Link.Component loginKey={loginKey} />;
}
