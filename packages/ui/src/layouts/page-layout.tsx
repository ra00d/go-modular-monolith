import * as React from "react";
import { cn } from "@/lib/utils";

export interface PageLayoutProps {
  children: React.ReactNode;
  className?: string;
}

export interface PageHeaderProps {
  children: React.ReactNode;
  className?: string;
}

export interface PageHeaderContentProps {
  children: React.ReactNode;
  className?: string;
}

export interface PageTitleProps {
  children: React.ReactNode;
  className?: string;
}

export interface PageDescriptionProps {
  children: React.ReactNode;
  className?: string;
}

export interface PageActionsProps {
  children: React.ReactNode;
  className?: string;
}

export interface PageBodyProps {
  children: React.ReactNode;
  className?: string;
}

export interface PageFooterProps {
  children: React.ReactNode;
  className?: string;
}

function PageLayout({ children, className }: PageLayoutProps) {
  return (
    <div className={cn("flex  min-h-0 flex-col", className)}>{children}</div>
  );
}

function PageHeader({ children, className }: PageHeaderProps) {
  return (
    <header
      className={cn(
        "flex shrink-0 items-start justify-between gap-6 border-b px-6 py-5",
        className,
      )}
    >
      {children}
    </header>
  );
}

function PageHeaderContent({ children, className }: PageHeaderContentProps) {
  return <div className={cn("min-w-0 flex-1", className)}>{children}</div>;
}

function PageTitle({ children, className }: PageTitleProps) {
  return (
    <h1 className={cn("text-2xl font-semibold tracking-tight", className)}>
      {children}
    </h1>
  );
}

function PageDescription({ children, className }: PageDescriptionProps) {
  return (
    <p className={cn("mt-1 text-sm text-muted-foreground", className)}>
      {children}
    </p>
  );
}

function PageActions({ children, className }: PageActionsProps) {
  return (
    <div className={cn("flex shrink-0 items-center gap-2", className)}>
      {children}
    </div>
  );
}

function PageBody({ children, className }: PageBodyProps) {
  return (
    <main className={cn("min-h-0 flex-1 overflow-auto p-6", className)}>
      {children}
    </main>
  );
}

function PageFooter({ children, className }: PageFooterProps) {
  return (
    <footer
      className={cn("flex shrink-0 items-center border-t px-6 py-4", className)}
    >
      {children}
    </footer>
  );
}

/* Compound components */
PageLayout.Header = PageHeader;
PageLayout.HeaderContent = PageHeaderContent;
PageLayout.Title = PageTitle;
PageLayout.Description = PageDescription;
PageLayout.Actions = PageActions;
PageLayout.Body = PageBody;
PageLayout.Footer = PageFooter;

export {
  PageLayout,
  PageHeader,
  PageHeaderContent,
  PageTitle,
  PageDescription,
  PageActions,
  PageBody,
  PageFooter,
};
