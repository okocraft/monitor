import { useState } from "react";
import { useLinkWithGoogle } from "../../../api/auth/auth.ts";
import { Button } from "../../../components/ui/Button";
import { Input } from "../../../components/ui/Input";
import { Text } from "../../../components/ui/Text";
import { Title } from "../../../components/ui/Title";
import { isValidInt64ForHex } from "../../../utils/number.ts";

export const Component = ({ initialLoginKey }: { initialLoginKey: string }) => {
    const [loginKey, setLoginKey] = useState(initialLoginKey);
    const isInvalid = !isValidInt64ForHex(loginKey);

    const [isClicked, setIsClicked] = useState(false);
    const [isNavigating, setIsNavigating] = useState(false);

    const { mutate } = useLinkWithGoogle({
        mutation: {
            onSuccess: (res) => {
                window.location.href = res.data.redirect_url;
                setIsNavigating(false);
            },
            onError: (error) => {
                console.error("Failed to get redirect URL:", error);
                setIsNavigating(false);
            },
        },
    });

    const handleClick = () => {
        setIsClicked(true);
        if (isInvalid) {
            return;
        }

        setIsNavigating(true);

        mutate({
            loginKey,
        });
    };

    return (
        <div>
            <Title type="large" className="m-3 mt-5">
                Link account with Google
            </Title>
            <Text type="base" className="ml-3">
                Enter the login key provided by the administrator.
            </Text>
            <div className="flex mx-3 my-2">
                <div className="flex my-auto">
                    <Input
                        id={"login-key"}
                        label={"Login key"}
                        hideLabel={true}
                        placeholder={"Login key"}
                        value={loginKey}
                        onChange={(e) => setLoginKey(e.target.value)}
                        onClick={(_) => setIsClicked(true)}
                        disabled={isNavigating}
                        variant={isNavigating ? "filled" : "outlined"}
                        className="w-72 h-10"
                    >
                        {loginKey !== "" && isInvalid && (
                            <Text color="red" className="mx-3 my-1">
                                Login key is not valid format.
                            </Text>
                        )}
                        {loginKey === "" && isClicked && (
                            <Text color="red" className="mx-3 my-1">
                                Login key cannot be empty.
                            </Text>
                        )}
                    </Input>
                    <Button
                        type={"button"}
                        content={"Go to login page"}
                        disabled={isInvalid || isNavigating}
                        variant="filled"
                        onClick={handleClick}
                        className="ml-2 mb-auto"
                    />
                </div>
            </div>
        </div>
    );
};
