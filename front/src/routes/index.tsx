import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
    component: Home,
});

function Home() {
    return (
        <>
            <h3>Welcome Home!</h3>
            <p className="text-red-500 font-bold">Hello Tailwind</p>
        </>
    );
}
