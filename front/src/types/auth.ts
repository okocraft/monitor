import { jwtDecode } from "jwt-decode";
import { useState } from "react";
import { logout, refreshAccessToken } from "../api/auth/auth.ts";
import { type Me, type MeState, createMeState } from "./me.ts";
import {
    type PagePermissionState,
    createPagePermissionState,
} from "./pagePermission.ts";

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

    pagePermission: PagePermissionState;
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
        getMe: async () => {
            return {
                nickname: "",
                uuid: "",
            } as Me;
        },
        setMe: (_) => {},
        refresh: async () => {
            return {
                nickname: "",
                uuid: "",
            } as Me;
        },
    } as MeState,

    pagePermission: {
        current: undefined,
        setPagePermissions: (_) => {},
    } as PagePermissionState,
} as AuthState;

export function createAuthState() {
    const [accessToken, setAccessToken] = useState<string>("");
    const [refreshed, setRefreshed] = useState<boolean>(false);
    const [isAuthSkipped, setSkipAuth] = useState(false);
    const meState = createMeState();
    const pagePermissionState = createPagePermissionState();

    const refresh = async () => {
        setRefreshed(true);

        try {
            const { data, status } = await refreshAccessToken({
                withCredentials: true,
            });
            if (status === 200) {
                setAccessToken(data.access_token);
                meState.setMe(data.me);
                pagePermissionState.setPagePermissions(data.page_permissions);
                return data.access_token;
            }
            setAccessToken("");
            meState.setMe(undefined);
            pagePermissionState.setPagePermissions(undefined);
            return "";
        } catch {
            setAccessToken("");
            meState.setMe(undefined);
            pagePermissionState.setPagePermissions(undefined);
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
        if (accessToken) {
            if (checkNotExpired(accessToken)) {
                return true;
            }
            setRefreshed(false);
        }

        if (refreshed) {
            return false;
        }

        const refreshedToken = await refresh();
        return !!refreshedToken;
    };

    const useLogout = async () => {
        try {
            setAccessToken("");
            meState.setMe(undefined);
            pagePermissionState.setPagePermissions(undefined);

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
        pagePermission: pagePermissionState,
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
