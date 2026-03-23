import { createFileRoute } from "@tanstack/react-router";
import { Header } from "../../../components/ui/Header";
import * as Link from "../../../pages/google/Link";
import { useSearchParam } from "../../../utils/searchParams.ts";

export const Route = createFileRoute("/google/link/")({
    component: Component,
});

function Component() {
    const loginKey = useSearchParam("loginKey");
    return (
        <>
            <Header />
            <Link.Component initialLoginKey={loginKey ? loginKey : ""} />
        </>
    );
}
