/**
 * Type helpers over the generated OpenAPI schema (`./api-schema`).
 *
 * Every helper is parameterized by a route path and an uppercase HTTP method:
 *
 * ```ts
 * type Query = QueryParams<"/api/admin/users", "GET">;
 * type CreateUserBody = RequestBody<"/api/admin/users", "POST">;
 * type User = ResponseData<"/api/admin/users/:id", "GET">;
 * type UsersOpts = QueryOptions<"/api/admin/users", "GET">;
 * ```
 */
import type { UseQueryOptions } from "@tanstack/react-query";
import type { FetchOptions } from "openapi-fetch";
import type { paths } from "./api-schema";

type Schema = paths;

/**
 * HTTP methods accepted by the helper types.
 *
 * Methods are always written in uppercase; the schema's lowercase
 * operation keys are mapped internally via {@link MethodKey}.
 */
export type HttpMethod =
	| "GET"
	| "POST"
	| "PUT"
	| "PATCH"
	| "DELETE"
	| "HEAD"
	| "OPTIONS";

type MethodKey = {
	GET: "get";
	POST: "post";
	PUT: "put";
	PATCH: "patch";
	DELETE: "delete";
	HEAD: "head";
	OPTIONS: "options";
};

/**
 * The operation object for a path/method pair, e.g.
 * `paths["/api/admin/users"]["get"]`.
 *
 * @typeParam P - Route path, e.g. `"/api/admin/users"`.
 * @typeParam M - Uppercase HTTP method, e.g. `"GET"`.
 */
type Operation<
	P extends keyof Schema & string,
	M extends HttpMethod,
> = NonNullable<Schema[P]>[MethodKey[M] & keyof NonNullable<Schema[P]>];

type Params<O> = O extends { parameters: infer P } ? P : never;

type Responses<O> = O extends { responses: infer R } ? R : never;

type Json<T> = T extends { content: { "application/json": infer B } }
	? B
	: never;

type SuccessCode = 200 | 201 | 202 | 203 | 204 | 205 | 206;

type SuccessJson<O> =
	Extract<keyof Responses<O>, SuccessCode> extends infer K
		? K extends keyof Responses<O>
			? Json<Responses<O>[K]>
			: never
		: never;

type ErrorJson<O> = {
	[K in keyof Responses<O>]: K extends SuccessCode
		? never
		: Json<Responses<O>[K]>;
}[keyof Responses<O>];

/** openapi-react-query's query key shape for a route. */
type KeyOf<P extends keyof Schema & string, M extends HttpMethod> =
	| readonly [Lowercase<M>, P]
	| readonly [Lowercase<M>, P, FetchOptions<Operation<P, M>>];

/** Raw JSON body of a success (2xx) response, envelope included.
 *
 * Extracts the `application/json` content of the first 2xx response.
 * Includes the full `{ data, message, status, meta }` envelope.
 *
 * @example
 * type ListUsersBody = ResponseBody<"/api/admin/users", "GET">;
 * // { data: users.User[]; message: string; meta: common.PaginationMeta; status: number }
 *
 * @typeParam P - Route path, e.g. `"/api/admin/users"`.
 * @typeParam M - Uppercase HTTP method, e.g. `"GET"`.
 */
export type ResponseBody<
	P extends keyof Schema & string = never,
	M extends HttpMethod = HttpMethod,
> = SuccessJson<Operation<P, M>>;

/** The `data` field of the success response envelope, or the body itself when
 * there is no envelope.
 *
 * @example
 * type User = ResponseData<"/api/admin/users/:id", "GET">;
 * // users.User
 * type Users = ResponseData<"/api/admin/users", "GET">;
 * // users.User[]
 *
 * @typeParam P - Route path, e.g. `"/api/admin/users"`.
 * @typeParam M - Uppercase HTTP method, e.g. `"GET"`.
 */
export type ResponseData<
	P extends keyof Schema & string = never,
	M extends HttpMethod = HttpMethod,
> = ResponseBody<P, M> extends { data: infer D } ? D : ResponseBody<P, M>;

/** Query params, including path params (the schema exposes them here).
 *
 * Resolves to `never` when the operation takes no query/path params.
 *
 * @example
 * type ListUsersQuery = QueryParams<"/api/admin/users", "GET">;
 * // { page?: number; perPage?: number; query?: string; sort?: http.SortParams }
 * type GetUserQuery = QueryParams<"/api/admin/users/:id", "GET">;
 * // { id?: string }
 *
 * @typeParam P - Route path, e.g. `"/api/admin/users"`.
 * @typeParam M - Uppercase HTTP method, e.g. `"GET"`.
 */
export type QueryParams<
	P extends keyof Schema & string = never,
	M extends HttpMethod = HttpMethod,
> =
	Params<Operation<P, M>> extends { query?: infer Q } ? NonNullable<Q> : never;

/** Request headers accepted by the operation.
 *
 * Resolves to `never` when the operation declares no headers.
 *
 * @example
 * type Headers = RequestHeaders<"/api/admin/users", "GET">;
 *
 * @typeParam P - Route path, e.g. `"/api/admin/users"`.
 * @typeParam M - Uppercase HTTP method, e.g. `"GET"`.
 */
export type RequestHeaders<
	P extends keyof Schema & string = never,
	M extends HttpMethod = HttpMethod,
> =
	Params<Operation<P, M>> extends { header?: infer H } ? NonNullable<H> : never;
/** JSON request body, `never` for body-less methods.
 *
 * @example
 * type CreateUserBody = RequestBody<"/api/admin/users", "POST">;
 * // users.CreateUserReq
 * type DeleteBody = RequestBody<"/api/admin/users/:id", "DELETE">;
 * // never
 *
 * @typeParam P - Route path, e.g. `"/api/admin/users"`.
 * @typeParam M - Uppercase HTTP method, e.g. `"POST"`.
 */
export type RequestBody<
	P extends keyof Schema & string = never,
	M extends HttpMethod = HttpMethod,
> =
	Operation<P, M> extends { requestBody?: infer RB }
		? Json<NonNullable<RB>>
		: never;

/** Options accepted by openapi-react-query's `useQuery` for this route, with
 * the data/error types taken from the schema.
 *
 * @example
 * type UsersOpts = QueryOptions<"/api/admin/users", "GET">;
 * // Omit<UseQueryOptions<ResponseBody<"/api/admin/users", "GET">,
 * //      ErrorJson<...>, never, KeyOf<...>>, "queryKey" | "queryFn">
 *
 * @typeParam P - Route path, e.g. `"/api/admin/users"`.
 * @typeParam M - Uppercase HTTP method, e.g. `"GET"`.
 */
export type QueryOptions<
	P extends keyof Schema & string = never,
	M extends HttpMethod = HttpMethod,
> = Omit<
	UseQueryOptions<
		ResponseBody<P, M>,
		ErrorJson<Operation<P, M>>,
		never,
		KeyOf<P, M>
	>,
	"queryKey" | "queryFn"
>;
