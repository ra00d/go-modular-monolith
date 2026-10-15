import { NotificationPanel } from "@/components/notification-sheet";
import { Globe } from "lucide-react";
import { SidebarTrigger } from "./components";
import { LanguageSwitcher } from "./components/language-switcher";
import { AppSidebar } from "./components/ui/app-sidebar";
import { ThemeSwitcher } from "./components/ui/theme-switcher";

export default function AppShell() {
  return (
    <main className="flex">
      <AppSidebar
        items={[
          {
            title: "test",
            icon: Globe,
            url: "/",
          },
          {
            title: "test-x",
            icon: Globe,
            url: "/",
          },
        ]}
      />
      <div className="flex gap-2 bg-card w-full">
        <SidebarTrigger />
        <ThemeSwitcher />
        <LanguageSwitcher />
        <NotificationPanel />
      </div>
    </main>
  );
}
