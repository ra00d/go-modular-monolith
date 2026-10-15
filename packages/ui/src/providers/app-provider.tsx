import type { ThemeProviderProps } from "next-themes";
import { ThemeProvider } from "./theme-provider";
import { useEffect, useState, type ReactNode } from "react";
import { DirectionProvider, SidebarProvider, Toaster } from "#/components";
import { getTextDirection } from "#/paraglide/runtime";
import type { TextDirection } from "@base-ui/react";
interface AppProviderProps {
  theme?: ThemeProviderProps;
  children: ReactNode;
}
export function AppProvider(opts: AppProviderProps) {
  const [dir, setDir] = useState<TextDirection>("rtl");
  useEffect(() => {
    setDir(getTextDirection());
  }, []);
  return (
    <DirectionProvider direction={dir}>
      <ThemeProvider {...opts.theme}>
        <SidebarProvider>{opts.children}</SidebarProvider>
      </ThemeProvider>
      <Toaster />
    </DirectionProvider>
  );
}
