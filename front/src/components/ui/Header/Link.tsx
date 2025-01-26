import { ChevronDownIcon } from "@heroicons/react/24/solid";
import { Link } from "@tanstack/react-router";
import type {
    HeaderLink,
    NestedHeaderLink,
} from "../../../types/header/link.ts";
import { DropDownMenu } from "../../common/DropDownMenu.tsx";

export type HeaderLinkProps = {
    link: HeaderLink;
};

const rootLinkStyle =
    "my-auto items-center justify-between flex px-2 py-2 rounded-lg hover:bg-gray-200 ease-in transition-colors duration-100";
const textStyle = "text-xl text-gray-800";

export const SingleHeaderLink = (props: HeaderLinkProps) => {
    return (
        <div className={rootLinkStyle}>
            <Link to={props.link.link}>
                <span className={textStyle}>{props.link.name}</span>
            </Link>
        </div>
    );
};

export type DropDownMenuLinkProps = {
    link: HeaderLink;
    nestedLinks: NestedHeaderLink[];
};

export const DropDownMenuLink = (props: DropDownMenuLinkProps) => {
    const display = (
        <div className={rootLinkStyle}>
            <Link to={props.link.link} className="flex">
                <span className={textStyle}>{props.link.name}</span>
                <ChevronDownIcon className="fill-gray-600 size-4 my-auto ml-1 relative top-0.5" />
            </Link>
        </div>
    );

    const elements = props.nestedLinks.map((link) => ({
        id: `header-links-${props.link.id}-${link.id}`,
        node: (
            <Link to={link.link}>
                <div className="w-full h-full px-4 py-2">
                    <span>{link.name}</span>
                </div>
            </Link>
        ),
    }));

    return (
        <DropDownMenu
            display={display}
            elements={elements}
            className="bg-white text-gray-800 shadow-lg rounded w-64 left-0"
        />
    );
};
