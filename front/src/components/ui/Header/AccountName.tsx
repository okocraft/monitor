import { UserCircleIcon } from "@heroicons/react/24/solid";
import { TruncatedText } from "../../common/TruncatedText.tsx";

export type AccountProps = {
    name: string;
};

export const AccountName = (props: AccountProps) => {
    return (
        <div className="flex items-center justify-between mx-3 my-1">
            <UserCircleIcon className="size-8 mr-2" />
            <div className="text-xl text-gray-600 w-30 overflow-hidden whitespace-nowrap text-ellipsis">
                <TruncatedText text={props.name} maxLength={16} />
            </div>
        </div>
    );
};
