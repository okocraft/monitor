import { createFileRoute } from "@tanstack/react-router";
import { Header } from "../components/ui/Header";
import { TopPage } from "../pages/TopPage.tsx";

export const Route = createFileRoute("/")({
    component: Component,
});

function Component() {
    return (
        <>
            <Header />
            <TopPage />
        </>
    );
}
