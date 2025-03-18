import { createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/settings/")({
    beforeLoad: async ({ context }) => {
        const isAuthenticated = await context.auth.isAuthenticated();
        if (!isAuthenticated) {
            throw redirect({
                to: "/",
            });
        }
    },
    component: Component,
});

function Component() {
    return <div>Hello "/settings/"!</div>;
}
