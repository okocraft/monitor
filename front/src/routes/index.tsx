import { createFileRoute } from "@tanstack/react-router";
import { TopPage } from "../pages/TopPage.tsx";

export const Route = createFileRoute("/")({
    component: Component,
});

function Component() {
    return <TopPage />;
}
