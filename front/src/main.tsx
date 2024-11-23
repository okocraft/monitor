import { RouterProvider, createRouter } from "@tanstack/react-router";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { routeTree } from "./routeTree.gen";
import "./index.css";

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

const root = document.getElementById("root");
if (!root) {
    throw Error("root not found");
}

createRoot(root).render(
    <StrictMode>
            <RouterProvider router={router} />
    </StrictMode>,
);
