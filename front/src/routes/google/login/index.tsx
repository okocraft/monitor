import { createFileRoute } from "@tanstack/react-router";
import * as Login from "../../../pages/google/Login";
import { getSearchParam } from "../../../utils/searchParams.ts";

export const Route = createFileRoute("/google/login/")({
    component: Component,
});

function Component() {
    const redirectTo = getSearchParam("redirectTo");
    return <Login.Component redirectTo={redirectTo} />;
}
