import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";

import type { AuthState, LoginCredentials, User, UserRole } from "#/types";

export const useAuthStore = create<AuthState>()(
  devtools(
    persist(
      (set, get) => ({
        user: null,
        status: "loading",
        error: null,
        setUser: (u) => {
          set({
            user: u,
            error: null,
            status: "authenticated",
          });
        },
        initialize: async () => {
          // Prevent duplicate initialization.
          if (get().status !== "loading") {
            return;
          }
          const user: User = {
            name: "ali",
            role: "admin",
            permissions: ["*"],
            id: "djd",
            email: "user@email",
          };
          try {
            set({
              user,
              status: user ? "authenticated" : "unauthenticated",
              error: null,
            });
          } catch (error) {
            set({
              user: null,
              status: "unauthenticated",
              error:
                error instanceof Error
                  ? error.message
                  : "Failed to initialize authentication.",
            });
          }
        },

        login: async (credentials: LoginCredentials): Promise<User> => {
          set({
            status: "loading",
            error: null,
          });

          try {
            const user: User = {
              name: "ali",
              role: "admin",
              permissions: ["*"],
              id: "djd",
              email: "user@email",
            };

            set({
              user,
              status: "authenticated",
              error: null,
            });

            return user;
          } catch (error) {
            const message =
              error instanceof Error ? error.message : "Login failed.";

            set({
              user: null,
              status: "unauthenticated",
              error: message,
            });

            throw error;
          }
        },

        logout: async () => {
          try {
            //         await authApi.logout();
          } finally {
            set({
              user: null,
              status: "unauthenticated",
              error: null,
            });
          }
        },

        refresh: async () => {
          try {
            const user: User = {
              name: "ali",
              role: "admin",
              permissions: ["*"],
              id: "djd",
              email: "user@email",
            };

            set({
              user,
              status: user ? "authenticated" : "unauthenticated",
              error: null,
            });

            return user;
          } catch (error) {
            set({
              user: null,
              status: "unauthenticated",
              error:
                error instanceof Error
                  ? error.message
                  : "Failed to refresh authentication.",
            });

            return null;
          }
        },

        clearError: () => {
          set({
            error: null,
          });
        },

        hasRole: (role: UserRole | UserRole[]) => {
          const user = get().user;

          if (!user) {
            return false;
          }

          const roles = Array.isArray(role) ? role : [role];

          return roles.includes(user.role);
        },

        hasPermission: (permission: string) => {
          const user = get().user;

          if (!user) {
            return false;
          }

          return user.permissions.includes(permission);
        },
      }),
      {
        name: "auth-store",
      },
    ),
  ),
);
