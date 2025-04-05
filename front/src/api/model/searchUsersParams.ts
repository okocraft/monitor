import type { SortType } from "./sortType";
import type { SortableUserDataType } from "./sortableUserDataType";
/**
 * Generated by orval v7.8.0 🍺
 * Do not edit manually.
 * Monitor API
 * OpenAPI spec version: 0.0.0
 */
import type { Uuid } from "./uuid";

export type SearchUsersParams = {
    nickname?: string;
    last_access_before?: string;
    last_access_after?: string;
    role_id?: Uuid;
    sorted_by?: SortableUserDataType;
    sort_type?: SortType;
};
