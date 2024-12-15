import { ArrowRightIcon } from "@heroicons/react/24/solid";

export const LoginButton = () => {
    return (
        <div className="flex items-center justify-between mx-2 my-1">
            <div className="text-xl text-gray-600 whitespace-nowrap">
                Log in
            </div>
            <ArrowRightIcon className="size-4 ml-1" />
        </div>
    );
};
