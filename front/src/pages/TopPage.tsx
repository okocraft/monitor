import { Link } from "@tanstack/react-router";
import { useAuth } from "../hooks/useAuth.ts";

export const TopPage = () => {
    const auth = useAuth();

    return (
        <>
            <h3>Welcome Home!</h3>
            {auth.hasAccessToken() ? (
                <>
                    <Link to="/example">Go to example</Link>
                    <br />
                    <button type="submit" onClick={() => auth.logout()}>
                        LogOut
                    </button>
                </>
            ) : (
                <Link
                    to="/google/login"
                    search={{
                        redirectTo: encodeURIComponent("/"),
                    }}
                >
                    Go to Google Login Page
                </Link>
            )}
        </>
    );
};
