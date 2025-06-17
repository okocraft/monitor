import type {
    AnyRouter,
    RegisteredRouter,
    ToPathOption,
} from "@tanstack/react-router";
import type { PagePermissions } from "../../api/model";
import type { Me } from "../me.ts";

export type HeaderLink<
    in out TRouter extends AnyRouter = RegisteredRouter,
    in out TFrom extends string = string,
    in out TTo extends string | undefined = ".",
> = {
    id: string;
    name: string;
    link?: ToPathOption<TRouter, TFrom, TTo> & {};
    nestedLinks?: NestedHeaderLink[];
    hideWhenNoNestedLinks?: boolean;
    canView: (
        me: Me | undefined,
        pagePermissions: PagePermissions | undefined,
    ) => boolean;
};

export type NestedHeaderLink<
    in out TRouter extends AnyRouter = RegisteredRouter,
    in out TFrom extends string = string,
    in out TTo extends string | undefined = ".",
> = {
    id: string;
    name: string;
    link?: ToPathOption<TRouter, TFrom, TTo> & {};
    canView: (
        me: Me | undefined,
        perms: PagePermissions | undefined,
    ) => boolean;
};

export function filterChildren(
    link: HeaderLink,
    me?: Me,
    perms?: PagePermissions,
): HeaderLink {
    const filtered = link.nestedLinks?.filter((child) =>
        child.canView(me, perms),
    );
    return {
        ...link,
        nestedLinks: filtered && 0 < filtered.length ? filtered : [],
    } as HeaderLink;
}
