import type { ReactNode } from "react";
import { HeaderLinkContext } from "../contexts/headerLinkContext.ts";
import type { HeaderLink } from "../types/header/link.ts";

export function HeaderLinkProvider({
    children,
    links,
}: {
    children: ReactNode;
    links: HeaderLink[];
}) {
    return <HeaderLinkContext value={links}>{children}</HeaderLinkContext>;
}
