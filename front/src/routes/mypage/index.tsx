import { createFileRoute, redirect } from "@tanstack/react-router";
import { Header } from "../../components/ui/Header";

export const Route = createFileRoute("/mypage/")({
    beforeLoad: async ({ context }) => {
        const isAuthenticated = await context.auth.isAuthenticated();
        if (!isAuthenticated) {
            throw redirect({
                to: "/",
            });
        }
    },
    component: MyPage,
});

function MyPage() {
    return (
        <>
            <Header />
            <div>Hello "/mypage/"!</div>
        </>
    );
}
