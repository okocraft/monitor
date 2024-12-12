import { type ReactNode, useEffect } from "react";
import { useAuth } from "./useAuth.ts";

export function SkipAuth({ children }: { children: ReactNode }) {
    const { skipAuth } = useAuth();

    useEffect(() => {
        skipAuth(true);
        return () => skipAuth(false);
    }, [skipAuth]);

    return <>{children}</>;
}
