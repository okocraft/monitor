import { useContext } from "react";
import { AuthContext } from "../contexts/authContext.ts";

export const useAuth = () => {
    const state = useContext(AuthContext);
    if (!state) throw new Error("useAuth must be used within an AuthProvider");
    return state;
};
