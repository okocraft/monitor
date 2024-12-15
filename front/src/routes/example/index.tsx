import { createFileRoute, redirect } from "@tanstack/react-router";
import { useAuth } from "../../hooks/useAuth.ts";
import {Header} from "../../components/ui/Header";

export const Route = createFileRoute("/example/")({
    beforeLoad: async ({ context }) => {
        const isAuthenticated = await context.auth.isAuthenticated();
        if (!isAuthenticated) {
            throw redirect({
                to: "/",
            });
        }
    },
    component: Example,
});

function Example() {
    const auth = useAuth();
    const handleRefreshButton = async () => {
        await auth.me.refresh();
    };
    return (
        <>
			<Header />
            <h3>Welcome Example!</h3>
            <p>Your uuid: {auth.me.current?.uuid}</p>
            <p>Your nickname: {auth.me.current?.nickname}</p>
            <button type={"button"} onClick={handleRefreshButton}>
                refresh
            </button>
        </>
    );
}
