import { useState } from "react";
import { getMe } from "../api/me/me.ts";

export type Me = {
    uuid: string;
    nickname: string;
};

export interface MeState {
    current: Me | undefined;
    getMe: () => Promise<Me>;
    setMe: (me: Me | undefined) => void;
    refresh: () => Promise<Me>;
}

export function createMeState() {
    const [me, setMe] = useState<Me | undefined>();

    const refreshMe = async () => {
        try {
            const { data, status } = await getMe();
            if (status === 200 && data) {
                setMe({
                    nickname: data.nickname,
                    uuid: data.uuid,
                } as Me);
                return;
            }
            setMe(undefined);
        } catch {
            setMe(undefined);
        }
        return me;
    };

    const getMeOrRefreshMe = async () => {
        if (!me) {
            await refreshMe();
        }
        return me;
    };

    return {
        current: me,
        getMe: getMeOrRefreshMe,
        setMe: setMe,
        refresh: refreshMe,
    } as MeState;
}
