import {
  createRouter as createTanStackRouter,
  ErrorComponent,
  type ErrorComponentProps,
} from "@tanstack/react-router";
import { routeTree } from "./routeTree.gen";

import { QueryClient } from "@tanstack/react-query";
import { setupRouterSsrQueryIntegration } from "@tanstack/react-router-ssr-query";
import { ErrorPage, NotFound } from "@workspace/ui/pages";
export const DefaultCatchBoundary = ({ error }: ErrorComponentProps) => {
  return (
    <div>
      <ErrorComponent error={error} />
      <ErrorPage />
    </div>
  );
};

export async function getRouter() {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        staleTime: 30_000,
      },
    },
  });

  const router = createTanStackRouter({
    routeTree,
    context: {
      queryClient,
    },
    scrollRestoration: true,
    defaultPreload: "intent",
    defaultPreloadStaleTime: 0,
    defaultErrorComponent: DefaultCatchBoundary,
    defaultNotFoundComponent: () => <NotFound />,
  });

  setupRouterSsrQueryIntegration({
    router,
    queryClient,
  });

  return router;
}

declare module "@tanstack/react-router" {
  interface Register {
    router: ReturnType<typeof getRouter>;
  }
}
