import { Link, useNavigate } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { Button } from "../../../components/ui/Button";
import { Text } from "../../../components/ui/Text";
import { Title } from "../../../components/ui/Title";

export type Props = {
    nickname: string;
};

export const Success = (props: Props) => {
    return (
        <div className="m-3">
            <Title>
                Successfully logged in to <b>{props.nickname}</b>!
            </Title>
            <div className="mt-3">
                <Link to="/mypage">
                    <Button
                        type="button"
                        variant="tonal"
                        content={<span className="px-5">Go to my page</span>}
                    />
                </Link>
            </div>
        </div>
    );
};

export type RedirectingProps = {
    nickname: string;
    url: string;
};

export const RedirectingSuccess = (props: RedirectingProps) => {
    const navigate = useNavigate();

    const waitingSecond = 3;
    const [countdown, setCountdown] = useState(waitingSecond);
    useEffect(() => {
        const interval = setInterval(() => {
            setCountdown((prev: number) => prev - 1);
        }, 1000);

        if (countdown === 0) {
            navigate({
                to: props.url,
            }).catch((err: Error) => {
                console.error("Navigation error:", err);
            });
        }

        return () => clearInterval(interval);
    }, [countdown, navigate, props.url]);

    return (
        <div className="m-3">
            <Title>
                Successfully logged in to <b>{props.nickname}</b>!
            </Title>
            <Text>
                Redirecting to <b>{props.url}</b> in {countdown} seconds...
            </Text>
        </div>
    );
};
