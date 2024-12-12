import { type ReactNode, useEffect } from "react";
import axios from "axios";
import { useAuth } from "./useAuth.ts";
import { checkNotExpired } from "../types/auth.ts";

export function AxiosClientProvider({ children }: { children: ReactNode }) {
    const { accessToken, refresh, shouldSkipAuth } = useAuth();

    useEffect(() => {
        const requestIntercept = axios.interceptors.request.use(
            async (config) => {
                if (
                    !config.url ||
                    config.url === "/auth/refresh" ||
                    shouldSkipAuth
                ) {
                    return config;
                }

                if (accessToken && checkNotExpired(accessToken)) {
                    config.headers.Authorization = `Bearer ${accessToken}`;
                    return config;
                }

                const refreshedToken = await refresh();
                if (refreshedToken) {
                    config.headers.Authorization = `Bearer ${refreshedToken}`;
                    return config;
                }

                // TODO: データ入力するページでは強制的に飛ばないようにしたい
                window.location.assign(
                    `/google/login?redirectTo=${encodeURIComponent(window.location.pathname)}`,
                );

                throw new axios.Cancel("No access token available.");
            },
            (error) => Promise.reject(error),
        );

        const responseIntercept = axios.interceptors.response.use(
            (response) => response,
            async (error) => {
                if (
                    error.response?.status === 401 &&
                    error.config.url !== "/auth/refresh"
                ) {
                    const refreshedToken = await refresh();
                    if (refreshedToken) {
                        error.config.headers.Authorization = `Bearer ${refreshedToken}`;
                        return axios(error.config);
                    }
                }
                return Promise.reject(error);
            },
        );

        return () => {
            axios.interceptors.request.eject(requestIntercept);
            axios.interceptors.response.eject(responseIntercept);
        };
    }, [accessToken, refresh, shouldSkipAuth]);

    return <>{children}</>;
}
