import { ChevronDownIcon } from "@heroicons/react/24/solid";
import { Link } from "@tanstack/react-router";
import { useState } from "react";
import type {
    HeaderLink,
    NestedHeaderLink,
} from "../../../types/header/link.ts";

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
                <div className={textStyle}>{props.link.name}</div>
            </Link>
        </div>
    );
};

export type DropDownMenuLinkProps = {
    link: HeaderLink;
    nestedLinks: NestedHeaderLink[];
};

export const DropDownMenuLink = (props: DropDownMenuLinkProps) => {
    const [isDropdownOpen, setIsDropdownOpen] = useState(false);
    let hoverTimeout: number;

    const handleMouseEnter = () => {
        clearTimeout(hoverTimeout);
        setIsDropdownOpen(true);
    };

    const handleMouseLeave = () => {
        hoverTimeout = setTimeout(() => {
            setIsDropdownOpen(false);
        }, 50);
    };

    return (
        <div
            className={`${rootLinkStyle} group`}
            onMouseEnter={handleMouseEnter}
            onMouseLeave={handleMouseLeave}
        >
            <div className="relative">
                <Link to={props.link.link} className="flex">
                    <div className={textStyle}>{props.link.name}</div>
                    <ChevronDownIcon className="fill-gray-600 size-4 my-auto ml-1 relative top-0.5" />
                </Link>
                {isDropdownOpen && (
                    <ul className="absolute left-0 mt-3 bg-white text-gray-800 shadow-lg rounded w-64">
                        {props.nestedLinks.map((link) => (
                            <li
                                key={link.id}
                                className="px-4 py-2 hover:bg-gray-100"
                            >
                                <Link to={link.link}>{link.name}</Link>
                            </li>
                        ))}
                    </ul>
                )}
            </div>
        </div>
    );
};
