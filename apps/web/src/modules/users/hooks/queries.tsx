import { $api } from "#/services/api.service";
import type { QueryOptions, QueryParams } from "#/types/helpers";

export type UsersQueryOptions = QueryOptions<"/api/admin/users", "GET">;

export const useUsersQuery = (
	params?: QueryParams<"/api/admin/users", "GET">,
	opts?: UsersQueryOptions,
) => $api.useQuery("get", "/api/admin/users", params, opts);
