import { ArrowTopRightOnSquareIcon } from "@heroicons/react/24/outline";
import { useLoginWithGoogle } from "../../../api/auth/auth.ts";
import { Button } from "../../../components/ui/Button";
import { Text } from "../../../components/ui/Text";
import { SkipAuth } from "../../../hooks/SkipAuth.tsx";

export const Component = ({ redirectTo }: { redirectTo?: string }) => {
    const { mutate, isPending } = useLoginWithGoogle({
        mutation: {
            onSuccess: (res) => {
                window.location.href = res.data.redirect_url;
            },
            onError: (error) => {
                console.error("Failed to get redirect URL:", error);
            },
        },
    });

    const handleLogin = () => {
        mutate({
            data: {
                url: redirectTo ? redirectTo : "",
            },
        });
    };

    return (
        <SkipAuth>
            <div className="m-5">
                <Text className="mb-2 block" type="large">
                    Our service uses social login to authenticate accounts.
                    <br />
                    For new accounts, please request the administrator to
                    provide a login URL.
                </Text>
                <Button
                    type={"button"}
                    variant={"filled"}
                    onClick={handleLogin}
                    disabled={isPending}
                    content={
                        <div className="ml-1 my-auto flex">
                            Login with Google account
                            <ArrowTopRightOnSquareIcon className="w-5 h-5 ml-1 my-auto" />
                        </div>
                    }
                />
            </div>
        </SkipAuth>
    );
};
