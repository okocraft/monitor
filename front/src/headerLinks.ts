import type { HeaderLink } from "./types/header/link.ts";

export const headerLinks: HeaderLink[] = [
    {
        id: "example",
        name: "Example",
        link: "/example",
        canView: (me, _) => !!me,
    },
    {
        id: "settings",
        name: "Settings",
        link: "/settings",
        hideWhenNoNestedLinks: true,
        nestedLinks: [
            {
                id: "users",
                name: "Users",
                link: "/settings/users",
                canView: (_, perms) => perms?.users ?? false,
            },
        ],
        canView: () => true,
    },
];
