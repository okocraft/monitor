import { Outlet, createRootRouteWithContext } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { ErrorOccurred } from "../pages/ErrorOccurred.tsx";
import { NotFound } from "../pages/NotFound.tsx";
import type { AuthState } from "../types/auth.ts";

interface RouterContext {
    auth: AuthState;
}

export const Route = createRootRouteWithContext<RouterContext>()({
    beforeLoad: async ({ context }) => {
        await context.auth.firstRefresh();
    },
    component: () => {
        return (
            <>
                <Outlet />
                <TanStackRouterDevtools />
            </>
        );
    },
    notFoundComponent: () => {
        return <NotFound />;
    },
    errorComponent: (props) => {
        return <ErrorOccurred props={props} />;
    },
});
