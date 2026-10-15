import { AppSidebar, type AppSidebarProps } from "#/components/app-sidebar";
import { useDirection } from "@/components/ui/direction";
import { SidebarInset } from "@/components/ui/sidebar";
import type { ReactNode } from "react";

export interface AppLayoutProps {
  children: ReactNode;
  TobBar: ReactNode;
  sidebarProps: AppSidebarProps;
}

export function AppLayout({ children, TobBar, sidebarProps }: AppLayoutProps) {
  const dir = useDirection();
  return (
    <>
      {dir === "ltr" ? (
        <>
          <SidebarInset>
            {TobBar}
            <div className="flex flex-1 flex-col gap-4 p-4">{children}</div>
          </SidebarInset>
          <AppSidebar {...sidebarProps} />
        </>
      ) : (
        <>
          <AppSidebar {...sidebarProps} />
          <SidebarInset>
            {TobBar}
            <div className="flex flex-1 flex-col gap-4 p-4">{children}</div>
          </SidebarInset>
        </>
      )}
    </>
  );
}
