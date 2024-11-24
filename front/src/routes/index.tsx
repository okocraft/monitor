import { createFileRoute } from "@tanstack/react-router";
import { usePing } from "../api/default/default.ts";

export const Route = createFileRoute("/")({
    component: Home,
});

function Home() {
    const { data: pong, refetch } = usePing({
        query: {
            enabled: false,
        },
    });

    return (
        <>
            <h3>Welcome Home!</h3>
            <button
                type="button"
                onClick={() => refetch()}
                className="bg-amber-500"
            >
                Click here
            </button>
            {pong && (
                <p className="text-blue-300 font-bold">
                    {pong.data.name}: {pong.data.message}
                </p>
            )}
        </>
    );
}
