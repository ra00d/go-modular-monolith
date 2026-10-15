import { UsersList } from "#/modules/users/pages/users-list";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/(app)/users/")({
  component: RouteComponent,
});

function RouteComponent() {
  return <UsersList />;
}
