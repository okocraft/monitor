import { createContext } from "react";
import { type AuthState, UnauthorizedState } from "../types/auth.ts";

export const AuthContext = createContext<AuthState>(UnauthorizedState);
