import { useMatchRoute, useRouterState } from "@tanstack/react-router";

interface BreadcrumbState {
  parts: string[];
  current: string;
}
export function useBreadcrumb(): BreadcrumbState {
  const router = useRouterState();
  return {
    parts:
      Array.from(new Set(router.matches.map((r) => r.pathname)))
        .at(-1)
        ?.split("/")
        .filter((r) => r.length > 0) || [],
    current: router.location.pathname,
  };
}
