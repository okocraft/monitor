import { type ReactNode, useState } from "react";

export type Props = {
    display: ReactNode;
    elements: {
        id: string;
        node: ReactNode;
        className?: string;
    }[];
    className?: string;
};

export const DropDownMenu = (props: Props) => {
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
            className={"group"}
            onMouseEnter={handleMouseEnter}
            onMouseLeave={handleMouseLeave}
            role="menu"
        >
            <div className="relative">
                {props.display}
                {isDropdownOpen && (
                    <ul className={`absolute ${props.className}`}>
                        {props.elements.map((e) => (
                            <li
                                key={e.id}
                                className={`hover:bg-gray-200 ${e.className}`}
                            >
                                {e.node}
                            </li>
                        ))}
                    </ul>
                )}
            </div>
        </div>
    );
};
