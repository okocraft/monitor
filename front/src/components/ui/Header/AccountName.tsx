import { UserCircleIcon } from "@heroicons/react/24/solid";
import { Link, useNavigate } from "@tanstack/react-router";
import { useAuth } from "../../../hooks/useAuth.ts";
import { DropDownMenu } from "../../common/DropDownMenu.tsx";
import { TruncatedText } from "../../common/TruncatedText.tsx";

export type AccountProps = {
    name: string;
};

export const AccountName = (props: AccountProps) => {
    const auth = useAuth();
    const navigate = useNavigate();

    const display = (
        <div className="flex items-center justify-between px-2 py-2 rounded-lg my-auto hover:bg-gray-200">
            <UserCircleIcon className="size-8 mr-2" />
            <span className="text-xl text-gray-600 max-w-30 overflow-hidden whitespace-nowrap text-ellipsis">
                <TruncatedText text={props.name} maxLength={16} />
            </span>
        </div>
    );

    const handleLogout = () => {
        auth.logout()
            .then(() => {
                navigate({
                    to: "/",
                }).catch((err: Error) => {
                    console.error("Navigation error:", err);
                });
            })
            .catch((err: Error) => {
                console.log(err); // TODO
            });
    };

    const elements = [
        {
            id: "account-name.my-page",
            node: (
                <Link to="/mypage">
                    <div className="w-full h-full px-4 py-2">
                        <span>My page</span>
                    </div>
                </Link>
            ),
        },
        {
            id: "account-name.logout",
            node: (
                <button
                    type="submit"
                    className="text-red-500 cursor-pointer w-full h-full px-4 py-2"
                    onClick={handleLogout}
                >
                    Logout
                </button>
            ),
        },
    ];

    return (
        <div className="mx-3">
            <DropDownMenu
                display={display}
                elements={elements}
                className="bg-white text-center text-gray-800 shadow-lg rounded w-36 right-0"
            />
        </div>
    );
};
