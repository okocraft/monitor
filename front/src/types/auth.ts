import { jwtDecode } from "jwt-decode";
import { useRef, useState } from "react";
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
                roleUuid: "",
                roleName: "",
            } as Me;
        },
        setMe: (_) => {},
        refresh: async () => {
            return {
                nickname: "",
                uuid: "",
                roleUuid: "",
                roleName: "",
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
    const accessTokenRef = useRef(accessToken);
    const refreshed = useRef(false);
    const [isAuthSkipped, setSkipAuth] = useState(false);
    const meState = createMeState();
    const pagePermissionState = createPagePermissionState();

    const updateAccessToken = (accessToken: string) => {
        setAccessToken(accessToken);
        accessTokenRef.current = accessToken;
        refreshed.current = accessToken !== "";
    };

    const refresh = async () => {
        refreshed.current = true;

        try {
            const { data, status } = await refreshAccessToken({
                withCredentials: true,
            });
            if (status === 200) {
                updateAccessToken(data.access_token);
                meState.setMe({
                    nickname: data.me.nickname,
                    uuid: data.me.uuid,
                    roleUuid: data.me.role_uuid,
                    roleName: data.me.role_name,
                } as Me);
                pagePermissionState.setPagePermissions(data.page_permissions);
                return data.access_token;
            }
            updateAccessToken("");
            meState.setMe(undefined);
            pagePermissionState.setPagePermissions(undefined);
            return "";
        } catch {
            updateAccessToken("");
            meState.setMe(undefined);
            pagePermissionState.setPagePermissions(undefined);
            return "";
        }
    };

    const firstRefresh = async () => {
        if (!refreshed.current) {
            await refresh();
        }
    };

    const hasAccessToken = () => {
        return !!accessTokenRef.current;
    };

    const isAuthenticated = async () => {
        if (accessTokenRef.current) {
            if (checkNotExpired(accessTokenRef.current)) {
                return true;
            }
            updateAccessToken("");
        }

        if (refreshed.current) {
            return false;
        }

        const refreshedToken = await refresh();
        return !!refreshedToken;
    };

    const useLogout = async () => {
        try {
            updateAccessToken("");
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
