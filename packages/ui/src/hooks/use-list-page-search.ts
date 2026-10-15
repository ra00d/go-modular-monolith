import { useCallback, useEffect, useState } from "react";
import { useNavigate, useSearch } from "@tanstack/react-router";

export interface ListUrlPatch {
  page?: number;
  pageSize?: number;
  q?: string;
  search?: {
    query?: string;
  };
  filter?: Record<string, any>;
  sort?: {
    field: string;
    direction?: "asc" | "desc";
  };
}

/**
 * Persists server-side list state (page / pageSize / search) in the URL so it
 * survives refreshes, is shareable, and works with browser back/forward.
 * Routes must declare `validateSearch: listPaginationSearch`.
 */
export function useListPageSearch() {
  const navigate = useNavigate();

  const params = useSearch({ strict: false }) as
    | Record<string, any>
    | undefined;
  const [page, setPage] = useState(Math.max(1, Number(params?.page) || 1));
  const [pageSize, setPageSize] = useState(
    Math.min(100, Math.max(1, Number(params?.pageSize) || 10)),
  );
  const [search, setSearch] = useState<ListUrlPatch["search"]>(
    typeof params?.search?.query === "string" ? params?.search : {},
  );
  const [filter, setFilter] = useState<ListUrlPatch["filter"]>(params?.filter);
  const [sort, setSort] = useState<ListUrlPatch["sort"]>(params?.filter);

  useEffect(() => {
    setSearch(params?.search);
    setPage(params?.page);
    setPageSize(params?.pageSize);
    setFilter(params?.filter);
    setSort(params?.sort);
  }, [params]);
  const update = useCallback(
    (patch: ListUrlPatch, opts?: { replace?: boolean; clear?: boolean }) => {
      // Cast needed: the typed router's ParamsReducer expects route-specific
      // search schemas; we merge generically over the current search params.
      navigate({
        search: (prev: Record<string, unknown>) => {
          let { q, ...rest } = prev;
          if (opts?.clear) rest = {};
          const url = {
            ...(rest ?? {}),
            ...patch,
          };

          return url;
        },
        replace: opts?.replace ?? false,
      } as any);
    },
    [navigate],
  );

  return {
    page,
    pageSize,
    search,
    filter,
    sort,
    update,
  };
}
