import { useState } from "react";
import { logout, refreshAccessToken } from "../api/auth/auth.ts";
import { jwtDecode } from "jwt-decode";
import { createMeState, EmptyMe, type MeState } from "./me.ts";

export interface AuthState {
    accessToken: string;
    hasAccessToken: () => boolean;
    refresh: () => Promise<string>;
    firstRefresh: () => Promise<void>;
    isAuthenticated: () => Promise<boolean>;

    logout: () => Promise<boolean>;

    shouldSkipAuth: boolean;
    skipAuth: (shouldSkip: boolean) => void;

    me: MeState;
}

export const UnauthorizedState = {
    accessToken: "",
    hasAccessToken: () => false,
    refresh: async () => "",
    firstRefresh: async () => {},
    isAuthenticated: async () => false,

    logout: async () => false,

    shouldSkipAuth: false,
    skipAuth: (_) => {},

    me: {
        current: undefined,
        getMe: async () => EmptyMe,
        setMe: (_) => {},
        refresh: async () => EmptyMe,
    } as MeState,
} as AuthState;

export function createAuthState() {
    const [accessToken, setAccessToken] = useState<string>("");
    const [refreshed, setRefreshed] = useState<boolean>(false);
    const [isAuthSkipped, setSkipAuth] = useState(false);
    const meState = createMeState();

    const refresh = async () => {
        setRefreshed(true);

        try {
            const { data, status } = await refreshAccessToken({
                withCredentials: true,
            });
            if (status === 200) {
                setAccessToken(data.access_token);
                meState.setMe(data.me);
                return data.access_token;
            }
            setAccessToken("");
            meState.setMe(EmptyMe);
            return "";
        } catch {
            setAccessToken("");
            meState.setMe(EmptyMe);
            return "";
        }
    };

    const firstRefresh = async () => {
        setRefreshed(true);
        if (!refreshed) {
            await refresh();
        }
    };

    const hasAccessToken = () => {
        return !!accessToken;
    };

    const isAuthenticated = async () => {
        if (accessToken && checkNotExpired(accessToken)) {
            return true;
        }

        const refreshedToken = await refresh();
        return !!refreshedToken;
    };

    const useLogout = async () => {
        try {
            setAccessToken("");
            meState.setMe(EmptyMe);

            const { status } = await logout({
                withCredentials: true,
            });
            if (status === 204) {
                return true;
            }
        } catch (e) {
            console.error(e);
        }
        return false;
    };

    const skipAuth = (shouldSkip: boolean) => {
        setSkipAuth(shouldSkip);
    };

    return {
        accessToken: accessToken,
        hasAccessToken: hasAccessToken,
        refresh: refresh,
        firstRefresh: firstRefresh,
        isAuthenticated: isAuthenticated,
        logout: useLogout,

        shouldSkipAuth: isAuthSkipped,
        skipAuth: skipAuth,

        me: meState,
    } as AuthState;
}

export function checkNotExpired(token: string) {
    try {
        const decoded = jwtDecode(token);
        return !!decoded.exp && Date.now() < decoded.exp * 1000;
    } catch (e) {
        return false;
    }
}
