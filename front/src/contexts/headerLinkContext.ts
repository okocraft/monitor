import { createContext } from "react";
import type { HeaderLink } from "../types/header/link.ts";

export const HeaderLinkContext = createContext<HeaderLink[]>([]);
