import { cn } from "@/lib/utils";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
  Separator,
  SidebarTrigger,
  useDirection,
} from "#/components";
import { Fragment, type ReactNode } from "react";
import { LanguageSwitcher } from "./language-switcher";
import { useBreadcrumb } from "#/hooks/use-breadcrumbs";

interface TobBarProps {
  className?: string;
  extras?: ReactNode[];
}
/*
Top bar is a component that is used to display the top bar of the application.
it contains a language switcher , a theme switcher, a notifications button, and a button to toggle sidebar
*/
function TobBar({ className, extras }: TobBarProps) {
  const dir = useDirection();
  const crumbs = useBreadcrumb();

  return (
    <header
      dir={dir === "rtl" ? "rtl" : "ltr"}
      className={cn(
        "flex h-16 shrink-0 items-center justify-between gap-2 border-b px-4",
        className,
      )}
    >
      <div className="flex gap-2">
        <SidebarTrigger className="-ml-1" />
        <Separator
          orientation="vertical"
          className="mr-2 data-[orientation=vertical]:h-4"
        />

        <Breadcrumb>
          <BreadcrumbList>
            {crumbs.parts.map((c, i) => (
              <Fragment key={c}>
                <BreadcrumbItem className="hidden md:block">
                  {i !== crumbs.parts.length - 1 ? (
                    <BreadcrumbLink href={c}>{c}</BreadcrumbLink>
                  ) : (
                    c
                  )}
                </BreadcrumbItem>
                {i !== crumbs.parts.length - 1 && (
                  <BreadcrumbSeparator className="hidden md:block" />
                )}
              </Fragment>
            ))}
          </BreadcrumbList>
        </Breadcrumb>
      </div>
      <div className="flex gap-2">
        {extras?.map((c, i) => (
          <Fragment key={i}>{c}</Fragment>
        ))}
        <LanguageSwitcher />
      </div>
    </header>
  );
}

export { TobBar };
