import { TanStackDevtools } from "@tanstack/react-devtools";
import {
  ErrorComponent,
  HeadContent,
  Outlet,
  Scripts,
  createRootRouteWithContext,
  type ErrorComponentProps,
} from "@tanstack/react-router";
import { TanStackRouterDevtoolsPanel } from "@tanstack/react-router-devtools";
import { AppProvider } from "@workspace/ui/providers/app-provider";

import TanStackQueryDevtools from "../integrations/tanstack-query/devtools";

import { getLocale, getTextDirection } from "#/paraglide/runtime";

//@ts-ignore
import appCss from "@workspace/ui/styles.css?url";

import { type QueryClient } from "@tanstack/react-query";
import { ErrorPage, NotFound } from "@workspace/ui/pages";

interface MyRouterContext {
  queryClient: QueryClient;
}

const THEME_INIT_SCRIPT = `(function(){try{var stored=window.localStorage.getItem('theme');var mode=(stored==='light'||stored==='dark'||stored==='auto')?stored:'auto';var prefersDark=window.matchMedia('(prefers-color-scheme: dark)').matches;var resolved=mode==='auto'?(prefersDark?'dark':'light'):mode;var root=document.documentElement;root.classList.remove('light','dark');root.classList.add(resolved);if(mode==='auto'){root.removeAttribute('data-theme')}else{root.setAttribute('data-theme',mode)}root.style.colorScheme=resolved;}catch(e){}})();`;
export const DefaultCatchBoundary = ({ error }: ErrorComponentProps) => {

  return (
    <div>
      <ErrorComponent error={error} />
      <ErrorPage />
    </div>
  );
};
export const Route = createRootRouteWithContext<MyRouterContext>()({
  // beforeLoad: async () => {
  //   // Other redirect strategies are possible; see
  //   // https://github.com/TanStack/router/tree/main/examples/react/i18n-paraglide#offline-redirect
  //   if (typeof document !== "undefined") {
  //     document.documentElement.setAttribute("lang", getLocale());
  //   }
  // },

  head: () => ({
    meta: [
      {
        charSet: "utf-8",
      },
      {
        name: "viewport",
        content: "width=device-width, initial-scale=1",
      },
      {
        title: "TanStack Start Starter",
      },
    ],
    links: [
      {
        rel: "stylesheet",
        href: appCss,
      },
    ],
  }),
  errorComponent: (props) => {
    return (
      <RootDocument>
        <DefaultCatchBoundary {...props} />
      </RootDocument>
    );
  },
  // shellComponent: RootDocument,
  component: RootComponent,
  notFoundComponent: NotFound,
});
function RootComponent() {
  return (
    <RootDocument>
      <Outlet />
    </RootDocument>
  );
}

function RootDocument({ children }: { children: React.ReactNode }) {
  return (
    <html lang={getLocale()} dir={getTextDirection()} suppressHydrationWarning>
      <head>
        <script dangerouslySetInnerHTML={{ __html: THEME_INIT_SCRIPT }} />
        <HeadContent />
      </head>
      <body className="w-full">
        <AppProvider>{children}</AppProvider>
        <TanStackDevtools
          config={{
            position: "bottom-right",
          }}
          plugins={[
            {
              name: "Tanstack Router",
              render: <TanStackRouterDevtoolsPanel />,
            },
            TanStackQueryDevtools,
          ]}
        />
        <Scripts />
      </body>
    </html>
  );
}
