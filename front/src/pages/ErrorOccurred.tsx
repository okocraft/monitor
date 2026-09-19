import type { ErrorComponentProps } from "@tanstack/react-router";

export const ErrorOccurred = ({ props }: { props: ErrorComponentProps }) => {
    const message =
        props.error instanceof Error
            ? props.error.message
            : String(props.error);

    return (
        <>
            <p>An error occurred</p>
            <p>Message: {message}</p>
        </>
    );
};
