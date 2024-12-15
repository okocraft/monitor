import type { ToPathOption } from "@tanstack/react-router";
import type { RoutePaths } from "@tanstack/react-router";
import type { AnyRouter, RegisteredRouter } from "@tanstack/react-router";
import type { Me } from "../me.ts";

export type HeaderLink<
    in out TRouter extends AnyRouter = RegisteredRouter,
    in out TFrom extends RoutePaths<TRouter["routeTree"]> | string = string,
    in out TTo extends string | undefined = ".",
> = {
    name: string;
    link?: ToPathOption<TRouter, TFrom, TTo> & {};
    nestedLinks?: NestedHeaderLink[];
    canView: (me: Me | undefined) => boolean;
};

export type NestedHeaderLink<
    in out TRouter extends AnyRouter = RegisteredRouter,
    in out TFrom extends RoutePaths<TRouter["routeTree"]> | string = string,
    in out TTo extends string | undefined = ".",
> = {
    id: string;
    name: string;
    link?: ToPathOption<TRouter, TFrom, TTo> & {};
    canView: (me: Me | undefined) => boolean;
};

export function filterChildren(link: HeaderLink, me?: Me): HeaderLink {
    const filtered = link.nestedLinks?.filter((child) => child.canView(me));
    return {
        name: link.name,
        link: link.link,
        canView: link.canView,
        nestedLinks: filtered && 0 < filtered.length ? filtered : [],
    } as HeaderLink;
}