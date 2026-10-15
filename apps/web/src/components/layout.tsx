import { sidebarItems } from "#/constants";
import { m } from "#/paraglide/messages";
import { NotificationPanel } from "@workspace/ui/components/notification-sheet";
import { TobBar } from "@workspace/ui/components/top-bar";
import { ThemeSwitcher } from "@workspace/ui/components/ui/theme-switcher";
import { AppLayout } from "@workspace/ui/layouts/app-layout";
import type { PropsWithChildren } from "react";
import { LogIn } from "lucide-react";

export const Layout = ({ children }: PropsWithChildren) => {
  return (
    <AppLayout
      sidebarProps={{
        label: m.nav_home(),
        items: [
          ...sidebarItems,
          {
            title: "form",
            // icon: Globe,
            url: "/form",
          },
          {
            title: m.error_not_found(),
            //icon: Globe,
            url: "/not-found",
          },
          {
            title: m.nav_login(),
            url: "/login",
            icon: LogIn,
          },
        ],
      }}
      TobBar={<TobBar extras={[<NotificationPanel />, <ThemeSwitcher />]} />}
    >
      <div className="max-h-[calc(100vh-64px)] overflow-hidden">{children}</div>
    </AppLayout>
  );
};
