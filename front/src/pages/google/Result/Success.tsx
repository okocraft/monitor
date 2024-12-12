import { useNavigate } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { getSearchParam } from "../../../utils/searchParams.ts";

export const Success = () => {
    const redirectTo = getSearchParam("redirectTo");
    if (redirectTo) {
        return <RedirectToPage url={redirectTo} />;
    }

    return (
        <>
            <p>Successfully logged in.</p>
        </>
    );
};

const waitingSecond = 3;

const RedirectToPage = ({ url }: { url: string }) => {
    const navigate = useNavigate();
    const [countdown, setCountdown] = useState(waitingSecond);
    useEffect(() => {
        const interval = setInterval(() => {
            setCountdown((prev: number) => prev - 1);
        }, 1000);

        if (countdown === 0) {
            navigate({
                to: url,
            }).catch((err) => {
                console.error("Navigation error:", err);
            });
        }

        return () => clearInterval(interval);
    }, [countdown, navigate, url]);

    return (
        <div>
            <p>
                Redirecting to {url} in {countdown} seconds...
            </p>
        </div>
    );
};
