export type UserRole = "admin" | "manager" | "user";

export type User = {
  id: string;
  email: string;
  name: string;
  avatarUrl?: string;
  role: UserRole;
  permissions: string[];
  roles?: string[];
};

export type LoginCredentials = {
  email: string;
  password: string;
  rememberMe?: boolean;
};

export type AuthStatus = "loading" | "authenticated" | "unauthenticated";

export type AuthState = {
  user: User;
  status: AuthStatus;
  error: string | null;

  initialize: () => Promise<void>;
  setUser: (u: User) => void;
  login: (credentials: LoginCredentials) => Promise<User>;
  logout: () => Promise<void>;
  refresh: () => Promise<User | null>;
  clearError: () => void;

  hasRole: (role: UserRole | UserRole[]) => boolean;
  hasPermission: (permission: string) => boolean;
};
