import type { ReactNode } from "react";
import styles from "./text.module.css";

export type Props = {
    type?: keyof typeof types;
    color?: keyof typeof colors;
    hoverColor?: keyof typeof hoverColors;
    className?: string;
    children: ReactNode;
};

export const Text = (props: Props) => {
    const type = props.type ? types[props.type] : types.base;
    const color = props.color ? colors[props.color] : colors.base;
    const hoverColor = props.hoverColor ? hoverColors[props.hoverColor] : "";
    const className = props.className ? props.className : "";
    return (
        <span className={`${type} ${color} ${hoverColor} ${className}`}>
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
    primary: styles.primaryColor,
} as const;

const hoverColors = {
    primaryDark: styles.primaryDarkHoverColor,
} as const;
