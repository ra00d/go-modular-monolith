import { useAuthStore } from "#/stores/auth-store";

export function useAuth() {
  return useAuthStore((state) => ({
    user: state.user,
    status: state.status,
    error: state.error,

    login: state.login,
    logout: state.logout,
    refresh: state.refresh,

    hasRole: state.hasRole,
    hasPermission: state.hasPermission,

    clearError: state.clearError,
  }));
}

export function useCurrentUser() {
  return useAuthStore((state) => state.user);
}

export function useIsAuthenticated() {
  return useAuthStore((state) => state.status === "authenticated");
}
