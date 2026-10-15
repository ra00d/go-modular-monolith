"use client";

import { Input } from "@/components/ui/input";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useAppTable } from "@/hooks/table";
import type { features } from "@/hooks/table-features";
import { cn } from "@/lib/utils";
import type { ExtendedColumnFilter } from "@/types";
import { useDebouncedCallback } from "@tanstack/react-pacer/debouncer";
import type { Column, ColumnDef, RowData } from "@tanstack/react-table";
import { Search } from "lucide-react";
import { useEffect, useState } from "react";
import { ScrollArea } from "../ui";

function getCommonPinningStyles<T extends RowData = any>(
  column: Column<typeof features, T>,
  isSelected = false,
): React.CSSProperties {
  const isPinned = column.getIsPinned();
  const isLastLeftPinnedColumn =
    isPinned === "start" && column.getIsLastColumn("start");
  const isFirstRightPinnedColumn =
    isPinned === "end" && column.getIsFirstColumn("end");

  return {
    boxShadow: isLastLeftPinnedColumn
      ? "-4px 0 4px -4px var(--border) inset"
      : isFirstRightPinnedColumn
        ? "4px 0 4px -4px var(--border) inset"
        : undefined,
    insetInlineStart:
      isPinned === "start" ? `${column.getStart("start")}px` : undefined,
    insetInlineEnd:
      isPinned === "end" ? `${column.getAfter("end")}px` : undefined,
    position: isPinned ? "sticky" : "relative",
    background: isSelected
      ? "var(--muted)"
      : isPinned
        ? "var(--background)"
        : undefined,
    zIndex: isPinned ? 1 : 0,
  };
}

export interface DataTableProps {
  table: any;
}
interface TableOptions<T extends RowData> {
  data: T[];
  columns: ColumnDef<typeof features, T, any>[];
}
export function DataTable<T extends RowData>(opts: TableOptions<T>) {
  const [columnFilters, setColumnFilters] = useState<
    Array<ExtendedColumnFilter>
  >([]);
  const [globalFilter, setGlobalFilter] = useState("");
  const table = useAppTable({
    //...baseTable.options,
    data: opts.data,
    columns: opts.columns,
    state: {
      columnFilters,
      globalFilter,
    },
    onColumnFiltersChange: setColumnFilters,
    onGlobalFilterChange: setGlobalFilter,
  });
  return (
    <table.AppTable>
      <div className=" container mx-auto p-4 flex flex-col gap-4">
        <div className="flex flex-col gap-4">
          <div className="flex items-center gap-2">
            <div className="relative w-full max-w-sm">
              <Search className="pointer-events-none absolute left-2.5 top-1/2 size-4 -translate-y-1/2 text-muted-foreground" />
              <DebouncedInput
                value={globalFilter}
                onChange={(value) => setGlobalFilter(String(value))}
                placeholder="Search all columns..."
                className="pl-8"
              />
            </div>
            <table.FilterList />
            <table.SortList />
            <table.ViewOptions />
          </div>
          <ScrollArea className=" h-150  rounded-md border">
            <Table className="max-h-full">
              <TableHeader className="bg-card sticky top-0 z-1">
                {table.getHeaderGroups().map((headerGroup: any) => (
                  <TableRow key={headerGroup.id}>
                    {headerGroup.headers
                      .filter((header: any) => header.column.getIsVisible())
                      .map((header: any) => (
                        <table.AppHeader header={header} key={header.id}>
                          {(h: any) => (
                            <TableHead
                              colSpan={h.colSpan}
                              className={cn("relative", {
                                "border-r": h.id !== "actions",
                                "text-center *:[[role=checkbox]]:mx-auto":
                                  h.column.id === "select",
                              })}
                              style={{
                                flexGrow: h.getSize(),
                                width: h.getSize(),
                                ...getCommonPinningStyles(h.column),
                              }}
                            >
                              {h.isPlaceholder ? null : <h.FlexRender />}
                              <h.ResizeHandle />
                            </TableHead>
                          )}
                        </table.AppHeader>
                      ))}
                  </TableRow>
                ))}
              </TableHeader>
              <TableBody>
                {table.getRowModel().rows.map((row: any) => (
                  <TableRow
                    key={row.id}
                    data-state={row.getIsSelected() ? "selected" : undefined}
                    aria-selected={row.getIsSelected()}
                  >
                    {row.getVisibleCells().map((cell: any) => (
                      <table.AppCell cell={cell} key={cell.id}>
                        {(c: any) => (
                          <TableCell
                            className={cn(
                              c.column.id === "actions" ? "" : "border-r",
                              c.column.id === "select" &&
                                "text-center *:[[role=checkbox]]:mx-auto",
                            )}
                            style={{
                              flexGrow: c.column.getSize(),
                              width: c.column.getSize(),
                              ...getCommonPinningStyles(
                                c.column,
                                row.getIsSelected(),
                              ),
                            }}
                          >
                            {c.getIsGrouped() ? (
                              <c.GroupedCell />
                            ) : (
                              <c.FlexRender />
                            )}
                          </TableCell>
                        )}
                      </table.AppCell>
                    ))}
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </ScrollArea>
          <table.Pagination />
        </div>
      </div>
    </table.AppTable>
  );
}

function DebouncedInput({
  value: initialValue,
  onChange,
  debounce = 300,
  ...props
}: {
  value: string | number;
  onChange: (value: string | number) => void;
  debounce?: number;
} & Omit<React.ComponentProps<typeof Input>, "onChange">) {
  const [value, setValue] = useState(initialValue);

  useEffect(() => {
    setValue(initialValue);
  }, [initialValue]);

  const debouncedOnChange = useDebouncedCallback(onChange, { wait: debounce });

  return (
    <Input
      {...props}
      value={value}
      onChange={(e) => {
        setValue(e.target.value);
        debouncedOnChange(e.target.value);
      }}
    />
  );
}
