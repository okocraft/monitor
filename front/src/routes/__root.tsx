import { Outlet, createRootRouteWithContext } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import type { AuthState } from "../types/auth.ts";
import { NotFound } from "../pages/NotFound.tsx";
import { ErrorOccurred } from "../pages/ErrorOccurred.tsx";

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
