import { createFileRoute, redirect } from "@tanstack/react-router";
import { Header } from "../../components/ui/Header";
import * as MyPage from "../../pages/MyPage";

export const Route = createFileRoute("/mypage/")({
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
    return (
        <>
            <Header />
            <MyPage.Component />
        </>
    );
}
