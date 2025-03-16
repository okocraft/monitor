import { createFileRoute } from "@tanstack/react-router";
import { Header } from "../../../components/ui/Header";
import * as Login from "../../../pages/google/Login";
import { getSearchParam } from "../../../utils/searchParams.ts";

export const Route = createFileRoute("/google/login/")({
    component: Component,
});

function Component() {
    const redirectTo = getSearchParam("redirectTo");
    return (
        <>
            <Header />
            <Login.Component redirectTo={redirectTo} />
        </>
    );
}
