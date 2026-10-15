import { Layout } from "#/components/layout";
import { createFileRoute, Outlet } from "@tanstack/react-router";
import { NotFound, ErrorPage } from "@workspace/ui/pages";

export const Route = createFileRoute("/(app)")({
  component: RouteComponent,
  notFoundComponent: NotFound,
  errorComponent: ErrorPage,
});

function RouteComponent() {
  return (
    <Layout>
      <Outlet />
    </Layout>
  );
}
