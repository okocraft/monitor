import { useLoginWithGoogle } from "../../../api/auth/auth.ts";
import { SkipAuth } from "../../../hooks/SkipAuth.tsx";

export const Component = ({ redirectTo }: { redirectTo?: string }) => {
    const { mutate } = useLoginWithGoogle({
        mutation: {
            onSuccess: (res) => {
                window.location.href = res.data.redirect_url;
            },
            onError: (error) => {
                console.error("Failed to get redirect URL:", error);
            },
        },
    });

    const handleLogin = () => {
        mutate({
            data: {
                url: redirectTo ? redirectTo : "",
            },
        });
    };

    return (
        <SkipAuth>
            <button
                type={"button"}
                onClick={handleLogin}
                style={{ padding: "10px", fontSize: "16px" }}
            >
                Login with Google
            </button>
        </SkipAuth>
    );
};
