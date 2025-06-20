import { createRouter, RouterProvider } from "@tanstack/react-router";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { routeTree } from "./routeTree.gen";
import "./index.css";
import { QueryClientProvider } from "@tanstack/react-query";
import axios from "axios";
import { queryClient } from "./client/api.ts";
import { headerLinks } from "./headerLinks.ts";
import { AuthProvider } from "./hooks/AuthProvider.tsx";
import { AxiosClientProvider } from "./hooks/AxiosClientProvider.tsx";
import { HeaderLinkProvider } from "./hooks/HeaderLinkProvider.tsx";
import { useAuth } from "./hooks/useAuth.ts";
import { UnauthorizedState } from "./types/auth.ts";

const root = document.getElementById("root");
if (!root) {
    throw Error("root not found");
}

const apiUrl = import.meta.env.VITE_API_URL as string;
if (!apiUrl) {
    throw Error("api url not set");
}

axios.defaults.baseURL = apiUrl;

const router = createRouter({
    routeTree,
    defaultPreload: "intent",
    defaultStaleTime: 5000,
    context: {
        auth: UnauthorizedState,
    },
});

declare module "@tanstack/react-router" {
    interface Register {
        router: typeof router;
    }
}

if (!root.innerHTML) {
    createRoot(root).render(
        <StrictMode>
            <AuthProvider>
                <QueryClientProvider client={queryClient}>
                    <AxiosClientProvider>
                        <HeaderLinkProvider links={headerLinks}>
                            <App />
                        </HeaderLinkProvider>
                    </AxiosClientProvider>
                </QueryClientProvider>
            </AuthProvider>
        </StrictMode>,
    );
}

function App() {
    const auth = useAuth();
    return <RouterProvider router={router} context={{ auth }} />;
}
