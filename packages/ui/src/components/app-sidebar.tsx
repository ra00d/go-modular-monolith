import { Link, useRouterState } from "@tanstack/react-router";
import { ChevronRight, type LucideIcon } from "lucide-react";

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
  SidebarFooter,
  SidebarMenuAction,
} from "@/components/ui/sidebar";
import { useDirection } from "@/components/ui/direction";
import {
  Avatar,
  AvatarFallback,
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "./ui";
import { useAuth, useCurrentUser } from "#/hooks/use-auth";
import { NavUser } from "./user-nav";

export type SidebarItem = {
  title: string;
  url?: string;
  icon?: LucideIcon;

  items?: SidebarItem[];

  /**
   * Optional metadata
   */
  permission?: string;
  badge?: string | number;
  disabled?: boolean;
};

export type AppSidebarProps = {
  items: SidebarItem[];
  label?: string;
  side?: "left" | "right";
  dir?: "rtl" | "ltr";
};

export function AppSidebar({ items, label = "Application" }: AppSidebarProps) {
  const dir = useDirection();
  const user = useCurrentUser();
  return (
    <Sidebar
      variant="inset"
      side={dir === "ltr" ? "left" : "right"}
      dir={dir}
      collapsible="icon"
    >
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>{label}</SidebarGroupLabel>

          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarItemRenderer key={item.title} item={item} />
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <NavUser
          user={{
            email: user?.email,
            name: user?.name,
            avatar: user?.avatarUrl,
          }}
        />
      </SidebarFooter>
    </Sidebar>
  );
}

function SidebarItemRenderer({ item }: { item: SidebarItem }) {
  const pathname = useRouterState({
    select: (state) => state.location.pathname,
  });
  const isActive =
    pathname === item.url ||
    (item.url !== "/" && pathname.startsWith(`${item.url}/`)) ||
    item?.items?.some(
      (c) =>
        !!c.url &&
        (pathname === c.url ||
          (c.url !== "/" && pathname.startsWith(`${c.url}/`))),
    );

  const hasChildren = !!item.items?.length;

  if (!item.url) {
    return (
      <Collapsible
        defaultOpen={isActive}
        render={
          <SidebarMenuItem>
            <SidebarMenuButton disabled={item.disabled}>
              {item.icon && <item.icon />}
              <span>{item.title}</span>

              {item.badge !== undefined && (
                <span className="ml-auto text-xs text-muted-foreground">
                  {item.badge}
                </span>
              )}
            </SidebarMenuButton>

            {hasChildren && (
              <>
                <CollapsibleTrigger
                  render={
                    <SidebarMenuAction className="data-[state=open]:rotate-90">
                      <ChevronRight />
                      <span className="sr-only">Toggle</span>
                    </SidebarMenuAction>
                  }
                />
                <CollapsibleContent>
                  <SidebarMenuSub>
                    {item.items!.map((child) => (
                      <SidebarSubItem key={child.title} item={child} />
                    ))}
                  </SidebarMenuSub>
                </CollapsibleContent>
              </>
            )}
          </SidebarMenuItem>
        }
      ></Collapsible>
    );
  }

  return (
    <Collapsible
      defaultOpen={isActive}
      render={
        <SidebarMenuItem className="text-start">
          <SidebarMenuButton
            isActive={isActive}
            tooltip={item.title}
            disabled={item.disabled}
            render={
              item.disabled ? (
                <>
                  {item.icon && <item.icon />}
                  <span>{item.title}</span>
                </>
              ) : (
                <Link to={item.url}>
                  {item.icon && <item.icon />}
                  <span>{item.title}</span>

                  {item.badge !== undefined && (
                    <span className="ml-auto text-xs text-muted-foreground">
                      {item.badge}
                    </span>
                  )}
                </Link>
              )
            }
          />

          {hasChildren && (
            <>
              <CollapsibleTrigger
                render={
                  <SidebarMenuAction className="data-[state=open]:rotate-90">
                    <ChevronRight />
                    <span className="sr-only">Toggle</span>
                  </SidebarMenuAction>
                }
              />
              <CollapsibleContent>
                <SidebarMenuSub>
                  {item.items!.map((child) => (
                    <SidebarSubItem key={child.title} item={child} />
                  ))}
                </SidebarMenuSub>
              </CollapsibleContent>
            </>
          )}
        </SidebarMenuItem>
      }
    ></Collapsible>
  );
}

function SidebarSubItem({ item }: { item: SidebarItem }) {
  const pathname = useRouterState({
    select: (state) => state.location.pathname,
  });

  const isActive =
    !!item.url &&
    (pathname === item.url ||
      (item.url !== "/" && pathname.startsWith(`${item.url}/`)));

  return (
    <SidebarMenuSubItem>
      <SidebarMenuSubButton
        isActive={isActive}
        className="text-start"
        render={
          item.disabled || !item.url ? (
            <>
              {item.icon && <item.icon />}
              <span>{item.title}</span>
            </>
          ) : (
            <Link to={item.url}>
              {item.icon && <item.icon />}
              <span>{item.title}</span>

              {item.badge !== undefined && (
                <span className="ml-auto text-xs text-muted-foreground">
                  {item.badge}
                </span>
              )}
            </Link>
          )
        }
      >
        {}
      </SidebarMenuSubButton>
    </SidebarMenuSubItem>
  );
}
