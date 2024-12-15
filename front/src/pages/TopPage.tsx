import { Link } from "@tanstack/react-router";
import { Header } from "../components/ui/Header";
import { useAuth } from "../hooks/useAuth.ts";

export const TopPage = () => {
    const auth = useAuth();

    return (
        <>
            <Header />
            <h3>Welcome Home!</h3>
            {auth.hasAccessToken() ? (
                <>
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
