import { createFileRoute } from "@tanstack/react-router";
import { Header } from "../../../components/ui/Header";
import * as Login from "../../../pages/google/Login";
import { useSearchParam } from "../../../utils/searchParams.ts";

export const Route = createFileRoute("/google/login/")({
    component: Component,
});

function Component() {
    const redirectTo = useSearchParam("redirectTo");
    return (
        <>
            <Header />
            <Login.Component redirectTo={redirectTo} />
        </>
    );
}
