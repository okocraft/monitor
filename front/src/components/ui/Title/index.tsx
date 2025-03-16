import type { ReactNode } from "react";
import styles from "./title.module.css";

export type Props = {
    type?: keyof typeof types;
    color?: keyof typeof colors;
    className?: string;
    children: ReactNode;
};

export const Title = (props: Props) => {
    const type = props.type ? types[props.type] : types.base;
    return type.tag(
        props.children,
        `${type.style} ${props.color ? colors[props.color] : colors.base} ${props.className}`,
    );
};

type type = {
    tag: (children: ReactNode, className?: string) => ReactNode;
    style: string;
};

const types = {
    large: {
        tag: (children, className) => <h1 className={className}>{children}</h1>,
        style: styles.largeSize,
    },
    base: {
        tag: (children, className) => <h2 className={className}>{children}</h2>,
        style: styles.smallSize,
    },
    small: {
        tag: (children, className) => <h3 className={className}>{children}</h3>,
        style: styles.smallSize,
    },
} as const satisfies Record<string, type>;

const colors = {
    base: styles.baseColor,
    red: styles.redColor,
};
