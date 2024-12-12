import type { ReactNode } from "react";
import { AuthContext } from "../contexts/authContext.ts";
import { createAuthState } from "../types/auth.ts";

export function AuthProvider({ children }: { children: ReactNode }) {
    const authState = createAuthState();
    return (
        <AuthContext.Provider value={authState}>
            {children}
        </AuthContext.Provider>
    );
}
