import { RouterProvider, createRouter } from "@tanstack/react-router";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { routeTree } from "./routeTree.gen";
import "./index.css";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import axios from "axios";

const router = createRouter({
    routeTree,
    defaultPreload: "intent",
    defaultStaleTime: 5000,
});

declare module "@tanstack/react-router" {
    interface Register {
        router: typeof router;
    }
}

const apiUrl = import.meta.env.VITE_API_URL as string;
if (!apiUrl) {
    throw Error("api url not set");
}

axios.defaults.baseURL = apiUrl;

const queryClient = new QueryClient();

const root = document.getElementById("root");
if (!root) {
    throw Error("root not found");
}

createRoot(root).render(
    <StrictMode>
        <QueryClientProvider client={queryClient}>
            <RouterProvider router={router} />
        </QueryClientProvider>
    </StrictMode>,
);
