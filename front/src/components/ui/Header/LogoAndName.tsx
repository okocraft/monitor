import { Link } from "@tanstack/react-router";
import icon from "../../../../assets/icon.png";

export const LogoAndName = () => {
    return (
        <div className="flex">
            <Link to="/">
                <div className="items-center justify-between flex p-1 rounded-lg hover:bg-gray-200 ease-in transition-colors duration-100">
                    <img src={icon} alt="icon" className="size-9 mr-3" />
                    <div className="text-2xl font-bold tracking-wide text-gray-600 my-auto">
                        Monitor
                    </div>
                </div>
            </Link>
        </div>
    );
};
