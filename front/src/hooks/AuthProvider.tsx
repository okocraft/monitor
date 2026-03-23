import type { ReactNode } from "react";
import { AuthContext } from "../contexts/authContext.ts";
import { useAuthState } from "../types/auth.ts";

export function AuthProvider({ children }: { children: ReactNode }) {
    const authState = useAuthState();
    return <AuthContext value={authState}>{children}</AuthContext>;
}
