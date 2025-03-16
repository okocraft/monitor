import type { ReactNode } from "react";
import styles from "./text.module.css";

export type Props = {
    type?: keyof typeof types;
    color?: keyof typeof colors;
    className?: string;
    children: ReactNode;
};

export const Text = (props: Props) => {
    return (
        <span
            className={`${props.type ? types[props.type] : types.base} ${props.color ? colors[props.color] : colors.base} ${props.className}`}
        >
            {props.children}
        </span>
    );
};

const types = {
    large: styles.largeSize,
    base: styles.baseSize,
    small: styles.smallSize,
} as const;

const colors = {
    base: styles.baseColor,
    red: styles.redColor,
} as const;
