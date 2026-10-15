export type paths = {
    "/api/account": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /** Soft delete the signed-in account */
        delete: {
            parameters: {
                query: {
                    /** @description the current password, proves possession */
                    password: string;
                    /** @description must be true, guards against accidental calls */
                    confirm?: boolean;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["account.DeleteAccountRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        /** Update the profile of the signed-in user */
        patch: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["account.UpdateProfileReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["account.UpdateProfileRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        trace?: never;
    };
    "/api/account/avatar": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Upload the avatar image */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "multipart/form-data": components["schemas"]["account.UploadAvatarReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["account.UploadAvatarRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/account/connections": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** List the social identities linked to the account */
        get: {
            parameters: {
                query?: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: account.ListConnectionsRes
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["users.Connection"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/account/connections/:provider": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /** Unlink a social identity */
        delete: {
            parameters: {
                query: {
                    /** @description the provider to unlink */
                    provider: "google" | "github" | "apple" | "microsoft";
                    /** @description the current password, proves possession */
                    password: string;
                    /** @description also revoke the tokens held at the provider */
                    revoke?: boolean;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["account.UnlinkConnectionRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/account/email": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        /** Request an email change; the address is only stored once verified */
        put: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["account.ChangeEmailReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["account.ChangeEmailRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/account/me": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Get the current user */
        get: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["account.MeResponse"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/account/phone": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        /** Request a phone change; the number is only stored once verified */
        put: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["account.ChangePhoneReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["account.ChangePhoneRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/account/preference": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Get the current user preference */
        get: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: account.GetPreferenceRes
                             * @description The returned data if the request is OK
                             */
                            data: {
                                [key: string]: components["schemas"]["interface"];
                            };
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /** Update the current user preference */
        patch: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["account.UpdatePreferenceReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: account.UpdatePreferenceRes
                             * @description The returned data if the request is OK
                             */
                            data: {
                                [key: string]: components["schemas"]["interface"];
                            };
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        trace?: never;
    };
    "/api/admin/api-keys": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** List the issued api keys */
        get: {
            parameters: {
                query?: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                    /** @description restrict to the keys of one user */
                    userId?: string;
                    /** @description restrict to the keys acting under one role */
                    roleId?: number;
                    /** @description false returns revoked or expired keys too */
                    isActive?: boolean;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: apikeys.ListAPIKeysRes
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["access.APIKey"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        /** Issue an api key */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["apikeys.CreateAPIKeyReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["apikeys.CreateAPIKeyRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/api-keys/:id": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Get one api key */
        get: {
            parameters: {
                query: {
                    /** @description the key id */
                    id: number;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["apikeys.GetAPIKeyRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        /** Revoke an api key */
        delete: {
            parameters: {
                query: {
                    /** @description the key id */
                    id: number;
                    /** @description why, written to the audit log */
                    reason: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["apikeys.RevokeAPIKeyRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        /** Update an api key, the secret itself cannot be changed */
        patch: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["apikeys.UpdateAPIKeyReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["apikeys.UpdateAPIKeyRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        trace?: never;
    };
    "/api/admin/api-keys/:id/rotate": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Replace the secret of a key, the old one stops working at once */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["apikeys.RotateAPIKeyReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["apikeys.RotateAPIKeyRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/audit-logs": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Search the audit trail */
        get: {
            parameters: {
                query?: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                    /** @description user id, api key or service name that acted */
                    actorId?: string;
                    /** @description action slug, e.g. user.login or role.update */
                    action?: string;
                    /** @description audit category */
                    category?: "auth" | "admin" | "data" | "security" | "general";
                    /** @description affected resource type */
                    resourceType?: string;
                    /** @description affected resource id */
                    resourceId?: string;
                    /** @description outcome of the audited action */
                    status?: "success" | "failure" | "denied" | "error";
                    /** @description session the action belongs to */
                    sessionId?: string;
                    /** @description only entries at or after this time */
                    from?: string;
                    /** @description only entries at or before this time */
                    to?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: audit.ListAuditLogsRes
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["access.AuditEntry"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/audit-logs/:id": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Get one audit entry */
        get: {
            parameters: {
                query: {
                    /** @description the audit entry id */
                    id: number;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["audit.GetAuditLogRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/permissions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: {
            parameters: {
                query?: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: []access.Permission
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["access.Permission"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        /** Create a permission */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["rbac.CreatePermissionReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["rbac.CreatePermissionRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/permissions/:id": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Get a permission */
        get: {
            parameters: {
                query: {
                    /** @description the permission id */
                    id: number;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["rbac.GetPermissionRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        /** Delete a permission, system permissions are rejected */
        delete: {
            parameters: {
                query: {
                    /** @description the permission id */
                    id: number;
                    /** @description why, written to the audit log */
                    reason?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["rbac.DeletePermissionRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        /** Update a permission */
        patch: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["rbac.UpdatePermissionReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["rbac.UpdatePermissionRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        trace?: never;
    };
    "/api/admin/policies": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** List the ABAC policies */
        get: {
            parameters: {
                query?: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: abac.ListPoliciesRes
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["access.Policy"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        /** Create an ABAC policy */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["abac.CreatePolicyReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["abac.CreatePolicyRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/policies/:id": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Get an ABAC policy */
        get: {
            parameters: {
                query: {
                    id: number;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["abac.GetPolicyRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        /** Delete an ABAC policy and drop its cached copy */
        delete: {
            parameters: {
                query: {
                    id: number;
                    /** @description why, written to the audit log */
                    reason?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["abac.DeletePolicyRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        /** Update an ABAC policy */
        patch: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["abac.UpdatePolicyReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["abac.UpdatePolicyRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        trace?: never;
    };
    "/api/admin/policies/:id/evaluate": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Dry-run a policy against a subject, resource and environment */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["abac.EvaluatePolicyReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["abac.EvaluatePolicyRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/roles": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: {
            parameters: {
                query?: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: []access.Role
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["access.Role"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        /** Create a role */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["rbac.CreateRoleReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["rbac.CreateRoleRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/roles/:id": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Get a role */
        get: {
            parameters: {
                query: {
                    /** @description the role id */
                    id: number;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["rbac.GetRoleRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        /** Delete a role, system roles are rejected */
        delete: {
            parameters: {
                query: {
                    /** @description the role id */
                    id: number;
                    /** @description why, written to the audit log */
                    reason?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["rbac.DeleteRoleRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        /** Update a role */
        patch: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["rbac.UpdateRoleReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["rbac.UpdateRoleRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        trace?: never;
    };
    "/api/admin/roles/:id/permissions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** List the permissions granted to a role */
        get: {
            parameters: {
                query: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                    /** @description the role id */
                    id: number;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: rbac.ListRolePermissionsRes
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["access.RolePermission"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        /** Grant a permission to a role */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["rbac.AssignRolePermissionReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["rbac.AssignRolePermissionRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/roles/:id/permissions/:permissionId": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /** Revoke a permission from a role */
        delete: {
            parameters: {
                query: {
                    /** @description the role id */
                    id: number;
                    /** @description the permission to revoke */
                    permissionId: number;
                    /** @description why, written to the audit log */
                    reason?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["rbac.RevokeRolePermissionRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/security-events": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Search the security events */
        get: {
            parameters: {
                query?: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                    /** @description filter on one event type */
                    eventType?: string;
                    /** @description filter on one severity */
                    severity?: "info" | "warning" | "critical";
                    /** @description events of one user */
                    userId?: string;
                    /** @description true returns only events nobody has resolved yet */
                    unresolved?: boolean;
                    /** @description only events at or after this time */
                    from?: string;
                    /** @description only events at or before this time */
                    to?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: audit.ListSecurityEventsRes
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["access.SecurityEvent"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/security-events/:id": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Get one security event */
        get: {
            parameters: {
                query: {
                    /** @description the event id */
                    id: number;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["audit.GetSecurityEventRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/security-events/:id/resolve": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Mark a security event as handled */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["audit.ResolveSecurityEventReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["audit.ResolveSecurityEventRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/users": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: {
            parameters: {
                query?: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: []users.User
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["users.User"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        /** Create new User */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["users.CreateUserReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.CreateUserRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/users/:id": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: {
            parameters: {
                query?: {
                    id?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.User"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        /** Soft delete a user and revoke its sessions */
        delete: {
            parameters: {
                query: {
                    id: string;
                    /** @description why the account is removed, written to the audit log */
                    reason: string;
                    /** @description revoke every active session of that user */
                    revokeSessions?: boolean;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.DeleteUserRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        /** Update a user */
        patch: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["users.UpdateUserReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.UpdateUserRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        trace?: never;
    };
    "/api/admin/users/:id/attributes": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** List the ABAC attributes of a user */
        get: {
            parameters: {
                query: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                    id: string;
                    /** @description restrict to one namespace */
                    namespace?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: users.ListUserAttributesRes
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["users.UserAttribute"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        /** Create or replace one ABAC attribute of a user */
        put: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["users.SetUserAttributeReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.SetUserAttributeRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/users/:id/attributes/:key": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /** Delete one ABAC attribute of a user */
        delete: {
            parameters: {
                query: {
                    id: string;
                    /** @description the attribute name */
                    key: string;
                    /** @description namespace the attribute lives in */
                    namespace?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.DeleteUserAttributeRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/users/:id/ban": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Ban a user until a given time or indefinitely */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["users.BanUserReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.BanUserRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/users/:id/impersonate": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Issue a short-lived token acting as a user, audited */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["users.ImpersonateUserReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.ImpersonateUserRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/users/:id/password": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Set a new password for a user */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["users.ResetUserPasswordReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.ResetUserPasswordRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/users/:id/roles": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** List the roles assigned to a user */
        get: {
            parameters: {
                query: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                    id: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: users.ListUserRolesRes
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["access.UserRole"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        /** Assign a role to a user */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["users.AssignUserRoleReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.AssignUserRoleRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/users/:id/roles/:roleId": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /** Revoke a role from a user */
        delete: {
            parameters: {
                query: {
                    id: string;
                    /** @description the role to revoke */
                    roleId: number;
                    /** @description why, written to the audit log */
                    reason?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.RevokeUserRoleRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/admin/users/:id/status": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /** Move a user to another account status */
        patch: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["users.SetUserStatusReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.SetUserStatusRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        trace?: never;
    };
    "/api/admin/users/:id/unban": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Lift a ban and clear the failed login counter */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["users.UnbanUserReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["users.UnbanUserRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/change-password": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Change the password of the signed-in user */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.ChangePasswordReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.ChangePasswordRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/forgot-password": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Request a password reset code */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.ForgotPasswordReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.ForgotPasswordRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/logout": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Revoke the current session */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.LogoutRequest"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.LogoutResponse"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/logout-all": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Sign out of every session */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.LogoutAllReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.LogoutAllRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/mfa/challenge": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Complete a sign-in that requires a second factor */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.ChallengeMFAReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.ChallengeMFARes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/mfa/confirm": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Confirm TOTP enrolment with a generated code */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.ConfirmMFAReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.ConfirmMFARes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/mfa/disable": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Turn off multi-factor authentication */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.DisableMFAReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.DisableMFARes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/mfa/enable": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Start TOTP enrolment and return the secret */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.EnableMFAReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.EnableMFARes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/oauth/:provider/authorize": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Redirect the client to the provider consent screen */
        get: {
            parameters: {
                query: {
                    /** @description the OAuth provider */
                    provider: "google" | "github" | "apple" | "microsoft";
                    /** @description callback url, must be pre-registered */
                    redirectUri?: string;
                    /** @description requested scopes, defaults to the provider's openid profile set */
                    scopes?: string[];
                    /** @description opaque value echoed back on the callback, used for CSRF protection */
                    state?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.OAuthAuthorizeRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/oauth/:provider/callback": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Exchange the provider code for a session */
        get: {
            parameters: {
                query: {
                    /** @description the OAuth provider */
                    provider: "google" | "github" | "apple" | "microsoft";
                    /** @description the authorization code returned by the provider */
                    code: string;
                    /** @description the state value issued by /oauth/:provider/authorize */
                    state: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.OAuthCallbackRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/refresh": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Refresh the session token */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.RefreshRequest"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.RefreshResponse"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/reset-password": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Set a new password with a valid reset code */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.ResetPasswordReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.ResetPasswordRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/reset-password/verify": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Check a password reset code before setting a new password */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.VerifyResetPasswordReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.VerifyResetPasswordRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/sessions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** List the active sessions of the signed-in user */
        get: {
            parameters: {
                query?: {
                    page?: number;
                    perPage?: number;
                    query?: string;
                    sort?: components["schemas"]["http.SortParams"];
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /**
                             * Format: authentication.ListSessionsRes
                             * @description The returned data if the request is OK
                             */
                            data: components["schemas"]["auth.SessionData"][];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        put?: never;
        post?: never;
        /** Revoke every session of the signed-in user */
        delete: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.RevokeAllSessionsRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/sessions/:id": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /** Revoke one session by id */
        delete: {
            parameters: {
                query: {
                    /** @description the session id (token id) */
                    id: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.RevokeSessionRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/sign-in": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Sign in with identifier and password */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.SignInRequest"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.SignInResponse"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/sign-up": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Register a new account */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.SignUpRequest"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.SignUpResponse"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/verification/resend": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Resend a verification code */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.ResendVerificationReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.ResendVerificationRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/verify-email": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Verify an email address with a one-time code */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.VerifyEmailReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.VerifyEmailRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/auth/verify-phone": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** Verify a phone number with a one-time code */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["authentication.VerifyPhoneReq"];
                };
            };
            responses: {
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": {
                            /** @description The returned data if the request is OK */
                            data: components["schemas"]["authentication.VerifyPhoneRes"];
                            /** Format: string */
                            message: string;
                            /** @description The pagination data for list response */
                            meta: components["schemas"]["common.PaginationMeta"];
                            /**
                             * Format: int
                             * @description The http status code of the response 200,300,400,500
                             */
                            status: number;
                        };
                    };
                };
                /** @description Invalid input */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.BadRequestResponse"];
                    };
                };
                /** @description User is not authenticated */
                401: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User is not authenticated */
                403: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.UnAuthenticatedResponse"];
                    };
                };
                /** @description User input are invalid and not processable */
                422: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InvalidInputErrorResponse"];
                    };
                };
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["http.InternalErrorResponse"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
};
export type webhooks = Record<string, never>;
export type components = {
    schemas: {
        "account.DeleteAccountReq": {
            /**
             * Format: string
             * @description the current password, proves possession
             */
            password: string;
            /**
             * Format: bool
             * @description must be true, guards against accidental calls
             */
            confirm?: boolean;
        };
        "account.DeleteAccountRes": {
            /**
             * Format: string
             * @description account deleted
             */
            message: string;
        };
        "common.PaginationMeta": {
            /** Format: int */
            count: number;
            cursor: components["schemas"]["interface"];
            /** Format: int */
            page: number;
            /** Format: int */
            perPage: number;
            /** Format: int */
            totalPages: number;
            /** Format: int */
            nextPage: number;
            /** Format: int */
            prevPage: number;
            /** Format: int */
            lastPage: number;
            /** Format: bool */
            hasNext: boolean;
        };
        interface: Record<string, never>;
        /** @description Invalid input */
        "http.BadRequestResponse": {
            /**
             * Format: string
             * @description A code describing the error if any
             */
            code: string;
            /**
             * Format: string
             * @description the error  message
             */
            message: string;
            /**
             * Format: int
             * @description The http status code of the response 400,500
             */
            status: number;
        };
        "http.InternalErrorResponse": {
            /**
             * Format: string
             * @description A code describing the error if any
             */
            code: string;
            /**
             * Format: string
             * @description the error  message
             */
            message: string;
            /**
             * Format: int
             * @description The http status code of the response 400,500
             */
            status: number;
        };
        /** @description User is not authenticated */
        "http.UnAuthenticatedResponse": {
            /**
             * Format: string
             * @description A code describing the error if any
             */
            code: string;
            /**
             * Format: string
             * @description the error  message
             */
            message: string;
            /**
             * Format: int
             * @description The http status code of the response 400,500
             */
            status: number;
        };
        /** @description User input are invalid and not processable */
        "http.InvalidInputErrorResponse": {
            /**
             * Format: string
             * @description A code describing the error if any
             */
            code: string;
            /**
             * Format: string
             * @description the error  message
             */
            message: string;
            /**
             * Format: int
             * @description The http status code of the response 400,500
             */
            status: number;
            /**
             * Format: map[string]interface {}
             * @description returned errors in post requests with body
             */
            errors: {
                [key: string]: components["schemas"]["interface"];
            };
        };
        "account.UpdateProfileReq": {
            /**
             * Format: *string
             * @description display name
             */
            name: string;
            /**
             * Format: *string
             * @description avatar url
             */
            image: string;
        };
        "account.UpdateProfileRes": {
            /**
             * Format: string
             * @description the stored display name
             */
            name: string;
            /**
             * Format: string
             * @description the stored avatar url
             */
            image: string;
            /**
             * Format: time.Time
             * @description when the profile was written
             */
            updatedAt: string;
        };
        "account.UploadAvatarReq": {
            /**
             * Format: *ghttp.UploadFile
             * @description the image, png or jpeg, within the configured size limit
             */
            file: Record<string, never>;
        };
        "account.UploadAvatarRes": {
            /**
             * Format: string
             * @description where the stored image can be fetched
             */
            url: string;
            /**
             * Format: string
             * @description the value to store on the user record
             */
            image: string;
            /**
             * Format: time.Time
             * @description when the image was stored
             */
            uploadedAt: string;
        };
        "account.ListConnectionsReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query: string;
            sort: components["schemas"]["http.SortParams"];
        };
        "http.SortParams": {
            /** Format: string */
            field: string;
            /**
             * Format: string
             * @default asc
             * @enum {string}
             */
            dir: "desc" | "asc";
        };
        "users.Connection": {
            /**
             * Format: int64
             * @description internal link id
             */
            id: number;
            /**
             * Format: string
             * @description google, github, apple, ...
             */
            provider: string;
            /**
             * Format: string
             * @description the identifier the provider assigned
             */
            providerAccountId: string;
            /**
             * Format: []string
             * @description granted scopes
             */
            scopes: string[];
            /**
             * Format: *time.Time
             * @description unix expiry of the stored access token
             */
            accessTokenExpires: string;
            /**
             * Format: time.Time
             * @description when the link was created
             */
            createdAt: string;
        };
        "account.UnlinkConnectionReq": {
            /**
             * Format: string
             * @description the provider to unlink
             * @enum {string}
             */
            provider: "google" | "github" | "apple" | "microsoft";
            /**
             * Format: string
             * @description the current password, proves possession
             */
            password: string;
            /**
             * Format: *bool
             * @description also revoke the tokens held at the provider
             * @default true
             */
            revoke: boolean;
        };
        "account.UnlinkConnectionRes": {
            /**
             * Format: string
             * @description the unlinked provider
             */
            provider: string;
            /**
             * Format: bool
             * @description always true on success
             */
            unlinked: boolean;
        };
        "account.ChangeEmailReq": {
            /**
             * Format: string
             * @description the new address
             */
            email: string;
            /**
             * Format: string
             * @description the current password, proves possession
             */
            password: string;
            /**
             * Format: *string
             * @description the code from the change-email mail, when the client already has it
             */
            code?: string;
        };
        "account.ChangeEmailRes": {
            /**
             * Format: string
             * @description verification mail sent to the new address
             */
            message: string;
            /**
             * Format: int64
             * @description unix expiry of the change code
             */
            expiresAt: number;
            /**
             * Format: int
             * @description seconds the previous address stays valid
             */
            pendingFor: number;
        };
        "account.MeRequest": Record<string, never>;
        "account.MeResponse": {
            /** Format: uuid.UUID */
            id: string;
            /** Format: string */
            name: string;
            /** Format: string */
            email: string;
            /** Format: string */
            phone: string;
            /** Format: string */
            username: string;
            /** Format: bool */
            email_verified: boolean;
            /** Format: bool */
            phone_verified: boolean;
            /** Format: string */
            image: string;
            /** Format: []string */
            roles: string[];
            /** Format: []string */
            permissions: string[];
            /** Format: string */
            status: string;
            /** Format: map[string]interface {} */
            preference: {
                [key: string]: components["schemas"]["interface"];
            };
            /** Format: bool */
            mfa_enabled: boolean;
            /** Format: time.Time */
            created_at: string;
            /** Format: time.Time */
            updated_at: string;
        };
        "account.ChangePhoneReq": {
            /**
             * Format: string
             * @description the new number in E.164 format
             */
            phone: string;
            /**
             * Format: string
             * @description the current password, proves possession
             */
            password: string;
            /**
             * Format: *string
             * @description the sms code, when the client already has it
             */
            code?: string;
        };
        "account.ChangePhoneRes": {
            /**
             * Format: string
             * @description verification sms sent to the new number
             */
            message: string;
            /**
             * Format: int64
             * @description unix expiry of the change code
             */
            expiresAt: number;
            /**
             * Format: int
             * @description seconds the previous number stays valid
             */
            pendingFor: number;
        };
        "account.GetPreferenceReq": Record<string, never>;
        "account.UpdatePreferenceReq": {
            /**
             * Format: map[string]interface {}
             * @description the json payload of the preference
             */
            payload: {
                [key: string]: components["schemas"]["interface"];
            };
        };
        "apikeys.ListAPIKeysReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query: string;
            sort: components["schemas"]["http.SortParams"];
            /**
             * Format: *uuid.UUID
             * @description restrict to the keys of one user
             */
            userId: string;
            /**
             * Format: *int32
             * @description restrict to the keys acting under one role
             */
            roleId: number;
            /**
             * Format: *bool
             * @description false returns revoked or expired keys too
             */
            isActive: boolean;
        };
        "access.APIKey": {
            /**
             * Format: int64
             * @description internal key id
             */
            id: number;
            /**
             * Format: string
             * @description human label
             */
            name: string;
            /**
             * Format: string
             * @description first characters of the key, enough to recognise it
             */
            keyPrefix: string;
            /**
             * Format: *uuid.UUID
             * @description owning user, nil for a service account
             */
            userId: string;
            /**
             * Format: *int32
             * @description role the key acts under
             */
            roleId: number;
            /**
             * Format: []string
             * @description permission slugs overriding the role grants
             */
            permissions: string[];
            /**
             * Format: *int
             * @description requests per minute
             */
            rateLimit: number;
            /**
             * Format: *time.Time
             * @description when the key stops working
             */
            expiresAt: string;
            /**
             * Format: *time.Time
             * @description last time the key authenticated a request
             */
            lastUsedAt: string;
            /**
             * Format: bool
             * @description false once revoked or expired
             */
            isActive: boolean;
            /**
             * Format: *uuid.UUID
             * @description who issued the key
             */
            createdBy: string;
            /**
             * Format: time.Time
             * @description when the key was issued
             */
            createdAt: string;
        };
        "apikeys.CreateAPIKeyReq": {
            /**
             * Format: string
             * @description human label, shows up in the audit log
             */
            name: string;
            /**
             * Format: *uuid.UUID
             * @description owning user, omit for a service account
             */
            userId?: string;
            /**
             * Format: *int32
             * @description role the key acts under
             */
            roleId?: number;
            /**
             * Format: []string
             * @description permission slugs, they narrow the role grants
             */
            permissions?: string[];
            /**
             * Format: *int
             * @description requests per minute
             */
            rateLimit?: number;
            /**
             * Format: *time.Time
             * @description when the key stops working, prefer a short life
             */
            expiresAt?: string;
        };
        "apikeys.CreateAPIKeyRes": {
            /**
             * Format: int64
             * @description internal key id
             */
            id: number;
            /**
             * Format: string
             * @description human label
             */
            name: string;
            /**
             * Format: string
             * @description first characters of the key, enough to recognise it
             */
            keyPrefix: string;
            /**
             * Format: *uuid.UUID
             * @description owning user, nil for a service account
             */
            userId: string;
            /**
             * Format: *int32
             * @description role the key acts under
             */
            roleId: number;
            /**
             * Format: []string
             * @description permission slugs overriding the role grants
             */
            permissions: string[];
            /**
             * Format: *int
             * @description requests per minute
             */
            rateLimit: number;
            /**
             * Format: *time.Time
             * @description when the key stops working
             */
            expiresAt: string;
            /**
             * Format: *time.Time
             * @description last time the key authenticated a request
             */
            lastUsedAt: string;
            /**
             * Format: bool
             * @description false once revoked or expired
             */
            isActive: boolean;
            /**
             * Format: *uuid.UUID
             * @description who issued the key
             */
            createdBy: string;
            /**
             * Format: time.Time
             * @description when the key was issued
             */
            createdAt: string;
            /**
             * Format: string
             * @description the plaintext key, returned only once and never recoverable
             */
            key: string;
            /**
             * Format: time.Time
             * @description when the key was issued
             */
            issuedAt: string;
        };
        "apikeys.RevokeAPIKeyReq": {
            /**
             * Format: int64
             * @description the key id
             */
            id: number;
            /**
             * Format: string
             * @description why, written to the audit log
             */
            reason: string;
        };
        "apikeys.RevokeAPIKeyRes": {
            /**
             * Format: int64
             * @description the revoked key id
             */
            id: number;
            /**
             * Format: bool
             * @description always true on success
             */
            revoked: boolean;
        };
        "apikeys.GetAPIKeyReq": {
            /**
             * Format: int64
             * @description the key id
             */
            id: number;
        };
        "apikeys.GetAPIKeyRes": {
            /**
             * Format: int64
             * @description internal key id
             */
            id: number;
            /**
             * Format: string
             * @description human label
             */
            name: string;
            /**
             * Format: string
             * @description first characters of the key, enough to recognise it
             */
            keyPrefix: string;
            /**
             * Format: *uuid.UUID
             * @description owning user, nil for a service account
             */
            userId: string;
            /**
             * Format: *int32
             * @description role the key acts under
             */
            roleId: number;
            /**
             * Format: []string
             * @description permission slugs overriding the role grants
             */
            permissions: string[];
            /**
             * Format: *int
             * @description requests per minute
             */
            rateLimit: number;
            /**
             * Format: *time.Time
             * @description when the key stops working
             */
            expiresAt: string;
            /**
             * Format: *time.Time
             * @description last time the key authenticated a request
             */
            lastUsedAt: string;
            /**
             * Format: bool
             * @description false once revoked or expired
             */
            isActive: boolean;
            /**
             * Format: *uuid.UUID
             * @description who issued the key
             */
            createdBy: string;
            /**
             * Format: time.Time
             * @description when the key was issued
             */
            createdAt: string;
        };
        "apikeys.UpdateAPIKeyReq": {
            /** Format: int64 */
            id: number;
            /**
             * Format: *string
             * @description human label
             */
            name?: string;
            /**
             * Format: *[]string
             * @description replaces the permission slugs
             */
            permissions?: string[][];
            /**
             * Format: *int
             * @description requests per minute
             */
            rateLimit?: number;
            /**
             * Format: *time.Time
             * @description when the key stops working
             */
            expiresAt?: string;
            /**
             * Format: *bool
             * @description suspend or resume the key
             */
            isActive?: boolean;
            /**
             * Format: string
             * @description why, written to the audit log
             */
            reason?: string;
        };
        "apikeys.UpdateAPIKeyRes": {
            /**
             * Format: int64
             * @description the updated key id
             */
            id: number;
        };
        "apikeys.RotateAPIKeyReq": {
            /**
             * Format: int64
             * @description the key id
             */
            id: number;
            /**
             * Format: string
             * @description why the secret is replaced, written to the audit log
             */
            reason?: string;
        };
        "apikeys.RotateAPIKeyRes": {
            /**
             * Format: int64
             * @description the rotated key id
             */
            id: number;
            /**
             * Format: string
             * @description the new plaintext key, returned only once
             */
            key: string;
            /**
             * Format: string
             * @description prefix of the new key
             */
            prefix: string;
            /**
             * Format: time.Time
             * @description when the secret was replaced
             */
            rotatedAt: string;
        };
        "audit.ListAuditLogsReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query: string;
            sort: components["schemas"]["http.SortParams"];
            /**
             * Format: string
             * @description user id, api key or service name that acted
             */
            actorId: string;
            /**
             * Format: string
             * @description action slug, e.g. user.login or role.update
             */
            action: string;
            /**
             * Format: string
             * @description audit category
             * @enum {string}
             */
            category: "auth" | "admin" | "data" | "security" | "general";
            /**
             * Format: string
             * @description affected resource type
             */
            resourceType: string;
            /**
             * Format: string
             * @description affected resource id
             */
            resourceId: string;
            /**
             * Format: string
             * @description outcome of the audited action
             * @enum {string}
             */
            status: "success" | "failure" | "denied" | "error";
            /**
             * Format: string
             * @description session the action belongs to
             */
            sessionId: string;
            /**
             * Format: *time.Time
             * @description only entries at or after this time
             */
            from: string;
            /**
             * Format: *time.Time
             * @description only entries at or before this time
             */
            to: string;
        };
        "access.AuditEntry": {
            /** Format: int32 */
            ID: number;
            /** Format: time.Time */
            CreatedAt: string;
            /** Format: string */
            ActorType: string;
            /** Format: string */
            ActorID: string;
            /** Format: string */
            ActorRole: string;
            /** Format: string */
            Action: string;
            /** Format: string */
            Category: string;
            /** Format: string */
            ResourceType: string;
            /** Format: string */
            ResourceID: string;
            /** Format: string */
            IPAddress: string;
            /** Format: string */
            UserAgent: string;
            /** Format: string */
            RequestID: string;
            /** Format: string */
            SessionID: string;
            /** Format: string */
            BeforeState: string;
            /** Format: string */
            AfterState: string;
            /** Format: string */
            Metadata: string;
            /** Format: access.AuditStatus */
            Status: string;
            /** Format: string */
            ErrorMessage: string;
        };
        "audit.GetAuditLogReq": {
            /**
             * Format: int32
             * @description the audit entry id
             */
            id: number;
        };
        "audit.GetAuditLogRes": {
            /** Format: int32 */
            ID: number;
            /** Format: time.Time */
            CreatedAt: string;
            /** Format: string */
            ActorType: string;
            /** Format: string */
            ActorID: string;
            /** Format: string */
            ActorRole: string;
            /** Format: string */
            Action: string;
            /** Format: string */
            Category: string;
            /** Format: string */
            ResourceType: string;
            /** Format: string */
            ResourceID: string;
            /** Format: string */
            IPAddress: string;
            /** Format: string */
            UserAgent: string;
            /** Format: string */
            RequestID: string;
            /** Format: string */
            SessionID: string;
            /** Format: string */
            BeforeState: string;
            /** Format: string */
            AfterState: string;
            /** Format: string */
            Metadata: string;
            /** Format: access.AuditStatus */
            Status: string;
            /** Format: string */
            ErrorMessage: string;
        };
        "rbac.ListPermissionReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query: string;
            sort: components["schemas"]["http.SortParams"];
        };
        "access.Permission": {
            /** Format: int32 */
            ID: number;
            /** Format: string */
            Name: string;
            /** Format: string */
            Slug: string;
            /** Format: string */
            Description: string;
            /** Format: string */
            Resource: string;
            /** Format: string */
            Action: string;
            /** Format: string */
            Category: string;
            /** Format: bool */
            IsSystem: boolean;
            /** Format: string */
            Conditions: string;
            /** Format: time.Time */
            CreatedAt: string;
            /** Format: time.Time */
            UpdatedAt: string;
        };
        "rbac.CreatePermissionReq": {
            /**
             * Format: string
             * @description display name
             */
            name: string;
            /**
             * Format: string
             * @description resource.action, e.g. invoice.export
             */
            slug: string;
            /**
             * Format: string
             * @description what the permission allows
             */
            description?: string;
            /**
             * Format: string
             * @description the resource the action applies to
             */
            resource: string;
            /**
             * Format: string
             * @description the action
             * @enum {string}
             */
            action: "create" | "read" | "update" | "delete" | "list" | "manage" | "assign" | "revoke" | "export";
            /**
             * Format: string
             * @description grouping used by the permission registry ui
             * @enum {string}
             */
            category?: "user-management" | "content" | "billing" | "system" | "admin";
            /**
             * Format: map[string]interface {}
             * @description ABAC conditions applied whenever the permission is granted
             */
            conditions?: {
                [key: string]: components["schemas"]["interface"];
            };
        };
        "rbac.CreatePermissionRes": {
            /**
             * Format: int32
             * @description the created permission id
             */
            id: number;
            /**
             * Format: string
             * @description the created permission slug
             */
            slug: string;
        };
        "rbac.DeletePermissionReq": {
            /**
             * Format: int32
             * @description the permission id
             */
            id: number;
            /**
             * Format: string
             * @description why, written to the audit log
             */
            reason?: string;
        };
        "rbac.DeletePermissionRes": {
            /**
             * Format: int32
             * @description the deleted permission
             */
            id: number;
            /**
             * Format: bool
             * @description always true on success
             */
            deleted: boolean;
        };
        "rbac.GetPermissionReq": {
            /**
             * Format: int32
             * @description the permission id
             */
            id: number;
        };
        "rbac.GetPermissionRes": {
            /** Format: int32 */
            ID: number;
            /** Format: string */
            Name: string;
            /** Format: string */
            Slug: string;
            /** Format: string */
            Description: string;
            /** Format: string */
            Resource: string;
            /** Format: string */
            Action: string;
            /** Format: string */
            Category: string;
            /** Format: bool */
            IsSystem: boolean;
            /** Format: string */
            Conditions: string;
            /** Format: time.Time */
            CreatedAt: string;
            /** Format: time.Time */
            UpdatedAt: string;
        };
        "rbac.UpdatePermissionReq": {
            /** Format: int32 */
            id: number;
            /**
             * Format: *string
             * @description display name
             */
            name?: string;
            /**
             * Format: *string
             * @description what the permission allows
             */
            description?: string;
            /**
             * Format: *string
             * @description grouping
             */
            category?: string;
            /**
             * Format: map[string]interface {}
             * @description replaces the ABAC conditions of the permission
             */
            conditions?: {
                [key: string]: components["schemas"]["interface"];
            };
        };
        "rbac.UpdatePermissionRes": {
            /**
             * Format: int32
             * @description the updated permission
             */
            id: number;
        };
        "abac.ListPoliciesReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query: string;
            sort: components["schemas"]["http.SortParams"];
        };
        "access.Policy": {
            /** Format: int32 */
            ID: number;
            /** Format: string */
            Name: string;
            /** Format: string */
            Slug: string;
            /** Format: string */
            Description: string;
            /** Format: string */
            Resource: string;
            /** Format: string */
            Action: string;
            /** Format: map[string]interface {} */
            Rules: {
                [key: string]: components["schemas"]["interface"];
            };
            /** Format: access.Effect */
            Effect: string;
            /** Format: int32 */
            Priority: number;
            /** Format: bool */
            IsActive: boolean;
            /** Format: *time.Time */
            ExpiresAt: string;
            /** Format: *uuid.UUID */
            CreatedBy: string;
            /** Format: time.Time */
            CreatedAt: string;
            /** Format: time.Time */
            UpdatedAt: string;
        };
        "abac.CreatePolicyReq": {
            /**
             * Format: string
             * @description display name
             */
            name: string;
            /**
             * Format: string
             * @description stable machine name, unique across policies
             */
            slug: string;
            /**
             * Format: string
             * @description what the policy decides
             */
            description?: string;
            /**
             * Format: string
             * @description resource type the policy guards
             */
            resource: string;
            /**
             * Format: string
             * @description action the policy guards
             */
            action: string;
            /**
             * Format: map[string]interface {}
             * @description subject, resource and environment rules
             */
            rules: {
                [key: string]: components["schemas"]["interface"];
            };
            /**
             * Format: access.Effect
             * @description allow or deny, deny wins on equal priority
             * @default allow
             * @enum {string}
             */
            effect: "allow" | "deny";
            /**
             * Format: int32
             * @description lower runs first
             * @default 100
             */
            priority: number;
            /**
             * Format: *time.Time
             * @description when the policy stops applying
             */
            expiresAt?: string;
        };
        "abac.CreatePolicyRes": {
            /**
             * Format: int32
             * @description the created policy id
             */
            id: number;
            /**
             * Format: string
             * @description the created policy slug
             */
            slug: string;
        };
        "abac.DeletePolicyReq": {
            /** Format: int32 */
            id: number;
            /**
             * Format: string
             * @description why, written to the audit log
             */
            reason?: string;
        };
        "abac.DeletePolicyRes": {
            /**
             * Format: int32
             * @description the deleted policy id
             */
            id: number;
            /**
             * Format: bool
             * @description always true on success
             */
            deleted: boolean;
        };
        "abac.GetPolicyReq": {
            /** Format: int32 */
            id: number;
        };
        "abac.GetPolicyRes": {
            /** Format: int32 */
            ID: number;
            /** Format: string */
            Name: string;
            /** Format: string */
            Slug: string;
            /** Format: string */
            Description: string;
            /** Format: string */
            Resource: string;
            /** Format: string */
            Action: string;
            /** Format: map[string]interface {} */
            Rules: {
                [key: string]: components["schemas"]["interface"];
            };
            /** Format: access.Effect */
            Effect: string;
            /** Format: int32 */
            Priority: number;
            /** Format: bool */
            IsActive: boolean;
            /** Format: *time.Time */
            ExpiresAt: string;
            /** Format: *uuid.UUID */
            CreatedBy: string;
            /** Format: time.Time */
            CreatedAt: string;
            /** Format: time.Time */
            UpdatedAt: string;
        };
        "abac.UpdatePolicyReq": {
            /** Format: int32 */
            id: number;
            /**
             * Format: *string
             * @description display name
             */
            name?: string;
            /**
             * Format: *string
             * @description what the policy decides
             */
            description?: string;
            /**
             * Format: map[string]interface {}
             * @description replaces the rules document
             */
            rules?: {
                [key: string]: components["schemas"]["interface"];
            };
            /**
             * Format: *access.Effect
             * @description allow or deny
             * @enum {string}
             */
            effect?: "allow" | "deny";
            /**
             * Format: *int32
             * @description lower runs first
             */
            priority?: number;
            /**
             * Format: *bool
             * @description switch the policy off without deleting it
             */
            isActive?: boolean;
            /**
             * Format: *time.Time
             * @description when the policy stops applying
             */
            expiresAt?: string;
        };
        "abac.UpdatePolicyRes": {
            /**
             * Format: int32
             * @description the updated policy id
             */
            id: number;
        };
        "abac.EvaluatePolicyReq": {
            /** Format: int32 */
            id: number;
            /**
             * Format: map[string]interface {}
             * @description subject attributes, e.g. {"role":"editor","id":"..."}
             */
            subject?: {
                [key: string]: components["schemas"]["interface"];
            };
            /**
             * Format: map[string]interface {}
             * @description resource attributes, e.g. {"owner_id":"{{subject.id}}"}
             */
            resource?: {
                [key: string]: components["schemas"]["interface"];
            };
            /**
             * Format: map[string]interface {}
             * @description environment attributes, e.g. time and ip_range
             */
            environment?: {
                [key: string]: components["schemas"]["interface"];
            };
        };
        "abac.EvaluatePolicyRes": {
            /**
             * Format: bool
             * @description the decision the engine would return
             */
            allowed: boolean;
            /**
             * Format: string
             * @description allow or deny
             */
            effect: string;
            /**
             * Format: string
             * @description why the engine decided this way
             */
            reason: string;
            /**
             * Format: int32
             * @description the policy priority that was applied
             */
            priority: number;
        };
        "rbac.ListRolesReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query: string;
            sort: components["schemas"]["http.SortParams"];
        };
        "access.Role": {
            /** Format: int32 */
            ID: number;
            /** Format: string */
            Name: string;
            /** Format: string */
            Slug: string;
            /** Format: string */
            Description: string;
            /** Format: *int32 */
            ParentRoleID: number;
            /** Format: int32 */
            Level: number;
            /** Format: bool */
            IsSystem: boolean;
            /** Format: string */
            Metadata: string;
            /** Format: time.Time */
            CreatedAt: string;
            /** Format: time.Time */
            UpdatedAt: string;
        };
        "rbac.CreateRoleReq": {
            /**
             * Format: string
             * @description display name
             */
            name: string;
            /**
             * Format: string
             * @description stable machine name, unique across roles
             */
            slug: string;
            /**
             * Format: string
             * @description what the role is for
             */
            description?: string;
            /**
             * Format: *int32
             * @description role this one inherits permissions from
             */
            parentRoleId?: number;
            /**
             * Format: map[string]interface {}
             * @description JSON attributes the policy engine can read for this role
             */
            metadata?: {
                [key: string]: components["schemas"]["interface"];
            };
        };
        "rbac.CreateRoleRes": {
            /**
             * Format: int32
             * @description the created role id
             */
            id: number;
            /**
             * Format: string
             * @description the created role slug
             */
            slug: string;
        };
        "rbac.DeleteRoleReq": {
            /**
             * Format: int32
             * @description the role id
             */
            id: number;
            /**
             * Format: string
             * @description why, written to the audit log
             */
            reason?: string;
        };
        "rbac.DeleteRoleRes": {
            /**
             * Format: int32
             * @description the deleted role
             */
            id: number;
            /**
             * Format: bool
             * @description always true on success
             */
            deleted: boolean;
        };
        "rbac.GetRoleReq": {
            /**
             * Format: int32
             * @description the role id
             */
            id: number;
        };
        "rbac.GetRoleRes": {
            /** Format: int32 */
            ID: number;
            /** Format: string */
            Name: string;
            /** Format: string */
            Slug: string;
            /** Format: string */
            Description: string;
            /** Format: *int32 */
            ParentRoleID: number;
            /** Format: int32 */
            Level: number;
            /** Format: bool */
            IsSystem: boolean;
            /** Format: string */
            Metadata: string;
            /** Format: time.Time */
            CreatedAt: string;
            /** Format: time.Time */
            UpdatedAt: string;
        };
        "rbac.UpdateRoleReq": {
            /** Format: int32 */
            id: number;
            /**
             * Format: *string
             * @description display name
             */
            name?: string;
            /**
             * Format: *string
             * @description what the role is for
             */
            description?: string;
            /**
             * Format: *int32
             * @description inheritance parent, nil detaches, must not create a cycle
             */
            parentRoleId?: number;
            /**
             * Format: map[string]interface {}
             * @description replaces the stored attributes
             */
            metadata?: {
                [key: string]: components["schemas"]["interface"];
            };
        };
        "rbac.UpdateRoleRes": {
            /**
             * Format: int32
             * @description the updated role
             */
            id: number;
        };
        "rbac.ListRolePermissionsReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query?: string;
            sort?: components["schemas"]["http.SortParams"];
            /**
             * Format: int32
             * @description the role id
             */
            id: number;
        };
        "access.RolePermission": {
            /** Format: int32 */
            ID: number;
            /** Format: int32 */
            RoleID: number;
            /** Format: int32 */
            PermissionID: number;
            /** Format: string */
            Conditions: string;
            /** Format: *time.Time */
            ExpiresAt: string;
            /** Format: *uuid.UUID */
            GrantedBy: string;
            /** Format: time.Time */
            CreatedAt: string;
        };
        "rbac.AssignRolePermissionReq": {
            /**
             * Format: int32
             * @description the role id
             */
            id: number;
            /**
             * Format: int32
             * @description the permission to grant
             */
            permissionId: number;
            /**
             * Format: map[string]interface {}
             * @description ABAC conditions narrowing the grant, e.g. {"owner_only":true}
             */
            conditions?: {
                [key: string]: components["schemas"]["interface"];
            };
            /**
             * Format: *time.Time
             * @description when the grant lapses, nil grants it permanently
             */
            expiresAt?: string;
        };
        "rbac.AssignRolePermissionRes": {
            /**
             * Format: int32
             * @description the role
             */
            roleId: number;
            /**
             * Format: int32
             * @description the granted permission
             */
            permissionId: number;
            /**
             * Format: bool
             * @description always true on success
             */
            granted: boolean;
        };
        "rbac.RevokeRolePermissionReq": {
            /**
             * Format: int32
             * @description the role id
             */
            id: number;
            /**
             * Format: int32
             * @description the permission to revoke
             */
            permissionId: number;
            /**
             * Format: string
             * @description why, written to the audit log
             */
            reason?: string;
        };
        "rbac.RevokeRolePermissionRes": {
            /**
             * Format: int32
             * @description the role
             */
            roleId: number;
            /**
             * Format: int32
             * @description the revoked permission
             */
            permissionId: number;
            /**
             * Format: bool
             * @description always true on success
             */
            revoked: boolean;
        };
        "audit.ListSecurityEventsReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query: string;
            sort: components["schemas"]["http.SortParams"];
            /**
             * Format: string
             * @description filter on one event type
             */
            eventType: string;
            /**
             * Format: string
             * @description filter on one severity
             * @enum {string}
             */
            severity: "info" | "warning" | "critical";
            /**
             * Format: *uuid.UUID
             * @description events of one user
             */
            userId: string;
            /**
             * Format: *bool
             * @description true returns only events nobody has resolved yet
             */
            unresolved: boolean;
            /**
             * Format: *time.Time
             * @description only events at or after this time
             */
            from: string;
            /**
             * Format: *time.Time
             * @description only events at or before this time
             */
            to: string;
        };
        "access.SecurityEvent": {
            /**
             * Format: int64
             * @description internal event id
             */
            id: number;
            /**
             * Format: string
             * @description login_success, login_failure, account_locked, mfa_enabled, ...
             */
            eventType: string;
            /**
             * Format: *uuid.UUID
             * @description the user the event belongs to, nil for system wide events
             */
            userId: string;
            /**
             * Format: string
             * @description info, warning or critical
             */
            severity: string;
            /**
             * Format: map[string]interface {}
             * @description JSON payload of the event
             */
            details: {
                [key: string]: components["schemas"]["interface"];
            };
            /**
             * Format: string
             * @description client address the event came from
             */
            ipAddress: string;
            /**
             * Format: *time.Time
             * @description when an operator resolved it
             */
            resolvedAt: string;
            /**
             * Format: *uuid.UUID
             * @description who resolved it
             */
            resolvedBy: string;
            /**
             * Format: time.Time
             * @description when the event was recorded
             */
            createdAt: string;
        };
        "audit.GetSecurityEventReq": {
            /**
             * Format: int64
             * @description the event id
             */
            id: number;
        };
        "audit.GetSecurityEventRes": {
            /**
             * Format: int64
             * @description internal event id
             */
            id: number;
            /**
             * Format: string
             * @description login_success, login_failure, account_locked, mfa_enabled, ...
             */
            eventType: string;
            /**
             * Format: *uuid.UUID
             * @description the user the event belongs to, nil for system wide events
             */
            userId: string;
            /**
             * Format: string
             * @description info, warning or critical
             */
            severity: string;
            /**
             * Format: map[string]interface {}
             * @description JSON payload of the event
             */
            details: {
                [key: string]: components["schemas"]["interface"];
            };
            /**
             * Format: string
             * @description client address the event came from
             */
            ipAddress: string;
            /**
             * Format: *time.Time
             * @description when an operator resolved it
             */
            resolvedAt: string;
            /**
             * Format: *uuid.UUID
             * @description who resolved it
             */
            resolvedBy: string;
            /**
             * Format: time.Time
             * @description when the event was recorded
             */
            createdAt: string;
        };
        "audit.ResolveSecurityEventReq": {
            /**
             * Format: int64
             * @description the event id
             */
            id: number;
            /**
             * Format: string
             * @description what was done about it, written to the audit log
             */
            note: string;
        };
        "audit.ResolveSecurityEventRes": {
            /**
             * Format: int64
             * @description the resolved event
             */
            id: number;
            /**
             * Format: time.Time
             * @description when it was resolved
             */
            resolvedAt: string;
            /**
             * Format: uuid.UUID
             * @description who resolved it
             */
            resolvedBy: string;
        };
        "users.GetAllUsersReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query: string;
            sort: components["schemas"]["http.SortParams"];
        };
        "users.User": {
            /** Format: uuid.UUID */
            id: string;
            /** Format: string */
            name: string;
            /** Format: string */
            email: string;
            /** Format: string */
            phone: string;
            /** Format: string */
            username: string;
            /** Format: bool */
            email_verified: boolean;
            /** Format: bool */
            phone_verified: boolean;
            /** Format: string */
            image: string;
            /** Format: []string */
            roles: string[];
            /** Format: []string */
            permissions: string[];
            /** Format: string */
            status: string;
            /** Format: map[string]interface {} */
            preference: {
                [key: string]: components["schemas"]["interface"];
            };
            /** Format: bool */
            mfa_enabled: boolean;
            /** Format: time.Time */
            created_at: string;
            /** Format: time.Time */
            updated_at: string;
        };
        "users.CreateUserReq": {
            /**
             * Format: string
             * @description Display name
             */
            name: string;
            /**
             * Format: *string
             * @description Email (optional if phone or username given)
             */
            email?: string;
            /**
             * Format: *string
             * @description Phone (optional if email or username given)
             */
            phone?: string;
            /**
             * Format: *string
             * @description Username (optional if email or phone given)
             */
            username?: string;
            /**
             * Format: string
             * @description Account password (subject to the configured policy)
             */
            password: string;
            /**
             * Format: int
             * @description the role id of the user
             * @default 1
             */
            role: number;
        };
        "users.CreateUserRes": {
            /** Format: uuid.UUID */
            id: string;
        };
        "users.DeleteUserReq": {
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: string
             * @description why the account is removed, written to the audit log
             */
            reason: string;
            /**
             * Format: *bool
             * @description revoke every active session of that user
             * @default true
             */
            revokeSessions: boolean;
        };
        "users.DeleteUserRes": {
            /**
             * Format: uuid.UUID
             * @description the soft deleted user
             */
            id: string;
            /**
             * Format: bool
             * @description always true on success
             */
            deleted: boolean;
        };
        "users.GetUserReq": {
            /** Format: uuid.UUID */
            id: string;
        };
        "users.UpdateUserReq": {
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: *string
             * @description display name
             */
            name?: string;
            /**
             * Format: *string
             * @description avatar url
             */
            image?: string;
            /**
             * Format: *string
             * @description replaces the address without verification, audited
             */
            email?: string;
            /**
             * Format: *string
             * @description replaces the number without verification, audited
             */
            phone?: string;
            /**
             * Format: *string
             * @description replaces the username without verification, audited
             */
            username?: string;
        };
        "users.UpdateUserRes": {
            /**
             * Format: uuid.UUID
             * @description the updated user
             */
            id: string;
        };
        "users.ListUserAttributesReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query?: string;
            sort?: components["schemas"]["http.SortParams"];
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: *string
             * @description restrict to one namespace
             */
            namespace?: string;
        };
        "users.UserAttribute": {
            /**
             * Format: int64
             * @description internal attribute id
             */
            id: number;
            /**
             * Format: uuid.UUID
             * @description the user the attribute belongs to
             */
            userId: string;
            /**
             * Format: string
             * @description org, project, system or default
             */
            namespace: string;
            /**
             * Format: string
             * @description attribute name
             */
            key: string;
            /**
             * Format: string
             * @description the raw stored value
             */
            value: string;
            /**
             * Format: string
             * @description string, number, boolean, json or date
             */
            valueType: string;
            /**
             * Format: time.Time
             * @description when the attribute was written
             */
            updatedAt: string;
        };
        "users.SetUserAttributeReq": {
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: string
             * @description org, project, system or default
             * @default default
             */
            namespace: string;
            /**
             * Format: string
             * @description attribute name
             */
            key: string;
            /**
             * Format: string
             * @description the value, stored as a string
             */
            value: string;
            /**
             * Format: string
             * @description how the value is parsed during policy evaluation
             * @default string
             * @enum {string}
             */
            valueType: "string" | "number" | "boolean" | "json" | "date";
        };
        "users.SetUserAttributeRes": {
            /**
             * Format: uuid.UUID
             * @description the user
             */
            userId: string;
            /**
             * Format: string
             * @description the namespace written
             */
            namespace: string;
            /**
             * Format: string
             * @description the attribute written
             */
            key: string;
        };
        "users.DeleteUserAttributeReq": {
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: string
             * @description the attribute name
             */
            key: string;
            /**
             * Format: *string
             * @description namespace the attribute lives in
             * @default default
             */
            namespace: string;
        };
        "users.DeleteUserAttributeRes": {
            /**
             * Format: uuid.UUID
             * @description the user
             */
            userId: string;
            /**
             * Format: string
             * @description the deleted attribute
             */
            key: string;
            /**
             * Format: bool
             * @description always true on success
             */
            deleted: boolean;
        };
        "users.BanUserReq": {
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: string
             * @description why the ban is applied, written to the audit log
             */
            reason: string;
            /**
             * Format: *time.Time
             * @description when the ban lifts, nil bans indefinitely
             */
            until?: string;
            /**
             * Format: *bool
             * @description revoke every active session of that user
             * @default true
             */
            revokeSessions: boolean;
        };
        "users.BanUserRes": {
            /**
             * Format: uuid.UUID
             * @description the banned user
             */
            id: string;
            /**
             * Format: string
             * @description the resulting user status
             */
            status: string;
            /**
             * Format: *time.Time
             * @description when the ban lifts
             */
            until: string;
        };
        "users.ImpersonateUserReq": {
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: int
             * @description token lifetime in minutes, capped by configuration
             * @default 15
             */
            duration: number;
            /**
             * Format: string
             * @description why the impersonation is needed, written to the audit log
             */
            reason: string;
        };
        "users.ImpersonateUserRes": {
            /**
             * Format: string
             * @description the impersonation token
             */
            token: string;
            /**
             * Format: string
             * @description always Bearer
             */
            tokenType: string;
            /**
             * Format: int64
             * @description unix expiry, short by design
             */
            expiresAt: number;
            /**
             * Format: time.Time
             * @description when the impersonation started
             */
            impersonatedAt: string;
        };
        "users.ResetUserPasswordReq": {
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: string
             * @description the new password, subject to the configured policy
             */
            password: string;
            /**
             * Format: *bool
             * @description force a change on next sign-in
             * @default true
             */
            mustChangePassword: boolean;
            /**
             * Format: *bool
             * @description revoke every active session of that user
             * @default true
             */
            revokeSessions: boolean;
        };
        "users.ResetUserPasswordRes": {
            /**
             * Format: uuid.UUID
             * @description the user whose password was reset
             */
            id: string;
            /**
             * Format: string
             * @description password reset
             */
            message: string;
        };
        "users.ListUserRolesReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query?: string;
            sort?: components["schemas"]["http.SortParams"];
            /** Format: uuid.UUID */
            id: string;
        };
        "access.UserRole": {
            /** Format: uuid.UUID */
            UserID: string;
            /** Format: int32 */
            RoleID: number;
            /** Format: *uuid.UUID */
            AssignedBy: string;
            /** Format: time.Time */
            AssignedAt: string;
            /** Format: *time.Time */
            ExpiresAt: string;
            /** Format: bool */
            IsPrimary: boolean;
        };
        "users.AssignUserRoleReq": {
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: int32
             * @description the role to assign
             */
            roleId: number;
            /**
             * Format: *bool
             * @description make this the user's primary role, demoting the previous one
             */
            isPrimary?: boolean;
            /**
             * Format: *time.Time
             * @description when the assignment lapses, nil grants it permanently
             */
            expiresAt?: string;
        };
        "users.AssignUserRoleRes": {
            /**
             * Format: uuid.UUID
             * @description the user
             */
            id: string;
            /**
             * Format: int32
             * @description the assigned role
             */
            roleId: number;
            /**
             * Format: bool
             * @description true when the role became primary
             */
            primary: boolean;
        };
        "users.RevokeUserRoleReq": {
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: int32
             * @description the role to revoke
             */
            roleId: number;
            /**
             * Format: string
             * @description why, written to the audit log
             */
            reason?: string;
        };
        "users.RevokeUserRoleRes": {
            /**
             * Format: uuid.UUID
             * @description the user
             */
            id: string;
            /**
             * Format: int32
             * @description the revoked role
             */
            roleId: number;
            /**
             * Format: bool
             * @description always true on success
             */
            revoked: boolean;
        };
        "users.SetUserStatusReq": {
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: string
             * @description the target status
             * @enum {string}
             */
            status: "active" | "suspended" | "locked" | "pending_verification";
            /**
             * Format: string
             * @description why, written to the audit log
             */
            reason: string;
        };
        "users.SetUserStatusRes": {
            /**
             * Format: uuid.UUID
             * @description the user
             */
            id: string;
            /**
             * Format: string
             * @description the resulting status
             */
            status: string;
        };
        "users.UnbanUserReq": {
            /** Format: uuid.UUID */
            id: string;
            /**
             * Format: string
             * @description why the ban is lifted, written to the audit log
             */
            reason?: string;
        };
        "users.UnbanUserRes": {
            /**
             * Format: uuid.UUID
             * @description the unbanned user
             */
            id: string;
            /**
             * Format: string
             * @description the resulting user status
             */
            status: string;
        };
        "authentication.ChangePasswordReq": {
            /**
             * Format: string
             * @description the current password, proves possession
             */
            currentPassword: string;
            /**
             * Format: string
             * @description the new password, subject to the configured policy
             */
            newPassword: string;
            /**
             * Format: *bool
             * @description revoke every other session after the change
             * @default true
             */
            revokeSessions: boolean;
        };
        "authentication.ChangePasswordRes": {
            /**
             * Format: string
             * @description password changed
             */
            message: string;
        };
        "authentication.ForgotPasswordReq": {
            /**
             * Format: string
             * @description email, phone or username of the account
             */
            identifier: string;
            /**
             * Format: string
             * @description channel the reset code is delivered on
             * @default email
             * @enum {string}
             */
            deliveryChannel: "email" | "sms";
        };
        "authentication.ForgotPasswordRes": {
            /**
             * Format: string
             * @description generic acknowledgement, never reveals whether the account exists
             */
            message: string;
            /**
             * Format: int64
             * @description unix expiry of the reset code
             */
            expiresAt: number;
            /**
             * Format: int
             * @description seconds to wait before requesting another code
             */
            retryAfter: number;
        };
        "authentication.LogoutRequest": Record<string, never>;
        "authentication.LogoutResponse": {
            /** Format: string */
            message: string;
        };
        "authentication.LogoutAllReq": Record<string, never>;
        "authentication.LogoutAllRes": {
            /**
             * Format: string
             * @description signed out of all sessions
             */
            message: string;
        };
        "authentication.ChallengeMFAReq": {
            /**
             * Format: string
             * @description short-lived token returned by sign-in when mfa is required
             */
            challengeToken: string;
            /**
             * Format: string
             * @description the TOTP code or one of the backup codes
             */
            code: string;
            /**
             * Format: string
             * @description which second factor produced the code
             * @default totp
             * @enum {string}
             */
            method: "totp" | "backup_code";
        };
        "authentication.ChallengeMFARes": {
            /**
             * Format: string
             * @description Bearer access token
             */
            token: string;
            /**
             * Format: string
             * @description Always "Bearer"
             */
            token_type: string;
            /**
             * Format: int64
             * @description Unix expiry of the access token
             */
            expires_at: number;
            /** Format: uuid.UUID */
            id: string;
            /** Format: string */
            name: string;
            /** Format: string */
            email: string;
            /** Format: string */
            phone: string;
            /** Format: string */
            username: string;
            /** Format: bool */
            email_verified: boolean;
            /** Format: bool */
            phone_verified: boolean;
            /** Format: string */
            image: string;
            /** Format: []string */
            roles: string[];
            /** Format: []string */
            permissions: string[];
            /** Format: string */
            status: string;
            /** Format: map[string]interface {} */
            preference: {
                [key: string]: components["schemas"]["interface"];
            };
            /** Format: bool */
            mfa_enabled: boolean;
            /** Format: time.Time */
            created_at: string;
            /** Format: time.Time */
            updated_at: string;
        };
        "authentication.ConfirmMFAReq": {
            /**
             * Format: string
             * @description a code generated from the freshly enrolled secret
             */
            code: string;
            /**
             * Format: string
             * @description the secret returned by /auth/mfa/enable, when the client holds it
             */
            secret?: string;
        };
        "authentication.ConfirmMFARes": {
            /**
             * Format: bool
             * @description true once the factor is active
             */
            enabled: boolean;
            /**
             * Format: int64
             * @description unix time the factor became active
             */
            enabledAt: number;
        };
        "authentication.DisableMFAReq": {
            /**
             * Format: string
             * @description the current password
             */
            password: string;
            /**
             * Format: string
             * @description a current TOTP code or a backup code
             */
            code: string;
        };
        "authentication.DisableMFARes": {
            /**
             * Format: string
             * @description multi-factor authentication disabled
             */
            message: string;
        };
        "authentication.EnableMFAReq": {
            /**
             * Format: string
             * @description the current password, proves possession before enrolling
             */
            password: string;
            /**
             * Format: *string
             * @description issuer shown in the authenticator app
             */
            issuer?: string;
        };
        "authentication.EnableMFARes": {
            /**
             * Format: string
             * @description base32 TOTP secret, shown once
             */
            secret: string;
            /**
             * Format: string
             * @description otpauth:// URI to render as a QR code
             */
            otpauthUrl: string;
            /**
             * Format: []string
             * @description single-use recovery codes, stored hashed
             */
            backupCodes: string[];
        };
        "authentication.OAuthAuthorizeReq": {
            /**
             * Format: authentication.OAuthProvider
             * @description the OAuth provider
             * @enum {string}
             */
            provider: "google" | "github" | "apple" | "microsoft";
            /**
             * Format: *string
             * @description callback url, must be pre-registered
             */
            redirectUri?: string;
            /**
             * Format: []string
             * @description requested scopes, defaults to the provider's openid profile set
             */
            scopes?: string[];
            /**
             * Format: *string
             * @description opaque value echoed back on the callback, used for CSRF protection
             */
            state?: string;
        };
        "authentication.OAuthAuthorizeRes": {
            /**
             * Format: string
             * @description where the user agent must be sent
             */
            authorizationUrl: string;
            /**
             * Format: string
             * @description the state value to send back on the callback
             */
            state: string;
            /**
             * Format: authentication.OAuthProvider
             * @description echo of the provider slug
             */
            provider: string;
            /**
             * Format: []string
             * @description the scopes actually requested
             */
            scopes: string[];
            /**
             * Format: *string
             * @description id_token nonce to validate on the callback
             */
            nonce: string;
        };
        "authentication.OAuthCallbackReq": {
            /**
             * Format: authentication.OAuthProvider
             * @description the OAuth provider
             * @enum {string}
             */
            provider: "google" | "github" | "apple" | "microsoft";
            /**
             * Format: string
             * @description the authorization code returned by the provider
             */
            code: string;
            /**
             * Format: string
             * @description the state value issued by /oauth/:provider/authorize
             */
            state: string;
        };
        "authentication.OAuthCallbackRes": {
            /**
             * Format: string
             * @description Bearer access token
             */
            token: string;
            /**
             * Format: string
             * @description Always "Bearer"
             */
            token_type: string;
            /**
             * Format: int64
             * @description Unix expiry of the access token
             */
            expires_at: number;
            /** Format: uuid.UUID */
            id: string;
            /** Format: string */
            name: string;
            /** Format: string */
            email: string;
            /** Format: string */
            phone: string;
            /** Format: string */
            username: string;
            /** Format: bool */
            email_verified: boolean;
            /** Format: bool */
            phone_verified: boolean;
            /** Format: string */
            image: string;
            /** Format: []string */
            roles: string[];
            /** Format: []string */
            permissions: string[];
            /** Format: string */
            status: string;
            /** Format: map[string]interface {} */
            preference: {
                [key: string]: components["schemas"]["interface"];
            };
            /** Format: bool */
            mfa_enabled: boolean;
            /** Format: time.Time */
            created_at: string;
            /** Format: time.Time */
            updated_at: string;
            /**
             * Format: bool
             * @description true when the provider identity created the local account
             */
            isNewUser: boolean;
        };
        "authentication.RefreshRequest": Record<string, never>;
        "authentication.RefreshResponse": {
            /** Format: string */
            token: string;
            /** Format: string */
            token_type: string;
            /** Format: int64 */
            expires_at: number;
        };
        "authentication.ResetPasswordReq": {
            /**
             * Format: string
             * @description email, phone or username of the account
             */
            identifier: string;
            /**
             * Format: string
             * @description the one-time reset code
             */
            code: string;
            /**
             * Format: string
             * @description the token issued by reset-password/verify
             */
            token?: string;
            /**
             * Format: string
             * @description the new password, subject to the configured policy
             */
            password: string;
            /**
             * Format: *bool
             * @description revoke every other session after the reset
             * @default true
             */
            revokeSessions: boolean;
        };
        "authentication.ResetPasswordRes": {
            /**
             * Format: string
             * @description password updated
             */
            message: string;
        };
        "authentication.VerifyResetPasswordReq": {
            /**
             * Format: string
             * @description email, phone or username the code was sent to
             */
            identifier: string;
            /**
             * Format: string
             * @description the one-time reset code
             */
            code: string;
            /**
             * Format: string
             * @description the short-lived token returned by the forgot-password call, when one is used
             */
            token?: string;
        };
        "authentication.VerifyResetPasswordRes": {
            /**
             * Format: bool
             * @description true when the code exists, is unexpired and has attempts left
             */
            valid: boolean;
            /**
             * Format: int64
             * @description unix expiry of the code
             */
            expiresAt: number;
            /**
             * Format: int
             * @description attempts already consumed
             */
            attempts: number;
        };
        "authentication.RevokeAllSessionsReq": Record<string, never>;
        "authentication.RevokeAllSessionsRes": {
            /**
             * Format: string
             * @description all sessions revoked
             */
            message: string;
        };
        "authentication.ListSessionsReq": {
            /**
             * Format: int
             * @default 1
             */
            page: number;
            /**
             * Format: int
             * @default 20
             */
            perPage: number;
            /** Format: string */
            query: string;
            sort: components["schemas"]["http.SortParams"];
        };
        "auth.SessionData": {
            /** Format: string */
            id: string;
            /** Format: string */
            user_id: string;
            /** Format: string */
            identifier: string;
            /** Format: string */
            status: string;
            /** Format: string */
            role: string;
            /** Format: []string */
            roles: string[];
            /** Format: []string */
            permissions: string[];
            /** Format: string */
            token: string;
            /** Format: time.Time */
            created_at: string;
            /** Format: time.Time */
            expires_at: string;
        };
        "authentication.RevokeSessionReq": {
            /**
             * Format: string
             * @description the session id (token id)
             */
            id: string;
        };
        "authentication.RevokeSessionRes": {
            /**
             * Format: string
             * @description the revoked session id
             */
            id: string;
            /**
             * Format: bool
             * @description always true on success
             */
            revoked: boolean;
        };
        "authentication.SignInRequest": {
            /**
             * Format: string
             * @description Email, phone, or username
             * @example raad@email.com
             */
            identifier: string;
            /**
             * Format: string
             * @description Account password
             * @example Password123#
             */
            password: string;
        };
        "authentication.SignInResponse": {
            /**
             * Format: string
             * @description Bearer access token
             */
            token: string;
            /**
             * Format: string
             * @description Always "Bearer"
             */
            token_type: string;
            /**
             * Format: int64
             * @description Unix expiry of the access token
             */
            expires_at: number;
            /** Format: uuid.UUID */
            id: string;
            /** Format: string */
            name: string;
            /** Format: string */
            email: string;
            /** Format: string */
            phone: string;
            /** Format: string */
            username: string;
            /** Format: bool */
            email_verified: boolean;
            /** Format: bool */
            phone_verified: boolean;
            /** Format: string */
            image: string;
            /** Format: []string */
            roles: string[];
            /** Format: []string */
            permissions: string[];
            /** Format: string */
            status: string;
            /** Format: map[string]interface {} */
            preference: {
                [key: string]: components["schemas"]["interface"];
            };
            /** Format: bool */
            mfa_enabled: boolean;
            /** Format: time.Time */
            created_at: string;
            /** Format: time.Time */
            updated_at: string;
        };
        "authentication.SignUpRequest": {
            /**
             * Format: string
             * @description Display name
             */
            name: string;
            /**
             * Format: *string
             * @description Email (optional if phone or username given)
             */
            email?: string;
            /**
             * Format: *string
             * @description Phone (optional if email or username given)
             */
            phone?: string;
            /**
             * Format: *string
             * @description Username (optional if email or phone given)
             */
            username?: string;
            /**
             * Format: string
             * @description Account password (subject to the configured policy)
             */
            password: string;
        };
        "authentication.SignUpResponse": {
            /** Format: string */
            message: string;
            /** Format: string */
            user_id: string;
        };
        "authentication.ResendVerificationReq": {
            /**
             * Format: string
             * @description email, phone or username the code is sent to
             */
            identifier: string;
            /**
             * Format: string
             * @description which identifier kind was supplied
             * @enum {string}
             */
            identifierType: "email" | "phone" | "username";
            /**
             * Format: string
             * @description what the code verifies
             * @enum {string}
             */
            purpose: "email_verification" | "phone_verification" | "account_recovery" | "email_change" | "phone_change";
            /**
             * Format: *string
             * @description required for the email_change and phone_change purposes
             */
            newIdentifier?: string;
        };
        "authentication.ResendVerificationRes": {
            /**
             * Format: string
             * @description always a generic acknowledgement, whether or not the identifier exists
             */
            message: string;
            /**
             * Format: int64
             * @description unix expiry of the freshly issued code
             */
            expiresAt: number;
            /**
             * Format: int
             * @description seconds to wait before requesting another code
             */
            retryAfter: number;
        };
        "authentication.VerifyEmailReq": {
            /**
             * Format: string
             * @description the email address the code was sent to
             */
            email: string;
            /**
             * Format: string
             * @description the one-time code
             */
            code: string;
        };
        "authentication.VerifyEmailRes": {
            /**
             * Format: bool
             * @description true when the address is now verified
             */
            verified: boolean;
            /**
             * Format: string
             * @description the verified address
             */
            email: string;
        };
        "authentication.VerifyPhoneReq": {
            /**
             * Format: string
             * @description the phone number in E.164 format the code was sent to
             */
            phone: string;
            /**
             * Format: string
             * @description the one-time code
             */
            code: string;
        };
        "authentication.VerifyPhoneRes": {
            /**
             * Format: bool
             * @description true when the number is now verified
             */
            verified: boolean;
            /**
             * Format: string
             * @description the verified number
             */
            phone: string;
        };
    };
    responses: never;
    parameters: never;
    requestBodies: never;
    headers: never;
    pathItems: never;
};
export type AccountDeleteAccountReq = components['schemas']['account.DeleteAccountReq'];
export type AccountDeleteAccountRes = components['schemas']['account.DeleteAccountRes'];
export type CommonPaginationMeta = components['schemas']['common.PaginationMeta'];
export type Interface = components['schemas']['interface'];
export type HttpBadRequestResponse = components['schemas']['http.BadRequestResponse'];
export type HttpInternalErrorResponse = components['schemas']['http.InternalErrorResponse'];
export type HttpUnAuthenticatedResponse = components['schemas']['http.UnAuthenticatedResponse'];
export type HttpInvalidInputErrorResponse = components['schemas']['http.InvalidInputErrorResponse'];
export type AccountUpdateProfileReq = components['schemas']['account.UpdateProfileReq'];
export type AccountUpdateProfileRes = components['schemas']['account.UpdateProfileRes'];
export type AccountUploadAvatarReq = components['schemas']['account.UploadAvatarReq'];
export type AccountUploadAvatarRes = components['schemas']['account.UploadAvatarRes'];
export type AccountListConnectionsReq = components['schemas']['account.ListConnectionsReq'];
export type HttpSortParams = components['schemas']['http.SortParams'];
export type UsersConnection = components['schemas']['users.Connection'];
export type AccountUnlinkConnectionReq = components['schemas']['account.UnlinkConnectionReq'];
export type AccountUnlinkConnectionRes = components['schemas']['account.UnlinkConnectionRes'];
export type AccountChangeEmailReq = components['schemas']['account.ChangeEmailReq'];
export type AccountChangeEmailRes = components['schemas']['account.ChangeEmailRes'];
export type AccountMeRequest = components['schemas']['account.MeRequest'];
export type AccountMeResponse = components['schemas']['account.MeResponse'];
export type AccountChangePhoneReq = components['schemas']['account.ChangePhoneReq'];
export type AccountChangePhoneRes = components['schemas']['account.ChangePhoneRes'];
export type AccountGetPreferenceReq = components['schemas']['account.GetPreferenceReq'];
export type AccountUpdatePreferenceReq = components['schemas']['account.UpdatePreferenceReq'];
export type ApikeysListApiKeysReq = components['schemas']['apikeys.ListAPIKeysReq'];
export type AccessApiKey = components['schemas']['access.APIKey'];
export type ApikeysCreateApiKeyReq = components['schemas']['apikeys.CreateAPIKeyReq'];
export type ApikeysCreateApiKeyRes = components['schemas']['apikeys.CreateAPIKeyRes'];
export type ApikeysRevokeApiKeyReq = components['schemas']['apikeys.RevokeAPIKeyReq'];
export type ApikeysRevokeApiKeyRes = components['schemas']['apikeys.RevokeAPIKeyRes'];
export type ApikeysGetApiKeyReq = components['schemas']['apikeys.GetAPIKeyReq'];
export type ApikeysGetApiKeyRes = components['schemas']['apikeys.GetAPIKeyRes'];
export type ApikeysUpdateApiKeyReq = components['schemas']['apikeys.UpdateAPIKeyReq'];
export type ApikeysUpdateApiKeyRes = components['schemas']['apikeys.UpdateAPIKeyRes'];
export type ApikeysRotateApiKeyReq = components['schemas']['apikeys.RotateAPIKeyReq'];
export type ApikeysRotateApiKeyRes = components['schemas']['apikeys.RotateAPIKeyRes'];
export type AuditListAuditLogsReq = components['schemas']['audit.ListAuditLogsReq'];
export type AccessAuditEntry = components['schemas']['access.AuditEntry'];
export type AuditGetAuditLogReq = components['schemas']['audit.GetAuditLogReq'];
export type AuditGetAuditLogRes = components['schemas']['audit.GetAuditLogRes'];
export type RbacListPermissionReq = components['schemas']['rbac.ListPermissionReq'];
export type AccessPermission = components['schemas']['access.Permission'];
export type RbacCreatePermissionReq = components['schemas']['rbac.CreatePermissionReq'];
export type RbacCreatePermissionRes = components['schemas']['rbac.CreatePermissionRes'];
export type RbacDeletePermissionReq = components['schemas']['rbac.DeletePermissionReq'];
export type RbacDeletePermissionRes = components['schemas']['rbac.DeletePermissionRes'];
export type RbacGetPermissionReq = components['schemas']['rbac.GetPermissionReq'];
export type RbacGetPermissionRes = components['schemas']['rbac.GetPermissionRes'];
export type RbacUpdatePermissionReq = components['schemas']['rbac.UpdatePermissionReq'];
export type RbacUpdatePermissionRes = components['schemas']['rbac.UpdatePermissionRes'];
export type AbacListPoliciesReq = components['schemas']['abac.ListPoliciesReq'];
export type AccessPolicy = components['schemas']['access.Policy'];
export type AbacCreatePolicyReq = components['schemas']['abac.CreatePolicyReq'];
export type AbacCreatePolicyRes = components['schemas']['abac.CreatePolicyRes'];
export type AbacDeletePolicyReq = components['schemas']['abac.DeletePolicyReq'];
export type AbacDeletePolicyRes = components['schemas']['abac.DeletePolicyRes'];
export type AbacGetPolicyReq = components['schemas']['abac.GetPolicyReq'];
export type AbacGetPolicyRes = components['schemas']['abac.GetPolicyRes'];
export type AbacUpdatePolicyReq = components['schemas']['abac.UpdatePolicyReq'];
export type AbacUpdatePolicyRes = components['schemas']['abac.UpdatePolicyRes'];
export type AbacEvaluatePolicyReq = components['schemas']['abac.EvaluatePolicyReq'];
export type AbacEvaluatePolicyRes = components['schemas']['abac.EvaluatePolicyRes'];
export type RbacListRolesReq = components['schemas']['rbac.ListRolesReq'];
export type AccessRole = components['schemas']['access.Role'];
export type RbacCreateRoleReq = components['schemas']['rbac.CreateRoleReq'];
export type RbacCreateRoleRes = components['schemas']['rbac.CreateRoleRes'];
export type RbacDeleteRoleReq = components['schemas']['rbac.DeleteRoleReq'];
export type RbacDeleteRoleRes = components['schemas']['rbac.DeleteRoleRes'];
export type RbacGetRoleReq = components['schemas']['rbac.GetRoleReq'];
export type RbacGetRoleRes = components['schemas']['rbac.GetRoleRes'];
export type RbacUpdateRoleReq = components['schemas']['rbac.UpdateRoleReq'];
export type RbacUpdateRoleRes = components['schemas']['rbac.UpdateRoleRes'];
export type RbacListRolePermissionsReq = components['schemas']['rbac.ListRolePermissionsReq'];
export type AccessRolePermission = components['schemas']['access.RolePermission'];
export type RbacAssignRolePermissionReq = components['schemas']['rbac.AssignRolePermissionReq'];
export type RbacAssignRolePermissionRes = components['schemas']['rbac.AssignRolePermissionRes'];
export type RbacRevokeRolePermissionReq = components['schemas']['rbac.RevokeRolePermissionReq'];
export type RbacRevokeRolePermissionRes = components['schemas']['rbac.RevokeRolePermissionRes'];
export type AuditListSecurityEventsReq = components['schemas']['audit.ListSecurityEventsReq'];
export type AccessSecurityEvent = components['schemas']['access.SecurityEvent'];
export type AuditGetSecurityEventReq = components['schemas']['audit.GetSecurityEventReq'];
export type AuditGetSecurityEventRes = components['schemas']['audit.GetSecurityEventRes'];
export type AuditResolveSecurityEventReq = components['schemas']['audit.ResolveSecurityEventReq'];
export type AuditResolveSecurityEventRes = components['schemas']['audit.ResolveSecurityEventRes'];
export type UsersGetAllUsersReq = components['schemas']['users.GetAllUsersReq'];
export type UsersUser = components['schemas']['users.User'];
export type UsersCreateUserReq = components['schemas']['users.CreateUserReq'];
export type UsersCreateUserRes = components['schemas']['users.CreateUserRes'];
export type UsersDeleteUserReq = components['schemas']['users.DeleteUserReq'];
export type UsersDeleteUserRes = components['schemas']['users.DeleteUserRes'];
export type UsersGetUserReq = components['schemas']['users.GetUserReq'];
export type UsersUpdateUserReq = components['schemas']['users.UpdateUserReq'];
export type UsersUpdateUserRes = components['schemas']['users.UpdateUserRes'];
export type UsersListUserAttributesReq = components['schemas']['users.ListUserAttributesReq'];
export type UsersUserAttribute = components['schemas']['users.UserAttribute'];
export type UsersSetUserAttributeReq = components['schemas']['users.SetUserAttributeReq'];
export type UsersSetUserAttributeRes = components['schemas']['users.SetUserAttributeRes'];
export type UsersDeleteUserAttributeReq = components['schemas']['users.DeleteUserAttributeReq'];
export type UsersDeleteUserAttributeRes = components['schemas']['users.DeleteUserAttributeRes'];
export type UsersBanUserReq = components['schemas']['users.BanUserReq'];
export type UsersBanUserRes = components['schemas']['users.BanUserRes'];
export type UsersImpersonateUserReq = components['schemas']['users.ImpersonateUserReq'];
export type UsersImpersonateUserRes = components['schemas']['users.ImpersonateUserRes'];
export type UsersResetUserPasswordReq = components['schemas']['users.ResetUserPasswordReq'];
export type UsersResetUserPasswordRes = components['schemas']['users.ResetUserPasswordRes'];
export type UsersListUserRolesReq = components['schemas']['users.ListUserRolesReq'];
export type AccessUserRole = components['schemas']['access.UserRole'];
export type UsersAssignUserRoleReq = components['schemas']['users.AssignUserRoleReq'];
export type UsersAssignUserRoleRes = components['schemas']['users.AssignUserRoleRes'];
export type UsersRevokeUserRoleReq = components['schemas']['users.RevokeUserRoleReq'];
export type UsersRevokeUserRoleRes = components['schemas']['users.RevokeUserRoleRes'];
export type UsersSetUserStatusReq = components['schemas']['users.SetUserStatusReq'];
export type UsersSetUserStatusRes = components['schemas']['users.SetUserStatusRes'];
export type UsersUnbanUserReq = components['schemas']['users.UnbanUserReq'];
export type UsersUnbanUserRes = components['schemas']['users.UnbanUserRes'];
export type AuthenticationChangePasswordReq = components['schemas']['authentication.ChangePasswordReq'];
export type AuthenticationChangePasswordRes = components['schemas']['authentication.ChangePasswordRes'];
export type AuthenticationForgotPasswordReq = components['schemas']['authentication.ForgotPasswordReq'];
export type AuthenticationForgotPasswordRes = components['schemas']['authentication.ForgotPasswordRes'];
export type AuthenticationLogoutRequest = components['schemas']['authentication.LogoutRequest'];
export type AuthenticationLogoutResponse = components['schemas']['authentication.LogoutResponse'];
export type AuthenticationLogoutAllReq = components['schemas']['authentication.LogoutAllReq'];
export type AuthenticationLogoutAllRes = components['schemas']['authentication.LogoutAllRes'];
export type AuthenticationChallengeMfaReq = components['schemas']['authentication.ChallengeMFAReq'];
export type AuthenticationChallengeMfaRes = components['schemas']['authentication.ChallengeMFARes'];
export type AuthenticationConfirmMfaReq = components['schemas']['authentication.ConfirmMFAReq'];
export type AuthenticationConfirmMfaRes = components['schemas']['authentication.ConfirmMFARes'];
export type AuthenticationDisableMfaReq = components['schemas']['authentication.DisableMFAReq'];
export type AuthenticationDisableMfaRes = components['schemas']['authentication.DisableMFARes'];
export type AuthenticationEnableMfaReq = components['schemas']['authentication.EnableMFAReq'];
export type AuthenticationEnableMfaRes = components['schemas']['authentication.EnableMFARes'];
export type AuthenticationOAuthAuthorizeReq = components['schemas']['authentication.OAuthAuthorizeReq'];
export type AuthenticationOAuthAuthorizeRes = components['schemas']['authentication.OAuthAuthorizeRes'];
export type AuthenticationOAuthCallbackReq = components['schemas']['authentication.OAuthCallbackReq'];
export type AuthenticationOAuthCallbackRes = components['schemas']['authentication.OAuthCallbackRes'];
export type AuthenticationRefreshRequest = components['schemas']['authentication.RefreshRequest'];
export type AuthenticationRefreshResponse = components['schemas']['authentication.RefreshResponse'];
export type AuthenticationResetPasswordReq = components['schemas']['authentication.ResetPasswordReq'];
export type AuthenticationResetPasswordRes = components['schemas']['authentication.ResetPasswordRes'];
export type AuthenticationVerifyResetPasswordReq = components['schemas']['authentication.VerifyResetPasswordReq'];
export type AuthenticationVerifyResetPasswordRes = components['schemas']['authentication.VerifyResetPasswordRes'];
export type AuthenticationRevokeAllSessionsReq = components['schemas']['authentication.RevokeAllSessionsReq'];
export type AuthenticationRevokeAllSessionsRes = components['schemas']['authentication.RevokeAllSessionsRes'];
export type AuthenticationListSessionsReq = components['schemas']['authentication.ListSessionsReq'];
export type AuthSessionData = components['schemas']['auth.SessionData'];
export type AuthenticationRevokeSessionReq = components['schemas']['authentication.RevokeSessionReq'];
export type AuthenticationRevokeSessionRes = components['schemas']['authentication.RevokeSessionRes'];
export type AuthenticationSignInRequest = components['schemas']['authentication.SignInRequest'];
export type AuthenticationSignInResponse = components['schemas']['authentication.SignInResponse'];
export type AuthenticationSignUpRequest = components['schemas']['authentication.SignUpRequest'];
export type AuthenticationSignUpResponse = components['schemas']['authentication.SignUpResponse'];
export type AuthenticationResendVerificationReq = components['schemas']['authentication.ResendVerificationReq'];
export type AuthenticationResendVerificationRes = components['schemas']['authentication.ResendVerificationRes'];
export type AuthenticationVerifyEmailReq = components['schemas']['authentication.VerifyEmailReq'];
export type AuthenticationVerifyEmailRes = components['schemas']['authentication.VerifyEmailRes'];
export type AuthenticationVerifyPhoneReq = components['schemas']['authentication.VerifyPhoneReq'];
export type AuthenticationVerifyPhoneRes = components['schemas']['authentication.VerifyPhoneRes'];
export type $defs = Record<string, never>;
export type operations = Record<string, never>;
