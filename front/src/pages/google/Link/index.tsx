import { useLinkWithGoogle } from "../../../api/auth/auth.ts";
import { SkipAuth } from "../../../hooks/SkipAuth.tsx";

export const Component = ({ loginKey }: { loginKey: string }) => {
    const { mutate } = useLinkWithGoogle({
        mutation: {
            onSuccess: (res) => {
                window.location.href = res.data.redirect_url;
            },
            onError: (error) => {
                console.error("Failed to get redirect URL:", error);
            },
        },
    });

    const handleClick = () => {
        mutate({
            loginKey,
        });
    };

    return (
        <SkipAuth>
            <button
                type={"button"}
                onClick={handleClick}
                style={{ padding: "10px", fontSize: "16px" }}
            >
                Login with Google
            </button>
        </SkipAuth>
    );
};
