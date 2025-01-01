import type { HeaderLink } from "./types/header/link.ts";

export const headerLinks: HeaderLink[] = [
    {
        id: "example",
        name: "Example",
        link: "/example",
        canView: (me, _) => !!me,
    },
];
