import type { ReactNode } from "react";
import { createAuthState } from "../types/auth.ts";
import { AuthContext } from "../contexts/authContext.ts";

export function AuthProvider({ children }: { children: ReactNode }) {
    const authState = createAuthState();
    return (
        <AuthContext.Provider value={authState}>
            {children}
        </AuthContext.Provider>
    );
}
