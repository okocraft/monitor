import type { ErrorComponentProps } from "@tanstack/react-router";

export const ErrorOccurred = ({ props }: { props: ErrorComponentProps }) => {
    return (
        <>
            <p>An error occurred</p>
            <p>Message: {props.error.message}</p>
        </>
    );
};
