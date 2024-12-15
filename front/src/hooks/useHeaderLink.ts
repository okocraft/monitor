import { useContext } from "react";
import { HeaderLinkContext } from "../contexts/headerLinkContext.ts";

export const useHeaderLink = () => {
    return useContext(HeaderLinkContext);
};
