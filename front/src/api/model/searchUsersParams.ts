/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Monitor API
 * OpenAPI spec version: 0.0.0
 */
import type { Uuid } from "./uuid";
import type { SortableUserDataType } from "./sortableUserDataType";
import type { SortType } from "./sortType";

export type SearchUsersParams = {
    nickname?: string;
    last_access_before?: string;
    last_access_after?: string;
    role_id?: Uuid;
    sorted_by?: SortableUserDataType;
    sort_type?: SortType;
};
