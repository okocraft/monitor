/**
 * Generated by orval v7.5.0 🍺
 * Do not edit manually.
 * Monitor API
 * OpenAPI spec version: 0.0.0
 */
import type { Uuid } from "./uuid";
import type { Role } from "./role";

export interface User {
    id: Uuid;
    nickname: string;
    last_access: string;
    created_at: string;
    updated_at: string;
    role: Role;
}
