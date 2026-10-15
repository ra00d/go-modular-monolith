import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useCellContext, useTableContext } from "@/hooks/table";
import { formatDate, toSentenceCase } from "@/lib/utils";
import { Subscribe } from "@tanstack/react-table";
import { ChevronDown, ChevronRight, MoreHorizontal } from "lucide-react";

export function SelectCell(): React.ReactNode {
  const cell = useCellContext();
  const table = useTableContext();
  const row = cell.row;

  return (
    <Subscribe source={table.atoms.rowSelection}>
      {() => (
        <Checkbox
          checked={row.getIsSelected()}
          indeterminate={!row.getIsSelected() && row.getIsSomeSelected()}
          onCheckedChange={(value) => row.toggleSelected(!!value)}
          aria-label="Select row"
          className="translate-y-0.5"
        />
      )}
    </Subscribe>
  );
}

export function TextCell(): React.ReactNode {
  const cell = useCellContext<string | undefined>();
  const value = cell.getValue();
  return value == null ? null : <>{String(value)}</>;
}

export function AgeCell(): React.ReactNode {
  const cell = useCellContext<number>();
  return <span>{String(cell.getValue())}</span>;
}

export function StatusCell<T extends string | number | symbol = string>(
  icons: Record<T, React.ReactNode>,
): React.ReactNode {
  const cell = useCellContext<T>();
  const status = cell.getValue();
  if (!status) return null;
  //
  // const icons: Record<T, React.ReactNode> = {
  //   active: <CheckCircle />,
  //   inactive: <XCircle />,
  //   pending: <Clock />,
  // };

  return (
    <Badge
      variant="outline"
      className="gap-1 w-fit [&>svg]:size-3.5 px-3 py-1 [&>svg]:shrink-0 rounded-full"
    >
      {icons[status]}
      <span className="truncate">{toSentenceCase(status.toString())}</span>
    </Badge>
  );
}
export function DateCell(): React.ReactNode {
  const cell = useCellContext<string>();
  return <>{formatDate(cell.getValue())}</>;
}

export function GroupedCell(): React.ReactNode {
  const cell = useCellContext();
  const table = useTableContext();
  const row = cell.row;

  return (
    <Subscribe source={table.atoms.expanded}>
      {() => (
        <Button
          variant="ghost"
          size="sm"
          className="-ml-2 h-7 gap-1 px-2"
          onClick={row.getToggleExpandedHandler()}
          disabled={!row.getCanExpand()}
          style={{
            paddingLeft: `calc(${row.depth} * 1.5rem + 0.5rem)`,
          }}
        >
          {row.getIsExpanded() ? (
            <ChevronDown className="size-4" />
          ) : (
            <ChevronRight className="size-4" />
          )}
          <cell.FlexRender />
          <span className="text-muted-foreground">({row.subRows.length})</span>
        </Button>
      )}
    </Subscribe>
  );
}
export interface TableAction<T = unknown> {
  label: React.ReactNode;
  action: (v: T) => any;
  withSeprator?: boolean;
}
export function ActionsCell<T = any>({
  actions,
}: {
  actions: TableAction<T>[];
}): React.ReactNode {
  const cell = useCellContext();
  const row = cell.row.original;

  return (
    <DropdownMenu>
      <DropdownMenuTrigger
        render={<Button variant="ghost" className="h-8 w-8 p-0" />}
      >
        <span className="sr-only">Open menu</span>
        <MoreHorizontal className="h-4 w-4" />
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuLabel>Actions</DropdownMenuLabel>
        {actions.map((a, index) => {
          return (
            <>
              <DropdownMenuItem
                key={index}
                onClick={() => {
                  a.action(row);
                }}
              >
                {a.label}
              </DropdownMenuItem>
              {a.withSeprator && <DropdownMenuSeparator />}
            </>
          );
        })}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
