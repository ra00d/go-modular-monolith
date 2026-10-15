import type { ResponseData } from "#/types/helpers";
import { DataTable, Empty, Spinner } from "@workspace/ui/components";
import { createAppColumnHelper } from "@workspace/ui/hooks/table";
import { PageHeader, PageLayout } from "@workspace/ui/layouts/page-layout";
import { useMemo } from "react";
import { useUsersQuery } from "../hooks/queries";
import { ErrorPage } from "@workspace/ui/pages";

export function UsersList() {
  const columns = useMemo(() => {
    const helper = createAppColumnHelper<ResponseData<"/api/admin/users">>();
    return helper.columns([
      helper.accessor("name", {}),
      helper.accessor("email", {}),
      helper.accessor("phone", {}),
      helper.accessor("status", {}),
    ]);
  }, []);
  const { data, isLoading, isError } = useUsersQuery();
  if (isLoading) return <Spinner />;
  if (isError) {
    return <ErrorPage />;
  }
  if (!data || data.data.length === 0) return <Empty />;

  return (
    <PageLayout>
      <PageHeader>Users</PageHeader>
      <DataTable columns={columns} data={data.data} />
    </PageLayout>
  );
}
