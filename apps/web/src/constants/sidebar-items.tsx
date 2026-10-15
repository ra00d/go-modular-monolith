import { m } from "#/paraglide/messages";
import type { SidebarItem } from "@workspace/ui/components";
import { Info, Package2, Plus, Users, Grid, List } from "lucide-react";

export const sidebarItems: SidebarItem[] = [
  {
    title: m.nav_about(),
    icon: Info,
    url: "/about",
  },
  {
    title: m.nav_users({
      count: "one",
    }),
    icon: Users,
    url: "/users",
    items: [
      {
        title: m.action_create(),
        icon: Plus,
        url: "/users/new",
      },
    ],
  },
  {
    title: m.nav_products(),
    icon: Package2,
    items: [
      {
        title: m.nav_products_list(),
        url: "/products",
        icon: List,
      },
      {
        title: m.nav_categories(),
        icon: Grid,
        url: "/products/categories",
      },

      {
        title: m.action_create(),
        icon: Plus,
        url: "/products/new",
      },
    ],
  },
];

export default sidebarItems;
