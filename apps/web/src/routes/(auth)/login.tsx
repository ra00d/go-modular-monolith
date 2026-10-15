import LoginPage from "#/modules/auth/pages/login";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/(auth)/login")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div className="w-full">
      <LoginPage />;
    </div>
  );
}
