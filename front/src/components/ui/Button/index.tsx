import type { MouseEventHandler, ReactNode } from "react";
import styles from "./button.module.css";

export type Props = {
    type: "submit" | "reset" | "button";
    onClick?: MouseEventHandler<HTMLButtonElement>;
    disabled?: boolean;

    icon?: SVGElement;
    content?: ReactNode;
    variant: keyof typeof variants;
    className?: string;
};

export const Button = (props: Props) => {
    return (
        <button
            type={props.type}
            onClick={props.onClick}
            disabled={props.disabled}
            className={`${styles.base} ${variants[props.variant]} ${props.className}`}
        >
            {!props.icon && props.icon}
            {props.content}
        </button>
    );
};

const variants = {
    filled: styles.filled,
    tonal: styles.tonal,
    outlined: styles.outlined,
} as const;
