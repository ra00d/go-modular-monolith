export interface paths {
  "/admin/branches/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    post?: never;
    /** Archive a branch */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.BranchResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a branch */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            zip_code?: string;
            /** Format: *string */
            landmark?: string;
            /** Format: *string */
            mobile?: string;
            /** Format: *string */
            alternate_number?: string;
            /** Format: *string */
            email?: string;
            /** Format: *string */
            website?: string;
            /** Format: *string */
            name?: string;
            /** Format: *string */
            state?: string;
            /** Format: *string */
            city?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.BranchResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/admin/subscriptions/features": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List feature gates */
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
            "application/json": components["schemas"]["qistas.internal.subscription.api.rest.ListFeaturesResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a feature gate */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.subscription.api.rest.CreateFeatureRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.subscription.api.rest.FeatureResponse"];
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
  "/admin/subscriptions/plans": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List plans */
    get: {
      parameters: {
        query?: {
          include_archived?: boolean;
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
            "application/json": components["schemas"]["qistas.internal.subscription.api.rest.ListPlansResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create or update a plan */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.subscription.api.rest.SavePlanRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.subscription.api.rest.PlanResponse"];
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
  "/admin/subscriptions/plans/{planId}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    post?: never;
    /** Archive a plan */
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
            "application/json": components["schemas"]["qistas.internal.subscription.api.rest.ArchivePlanResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/admin/subscriptions/tenants/{tenantId}/subscription": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    /** Change a tenant's plan */
    put: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.subscription.api.rest.ChangeSubscriptionRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.subscription.api.rest.SubscriptionResponse"];
          };
        };
      };
    };
    /** Assign a plan to a tenant */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.subscription.api.rest.AssignSubscriptionRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.subscription.api.rest.SubscriptionResponse"];
          };
        };
      };
    };
    /** Cancel a tenant's subscription */
    delete: {
      parameters: {
        query?: {
          now?: boolean;
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
            "application/json": components["schemas"]["qistas.internal.subscription.api.rest.SubscriptionResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/admin/tenants": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List tenants */
    get: {
      parameters: {
        query?: {
          offset?: number;
          limit?: number;
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
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.ListTenantsResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a tenant */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.tenants.api.rest.CreateTenantRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.TenantResponse"];
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
  "/admin/tenants/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a tenant */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.GetTenantResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Archive a tenant */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.TenantResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a tenant */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            legal_name?: string;
            /** Format: *string */
            tax_number_1?: string;
            /** Format: *string */
            tax_number_2?: string;
            /** Format: *string */
            tax_label_2?: string;
            /** Format: *string */
            code_2?: string;
            /** Format: *string */
            currency?: string;
            /** Format: *string */
            name?: string;
            /** Format: *string */
            tax_label_1?: string;
            /** Format: *string */
            code_1?: string;
            /** Format: *string */
            country?: string;
            /** Format: *string */
            code_label_1?: string;
            /** Format: *string */
            date_format?: string;
            /** Format: *string */
            logo?: string;
            /** Format: *string */
            code_label_2?: string;
            /** Format: *string */
            timezone?: string;
            /**
             * Format: *string
             * @enum {string}
             */
            time_format?: 12 | 24;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.TenantResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/admin/tenants/{id}/activate": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Activate a tenant */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.TenantResponse"];
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
  "/admin/tenants/{id}/branches": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List a tenant's branches */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.ListBranchesResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a branch */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            website?: string;
            /** Format: string */
            name: string;
            /** Format: *string */
            state?: string;
            /** Format: *string */
            city?: string;
            /** Format: *string */
            zip_code?: string;
            /** Format: *string */
            code?: string;
            /** Format: string */
            country: string;
            /** Format: *string */
            landmark?: string;
            /** Format: *string */
            mobile?: string;
            /** Format: *string */
            alternate_number?: string;
            /** Format: *string */
            email?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.BranchResponse"];
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
  "/admin/tenants/{id}/suspend": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Suspend a tenant */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.TenantResponse"];
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
  "/auth/logout": {
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
          "application/json": components["schemas"]["qistas.internal.auth.api.rest.LogoutRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.auth.api.rest.LogoutResponse"];
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
  "/auth/me": {
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
            "application/json": components["schemas"]["qistas.internal.auth.api.rest.MeResponse.qistas.internal.auth.domain.users.User."];
          };
        };
        400: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.BadRequestResponse"];
          };
        };
        /** @description User is not authenticated */
        401: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.UnAuthenticatedResponse"];
          };
        };
        /** @description User is not authenticated */
        403: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.UnAuthenticatedResponse"];
          };
        };
        422: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.InvalidInputErrorResponse"];
          };
        };
        500: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.InternalErrorResponse"];
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
  "/auth/polcies": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
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
            "application/json": components["schemas"]["qistas.internal.auth.api.rest.ListPoliciesRes"];
          };
        };
        400: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.BadRequestResponse"];
          };
        };
        /** @description User is not authenticated */
        401: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.UnAuthenticatedResponse"];
          };
        };
        /** @description User is not authenticated */
        403: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.UnAuthenticatedResponse"];
          };
        };
        422: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.InvalidInputErrorResponse"];
          };
        };
        500: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.InternalErrorResponse"];
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
  "/auth/polcies/one": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
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
            "application/json": components["schemas"]["qistas.internal.auth.api.rest.ListPoliciesRes"];
          };
        };
        400: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.BadRequestResponse"];
          };
        };
        /** @description User is not authenticated */
        401: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.UnAuthenticatedResponse"];
          };
        };
        /** @description User is not authenticated */
        403: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.UnAuthenticatedResponse"];
          };
        };
        422: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.InvalidInputErrorResponse"];
          };
        };
        500: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.InternalErrorResponse"];
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
  "/auth/refresh": {
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
          "application/json": components["schemas"]["qistas.internal.auth.api.rest.RefreshRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.auth.api.rest.RefreshResponse"];
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
  "/auth/sign-in": {
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
          "application/json": components["schemas"]["qistas.internal.auth.api.rest.SignInRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.http.Response.qistas.internal.auth.api.rest.SignInResponse."];
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
  "/contacts": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List contacts */
    get: {
      parameters: {
        query?: {
          type?: string;
          type_mode?: string;
          status?: string;
          customer_group_id?: string;
          assigned_to?: string;
          q?: string;
          offset?: number;
          limit?: number;
          sort?: string;
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
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.ListContactsResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a contact */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.contacts.api.rest.CreateContactRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.ContactResponse"];
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
  "/contacts/check-contact-id": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Check if a contact_id is taken */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.contacts.api.rest.CheckContactIDRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.CheckContactIDResponse"];
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
  "/contacts/check-mobile": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Check if a mobile is taken */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.contacts.api.rest.CheckMobileRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.CheckMobileResponse"];
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
  "/contacts/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a contact */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.ContactResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Soft-delete a contact */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.EmptyResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a contact */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            name?: string;
            /** Format: *string */
            shipping_address?: string;
            /** Format: *string */
            custom_field_10?: string;
            /** Format: *string */
            contact_id?: string;
            /** Format: *string */
            tax_number?: string;
            /** Format: *string */
            landline?: string;
            /** Format: *string */
            export_custom_field_3?: string;
            /** Format: *string */
            address_line_1?: string;
            /** Format: *string */
            zip_code?: string;
            /** Format: *bool */
            is_default?: boolean;
            /** Format: *string */
            shipping_custom_field_details?: string;
            /** Format: *string */
            export_custom_field_2?: string;
            /** Format: *string */
            export_custom_field_4?: string;
            /** Format: *int */
            pay_term_number?: number;
            /** Format: *string */
            credit_limit?: string;
            /** Format: *bool */
            is_export?: boolean;
            /** Format: *string */
            export_custom_field_5?: string;
            /** Format: *string */
            customer_group_id?: string;
            /** Format: *string */
            custom_field_4?: string;
            /** Format: *string */
            custom_field_5?: string;
            /** Format: *string */
            custom_field_9?: string;
            /** Format: *string */
            email?: string;
            /** Format: *string */
            address_line_2?: string;
            /** Format: *string */
            country?: string;
            /** Format: *string */
            dob?: string;
            /** Format: *string */
            position?: string;
            /** Format: *string */
            crm_source?: string;
            /** Format: *string */
            prefix?: string;
            /** Format: *string */
            middle_name?: string;
            /** Format: *string */
            city?: string;
            /** Format: *string */
            alternate_number?: string;
            /** Format: *string */
            pay_term_type?: string;
            /** Format: *string */
            export_custom_field_6?: string;
            /** Format: *string */
            custom_field_8?: string;
            /** Format: *string */
            crm_life_stage?: string;
            /** Format: *string */
            first_name?: string;
            /** Format: *string */
            state?: string;
            /** Format: *string */
            export_custom_field_1?: string;
            /** Format: *string */
            custom_field_1?: string;
            /** Format: *string */
            custom_field_2?: string;
            /** Format: *string */
            custom_field_3?: string;
            /** Format: *string */
            custom_field_6?: string;
            /** Format: *string */
            last_name?: string;
            /** Format: *string */
            custom_field_7?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.ContactResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/contacts/{id}/access": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List users with access */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.AccessResponse"];
          };
        };
      };
    };
    /** Replace contact access list */
    put: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: []string */
            user_ids?: string[];
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.AccessResponse"];
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
  "/contacts/{id}/status": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Change contact status */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /**
             * Format: string
             * @enum {string}
             */
            to: "active" | "inactive";
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.ContactResponse"];
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
  "/customer-groups": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List customer groups */
    get: {
      parameters: {
        query?: {
          offset?: number;
          limit?: number;
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
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.ListGroupsResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a customer group */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.contacts.api.rest.CreateGroupRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.GroupResponse"];
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
  "/customer-groups/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a customer group */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.GroupResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Delete a customer group */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.lib.common.EmptyResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a customer group */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            amount?: string;
            /** Format: *string */
            price_calculation_type?: string;
            /** Format: *string */
            selling_price_group_id?: string;
            /** Format: *string */
            name?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.contacts.api.rest.GroupResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/onboarding/tenant": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Create a tenant for the calling user */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.tenants.api.rest.CreateMyTenantRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.CreateMyTenantResponse"];
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
  "/products/v1/brands": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List brands */
    get: {
      parameters: {
        query?: {
          search?: string;
          page?: number;
          perPage?: number;
          sort?: string;
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
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ListBrandsResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a brand */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.products.api.rest.CreateBrandRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.BrandResponse"];
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
  "/products/v1/brands/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a brand */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.GetBrandResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Soft-delete a brand */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.BrandResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a brand */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            name?: string;
            /** Format: *string */
            description?: string;
            /** Format: *bool */
            use_for_repair?: boolean;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.BrandResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/products/v1/categories": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List categories */
    get: {
      parameters: {
        query?: {
          category_type?: string;
          parent_id?: string;
          search?: string;
          page?: number;
          perPage?: number;
          sort?: string;
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
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ListCategoriesResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a category */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.products.api.rest.CreateCategoryRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.CategoryResponse"];
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
  "/products/v1/categories/sub": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Sub-categories of a parent */
    get: {
      parameters: {
        query: {
          parent_id: string;
          category_type: "product" | "expense" | "both";
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
            "application/json": components["schemas"]["qistas.internal.products.api.rest.SubCategoriesResponse"];
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
  "/products/v1/categories/taxonomies": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Nested taxonomy */
    get: {
      parameters: {
        query: {
          category_type: "product" | "expense" | "both";
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
            "application/json": components["schemas"]["qistas.internal.products.api.rest.TaxonomiesResponse"];
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
  "/products/v1/categories/taxonomies/index": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Flat taxonomy index */
    get: {
      parameters: {
        query: {
          category_type: "product" | "expense" | "both";
          search?: string;
          page?: number;
          perPage?: number;
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
            "application/json": components["schemas"]["qistas.internal.products.api.rest.TaxonomiesIndexResponse"];
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
  "/products/v1/categories/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a category */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.GetCategoryResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Soft-delete a category */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.CategoryResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a category */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            name?: string;
            /** Format: *string */
            short_code?: string;
            /** Format: *string */
            slug?: string;
            /** Format: *string */
            category_type?: string;
            /** Format: *string */
            description?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.CategoryResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/products/v1/categories/{id}/move": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Move a category */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            new_parent_id?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.CategoryResponse"];
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
  "/products/v1/expense-categories": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List expense categories */
    get: {
      parameters: {
        query?: {
          search?: string;
          parent_id?: string;
          page?: number;
          perPage?: number;
          sort?: string;
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
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ListExpenseCategoriesResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create an expense category */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.products.api.rest.CreateExpenseCategoryRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ExpenseCategoryResponse"];
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
  "/products/v1/expense-categories/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get an expense category */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.GetExpenseCategoryResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Soft-delete an expense category */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ExpenseCategoryResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update an expense category */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            name?: string;
            /** Format: *string */
            code?: string;
            /** Format: *string */
            account_id?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ExpenseCategoryResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/products/v1/expense-categories/{id}/move": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Move an expense category */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            new_parent_id?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ExpenseCategoryResponse"];
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
  "/products/v1/import": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Queue an XLSX import job */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "multipart/form-data": components["schemas"]["qistas.internal.products.api.rest.EnqueueImportRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.EnqueueImportResponse"];
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
  "/products/v1/import/{job_id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Import job status */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          job_id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.StatusImportResponse"];
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
  "/products/v1/products": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List products */
    get: {
      parameters: {
        query?: {
          search?: string;
          sku?: string;
          type?: string;
          brand_id?: string;
          category_id?: string;
          is_inactive?: boolean;
          page?: number;
          perPage?: number;
          sort?: string;
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
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ListProductsResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a product */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.products.api.rest.CreateProductRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ProductResponse"];
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
  "/products/v1/products/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a product */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ProductDetailResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Soft-delete a product */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ProductResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a product */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            brand_id?: string;
            /** Format: *string */
            expiry_period?: string;
            /** Format: *string */
            warranty_id?: string;
            /** Format: []rest.UpdateProductVariationDTO */
            variations?: components["schemas"]["qistas.internal.products.api.rest.UpdateProductVariationDTO"][];
            /** Format: *string */
            sub_category_id?: string;
            /** Format: *string */
            alert_quantity?: string;
            /** Format: *string */
            unit_id?: string;
            /** Format: *bool */
            enable_sr_no?: boolean;
            /** Format: *string */
            name?: string;
            /** Format: *string */
            type?: string;
            /** Format: *string */
            category_id?: string;
            /** Format: *string */
            secondary_unit_id?: string;
            /** Format: *string */
            tax_type?: string;
            /** Format: *bool */
            enable_stock?: boolean;
            /** Format: *int */
            preparation_time_in_minutes?: number;
            /** Format: *string */
            image_media_id?: string;
            /** Format: *string */
            expiry_period_type?: string;
            /** Format: bool */
            replace_variations?: boolean;
            /** Format: bool */
            replace_racks?: boolean;
            /** Format: *string */
            sku?: string;
            /** Format: *[]string */
            sub_unit_ids?: string[][];
            /** Format: *map[string]string */
            custom_fields?: {
              [key: string]: string;
            };
            /** Format: *string */
            description?: string;
            /** Format: *string */
            tax_id?: string;
            /** Format: *string */
            weight?: string;
            /** Format: *bool */
            not_for_selling?: boolean;
            /** Format: *string */
            barcode_type?: string;
            /** Format: *bool */
            is_inactive?: boolean;
            /** Format: []rest.CreateRackDTO */
            racks?: components["schemas"]["qistas.internal.products.api.rest.CreateRackDTO"][];
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ProductResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/products/v1/products/{id}/racks": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List racks for a product */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ListRacksResponse"];
          };
        };
      };
    };
    /** Replace the rack set */
    put: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: []rest.CreateRackDTO */
            racks?: components["schemas"]["qistas.internal.products.api.rest.CreateRackDTO"][];
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ProductResponse"];
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
  "/products/v1/products/{id}/variations": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List variations for a product */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ListVariationsResponse"];
          };
        };
      };
    };
    put?: never;
    /** Add a variation to a product */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: string */
            axis_name: string;
            variation?: components["schemas"]["qistas.internal.products.api.rest.CreateVariationDTO"];
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.AddVariationResponse"];
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
  "/products/v1/products/{id}/variations/{vid}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    post?: never;
    /** Soft-delete a variation */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
          vid: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.VariationResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a single variation */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
          vid: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            default_sell_price?: string;
            /** Format: *string */
            sell_price_inc_tax?: string;
            /** Format: *string */
            profit_percent?: string;
            /** Format: *string */
            name?: string;
            /** Format: *string */
            sub_sku?: string;
            /** Format: *string */
            default_purchase_price?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.VariationResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/products/v1/units": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List units */
    get: {
      parameters: {
        query?: {
          search?: string;
          page?: number;
          perPage?: number;
          sort?: string;
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
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ListUnitsResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a unit */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.products.api.rest.CreateUnitRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.UnitResponse"];
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
  "/products/v1/units/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a unit */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.GetUnitResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Soft-delete a unit */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.UnitResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a unit */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *bool */
            allow_decimal?: boolean;
            /** Format: *string */
            actual_name?: string;
            /** Format: *string */
            short_name?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.UnitResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/products/v1/units/{id}/base": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Set or clear base unit */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            base_unit_id?: string;
            /** Format: *string */
            base_unit_multiplier?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.UnitResponse"];
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
  "/products/v1/variation-templates": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List variation templates */
    get: {
      parameters: {
        query?: {
          search?: string;
          page?: number;
          perPage?: number;
          sort?: string;
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
            "application/json": components["schemas"]["qistas.internal.products.api.rest.ListVariationTemplatesResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a variation template */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.products.api.rest.CreateVariationTemplateRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.TemplateResponse"];
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
  "/products/v1/variation-templates/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a template */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.GetVariationTemplateResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Soft-delete a template */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.TemplateResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a template */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            name?: string;
            /** Format: *[]rest.CreateValueDTO */
            values?: components["schemas"]["qistas.internal.products.api.rest.CreateValueDTO"][][];
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.products.api.rest.TemplateResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/purchases/v1/locations/locations/{location_id}/purchase-requisitions": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List purchase requisitions for a location */
    get: {
      parameters: {
        query?: {
          page?: number;
          perPage?: number;
          sort?: string;
        };
        header?: never;
        path: {
          location_id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.ListPurchaseRequisitionsResponse"];
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
  "/purchases/v1/purchase-orders": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List purchase orders */
    get: {
      parameters: {
        query?: {
          status?: string;
          payment_status?: string;
          shipping_status?: string;
          supplier_id?: string;
          location_id?: string;
          branch_id?: string;
          date_from?: string;
          date_to?: string;
          ref_no?: string;
          page?: number;
          perPage?: number;
          sort?: string;
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
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.ListPurchaseOrdersResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a purchase receipt */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.purchases.api.rest.CreatePurchaseRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
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
  "/purchases/v1/purchase-orders/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a purchase order */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Soft-delete a purchase order */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.DeletePurchaseResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a purchase receipt */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *string */
            custom_field_3?: string;
            /** Format: *uuid.UUID */
            branch_id?: string;
            /** Format: *string */
            transaction_date?: string;
            /** Format: *int */
            pay_term_number?: number;
            /** Format: *uuid.UUID */
            location_id?: string;
            /** Format: *[]string */
            purchase_order_ids?: string[][];
            /** Format: *string */
            custom_field_4?: string;
            /** Format: *string */
            exchange_rate?: string;
            /** Format: *string */
            discount_type?: string;
            /** Format: *string */
            shipping_details?: string;
            /** Format: *string */
            round_off_amount?: string;
            /** Format: *[]rest.CreateLineRequest */
            lines?: components["schemas"]["qistas.internal.purchases.api.rest.CreateLineRequest"][][];
            /** Format: *string */
            invoice_no?: string;
            /** Format: *string */
            delivery_date?: string;
            /** Format: *string */
            tax_id?: string;
            /** Format: *string */
            discount_amount?: string;
            /** Format: *string */
            pay_term_type?: string;
            /** Format: *string */
            document_path?: string;
            /** Format: *string */
            custom_field_1?: string;
            /** Format: *string */
            custom_field_2?: string;
            /** Format: *uuid.UUID */
            supplier_id?: string;
            /** Format: *string */
            currency?: string;
            /** Format: *string */
            shipping_charges?: string;
            /** Format: *string */
            additional_notes?: string;
            /** Format: *string */
            staff_note?: string;
            /** Format: *[]string */
            purchase_requisition_ids?: string[][];
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/purchases/v1/purchase-orders/{id}/status": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    /** Transition a purchase order */
    put: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /**
             * Format: string
             * @enum {string}
             */
            status:
              | "draft"
              | "ordered"
              | "partial"
              | "received"
              | "pending"
              | "closed"
              | "cancelled"
              | "returned";
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
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
  "/purchases/v1/purchase-requisitions": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List purchase requisitions */
    get: {
      parameters: {
        query?: {
          status?: string;
          supplier_id?: string;
          location_id?: string;
          branch_id?: string;
          date_from?: string;
          date_to?: string;
          ref_no?: string;
          page?: number;
          perPage?: number;
          sort?: string;
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
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.ListPurchaseRequisitionsResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a purchase receipt */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.purchases.api.rest.CreatePurchaseRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
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
  "/purchases/v1/purchase-requisitions/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a purchase requisition */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Soft-delete a purchase requisition */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.DeletePurchaseResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a purchase receipt */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *[]string */
            purchase_order_ids?: string[][];
            /** Format: *string */
            tax_id?: string;
            /** Format: *[]string */
            purchase_requisition_ids?: string[][];
            /** Format: *string */
            custom_field_1?: string;
            /** Format: *string */
            custom_field_3?: string;
            /** Format: *string */
            custom_field_4?: string;
            /** Format: *string */
            discount_amount?: string;
            /** Format: *string */
            shipping_details?: string;
            /** Format: *string */
            additional_notes?: string;
            /** Format: *string */
            staff_note?: string;
            /** Format: *int */
            pay_term_number?: number;
            /** Format: *string */
            pay_term_type?: string;
            /** Format: *uuid.UUID */
            location_id?: string;
            /** Format: *string */
            custom_field_2?: string;
            /** Format: *uuid.UUID */
            supplier_id?: string;
            /** Format: *string */
            transaction_date?: string;
            /** Format: *string */
            delivery_date?: string;
            /** Format: *string */
            exchange_rate?: string;
            /** Format: *[]rest.CreateLineRequest */
            lines?: components["schemas"]["qistas.internal.purchases.api.rest.CreateLineRequest"][][];
            /** Format: *uuid.UUID */
            branch_id?: string;
            /** Format: *string */
            invoice_no?: string;
            /** Format: *string */
            currency?: string;
            /** Format: *string */
            discount_type?: string;
            /** Format: *string */
            shipping_charges?: string;
            /** Format: *string */
            round_off_amount?: string;
            /** Format: *string */
            document_path?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/purchases/v1/purchase-requisitions/{id}/lines": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List the lines of a purchase requisition */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.ListPurchaseRequisitionLinesResponse"];
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
  "/purchases/v1/purchase-returns": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List purchase returns */
    get: {
      parameters: {
        query?: {
          status?: string;
          payment_status?: string;
          supplier_id?: string;
          location_id?: string;
          branch_id?: string;
          date_from?: string;
          date_to?: string;
          ref_no?: string;
          page?: number;
          perPage?: number;
          sort?: string;
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
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.ListPurchaseReturnsResponse"];
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
  "/purchases/v1/purchase-returns/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a purchase return */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Soft-delete a purchase return */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.DeletePurchaseResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a purchase receipt */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *[]string */
            purchase_order_ids?: string[][];
            /** Format: *string */
            custom_field_2?: string;
            /** Format: *string */
            transaction_date?: string;
            /** Format: *string */
            discount_type?: string;
            /** Format: *string */
            discount_amount?: string;
            /** Format: *string */
            shipping_charges?: string;
            /** Format: *string */
            staff_note?: string;
            /** Format: *string */
            pay_term_type?: string;
            /** Format: *string */
            custom_field_3?: string;
            /** Format: *string */
            custom_field_4?: string;
            /** Format: *string */
            tax_id?: string;
            /** Format: *string */
            shipping_details?: string;
            /** Format: *string */
            additional_notes?: string;
            /** Format: *string */
            round_off_amount?: string;
            /** Format: *[]string */
            purchase_requisition_ids?: string[][];
            /** Format: *string */
            invoice_no?: string;
            /** Format: *string */
            currency?: string;
            /** Format: *int */
            pay_term_number?: number;
            /** Format: *string */
            custom_field_1?: string;
            /** Format: *[]rest.CreateLineRequest */
            lines?: components["schemas"]["qistas.internal.purchases.api.rest.CreateLineRequest"][][];
            /** Format: *uuid.UUID */
            branch_id?: string;
            /** Format: *uuid.UUID */
            supplier_id?: string;
            /** Format: *string */
            delivery_date?: string;
            /** Format: *string */
            exchange_rate?: string;
            /** Format: *uuid.UUID */
            location_id?: string;
            /** Format: *string */
            document_path?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/purchases/v1/purchases": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List purchase receipts */
    get: {
      parameters: {
        query?: {
          status?: string;
          payment_status?: string;
          shipping_status?: string;
          supplier_id?: string;
          location_id?: string;
          branch_id?: string;
          date_from?: string;
          date_to?: string;
          ref_no?: string;
          page?: number;
          perPage?: number;
          sort?: string;
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
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.ListPurchasesResponse"];
          };
        };
      };
    };
    put?: never;
    /** Create a purchase receipt */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path?: never;
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": components["schemas"]["qistas.internal.purchases.api.rest.CreatePurchaseRequest"];
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
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
  "/purchases/v1/purchases/purchases/{id}/returns": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Start a purchase return from an existing purchase */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: string */
            transaction_date: string;
            /** Format: *string */
            exchange_rate?: string;
            /** Format: *string */
            discount_type?: string;
            /** Format: *string */
            custom_field_3?: string;
            /** Format: *string */
            delivery_date?: string;
            /** Format: *string */
            discount_amount?: string;
            /** Format: *int */
            pay_term_number?: number;
            /** Format: *string */
            pay_term_type?: string;
            /** Format: *string */
            document_path?: string;
            /** Format: *string */
            custom_field_1?: string;
            /** Format: []rest.CreateLineRequest */
            lines: components["schemas"]["qistas.internal.purchases.api.rest.CreateLineRequest"][];
            /** Format: *string */
            shipping_details?: string;
            /** Format: *string */
            shipping_charges?: string;
            /** Format: *string */
            additional_notes?: string;
            /** Format: *string */
            custom_field_2?: string;
            /** Format: *string */
            invoice_no?: string;
            /** Format: string */
            currency?: string;
            /** Format: *string */
            tax_id?: string;
            /** Format: *string */
            staff_note?: string;
            /** Format: *string */
            round_off_amount?: string;
            /** Format: *string */
            custom_field_4?: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
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
  "/purchases/v1/purchases/ref-no/exists": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Check if a ref_no is taken */
    get: {
      parameters: {
        query: {
          ref_no: string;
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
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.RefNoExistsResponse"];
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
  "/purchases/v1/purchases/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get a purchase receipt */
    get: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
          };
        };
      };
    };
    put?: never;
    post?: never;
    /** Soft-delete a purchase receipt */
    delete: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.DeletePurchaseResponse"];
          };
        };
      };
    };
    options?: never;
    head?: never;
    /** Update a purchase receipt */
    patch: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: *uuid.UUID */
            branch_id?: string;
            /** Format: *uuid.UUID */
            supplier_id?: string;
            /** Format: *string */
            invoice_no?: string;
            /** Format: *string */
            transaction_date?: string;
            /** Format: *string */
            currency?: string;
            /** Format: *string */
            exchange_rate?: string;
            /** Format: *string */
            tax_id?: string;
            /** Format: *string */
            shipping_details?: string;
            /** Format: *string */
            shipping_charges?: string;
            /** Format: *string */
            additional_notes?: string;
            /** Format: *int */
            pay_term_number?: number;
            /** Format: *uuid.UUID */
            location_id?: string;
            /** Format: *[]string */
            purchase_requisition_ids?: string[][];
            /** Format: *string */
            custom_field_1?: string;
            /** Format: *string */
            custom_field_3?: string;
            /** Format: *string */
            custom_field_4?: string;
            /** Format: *string */
            discount_type?: string;
            /** Format: *string */
            pay_term_type?: string;
            /** Format: *string */
            document_path?: string;
            /** Format: *[]string */
            purchase_order_ids?: string[][];
            /** Format: *string */
            delivery_date?: string;
            /** Format: *string */
            discount_amount?: string;
            /** Format: *string */
            staff_note?: string;
            /** Format: *string */
            round_off_amount?: string;
            /** Format: *string */
            custom_field_2?: string;
            /** Format: *[]rest.CreateLineRequest */
            lines?: components["schemas"]["qistas.internal.purchases.api.rest.CreateLineRequest"][][];
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
          };
        };
      };
    };
    trace?: never;
  };
  "/purchases/v1/purchases/{id}/payments": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Record a payment against a purchase */
    post: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /** Format: string */
            amount: string;
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
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
  "/purchases/v1/purchases/{id}/status": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    /** Transition a purchase document */
    put: {
      parameters: {
        query?: never;
        header?: never;
        path: {
          id: string;
        };
        cookie?: never;
      };
      requestBody?: {
        content: {
          "application/json": {
            /**
             * Format: string
             * @enum {string}
             */
            status:
              | "draft"
              | "ordered"
              | "partial"
              | "received"
              | "pending"
              | "closed"
              | "cancelled"
              | "returned";
          };
        };
      };
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"];
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
  "/purchases/v1/suppliers/suppliers/{supplier_id}/purchase-orders": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** List purchase orders for a supplier */
    get: {
      parameters: {
        query?: {
          page?: number;
          perPage?: number;
          sort?: string;
        };
        header?: never;
        path: {
          supplier_id: string;
        };
        cookie?: never;
      };
      requestBody?: never;
      responses: {
        200: {
          headers: {
            [name: string]: unknown;
          };
          content: {
            "application/json": components["schemas"]["qistas.internal.purchases.api.rest.ListPurchaseOrdersResponse"];
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
  "/subscription": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get the tenant's current subscription */
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
            "application/json": components["schemas"]["qistas.internal.subscription.api.rest.GetMySubscriptionResponse"];
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
  "/tenant/resolve": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Resolve the current tenant */
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
            "application/json": components["schemas"]["qistas.internal.tenants.api.rest.ResolveResponse"];
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
}
export type webhooks = Record<string, never>;
export interface components {
  schemas: {
    "qistas.internal.tenants.api.rest.ArchiveBranchRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.tenants.api.rest.BranchResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      tenant_id?: string;
      /** Format: *string */
      code?: string;
      /** Format: string */
      name?: string;
      /** Format: string */
      country?: string;
      /** Format: *string */
      state?: string;
      /** Format: *string */
      city?: string;
      /** Format: *string */
      zip_code?: string;
      /** Format: *string */
      landmark?: string;
      /** Format: *string */
      mobile?: string;
      /** Format: *string */
      alternate_number?: string;
      /** Format: *string */
      email?: string;
      /** Format: *string */
      website?: string;
      /** Format: bool */
      is_active?: boolean;
      /** Format: string */
      created_at?: string;
      /** Format: string */
      updated_at?: string;
    };
    "qistas.internal.tenants.api.rest.UpdateBranchRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      name?: string;
      /** Format: *string */
      state?: string;
      /** Format: *string */
      city?: string;
      /** Format: *string */
      zip_code?: string;
      /** Format: *string */
      landmark?: string;
      /** Format: *string */
      mobile?: string;
      /** Format: *string */
      alternate_number?: string;
      /** Format: *string */
      email?: string;
      /** Format: *string */
      website?: string;
    };
    "qistas.internal.subscription.api.rest.ListFeaturesRequest": Record<
      string,
      never
    >;
    "qistas.internal.subscription.api.rest.ListFeaturesResponse": {
      /** Format: []rest.FeatureResponse */
      features?: components["schemas"]["qistas.internal.subscription.api.rest.FeatureResponse"][];
    };
    "qistas.internal.subscription.api.rest.FeatureResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      key?: string;
      /** Format: string */
      name?: string;
      /** Format: string */
      kind?: string;
    };
    "qistas.internal.subscription.api.rest.CreateFeatureRequest": {
      /** Format: string */
      key: string;
      /** Format: string */
      name: string;
      /**
       * Format: string
       * @enum {string}
       */
      kind?: "toggle" | "limit";
    };
    "qistas.internal.subscription.api.rest.ListPlansRequest": {
      /** Format: bool */
      include_archived?: boolean;
    };
    "qistas.internal.subscription.api.rest.ListPlansResponse": {
      /** Format: []rest.PlanResponse */
      plans?: components["schemas"]["qistas.internal.subscription.api.rest.PlanResponse"][];
    };
    "qistas.internal.subscription.api.rest.PlanResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      code?: string;
      /** Format: string */
      name?: string;
      /** Format: string */
      status?: string;
      /** Format: string */
      interval?: string;
      /** Format: int */
      trial_days?: number;
      /** Format: *int64 */
      price_cents?: number;
      /** Format: *string */
      currency?: string;
      /** Format: bool */
      is_default?: boolean;
      /** Format: []rest.PlanFeatureDTO */
      features?: components["schemas"]["qistas.internal.subscription.api.rest.PlanFeatureDTO"][];
    };
    "qistas.internal.subscription.api.rest.PlanFeatureDTO": {
      /** Format: string */
      feature_key: string;
      /** Format: string */
      value?: string;
    };
    "qistas.internal.subscription.api.rest.SavePlanRequest": {
      /** Format: string */
      id?: string;
      /** Format: string */
      code: string;
      /** Format: string */
      name: string;
      /**
       * Format: string
       * @enum {string}
       */
      status?: "draft" | "active" | "archived";
      /**
       * Format: string
       * @enum {string}
       */
      interval?: "monthly" | "yearly";
      /** Format: int */
      trial_days?: number;
      /** Format: *int64 */
      price_cents?: number;
      /** Format: *string */
      currency?: string;
      /** Format: bool */
      is_default?: boolean;
      /** Format: []rest.PlanFeatureDTO */
      features?: components["schemas"]["qistas.internal.subscription.api.rest.PlanFeatureDTO"][];
    };
    "qistas.internal.subscription.api.rest.ArchivePlanRequest": Record<
      string,
      never
    >;
    "qistas.internal.subscription.api.rest.ArchivePlanResponse": {
      /** Format: bool */
      ack?: boolean;
    };
    "qistas.internal.subscription.api.rest.CancelSubscriptionRequest": {
      /** Format: bool */
      now?: boolean;
    };
    "qistas.internal.subscription.api.rest.SubscriptionResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      tenant_id?: string;
      /** Format: string */
      plan_id?: string;
      /** Format: string */
      status?: string;
      /** Format: *string */
      period_start?: string;
      /** Format: *string */
      period_end?: string;
      /** Format: *string */
      trial_ends_at?: string;
      /** Format: bool */
      cancel_at_period_end?: boolean;
    };
    "qistas.internal.subscription.api.rest.AssignSubscriptionRequest": {
      /** Format: string */
      plan_id: string;
      /** Format: bool */
      force?: boolean;
    };
    "qistas.internal.subscription.api.rest.ChangeSubscriptionRequest": {
      /** Format: string */
      plan_id: string;
    };
    "qistas.internal.tenants.api.rest.ListTenantsRequest": {
      /** Format: int */
      offset?: number;
      /** Format: int */
      limit?: number;
    };
    "qistas.internal.tenants.api.rest.ListTenantsResponse": {
      /** Format: []rest.TenantResponse */
      tenants?: components["schemas"]["qistas.internal.tenants.api.rest.TenantResponse"][];
    };
    "qistas.internal.tenants.api.rest.TenantResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      name?: string;
      /** Format: string */
      slug?: string;
      /** Format: string */
      status?: string;
      /** Format: bool */
      is_active?: boolean;
      /** Format: *string */
      legal_name?: string;
      /** Format: *string */
      tax_number_1?: string;
      /** Format: *string */
      tax_label_1?: string;
      /** Format: *string */
      tax_number_2?: string;
      /** Format: *string */
      tax_label_2?: string;
      /** Format: *string */
      code_label_1?: string;
      /** Format: *string */
      code_1?: string;
      /** Format: *string */
      code_label_2?: string;
      /** Format: *string */
      code_2?: string;
      /** Format: *string */
      country?: string;
      /** Format: *string */
      currency?: string;
      /** Format: string */
      timezone?: string;
      /** Format: string */
      date_format?: string;
      /** Format: string */
      time_format?: string;
      /** Format: *string */
      start_date?: string;
      /** Format: int */
      fy_start_month?: number;
      /** Format: string */
      accounting_method?: string;
      /** Format: string */
      sell_price_tax?: string;
      /** Format: float64 */
      default_profit_percent?: number;
      /** Format: *string */
      currency_id?: string;
      /** Format: *string */
      owner_id?: string;
      /** Format: *string */
      logo?: string;
      /** Format: string */
      created_at?: string;
      /** Format: string */
      updated_at?: string;
    };
    "qistas.internal.tenants.api.rest.CreateTenantRequest": {
      /** Format: string */
      name: string;
      /** Format: string */
      slug: string;
      /** Format: *string */
      legal_name?: string;
      /** Format: *string */
      tax_number_1?: string;
      /** Format: *string */
      tax_label_1?: string;
      /** Format: *string */
      tax_number_2?: string;
      /** Format: *string */
      tax_label_2?: string;
      /** Format: *string */
      code_label_1?: string;
      /** Format: *string */
      code_1?: string;
      /** Format: *string */
      code_label_2?: string;
      /** Format: *string */
      code_2?: string;
      /** Format: *string */
      country?: string;
      /** Format: *string */
      currency?: string;
      /** Format: string */
      timezone?: string;
      /** Format: string */
      date_format?: string;
      /**
       * Format: string
       * @enum {string}
       */
      time_format?: 12 | 24;
      /** Format: int */
      fy_start_month?: number;
      /**
       * Format: string
       * @enum {string}
       */
      accounting_method?: "fifo" | "lifo" | "avco";
      /**
       * Format: string
       * @enum {string}
       */
      sell_price_tax?: "includes" | "excludes";
      /** Format: float64 */
      default_profit_percent?: number;
      /** Format: *string */
      logo?: string;
    };
    "qistas.internal.tenants.api.rest.ArchiveTenantRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.tenants.api.rest.GetTenantRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.tenants.api.rest.GetTenantResponse": {
      tenant?: components["schemas"]["qistas.internal.tenants.api.rest.TenantResponse"];
    };
    "qistas.internal.tenants.api.rest.UpdateTenantRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      name?: string;
      /** Format: *string */
      legal_name?: string;
      /** Format: *string */
      tax_number_1?: string;
      /** Format: *string */
      tax_label_1?: string;
      /** Format: *string */
      tax_number_2?: string;
      /** Format: *string */
      tax_label_2?: string;
      /** Format: *string */
      code_label_1?: string;
      /** Format: *string */
      code_1?: string;
      /** Format: *string */
      code_label_2?: string;
      /** Format: *string */
      code_2?: string;
      /** Format: *string */
      country?: string;
      /** Format: *string */
      currency?: string;
      /** Format: *string */
      timezone?: string;
      /** Format: *string */
      date_format?: string;
      /**
       * Format: *string
       * @enum {string}
       */
      time_format?: 12 | 24;
      /** Format: *string */
      logo?: string;
    };
    "qistas.internal.tenants.api.rest.ActivateTenantRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.tenants.api.rest.ListBranchesRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.tenants.api.rest.ListBranchesResponse": {
      /** Format: []rest.BranchResponse */
      branches?: components["schemas"]["qistas.internal.tenants.api.rest.BranchResponse"][];
    };
    "qistas.internal.tenants.api.rest.CreateBranchRequest": {
      /** Format: string */
      id: string;
      /** Format: string */
      name: string;
      /** Format: *string */
      code?: string;
      /** Format: string */
      country: string;
      /** Format: *string */
      state?: string;
      /** Format: *string */
      city?: string;
      /** Format: *string */
      zip_code?: string;
      /** Format: *string */
      landmark?: string;
      /** Format: *string */
      mobile?: string;
      /** Format: *string */
      alternate_number?: string;
      /** Format: *string */
      email?: string;
      /** Format: *string */
      website?: string;
    };
    "qistas.internal.tenants.api.rest.SuspendTenantRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.auth.api.rest.LogoutRequest": Record<string, never>;
    "qistas.internal.auth.api.rest.LogoutResponse": {
      /** Format: string */
      message?: string;
    };
    "qistas.internal.auth.api.rest.MeRequest": Record<string, never>;
    "qistas.internal.auth.api.rest.MeResponse.qistas.internal.auth.domain.users.User.": {
      user?: components["schemas"]["qistas.internal.auth.domain.users.User"];
      Data?: components["schemas"]["qistas.internal.auth.domain.users.User"];
    };
    "qistas.internal.auth.domain.users.User": {
      /** Format: uuid.UUID */
      id?: string;
      /** Format: string */
      name?: string;
      /** Format: string */
      email?: string;
      /** Format: string */
      phone?: string;
      /** Format: string */
      username?: string;
      /** Format: bool */
      email_verified?: boolean;
      /** Format: bool */
      phone_verified?: boolean;
      /** Format: string */
      image?: string;
      /** Format: []string */
      roles?: string[];
      /** Format: []string */
      permissions?: string[];
      /** Format: string */
      status?: string;
      /** Format: bool */
      mfa_enabled?: boolean;
      /** Format: time.Time */
      created_at?: string;
      /** Format: time.Time */
      updated_at?: string;
    };
    "qistas.lib.common.BadRequestResponse": {
      Code?: components["schemas"]["interface"];
      /** Format: string */
      Error?: string;
      Data?: components["schemas"]["interface"];
    };
    interface: Record<string, never>;
    "qistas.lib.common.InternalErrorResponse": {
      Code?: components["schemas"]["interface"];
      /** Format: string */
      Error?: string;
      Data?: components["schemas"]["interface"];
    };
    /** @description User is not authenticated */
    "qistas.lib.common.UnAuthenticatedResponse": {
      Code?: components["schemas"]["interface"];
      /** Format: string */
      Error?: string;
      Data?: components["schemas"]["interface"];
    };
    "qistas.lib.common.InvalidInputErrorResponse": {
      Code?: components["schemas"]["interface"];
      /** Format: string */
      Error?: string;
      Data?: components["schemas"]["interface"];
      Errors?: components["schemas"]["interface"];
    };
    "qistas.internal.auth.api.rest.ListPoliciesReq": Record<string, never>;
    "qistas.internal.auth.api.rest.ListPoliciesRes": components["schemas"]["qistas.internal.auth.domain.access.Policy"][];
    "qistas.internal.auth.domain.access.Policy": {
      /** Format: int32 */
      ID?: number;
      /** Format: string */
      Name?: string;
      /** Format: string */
      Slug?: string;
      /** Format: string */
      Description?: string;
      /** Format: string */
      Resource?: string;
      /** Format: string */
      Action?: string;
      /** Format: string */
      Rules?: string;
      /** Format: access.Effect */
      Effect?: string;
      /** Format: int32 */
      Priority?: number;
      /** Format: bool */
      IsActive?: boolean;
      /** Format: *time.Time */
      ExpiresAt?: string;
      /** Format: *uuid.UUID */
      CreatedBy?: string;
      /** Format: time.Time */
      CreatedAt?: string;
      /** Format: time.Time */
      UpdatedAt?: string;
    };
    "qistas.internal.auth.api.rest.OnePoliciesReq": Record<string, never>;
    "qistas.internal.auth.api.rest.RefreshRequest": Record<string, never>;
    "qistas.internal.auth.api.rest.RefreshResponse": {
      /** Format: string */
      token?: string;
      /** Format: string */
      token_type?: string;
      /** Format: int64 */
      expires_at?: number;
    };
    "qistas.internal.auth.api.rest.SignInRequest": {
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
    "qistas.lib.http.Response.qistas.internal.auth.api.rest.SignInResponse.": {
      /** @description The returned data if the request is OK */
      data?: components["schemas"]["qistas.internal.auth.api.rest.SignInResponse"];
      /** Format: string */
      message?: string;
      error?: components["schemas"]["qistas.lib.http.ErrorResponse"];
    };
    "qistas.internal.auth.api.rest.SignInResponse": {
      /**
       * Format: string
       * @description Bearer access token
       */
      token?: string;
      /**
       * Format: string
       * @description Always "Bearer"
       */
      token_type?: string;
      /**
       * Format: int64
       * @description Unix expiry of the access token
       */
      expires_at?: number;
      /** @description Authenticated user */
      user?: components["schemas"]["qistas.internal.auth.domain.users.User"];
    };
    "qistas.lib.http.ErrorResponse": {
      /**
       * Format: string
       * @description the error code
       */
      code?: string;
      /** Format: string */
      message?: string;
    };
    "qistas.internal.contacts.api.rest.ListContactsRequest": {
      /** Format: string */
      type?: string;
      /** Format: string */
      type_mode?: string;
      /** Format: string */
      status?: string;
      /** Format: string */
      customer_group_id?: string;
      /** Format: string */
      assigned_to?: string;
      /** Format: string */
      q?: string;
      /** Format: int */
      offset?: number;
      /** Format: int */
      limit?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.contacts.api.rest.ListContactsResponse": {
      /** Format: []rest.ContactResponse */
      contacts?: components["schemas"]["qistas.internal.contacts.api.rest.ContactResponse"][];
      /** Format: int */
      total?: number;
    };
    "qistas.internal.contacts.api.rest.ContactResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      type?: string;
      /** Format: *string */
      contact_type?: string;
      /** Format: *string */
      supplier_tenant_name?: string;
      /** Format: *string */
      prefix?: string;
      /** Format: *string */
      first_name?: string;
      /** Format: *string */
      middle_name?: string;
      /** Format: *string */
      last_name?: string;
      /** Format: string */
      name?: string;
      /** Format: *string */
      email?: string;
      /** Format: *string */
      contact_id?: string;
      /** Format: string */
      contact_status?: string;
      /** Format: *string */
      tax_number?: string;
      /** Format: *string */
      address_line_1?: string;
      /** Format: *string */
      address_line_2?: string;
      /** Format: *string */
      city?: string;
      /** Format: *string */
      state?: string;
      /** Format: *string */
      country?: string;
      /** Format: *string */
      zip_code?: string;
      /** Format: *string */
      dob?: string;
      /** Format: string */
      mobile?: string;
      /** Format: *string */
      landline?: string;
      /** Format: *string */
      alternate_number?: string;
      /** Format: *int */
      pay_term_number?: number;
      /** Format: *string */
      pay_term_type?: string;
      /** Format: *string */
      credit_limit?: string;
      /** Format: string */
      opening_balance?: string;
      /** Format: string */
      balance?: string;
      /** Format: bool */
      is_default?: boolean;
      /** Format: *string */
      shipping_address?: string;
      /** Format: *string */
      shipping_custom_field_details?: string;
      /** Format: bool */
      is_export?: boolean;
      /** Format: *string */
      export_custom_field_1?: string;
      /** Format: *string */
      export_custom_field_2?: string;
      /** Format: *string */
      export_custom_field_3?: string;
      /** Format: *string */
      export_custom_field_4?: string;
      /** Format: *string */
      export_custom_field_5?: string;
      /** Format: *string */
      export_custom_field_6?: string;
      /** Format: *string */
      position?: string;
      /** Format: *string */
      customer_group_id?: string;
      /** Format: *string */
      crm_source?: string;
      /** Format: *string */
      crm_life_stage?: string;
      /** Format: *string */
      custom_field_1?: string;
      /** Format: *string */
      custom_field_2?: string;
      /** Format: *string */
      custom_field_3?: string;
      /** Format: *string */
      custom_field_4?: string;
      /** Format: *string */
      custom_field_5?: string;
      /** Format: *string */
      custom_field_6?: string;
      /** Format: *string */
      custom_field_7?: string;
      /** Format: *string */
      custom_field_8?: string;
      /** Format: *string */
      custom_field_9?: string;
      /** Format: *string */
      custom_field_10?: string;
      /** Format: string */
      created_at?: string;
      /** Format: string */
      updated_at?: string;
    };
    "qistas.internal.contacts.api.rest.CreateContactRequest": {
      /**
       * Format: string
       * @enum {string}
       */
      type: "customer" | "supplier" | "both";
      /** Format: *string */
      contact_type?: string;
      /** Format: *string */
      supplier_tenant_name?: string;
      /** Format: *string */
      prefix?: string;
      /** Format: *string */
      first_name?: string;
      /** Format: *string */
      middle_name?: string;
      /** Format: *string */
      last_name?: string;
      /** Format: *string */
      name?: string;
      /** Format: *string */
      email?: string;
      /** Format: *string */
      contact_id?: string;
      /** Format: *string */
      tax_number?: string;
      /** Format: string */
      mobile: string;
      /** Format: *string */
      landline?: string;
      /** Format: *string */
      alternate_number?: string;
      /** Format: *string */
      address_line_1?: string;
      /** Format: *string */
      address_line_2?: string;
      /** Format: *string */
      city?: string;
      /** Format: *string */
      state?: string;
      /** Format: *string */
      country?: string;
      /** Format: *string */
      zip_code?: string;
      /** Format: *string */
      dob?: string;
      /** Format: *int */
      pay_term_number?: number;
      /** Format: *string */
      pay_term_type?: string;
      /** Format: *string */
      credit_limit?: string;
      /** Format: *string */
      opening_balance?: string;
      /** Format: *bool */
      is_default?: boolean;
      /** Format: *string */
      shipping_address?: string;
      /** Format: *string */
      shipping_custom_field_details?: string;
      /** Format: *bool */
      is_export?: boolean;
      /** Format: *string */
      export_custom_field_1?: string;
      /** Format: *string */
      export_custom_field_2?: string;
      /** Format: *string */
      export_custom_field_3?: string;
      /** Format: *string */
      export_custom_field_4?: string;
      /** Format: *string */
      export_custom_field_5?: string;
      /** Format: *string */
      export_custom_field_6?: string;
      /** Format: *string */
      position?: string;
      /** Format: *string */
      customer_group_id?: string;
      /** Format: []string */
      assigned_to_users?: string[];
      /** Format: *string */
      custom_field_1?: string;
      /** Format: *string */
      custom_field_2?: string;
      /** Format: *string */
      custom_field_3?: string;
      /** Format: *string */
      custom_field_4?: string;
      /** Format: *string */
      custom_field_5?: string;
      /** Format: *string */
      custom_field_6?: string;
      /** Format: *string */
      custom_field_7?: string;
      /** Format: *string */
      custom_field_8?: string;
      /** Format: *string */
      custom_field_9?: string;
      /** Format: *string */
      custom_field_10?: string;
      /** Format: *string */
      crm_source?: string;
      /** Format: *string */
      crm_life_stage?: string;
    };
    "qistas.internal.contacts.api.rest.CheckContactIDRequest": {
      /** Format: string */
      code: string;
      /** Format: string */
      exclude_id?: string;
    };
    "qistas.internal.contacts.api.rest.CheckContactIDResponse": {
      /** Format: bool */
      exists?: boolean;
    };
    "qistas.internal.contacts.api.rest.CheckMobileRequest": {
      /** Format: string */
      mobile: string;
      /** Format: string */
      exclude_id?: string;
    };
    "qistas.internal.contacts.api.rest.CheckMobileResponse": {
      /** Format: bool */
      exists?: boolean;
      /** Format: []string */
      names?: string[];
    };
    "qistas.internal.contacts.api.rest.DeleteContactRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.lib.common.EmptyResponse": Record<string, never>;
    "qistas.internal.contacts.api.rest.GetContactRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.contacts.api.rest.UpdateContactRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      prefix?: string;
      /** Format: *string */
      first_name?: string;
      /** Format: *string */
      middle_name?: string;
      /** Format: *string */
      last_name?: string;
      /** Format: *string */
      name?: string;
      /** Format: *string */
      email?: string;
      /** Format: *string */
      contact_id?: string;
      /** Format: *string */
      tax_number?: string;
      /** Format: *string */
      address_line_1?: string;
      /** Format: *string */
      address_line_2?: string;
      /** Format: *string */
      city?: string;
      /** Format: *string */
      state?: string;
      /** Format: *string */
      country?: string;
      /** Format: *string */
      zip_code?: string;
      /** Format: *string */
      dob?: string;
      /** Format: *string */
      landline?: string;
      /** Format: *string */
      alternate_number?: string;
      /** Format: *int */
      pay_term_number?: number;
      /** Format: *string */
      pay_term_type?: string;
      /** Format: *string */
      credit_limit?: string;
      /** Format: *bool */
      is_default?: boolean;
      /** Format: *string */
      shipping_address?: string;
      /** Format: *string */
      shipping_custom_field_details?: string;
      /** Format: *bool */
      is_export?: boolean;
      /** Format: *string */
      export_custom_field_1?: string;
      /** Format: *string */
      export_custom_field_2?: string;
      /** Format: *string */
      export_custom_field_3?: string;
      /** Format: *string */
      export_custom_field_4?: string;
      /** Format: *string */
      export_custom_field_5?: string;
      /** Format: *string */
      export_custom_field_6?: string;
      /** Format: *string */
      position?: string;
      /** Format: *string */
      customer_group_id?: string;
      /** Format: *string */
      custom_field_1?: string;
      /** Format: *string */
      custom_field_2?: string;
      /** Format: *string */
      custom_field_3?: string;
      /** Format: *string */
      custom_field_4?: string;
      /** Format: *string */
      custom_field_5?: string;
      /** Format: *string */
      custom_field_6?: string;
      /** Format: *string */
      custom_field_7?: string;
      /** Format: *string */
      custom_field_8?: string;
      /** Format: *string */
      custom_field_9?: string;
      /** Format: *string */
      custom_field_10?: string;
      /** Format: *string */
      crm_source?: string;
      /** Format: *string */
      crm_life_stage?: string;
    };
    "qistas.internal.contacts.api.rest.ListAccessRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.contacts.api.rest.AccessResponse": {
      /** Format: []string */
      user_ids?: string[];
    };
    "qistas.internal.contacts.api.rest.AccessRequest": {
      /** Format: string */
      id: string;
      /** Format: []string */
      user_ids?: string[];
    };
    "qistas.internal.contacts.api.rest.ChangeStatusRequest": {
      /** Format: string */
      id: string;
      /**
       * Format: string
       * @enum {string}
       */
      to: "active" | "inactive";
    };
    "qistas.internal.contacts.api.rest.ListGroupsRequest": {
      /** Format: int */
      offset?: number;
      /** Format: int */
      limit?: number;
    };
    "qistas.internal.contacts.api.rest.ListGroupsResponse": {
      /** Format: []rest.GroupResponse */
      groups?: components["schemas"]["qistas.internal.contacts.api.rest.GroupResponse"][];
      /** Format: int */
      total?: number;
    };
    "qistas.internal.contacts.api.rest.GroupResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      name?: string;
      /** Format: string */
      amount?: string;
      /** Format: string */
      price_calculation_type?: string;
      /** Format: *string */
      selling_price_group_id?: string;
      /** Format: string */
      created_at?: string;
      /** Format: string */
      updated_at?: string;
    };
    "qistas.internal.contacts.api.rest.CreateGroupRequest": {
      /** Format: string */
      name: string;
      /** Format: *string */
      amount?: string;
      /** Format: string */
      price_calculation_type?: string;
      /** Format: *string */
      selling_price_group_id?: string;
    };
    "qistas.internal.contacts.api.rest.DeleteGroupRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.contacts.api.rest.GetGroupRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.contacts.api.rest.UpdateGroupRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      name?: string;
      /** Format: *string */
      amount?: string;
      /** Format: *string */
      price_calculation_type?: string;
      /** Format: *string */
      selling_price_group_id?: string;
    };
    "qistas.internal.tenants.api.rest.CreateMyTenantRequest": {
      /** Format: string */
      name: string;
      /** Format: string */
      slug: string;
      /** Format: *string */
      legal_name?: string;
      /** Format: *string */
      country?: string;
      /** Format: *string */
      currency?: string;
      /** Format: string */
      timezone?: string;
      /** Format: string */
      date_format?: string;
      /** Format: string */
      time_format?: string;
    };
    "qistas.internal.tenants.api.rest.CreateMyTenantResponse": {
      tenant?: components["schemas"]["qistas.internal.tenants.api.rest.TenantResponse"];
    };
    "qistas.internal.products.api.rest.ListBrandsRequest": {
      /** Format: string */
      search?: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.products.api.rest.ListBrandsResponse": {
      /** Format: []rest.BrandResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.BrandResponse"][];
      Page?: components["schemas"]["qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.BrandResponse."];
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.BrandResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      name?: string;
      /** Format: *string */
      description?: string;
      /** Format: bool */
      use_for_repair?: boolean;
      /** Format: string */
      created_at?: string;
      /** Format: string */
      updated_at?: string;
    };
    "qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.BrandResponse.": {
      /** Format: []rest.BrandResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.BrandResponse"][];
      /** Format: int */
      page?: number;
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.CreateBrandRequest": {
      /** Format: string */
      name: string;
      /** Format: *string */
      description?: string;
      /** Format: bool */
      use_for_repair?: boolean;
    };
    "qistas.internal.products.api.rest.DeleteBrandRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.GetBrandRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.GetBrandResponse": {
      brand?: components["schemas"]["qistas.internal.products.api.rest.BrandResponse"];
    };
    "qistas.internal.products.api.rest.UpdateBrandRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      name?: string;
      /** Format: *string */
      description?: string;
      /** Format: *bool */
      use_for_repair?: boolean;
    };
    "qistas.internal.products.api.rest.ListCategoriesRequest": {
      /** Format: string */
      category_type?: string;
      /** Format: *string */
      parent_id?: string;
      /** Format: string */
      search?: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.products.api.rest.ListCategoriesResponse": {
      /** Format: []rest.CategoryResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.CategoryResponse"][];
      Page?: components["schemas"]["qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.CategoryResponse."];
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.CategoryResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      name?: string;
      /** Format: *string */
      short_code?: string;
      /** Format: *string */
      slug?: string;
      /** Format: *string */
      parent_id?: string;
      /** Format: string */
      category_type?: string;
      /** Format: *string */
      description?: string;
      /** Format: string */
      created_at?: string;
      /** Format: string */
      updated_at?: string;
    };
    "qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.CategoryResponse.": {
      /** Format: []rest.CategoryResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.CategoryResponse"][];
      /** Format: int */
      page?: number;
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.CreateCategoryRequest": {
      /** Format: string */
      name: string;
      /** Format: *string */
      short_code?: string;
      /** Format: *string */
      slug?: string;
      /** Format: *string */
      parent_id?: string;
      /**
       * Format: string
       * @enum {string}
       */
      category_type: "product" | "expense" | "both";
      /** Format: *string */
      description?: string;
    };
    "qistas.internal.products.api.rest.SubCategoriesRequest": {
      /** Format: string */
      parent_id: string;
      /**
       * Format: string
       * @enum {string}
       */
      category_type: "product" | "expense" | "both";
    };
    "qistas.internal.products.api.rest.SubCategoriesResponse": {
      /** Format: []rest.CategoryResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.CategoryResponse"][];
    };
    "qistas.internal.products.api.rest.TaxonomiesRequest": {
      /**
       * Format: string
       * @enum {string}
       */
      category_type: "product" | "expense" | "both";
    };
    "qistas.internal.products.api.rest.TaxonomiesResponse": {
      /** Format: []services.TreeNode */
      data?: components["schemas"]["qistas.internal.products.app.services.TreeNode"][];
    };
    "qistas.internal.products.app.services.TreeNode": {
      /** Format: uuid.UUID */
      id?: string;
      /** Format: string */
      name?: string;
      /** Format: *string */
      short_code?: string;
      /** Format: *string */
      slug?: string;
      /** Format: category.CategoryType */
      category_type?: string;
      /** Format: *string */
      description?: string;
      /** Format: []services.TreeNode */
      children?: components["schemas"]["qistas.internal.products.app.services.TreeNode"][];
    };
    "qistas.internal.products.api.rest.TaxonomiesIndexRequest": {
      /**
       * Format: string
       * @enum {string}
       */
      category_type: "product" | "expense" | "both";
      /** Format: string */
      search?: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
    };
    "qistas.internal.products.api.rest.TaxonomiesIndexResponse": {
      /** Format: []rest.CategoryResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.CategoryResponse"][];
      Page?: components["schemas"]["qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.CategoryResponse."];
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.DeleteCategoryRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.GetCategoryRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.GetCategoryResponse": {
      category?: components["schemas"]["qistas.internal.products.api.rest.CategoryResponse"];
    };
    "qistas.internal.products.api.rest.UpdateCategoryRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      name?: string;
      /** Format: *string */
      short_code?: string;
      /** Format: *string */
      slug?: string;
      /** Format: *string */
      category_type?: string;
      /** Format: *string */
      description?: string;
    };
    "qistas.internal.products.api.rest.MoveCategoryRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      new_parent_id?: string;
    };
    "qistas.internal.products.api.rest.ListExpenseCategoriesRequest": {
      /** Format: string */
      search?: string;
      /** Format: *string */
      parent_id?: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.products.api.rest.ListExpenseCategoriesResponse": {
      /** Format: []rest.ExpenseCategoryResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.ExpenseCategoryResponse"][];
      Page?: components["schemas"]["qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.ExpenseCategoryResponse."];
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.ExpenseCategoryResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      name?: string;
      /** Format: *string */
      code?: string;
      /** Format: *string */
      parent_id?: string;
      /** Format: *string */
      account_id?: string;
      /** Format: string */
      created_at?: string;
      /** Format: string */
      updated_at?: string;
    };
    "qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.ExpenseCategoryResponse.": {
      /** Format: []rest.ExpenseCategoryResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.ExpenseCategoryResponse"][];
      /** Format: int */
      page?: number;
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.CreateExpenseCategoryRequest": {
      /** Format: string */
      name: string;
      /** Format: *string */
      code?: string;
      /** Format: *string */
      parent_id?: string;
      /** Format: *string */
      account_id?: string;
    };
    "qistas.internal.products.api.rest.DeleteExpenseCategoryRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.GetExpenseCategoryRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.GetExpenseCategoryResponse": {
      category?: components["schemas"]["qistas.internal.products.api.rest.ExpenseCategoryResponse"];
    };
    "qistas.internal.products.api.rest.UpdateExpenseCategoryRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      name?: string;
      /** Format: *string */
      code?: string;
      /** Format: *string */
      account_id?: string;
    };
    "qistas.internal.products.api.rest.MoveExpenseCategoryRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      new_parent_id?: string;
    };
    "qistas.internal.products.api.rest.EnqueueImportRequest": Record<
      string,
      never
    >;
    "qistas.internal.products.api.rest.EnqueueImportResponse": {
      /** Format: string */
      job_id?: string;
      /** Format: string */
      status?: string;
    };
    "qistas.internal.products.api.rest.StatusImportRequest": {
      /** Format: string */
      job_id: string;
    };
    "qistas.internal.products.api.rest.StatusImportResponse": {
      job?: components["schemas"]["qistas.internal.products.api.rest.importJob"];
    };
    "qistas.internal.products.api.rest.importJob": {
      /** Format: string */
      job_id?: string;
      /** Format: string */
      filename?: string;
      /** Format: string */
      status?: string;
      /** Format: time.Time */
      started_at?: string;
      /** Format: *time.Time */
      finished_at?: string;
      /** Format: int */
      total?: number;
      /** Format: int */
      succeeded?: number;
      /** Format: int */
      failed?: number;
      /** Format: []rest.importRowError */
      errors?: components["schemas"]["qistas.internal.products.api.rest.importRowError"][];
    };
    "qistas.internal.products.api.rest.importRowError": {
      /** Format: int */
      row?: number;
      /** Format: string */
      message?: string;
    };
    "qistas.internal.products.api.rest.ListProductsRequest": {
      /** Format: string */
      search?: string;
      /** Format: string */
      sku?: string;
      /** Format: *string */
      type?: string;
      /** Format: *string */
      brand_id?: string;
      /** Format: *string */
      category_id?: string;
      /** Format: *bool */
      is_inactive?: boolean;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.products.api.rest.ListProductsResponse": {
      /** Format: []rest.ProductResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.ProductResponse"][];
      Page?: components["schemas"]["qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.ProductResponse."];
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.ProductResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      name?: string;
      /** Format: string */
      type?: string;
      /** Format: *string */
      sku?: string;
      /** Format: *string */
      barcode_type?: string;
      /** Format: *string */
      description?: string;
      /** Format: *string */
      image_media_id?: string;
      /** Format: *string */
      brand_id?: string;
      /** Format: *string */
      category_id?: string;
      /** Format: *string */
      sub_category_id?: string;
      /** Format: *string */
      unit_id?: string;
      /** Format: *string */
      secondary_unit_id?: string;
      /** Format: []string */
      sub_unit_ids?: string[];
      /** Format: *string */
      tax_id?: string;
      /** Format: string */
      tax_type?: string;
      /** Format: bool */
      enable_stock?: boolean;
      /** Format: *string */
      alert_quantity?: string;
      /** Format: *string */
      expiry_period?: string;
      /** Format: *string */
      expiry_period_type?: string;
      /** Format: bool */
      enable_sr_no?: boolean;
      /** Format: *string */
      weight?: string;
      /** Format: *int */
      preparation_time_in_minutes?: number;
      /** Format: *string */
      warranty_id?: string;
      /** Format: bool */
      is_inactive?: boolean;
      /** Format: bool */
      not_for_selling?: boolean;
      /** Format: map[string]string */
      custom_fields?: {
        [key: string]: string;
      };
      /** Format: string */
      created_at?: string;
      /** Format: string */
      updated_at?: string;
    };
    "qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.ProductResponse.": {
      /** Format: []rest.ProductResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.ProductResponse"][];
      /** Format: int */
      page?: number;
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.CreateProductRequest": {
      /** Format: string */
      name: string;
      /**
       * Format: string
       * @enum {string}
       */
      type: "single" | "variable" | "combo";
      /** Format: *string */
      sku?: string;
      /** Format: *string */
      barcode_type?: string;
      /** Format: *string */
      description?: string;
      /** Format: *string */
      image_media_id?: string;
      /** Format: *string */
      brand_id?: string;
      /** Format: *string */
      category_id?: string;
      /** Format: *string */
      sub_category_id?: string;
      /** Format: *string */
      unit_id?: string;
      /** Format: *string */
      secondary_unit_id?: string;
      /** Format: []string */
      sub_unit_ids?: string[];
      /** Format: *string */
      tax_id?: string;
      /**
       * Format: string
       * @enum {string}
       */
      tax_type: "inclusive" | "exclusive";
      /** Format: bool */
      enable_stock?: boolean;
      /** Format: *string */
      alert_quantity?: string;
      /** Format: *string */
      expiry_period?: string;
      /** Format: *string */
      expiry_period_type?: string;
      /** Format: bool */
      enable_sr_no?: boolean;
      /** Format: *string */
      weight?: string;
      /** Format: *int */
      preparation_time_in_minutes?: number;
      /** Format: *string */
      warranty_id?: string;
      /** Format: bool */
      is_inactive?: boolean;
      /** Format: bool */
      not_for_selling?: boolean;
      /** Format: map[string]string */
      custom_fields?: {
        [key: string]: string;
      };
      /** Format: []rest.CreateProductVariationDTO */
      product_variations?: components["schemas"]["qistas.internal.products.api.rest.CreateProductVariationDTO"][];
      /** Format: []rest.CreateRackDTO */
      racks?: components["schemas"]["qistas.internal.products.api.rest.CreateRackDTO"][];
    };
    "qistas.internal.products.api.rest.CreateProductVariationDTO": {
      /** Format: *string */
      variation_template_id?: string;
      /** Format: string */
      name: string;
      /** Format: bool */
      is_dummy?: boolean;
      /** Format: []rest.CreateVariationDTO */
      values?: components["schemas"]["qistas.internal.products.api.rest.CreateVariationDTO"][];
    };
    "qistas.internal.products.api.rest.CreateVariationDTO": {
      /** Format: string */
      name: string;
      /** Format: *string */
      sub_sku?: string;
      /** Format: *string */
      variation_value_id?: string;
      /** Format: *string */
      default_purchase_price?: string;
      /** Format: *string */
      dpp_inc_tax?: string;
      /** Format: *string */
      profit_percent?: string;
      /** Format: *string */
      default_sell_price?: string;
      /** Format: *string */
      sell_price_inc_tax?: string;
      /** Format: []product.ComboLine */
      combo_variations?: components["schemas"]["qistas.internal.products.domain.product.ComboLine"][];
    };
    "qistas.internal.products.domain.product.ComboLine": {
      /** Format: uuid.UUID */
      variation_id?: string;
      quantity?: components["schemas"]["github.com.shopspring.decimal.Decimal"];
      /** Format: uuid.UUID */
      unit_id?: string;
    };
    "github.com.shopspring.decimal.Decimal": Record<string, never>;
    "qistas.internal.products.api.rest.CreateRackDTO": {
      /** Format: string */
      branch_id: string;
      /** Format: *string */
      rack?: string;
      /** Format: *string */
      row?: string;
      /** Format: *string */
      position?: string;
    };
    "qistas.internal.products.api.rest.DeleteProductRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.GetProductRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.ProductDetailResponse": {
      product?: components["schemas"]["qistas.internal.products.api.rest.ProductResponse"];
      /** Format: []rest.ProductVariationResponse */
      product_variations?: components["schemas"]["qistas.internal.products.api.rest.ProductVariationResponse"][];
      /** Format: []rest.VariationResponse */
      variations?: components["schemas"]["qistas.internal.products.api.rest.VariationResponse"][];
      /** Format: []rest.VariationGroupPriceResponse */
      group_prices?: components["schemas"]["qistas.internal.products.api.rest.VariationGroupPriceResponse"][];
      /** Format: []rest.RackResponse */
      racks?: components["schemas"]["qistas.internal.products.api.rest.RackResponse"][];
    };
    "qistas.internal.products.api.rest.ProductVariationResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      name?: string;
      /** Format: *string */
      variation_template_id?: string;
      /** Format: bool */
      is_dummy?: boolean;
    };
    "qistas.internal.products.api.rest.VariationResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      product_variation_id?: string;
      /** Format: string */
      name?: string;
      /** Format: *string */
      sub_sku?: string;
      /** Format: *string */
      variation_value_id?: string;
      /** Format: *string */
      default_purchase_price?: string;
      /** Format: string */
      dpp_inc_tax?: string;
      /** Format: string */
      profit_percent?: string;
      /** Format: *string */
      default_sell_price?: string;
      /** Format: *string */
      sell_price_inc_tax?: string;
      /** Format: []product.ComboLine */
      combo_variations?: components["schemas"]["qistas.internal.products.domain.product.ComboLine"][];
    };
    "qistas.internal.products.api.rest.VariationGroupPriceResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      variation_id?: string;
      /** Format: string */
      price_group_id?: string;
      /** Format: string */
      price_inc_tax?: string;
      /** Format: string */
      price_type?: string;
    };
    "qistas.internal.products.api.rest.RackResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      branch_id?: string;
      /** Format: *string */
      rack?: string;
      /** Format: *string */
      row?: string;
      /** Format: *string */
      position?: string;
    };
    "qistas.internal.products.api.rest.UpdateProductRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      name?: string;
      /** Format: *string */
      type?: string;
      /** Format: *string */
      sku?: string;
      /** Format: *string */
      barcode_type?: string;
      /** Format: *string */
      description?: string;
      /** Format: *string */
      image_media_id?: string;
      /** Format: *string */
      brand_id?: string;
      /** Format: *string */
      category_id?: string;
      /** Format: *string */
      sub_category_id?: string;
      /** Format: *string */
      unit_id?: string;
      /** Format: *string */
      secondary_unit_id?: string;
      /** Format: *[]string */
      sub_unit_ids?: string[][];
      /** Format: *string */
      tax_id?: string;
      /** Format: *string */
      tax_type?: string;
      /** Format: *bool */
      enable_stock?: boolean;
      /** Format: *string */
      alert_quantity?: string;
      /** Format: *string */
      expiry_period?: string;
      /** Format: *string */
      expiry_period_type?: string;
      /** Format: *bool */
      enable_sr_no?: boolean;
      /** Format: *string */
      weight?: string;
      /** Format: *int */
      preparation_time_in_minutes?: number;
      /** Format: *string */
      warranty_id?: string;
      /** Format: *bool */
      is_inactive?: boolean;
      /** Format: *bool */
      not_for_selling?: boolean;
      /** Format: *map[string]string */
      custom_fields?: {
        [key: string]: string;
      };
      /** Format: bool */
      replace_variations?: boolean;
      /** Format: []rest.UpdateProductVariationDTO */
      variations?: components["schemas"]["qistas.internal.products.api.rest.UpdateProductVariationDTO"][];
      /** Format: bool */
      replace_racks?: boolean;
      /** Format: []rest.CreateRackDTO */
      racks?: components["schemas"]["qistas.internal.products.api.rest.CreateRackDTO"][];
    };
    "qistas.internal.products.api.rest.UpdateProductVariationDTO": {
      /** Format: *string */
      variation_template_id?: string;
      /** Format: string */
      name: string;
      /** Format: bool */
      is_dummy?: boolean;
      /** Format: []rest.UpdateVariationDTO */
      values?: components["schemas"]["qistas.internal.products.api.rest.UpdateVariationDTO"][];
    };
    "qistas.internal.products.api.rest.UpdateVariationDTO": {
      /** Format: *string */
      id?: string;
      /** Format: string */
      name: string;
      /** Format: *string */
      sub_sku?: string;
      /** Format: *string */
      variation_value_id?: string;
      /** Format: *string */
      default_purchase_price?: string;
      /** Format: *string */
      dpp_inc_tax?: string;
      /** Format: *string */
      profit_percent?: string;
      /** Format: *string */
      default_sell_price?: string;
      /** Format: *string */
      sell_price_inc_tax?: string;
      /** Format: []product.ComboLine */
      combo_variations?: components["schemas"]["qistas.internal.products.domain.product.ComboLine"][];
    };
    "qistas.internal.products.api.rest.ListRacksRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.ListRacksResponse": {
      /** Format: []rest.RackResponse */
      racks?: components["schemas"]["qistas.internal.products.api.rest.RackResponse"][];
    };
    "qistas.internal.products.api.rest.ReplaceRacksRequest": {
      /** Format: string */
      id: string;
      /** Format: []rest.CreateRackDTO */
      racks?: components["schemas"]["qistas.internal.products.api.rest.CreateRackDTO"][];
    };
    "qistas.internal.products.api.rest.ListVariationsRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.ListVariationsResponse": {
      /** Format: []rest.VariationResponse */
      variations?: components["schemas"]["qistas.internal.products.api.rest.VariationResponse"][];
      /** Format: []rest.VariationGroupPriceResponse */
      group_prices?: components["schemas"]["qistas.internal.products.api.rest.VariationGroupPriceResponse"][];
    };
    "qistas.internal.products.api.rest.AddVariationRequest": {
      /** Format: string */
      id: string;
      /** Format: string */
      axis_name: string;
      variation?: components["schemas"]["qistas.internal.products.api.rest.CreateVariationDTO"];
    };
    "qistas.internal.products.api.rest.AddVariationResponse": {
      variation?: components["schemas"]["qistas.internal.products.api.rest.VariationResponse"];
    };
    "qistas.internal.products.api.rest.DeleteVariationRequest": {
      /** Format: string */
      id: string;
      /** Format: string */
      vid: string;
    };
    "qistas.internal.products.api.rest.UpdateVariationRequest": {
      /** Format: string */
      id: string;
      /** Format: string */
      vid: string;
      /** Format: *string */
      name?: string;
      /** Format: *string */
      sub_sku?: string;
      /** Format: *string */
      default_purchase_price?: string;
      /** Format: *string */
      default_sell_price?: string;
      /** Format: *string */
      sell_price_inc_tax?: string;
      /** Format: *string */
      profit_percent?: string;
    };
    "qistas.internal.products.api.rest.ListUnitsRequest": {
      /** Format: string */
      search?: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.products.api.rest.ListUnitsResponse": {
      /** Format: []rest.UnitResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.UnitResponse"][];
      Page?: components["schemas"]["qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.UnitResponse."];
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.UnitResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      actual_name?: string;
      /** Format: string */
      short_name?: string;
      /** Format: bool */
      allow_decimal?: boolean;
      /** Format: *string */
      base_unit_id?: string;
      /** Format: *string */
      base_unit_multiplier?: string;
      /** Format: string */
      created_at?: string;
      /** Format: string */
      updated_at?: string;
    };
    "qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.UnitResponse.": {
      /** Format: []rest.UnitResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.UnitResponse"][];
      /** Format: int */
      page?: number;
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.CreateUnitRequest": {
      /** Format: string */
      actual_name: string;
      /** Format: string */
      short_name: string;
      /** Format: bool */
      allow_decimal?: boolean;
      /** Format: *string */
      base_unit_id?: string;
      /** Format: *string */
      base_unit_multiplier?: string;
    };
    "qistas.internal.products.api.rest.DeleteUnitRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.GetUnitRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.GetUnitResponse": {
      unit?: components["schemas"]["qistas.internal.products.api.rest.UnitResponse"];
    };
    "qistas.internal.products.api.rest.UpdateUnitRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      actual_name?: string;
      /** Format: *string */
      short_name?: string;
      /** Format: *bool */
      allow_decimal?: boolean;
    };
    "qistas.internal.products.api.rest.SetBaseUnitRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      base_unit_id?: string;
      /** Format: *string */
      base_unit_multiplier?: string;
    };
    "qistas.internal.products.api.rest.ListVariationTemplatesRequest": {
      /** Format: string */
      search?: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.products.api.rest.ListVariationTemplatesResponse": {
      /** Format: []rest.TemplateResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.TemplateResponse"][];
      Page?: components["schemas"]["qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.TemplateResponse."];
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.TemplateResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      name?: string;
      /** Format: []rest.ValueResponse */
      values?: components["schemas"]["qistas.internal.products.api.rest.ValueResponse"][];
      /** Format: string */
      created_at?: string;
      /** Format: string */
      updated_at?: string;
    };
    "qistas.internal.products.api.rest.ValueResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      name?: string;
    };
    "qistas.internal.products.api.rest.Page.qistas.internal.products.api.rest.TemplateResponse.": {
      /** Format: []rest.TemplateResponse */
      data?: components["schemas"]["qistas.internal.products.api.rest.TemplateResponse"][];
      /** Format: int */
      page?: number;
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.products.api.rest.CreateVariationTemplateRequest": {
      /** Format: string */
      name: string;
      /** Format: []rest.CreateValueDTO */
      values?: components["schemas"]["qistas.internal.products.api.rest.CreateValueDTO"][];
    };
    "qistas.internal.products.api.rest.CreateValueDTO": {
      /** Format: string */
      name: string;
    };
    "qistas.internal.products.api.rest.DeleteVariationTemplateRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.GetVariationTemplateRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.products.api.rest.GetVariationTemplateResponse": {
      template?: components["schemas"]["qistas.internal.products.api.rest.TemplateResponse"];
    };
    "qistas.internal.products.api.rest.UpdateVariationTemplateRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      name?: string;
      /** Format: *[]rest.CreateValueDTO */
      values?: components["schemas"]["qistas.internal.products.api.rest.CreateValueDTO"][][];
    };
    "qistas.internal.purchases.api.rest.ListPurchaseRequisitionsForLocationRequest": {
      /** Format: string */
      location_id: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.purchases.api.rest.ListPurchaseRequisitionsResponse": {
      /** Format: []rest.PurchaseResponse */
      data?: components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"][];
      Page?: components["schemas"]["qistas.internal.purchases.api.rest.Page.qistas.internal.purchases.api.rest.PurchaseResponse."];
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.purchases.api.rest.PurchaseResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      type?: string;
      /** Format: string */
      status?: string;
      /** Format: string */
      payment_status?: string;
      /** Format: string */
      shipping_status?: string;
      /** Format: string */
      ref_no?: string;
      /** Format: *string */
      invoice_no?: string;
      /** Format: string */
      branch_id?: string;
      /** Format: string */
      supplier_id?: string;
      /** Format: string */
      location_id?: string;
      /** Format: string */
      transaction_date?: string;
      /** Format: *string */
      delivery_date?: string;
      /** Format: string */
      currency?: string;
      /** Format: string */
      exchange_rate?: string;
      /** Format: string */
      total_before_tax?: string;
      /** Format: *string */
      tax_id?: string;
      /** Format: string */
      tax_amount?: string;
      /** Format: *string */
      discount_type?: string;
      /** Format: string */
      discount_amount?: string;
      /** Format: *string */
      shipping_details?: string;
      /** Format: string */
      shipping_charges?: string;
      /** Format: *string */
      additional_notes?: string;
      /** Format: *string */
      staff_note?: string;
      /** Format: string */
      round_off_amount?: string;
      /** Format: string */
      final_total?: string;
      /** Format: *int */
      pay_term_number?: number;
      /** Format: *string */
      pay_term_type?: string;
      /** Format: *string */
      document_url?: string;
      /** Format: *string */
      return_parent_id?: string;
      /** Format: []string */
      purchase_order_ids?: string[];
      /** Format: []string */
      purchase_requisition_ids?: string[];
      /** Format: *string */
      custom_field_1?: string;
      /** Format: *string */
      custom_field_2?: string;
      /** Format: *string */
      custom_field_3?: string;
      /** Format: *string */
      custom_field_4?: string;
      /** Format: []rest.PurchaseLineResponse */
      lines?: components["schemas"]["qistas.internal.purchases.api.rest.PurchaseLineResponse"][];
      /** Format: string */
      created_by?: string;
      /** Format: string */
      created_at?: string;
      /** Format: string */
      updated_at?: string;
    };
    "qistas.internal.purchases.api.rest.PurchaseLineResponse": {
      /** Format: string */
      id?: string;
      /** Format: string */
      product_id?: string;
      /** Format: string */
      variation_id?: string;
      /** Format: string */
      quantity?: string;
      /** Format: string */
      secondary_unit_quantity?: string;
      /** Format: string */
      pp_without_discount?: string;
      /** Format: string */
      discount_percent?: string;
      /** Format: string */
      purchase_price?: string;
      /** Format: string */
      purchase_price_inc_tax?: string;
      /** Format: string */
      item_tax?: string;
      /** Format: string */
      item_tobacco_tax?: string;
      /** Format: *string */
      tax_id?: string;
      /** Format: *string */
      purchase_requisition_line_id?: string;
      /** Format: *string */
      purchase_order_line_id?: string;
      /** Format: string */
      quantity_sold?: string;
      /** Format: string */
      quantity_adjusted?: string;
      /** Format: string */
      quantity_returned?: string;
      /** Format: string */
      po_quantity_purchased?: string;
      /** Format: string */
      mfg_quantity_used?: string;
      /** Format: *string */
      mfg_date?: string;
      /** Format: *string */
      exp_date?: string;
      /** Format: *string */
      lot_number?: string;
      /** Format: *string */
      sub_unit_id?: string;
    };
    "qistas.internal.purchases.api.rest.Page.qistas.internal.purchases.api.rest.PurchaseResponse.": {
      /** Format: []rest.PurchaseResponse */
      data?: components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"][];
      /** Format: int */
      page?: number;
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.purchases.api.rest.ListPurchaseOrdersRequest": {
      /** Format: *string */
      status?: string;
      /** Format: *string */
      payment_status?: string;
      /** Format: *string */
      shipping_status?: string;
      /** Format: *string */
      supplier_id?: string;
      /** Format: *string */
      location_id?: string;
      /** Format: *string */
      branch_id?: string;
      /** Format: *string */
      date_from?: string;
      /** Format: *string */
      date_to?: string;
      /** Format: string */
      ref_no?: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.purchases.api.rest.ListPurchaseOrdersResponse": {
      /** Format: []rest.PurchaseResponse */
      data?: components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"][];
      Page?: components["schemas"]["qistas.internal.purchases.api.rest.Page.qistas.internal.purchases.api.rest.PurchaseResponse."];
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.purchases.api.rest.CreatePurchaseRequest": {
      /** Format: uuid.UUID */
      branch_id: string;
      /** Format: uuid.UUID */
      supplier_id: string;
      /** Format: string */
      ref_no?: string;
      /** Format: *string */
      invoice_no?: string;
      /** Format: string */
      transaction_date: string;
      /** Format: *string */
      delivery_date?: string;
      /** Format: string */
      currency?: string;
      /** Format: *string */
      exchange_rate?: string;
      /** Format: *string */
      tax_id?: string;
      /** Format: *string */
      discount_type?: string;
      /** Format: *string */
      discount_amount?: string;
      /** Format: *string */
      shipping_details?: string;
      /** Format: *string */
      shipping_charges?: string;
      /** Format: *string */
      additional_notes?: string;
      /** Format: *string */
      staff_note?: string;
      /** Format: *string */
      round_off_amount?: string;
      /** Format: *int */
      pay_term_number?: number;
      /** Format: *string */
      pay_term_type?: string;
      /** Format: uuid.UUID */
      location_id: string;
      /** Format: *string */
      document_path?: string;
      /** Format: []string */
      purchase_order_ids?: string[];
      /** Format: []string */
      purchase_requisition_ids?: string[];
      /** Format: *string */
      custom_field_1?: string;
      /** Format: *string */
      custom_field_2?: string;
      /** Format: *string */
      custom_field_3?: string;
      /** Format: *string */
      custom_field_4?: string;
      /** Format: *string */
      status?: string;
      /** Format: []rest.CreateLineRequest */
      lines: components["schemas"]["qistas.internal.purchases.api.rest.CreateLineRequest"][];
    };
    "qistas.internal.purchases.api.rest.CreateLineRequest": {
      /** Format: uuid.UUID */
      product_id: string;
      /** Format: uuid.UUID */
      variation_id: string;
      /** Format: *string */
      quantity: string;
      /** Format: *string */
      secondary_unit_quantity?: string;
      /** Format: *string */
      pp_without_discount: string;
      /** Format: *string */
      discount_percent?: string;
      /** Format: *string */
      purchase_price: string;
      /** Format: *string */
      purchase_price_inc_tax?: string;
      /** Format: *string */
      item_tax?: string;
      /** Format: *string */
      item_tobacco_tax?: string;
      /** Format: *string */
      tax_id?: string;
      /** Format: *string */
      purchase_requisition_line_id?: string;
      /** Format: *string */
      purchase_order_line_id?: string;
      /** Format: *string */
      mfg_date?: string;
      /** Format: *string */
      exp_date?: string;
      /** Format: *string */
      lot_number?: string;
      /** Format: *string */
      sub_unit_id?: string;
    };
    "qistas.internal.purchases.api.rest.DeletePurchaseOrderRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.purchases.api.rest.DeletePurchaseResponse": {
      /** Format: string */
      id?: string;
    };
    "qistas.internal.purchases.api.rest.GetPurchaseOrderRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.purchases.api.rest.UpdatePurchaseRequest": {
      /** Format: string */
      id: string;
      /** Format: *uuid.UUID */
      branch_id?: string;
      /** Format: *uuid.UUID */
      supplier_id?: string;
      /** Format: *string */
      invoice_no?: string;
      /** Format: *string */
      transaction_date?: string;
      /** Format: *string */
      delivery_date?: string;
      /** Format: *string */
      currency?: string;
      /** Format: *string */
      exchange_rate?: string;
      /** Format: *string */
      tax_id?: string;
      /** Format: *string */
      discount_type?: string;
      /** Format: *string */
      discount_amount?: string;
      /** Format: *string */
      shipping_details?: string;
      /** Format: *string */
      shipping_charges?: string;
      /** Format: *string */
      additional_notes?: string;
      /** Format: *string */
      staff_note?: string;
      /** Format: *string */
      round_off_amount?: string;
      /** Format: *int */
      pay_term_number?: number;
      /** Format: *string */
      pay_term_type?: string;
      /** Format: *uuid.UUID */
      location_id?: string;
      /** Format: *string */
      document_path?: string;
      /** Format: *[]string */
      purchase_order_ids?: string[][];
      /** Format: *[]string */
      purchase_requisition_ids?: string[][];
      /** Format: *string */
      custom_field_1?: string;
      /** Format: *string */
      custom_field_2?: string;
      /** Format: *string */
      custom_field_3?: string;
      /** Format: *string */
      custom_field_4?: string;
      /** Format: *[]rest.CreateLineRequest */
      lines?: components["schemas"]["qistas.internal.purchases.api.rest.CreateLineRequest"][][];
    };
    "qistas.internal.purchases.api.rest.UpdatePurchaseOrderStatusRequest": {
      /** Format: string */
      id: string;
      /**
       * Format: string
       * @enum {string}
       */
      status:
        | "draft"
        | "ordered"
        | "partial"
        | "received"
        | "pending"
        | "closed"
        | "cancelled"
        | "returned";
    };
    "qistas.internal.purchases.api.rest.ListPurchaseRequisitionsRequest": {
      /** Format: *string */
      status?: string;
      /** Format: *string */
      supplier_id?: string;
      /** Format: *string */
      location_id?: string;
      /** Format: *string */
      branch_id?: string;
      /** Format: *string */
      date_from?: string;
      /** Format: *string */
      date_to?: string;
      /** Format: string */
      ref_no?: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.purchases.api.rest.DeletePurchaseRequisitionRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.purchases.api.rest.GetPurchaseRequisitionRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.purchases.api.rest.ListPurchaseRequisitionLinesRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.purchases.api.rest.ListPurchaseRequisitionLinesResponse": {
      /** Format: []rest.PurchaseLineResponse */
      lines?: components["schemas"]["qistas.internal.purchases.api.rest.PurchaseLineResponse"][];
    };
    "qistas.internal.purchases.api.rest.ListPurchaseReturnsRequest": {
      /** Format: *string */
      status?: string;
      /** Format: *string */
      payment_status?: string;
      /** Format: *string */
      supplier_id?: string;
      /** Format: *string */
      location_id?: string;
      /** Format: *string */
      branch_id?: string;
      /** Format: *string */
      date_from?: string;
      /** Format: *string */
      date_to?: string;
      /** Format: string */
      ref_no?: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.purchases.api.rest.ListPurchaseReturnsResponse": {
      /** Format: []rest.PurchaseResponse */
      data?: components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"][];
      Page?: components["schemas"]["qistas.internal.purchases.api.rest.Page.qistas.internal.purchases.api.rest.PurchaseResponse."];
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.purchases.api.rest.DeletePurchaseReturnRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.purchases.api.rest.GetPurchaseReturnRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.purchases.api.rest.ListPurchasesRequest": {
      /** Format: *string */
      status?: string;
      /** Format: *string */
      payment_status?: string;
      /** Format: *string */
      shipping_status?: string;
      /** Format: *string */
      supplier_id?: string;
      /** Format: *string */
      location_id?: string;
      /** Format: *string */
      branch_id?: string;
      /** Format: *string */
      date_from?: string;
      /** Format: *string */
      date_to?: string;
      /** Format: string */
      ref_no?: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.purchases.api.rest.ListPurchasesResponse": {
      /** Format: []rest.PurchaseResponse */
      data?: components["schemas"]["qistas.internal.purchases.api.rest.PurchaseResponse"][];
      Page?: components["schemas"]["qistas.internal.purchases.api.rest.Page.qistas.internal.purchases.api.rest.PurchaseResponse."];
      /** Format: int */
      per_page?: number;
      /** Format: int64 */
      total?: number;
      /** Format: int */
      last_page?: number;
    };
    "qistas.internal.purchases.api.rest.AddReturnRequest": {
      /** Format: string */
      id: string;
      /** Format: *string */
      invoice_no?: string;
      /** Format: string */
      transaction_date: string;
      /** Format: *string */
      delivery_date?: string;
      /** Format: string */
      currency?: string;
      /** Format: *string */
      exchange_rate?: string;
      /** Format: *string */
      tax_id?: string;
      /** Format: *string */
      discount_type?: string;
      /** Format: *string */
      discount_amount?: string;
      /** Format: *string */
      shipping_details?: string;
      /** Format: *string */
      shipping_charges?: string;
      /** Format: *string */
      additional_notes?: string;
      /** Format: *string */
      staff_note?: string;
      /** Format: *string */
      round_off_amount?: string;
      /** Format: *int */
      pay_term_number?: number;
      /** Format: *string */
      pay_term_type?: string;
      /** Format: *string */
      document_path?: string;
      /** Format: *string */
      custom_field_1?: string;
      /** Format: *string */
      custom_field_2?: string;
      /** Format: *string */
      custom_field_3?: string;
      /** Format: *string */
      custom_field_4?: string;
      /** Format: []rest.CreateLineRequest */
      lines: components["schemas"]["qistas.internal.purchases.api.rest.CreateLineRequest"][];
    };
    "qistas.internal.purchases.api.rest.RefNoExistsRequest": {
      /** Format: string */
      ref_no: string;
    };
    "qistas.internal.purchases.api.rest.RefNoExistsResponse": {
      /** Format: bool */
      exists?: boolean;
    };
    "qistas.internal.purchases.api.rest.DeletePurchaseRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.purchases.api.rest.GetPurchaseRequest": {
      /** Format: string */
      id: string;
    };
    "qistas.internal.purchases.api.rest.AddPaymentRequest": {
      /** Format: string */
      id: string;
      /** Format: string */
      amount: string;
    };
    "qistas.internal.purchases.api.rest.UpdatePurchaseStatusRequest": {
      /** Format: string */
      id: string;
      /**
       * Format: string
       * @enum {string}
       */
      status:
        | "draft"
        | "ordered"
        | "partial"
        | "received"
        | "pending"
        | "closed"
        | "cancelled"
        | "returned";
    };
    "qistas.internal.purchases.api.rest.ListPurchaseOrdersForContactRequest": {
      /** Format: string */
      supplier_id: string;
      /** Format: int */
      page?: number;
      /** Format: int */
      perPage?: number;
      /** Format: string */
      sort?: string;
    };
    "qistas.internal.subscription.api.rest.GetMySubscriptionRequest": Record<
      string,
      never
    >;
    "qistas.internal.subscription.api.rest.GetMySubscriptionResponse": {
      /** Format: bool */
      subscribed?: boolean;
      subscription?: components["schemas"]["qistas.internal.subscription.api.rest.SubscriptionResponse"];
    };
    "qistas.internal.tenants.api.rest.ResolveRequest": Record<string, never>;
    "qistas.internal.tenants.api.rest.ResolveResponse": {
      /** Format: string */
      tenant_id?: string;
      /** Format: string */
      slug?: string;
      /** Format: string */
      status?: string;
    };
  };
  responses: never;
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}
export type $defs = Record<string, never>;
export type operations = Record<string, never>;
