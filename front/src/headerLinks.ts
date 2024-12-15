import type { HeaderLink } from "./types/header/link.ts";

export const headerLinks: HeaderLink[] = [
    {
        name: "Example",
        link: "/example",
        canView: (me) => !!me,
    },
];
