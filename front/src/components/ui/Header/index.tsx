import { Link } from "@tanstack/react-router";
import { useAuth } from "../../../hooks/useAuth.ts";
import { useHeaderLink } from "../../../hooks/useHeaderLink.ts";
import { type HeaderLink, filterChildren } from "../../../types/header/link.ts";
import type { Me } from "../../../types/me.ts";
import { AccountName } from "./AccountName.tsx";
import { DropDownMenuLink, SingleHeaderLink } from "./Link.tsx";
import { LoginButton } from "./LoginButton.tsx";
import { LogoAndName } from "./LogoAndName.tsx";

export const Header = () => {
    const auth = useAuth();
    const headerLinks = useHeaderLink();
    return (
        <header className="bg-gray-100 shadow flex py-1.5">
            <div className="ml-3 my-auto">
                <LogoAndName />
            </div>

            <nav className="my-auto flex ml-3">
                {headerLinks
                    .map((link) => createHeaderLink(auth.me.current, link))
                    .filter((e) => !!e)}
            </nav>

            <div className="ml-auto my-auto mr-5">
                {auth.me.current ? (
                    <AccountName name={auth.me.current.nickname} />
                ) : (
                    <Link to="/google/login">
                        <LoginButton />
                    </Link>
                )}
            </div>
        </header>
    );
};

function createHeaderLink(me: Me | undefined, link: HeaderLink) {
    if (!link.canView(me)) {
        return undefined;
    }

    const filtered = filterChildren(link, me);
    return (
        <div key={`header-links-${link.id}`} className="ml-2">
            {filtered.nestedLinks && 0 < filtered.nestedLinks.length ? (
                <DropDownMenuLink
                    link={filtered}
                    nestedLinks={filtered.nestedLinks}
                />
            ) : (
                <SingleHeaderLink link={filtered} />
            )}
        </div>
    );
}
