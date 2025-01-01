/**
 * Generated by orval v7.3.0 🍺
 * Do not edit manually.
 * Monitor API
 * OpenAPI spec version: 0.0.0
 */
import { useMutation, useQuery } from "@tanstack/react-query";
import type {
    DataTag,
    DefinedInitialDataOptions,
    DefinedUseQueryResult,
    MutationFunction,
    QueryFunction,
    QueryKey,
    UndefinedInitialDataOptions,
    UseMutationOptions,
    UseMutationResult,
    UseQueryOptions,
    UseQueryResult,
} from "@tanstack/react-query";
import * as axios from "axios";
import type { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import type {
    AccessTokenWithMeAndPagePermissions,
    CurrentPage,
    GoogleLoginURL,
} from ".././model";

/**
 * Callback from Google
 */
export const callbackFromGoogle = (
    options?: AxiosRequestConfig,
): Promise<AxiosResponse<unknown>> => {
    return axios.default.get(`/auth/google/callback`, options);
};

export const getCallbackFromGoogleQueryKey = () => {
    return [`/auth/google/callback`] as const;
};

export const getCallbackFromGoogleQueryOptions = <
    TData = Awaited<ReturnType<typeof callbackFromGoogle>>,
    TError = AxiosError<void>,
>(options?: {
    query?: Partial<
        UseQueryOptions<
            Awaited<ReturnType<typeof callbackFromGoogle>>,
            TError,
            TData
        >
    >;
    axios?: AxiosRequestConfig;
}) => {
    const { query: queryOptions, axios: axiosOptions } = options ?? {};

    const queryKey = queryOptions?.queryKey ?? getCallbackFromGoogleQueryKey();

    const queryFn: QueryFunction<
        Awaited<ReturnType<typeof callbackFromGoogle>>
    > = ({ signal }) => callbackFromGoogle({ signal, ...axiosOptions });

    return { queryKey, queryFn, ...queryOptions } as UseQueryOptions<
        Awaited<ReturnType<typeof callbackFromGoogle>>,
        TError,
        TData
    > & { queryKey: DataTag<QueryKey, TData> };
};

export type CallbackFromGoogleQueryResult = NonNullable<
    Awaited<ReturnType<typeof callbackFromGoogle>>
>;
export type CallbackFromGoogleQueryError = AxiosError<void>;

export function useCallbackFromGoogle<
    TData = Awaited<ReturnType<typeof callbackFromGoogle>>,
    TError = AxiosError<void>,
>(options: {
    query: Partial<
        UseQueryOptions<
            Awaited<ReturnType<typeof callbackFromGoogle>>,
            TError,
            TData
        >
    > &
        Pick<
            DefinedInitialDataOptions<
                Awaited<ReturnType<typeof callbackFromGoogle>>,
                TError,
                TData
            >,
            "initialData"
        >;
    axios?: AxiosRequestConfig;
}): DefinedUseQueryResult<TData, TError> & {
    queryKey: DataTag<QueryKey, TData>;
};
export function useCallbackFromGoogle<
    TData = Awaited<ReturnType<typeof callbackFromGoogle>>,
    TError = AxiosError<void>,
>(options?: {
    query?: Partial<
        UseQueryOptions<
            Awaited<ReturnType<typeof callbackFromGoogle>>,
            TError,
            TData
        >
    > &
        Pick<
            UndefinedInitialDataOptions<
                Awaited<ReturnType<typeof callbackFromGoogle>>,
                TError,
                TData
            >,
            "initialData"
        >;
    axios?: AxiosRequestConfig;
}): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData> };
export function useCallbackFromGoogle<
    TData = Awaited<ReturnType<typeof callbackFromGoogle>>,
    TError = AxiosError<void>,
>(options?: {
    query?: Partial<
        UseQueryOptions<
            Awaited<ReturnType<typeof callbackFromGoogle>>,
            TError,
            TData
        >
    >;
    axios?: AxiosRequestConfig;
}): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData> };

export function useCallbackFromGoogle<
    TData = Awaited<ReturnType<typeof callbackFromGoogle>>,
    TError = AxiosError<void>,
>(options?: {
    query?: Partial<
        UseQueryOptions<
            Awaited<ReturnType<typeof callbackFromGoogle>>,
            TError,
            TData
        >
    >;
    axios?: AxiosRequestConfig;
}): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData> } {
    const queryOptions = getCallbackFromGoogleQueryOptions(options);

    const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & {
        queryKey: DataTag<QueryKey, TData>;
    };

    query.queryKey = queryOptions.queryKey;

    return query;
}

/**
 * First login with Google Account
 */
export const linkWithGoogle = (
    loginKey: string,
    options?: AxiosRequestConfig,
): Promise<AxiosResponse<GoogleLoginURL>> => {
    return axios.default.post(
        `/auth/google/link/${loginKey}`,
        undefined,
        options,
    );
};

export const getLinkWithGoogleMutationOptions = <
    TError = AxiosError<void>,
    TContext = unknown,
>(options?: {
    mutation?: UseMutationOptions<
        Awaited<ReturnType<typeof linkWithGoogle>>,
        TError,
        { loginKey: string },
        TContext
    >;
    axios?: AxiosRequestConfig;
}): UseMutationOptions<
    Awaited<ReturnType<typeof linkWithGoogle>>,
    TError,
    { loginKey: string },
    TContext
> => {
    const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

    const mutationFn: MutationFunction<
        Awaited<ReturnType<typeof linkWithGoogle>>,
        { loginKey: string }
    > = (props) => {
        const { loginKey } = props ?? {};

        return linkWithGoogle(loginKey, axiosOptions);
    };

    return { mutationFn, ...mutationOptions };
};

export type LinkWithGoogleMutationResult = NonNullable<
    Awaited<ReturnType<typeof linkWithGoogle>>
>;

export type LinkWithGoogleMutationError = AxiosError<void>;

export const useLinkWithGoogle = <
    TError = AxiosError<void>,
    TContext = unknown,
>(options?: {
    mutation?: UseMutationOptions<
        Awaited<ReturnType<typeof linkWithGoogle>>,
        TError,
        { loginKey: string },
        TContext
    >;
    axios?: AxiosRequestConfig;
}): UseMutationResult<
    Awaited<ReturnType<typeof linkWithGoogle>>,
    TError,
    { loginKey: string },
    TContext
> => {
    const mutationOptions = getLinkWithGoogleMutationOptions(options);

    return useMutation(mutationOptions);
};
/**
 * Login with Google Account
 */
export const loginWithGoogle = (
    currentPage: CurrentPage,
    options?: AxiosRequestConfig,
): Promise<AxiosResponse<GoogleLoginURL>> => {
    return axios.default.post(`/auth/google/login`, currentPage, options);
};

export const getLoginWithGoogleMutationOptions = <
    TError = AxiosError<void>,
    TContext = unknown,
>(options?: {
    mutation?: UseMutationOptions<
        Awaited<ReturnType<typeof loginWithGoogle>>,
        TError,
        { data: CurrentPage },
        TContext
    >;
    axios?: AxiosRequestConfig;
}): UseMutationOptions<
    Awaited<ReturnType<typeof loginWithGoogle>>,
    TError,
    { data: CurrentPage },
    TContext
> => {
    const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

    const mutationFn: MutationFunction<
        Awaited<ReturnType<typeof loginWithGoogle>>,
        { data: CurrentPage }
    > = (props) => {
        const { data } = props ?? {};

        return loginWithGoogle(data, axiosOptions);
    };

    return { mutationFn, ...mutationOptions };
};

export type LoginWithGoogleMutationResult = NonNullable<
    Awaited<ReturnType<typeof loginWithGoogle>>
>;
export type LoginWithGoogleMutationBody = CurrentPage;
export type LoginWithGoogleMutationError = AxiosError<void>;

export const useLoginWithGoogle = <
    TError = AxiosError<void>,
    TContext = unknown,
>(options?: {
    mutation?: UseMutationOptions<
        Awaited<ReturnType<typeof loginWithGoogle>>,
        TError,
        { data: CurrentPage },
        TContext
    >;
    axios?: AxiosRequestConfig;
}): UseMutationResult<
    Awaited<ReturnType<typeof loginWithGoogle>>,
    TError,
    { data: CurrentPage },
    TContext
> => {
    const mutationOptions = getLoginWithGoogleMutationOptions(options);

    return useMutation(mutationOptions);
};
/**
 * Invalidate refresh_token and access_token
 */
export const logout = (
    options?: AxiosRequestConfig,
): Promise<AxiosResponse<void>> => {
    return axios.default.post(`/auth/logout`, undefined, options);
};

export const getLogoutMutationOptions = <
    TError = AxiosError<void>,
    TContext = unknown,
>(options?: {
    mutation?: UseMutationOptions<
        Awaited<ReturnType<typeof logout>>,
        TError,
        void,
        TContext
    >;
    axios?: AxiosRequestConfig;
}): UseMutationOptions<
    Awaited<ReturnType<typeof logout>>,
    TError,
    void,
    TContext
> => {
    const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

    const mutationFn: MutationFunction<
        Awaited<ReturnType<typeof logout>>,
        void
    > = () => {
        return logout(axiosOptions);
    };

    return { mutationFn, ...mutationOptions };
};

export type LogoutMutationResult = NonNullable<
    Awaited<ReturnType<typeof logout>>
>;

export type LogoutMutationError = AxiosError<void>;

export const useLogout = <
    TError = AxiosError<void>,
    TContext = unknown,
>(options?: {
    mutation?: UseMutationOptions<
        Awaited<ReturnType<typeof logout>>,
        TError,
        void,
        TContext
    >;
    axios?: AxiosRequestConfig;
}): UseMutationResult<
    Awaited<ReturnType<typeof logout>>,
    TError,
    void,
    TContext
> => {
    const mutationOptions = getLogoutMutationOptions(options);

    return useMutation(mutationOptions);
};
/**
 * Refresh access token
 */
export const refreshAccessToken = (
    options?: AxiosRequestConfig,
): Promise<AxiosResponse<AccessTokenWithMeAndPagePermissions>> => {
    return axios.default.post(`/auth/refresh`, undefined, options);
};

export const getRefreshAccessTokenMutationOptions = <
    TError = AxiosError<void>,
    TContext = unknown,
>(options?: {
    mutation?: UseMutationOptions<
        Awaited<ReturnType<typeof refreshAccessToken>>,
        TError,
        void,
        TContext
    >;
    axios?: AxiosRequestConfig;
}): UseMutationOptions<
    Awaited<ReturnType<typeof refreshAccessToken>>,
    TError,
    void,
    TContext
> => {
    const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

    const mutationFn: MutationFunction<
        Awaited<ReturnType<typeof refreshAccessToken>>,
        void
    > = () => {
        return refreshAccessToken(axiosOptions);
    };

    return { mutationFn, ...mutationOptions };
};

export type RefreshAccessTokenMutationResult = NonNullable<
    Awaited<ReturnType<typeof refreshAccessToken>>
>;

export type RefreshAccessTokenMutationError = AxiosError<void>;

export const useRefreshAccessToken = <
    TError = AxiosError<void>,
    TContext = unknown,
>(options?: {
    mutation?: UseMutationOptions<
        Awaited<ReturnType<typeof refreshAccessToken>>,
        TError,
        void,
        TContext
    >;
    axios?: AxiosRequestConfig;
}): UseMutationResult<
    Awaited<ReturnType<typeof refreshAccessToken>>,
    TError,
    void,
    TContext
> => {
    const mutationOptions = getRefreshAccessTokenMutationOptions(options);

    return useMutation(mutationOptions);
};
