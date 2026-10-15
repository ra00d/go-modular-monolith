"use client";

import { Subscribe } from "@tanstack/react-table";
import { useHeaderContext, useTableContext } from "@/hooks/table";
import { Checkbox } from "@/components/ui/checkbox";
import { cn } from "@/lib/utils";

export function SelectAllHeader(): React.ReactNode {
  const table = useTableContext();

  return (
    <Subscribe source={table.atoms.rowSelection}>
      {() => (
        <Checkbox
          checked={table.getIsAllPageRowsSelected()}
          indeterminate={
            !table.getIsAllPageRowsSelected() &&
            table.getIsSomePageRowsSelected()
          }
          onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
          aria-label="Select all"
          className="translate-y-0.5"
        />
      )}
    </Subscribe>
  );
}

export function ResizeHandle(): React.ReactNode {
  const header = useHeaderContext();
  const table = useTableContext();

  if (!header.column.getCanResize()) return null;

  return (
    <Subscribe source={table.atoms.columnResizing}>
      {() => (
        <div
          onDoubleClick={() => header.column.resetSize()}
          onMouseDown={header.getResizeHandler()}
          onTouchStart={header.getResizeHandler()}
          className={cn(
            'absolute -right-0.5 z-10 top-1/2 h-6 w-0.75 -translate-y-1/2 cursor-e-resize select-none touch-none rounded-md transition-colors hover:bg-blue-600 before:absolute before:-left-1 before:-right-1 before:top-0 before:h-full before:content-[""]',
            header.column.getIsResizing() && "bg-blue-600",
          )}
        />
      )}
    </Subscribe>
  );
}
