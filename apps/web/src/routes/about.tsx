import { Layout } from "#/components/layout";
import { createFileRoute } from "@tanstack/react-router";
import { Button, ColumnHeader, DataTable } from "@workspace/ui/components";
import { createAppColumnHelper } from "@workspace/ui/hooks/table";
import { PageLayout } from "@workspace/ui/layouts/page-layout";
import { useMemo } from "react";

export const Route = createFileRoute("/about")({
  component: RouteComponent,
});
type Person = {
  id: string;
  name: string;
};
function RouteComponent() {
  const columns = useMemo(() => {
    const columnHelper = createAppColumnHelper<Person>();
    return columnHelper.columns([
      // columnHelper.display({
      //   id: "select",
      //   // header: SelectCell,
      //   // cell: SelectCell,
      // }),
      columnHelper.accessor("id", {
        id: "id",
        header: () => ColumnHeader({ title: "ID" }),
        cell: (info) => info.getValue(),
      }),
      columnHelper.accessor("name", {
        id: "name",
        header: () => ColumnHeader({ title: "Name" }),
        cell: (info) => info.getValue(),
      }),
    ]);
  }, []);
  return (
    <Layout>
      <PageLayout>
        <PageLayout.Header>
          <PageLayout.HeaderContent>
            <PageLayout.Title>Products</PageLayout.Title>

            <PageLayout.Description>
              Manage your products and inventory.
            </PageLayout.Description>
          </PageLayout.HeaderContent>

          <PageLayout.Actions>
            <Button variant="outline">Export</Button>

            <Button>Add Product</Button>
          </PageLayout.Actions>
        </PageLayout.Header>

        <PageLayout.Body>
          <DataTable
            columns={columns}
            data={Array.from({ length: 30 }, () => ({
              id: Math.random().toString(),
              name: "Raad " + Math.random(),
            }))}
          />
        </PageLayout.Body>
      </PageLayout>
    </Layout>
  );
}
