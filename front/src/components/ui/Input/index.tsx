import type {
    ChangeEventHandler,
    HTMLInputTypeAttribute,
    MouseEventHandler,
    ReactNode,
} from "react";
import styles from "./input.module.css";

export type Props = {
    id: string;
    label: string;
    hideLabel?: boolean;
    placeholder: string;
    type?: HTMLInputTypeAttribute;

    value: string;
    defaultValue?: string;
    onChange?: ChangeEventHandler<HTMLInputElement>;
    onClick?: MouseEventHandler<HTMLInputElement>;
    disabled?: boolean;

    variant: keyof typeof variants;
    className?: string;
    children?: ReactNode;
};

export const Input = (props: Props) => {
    return (
        <div className={styles.inputContainer}>
            <label htmlFor={props.id} className={styles.inputLabel}>
                {props.hideLabel ? "" : props.label}
            </label>
            <input
                id={props.id}
                aria-label={props.label}
                placeholder={props.placeholder}
                type={props.type}
                value={props.value}
                defaultValue={props.defaultValue}
                onChange={props.onChange}
                onClick={props.onClick}
                disabled={props.disabled}
                className={`${styles.base} ${variants[props.variant]} ${props.className}`}
            />
            {props.children}
        </div>
    );
};

const variants = {
    filled: styles.filled,
    outlined: styles.outlined,
} as const;
