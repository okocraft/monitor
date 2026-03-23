import { useLocation } from "@tanstack/react-router";

export function useSearchParam(key: string) {
    const loc = useLocation();
    const param = new URLSearchParams(loc.search).get(key);
    return param ? decodeURIComponent(param) : undefined;
}
