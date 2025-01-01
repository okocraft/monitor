import { useState } from "react";
import type { PagePermissions } from "../api/model";

export interface PagePermissionState {
    current: PagePermissions | undefined;
    setPagePermissions: (pagePermissions?: PagePermissions) => void;
}

export function createPagePermissionState() {
    const [pagePermissions, setPagePermissions] = useState<
        PagePermissions | undefined
    >();

    return {
        current: pagePermissions,
        setPagePermissions: setPagePermissions,
    } as PagePermissionState;
}
