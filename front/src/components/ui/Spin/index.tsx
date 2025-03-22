import styles from "./spin.module.css";

export type Props = {
    className?: string;
    size?: keyof typeof sizes;
    color?: keyof typeof colors;
    style?: keyof typeof borderStyles;
};

export const Spin = (props: Props) => {
    const size = props.size ? sizes[props.size] : sizes.base;
    const color = props.color ? colors[props.color] : colors.primary;
    const style = props.style ? borderStyles[props.style] : borderStyles.solid;

    return (
        <div
            className={`${styles.base} ${size} ${color} ${style} animate-spin ${props.className ?? ""}`}
        />
    );
};

const sizes = {
    base: styles.baseSize,
} as const;

const colors = {
    primary: styles.primaryColor,
    red: styles.redColor,
} as const;

const borderStyles = {
    solid: styles.solidStyle,
} as const;
