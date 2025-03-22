import {
    createColumnHelper,
    flexRender,
    getCoreRowModel,
    useReactTable,
} from "@tanstack/react-table";
import type { Role } from "../../../api/model";
import { useGetRoles } from "../../../api/role/role.ts";
import { Loading } from "../../../components/ui/Loading";

export const Component = () => {
    const { data, isLoading, error } = useGetRoles();

    if (isLoading) {
        return <Loading />;
    }

    if (error) {
        throw new Error(error.message);
    }

    const roles = data?.data ?? [];

    return <RoleTable roles={roles} />;
};

const columnHelper = createColumnHelper<Role>();
const columns = [
    columnHelper.accessor("name", {
        header: () => "Name",
        cell: (info) => info.getValue(),
    }),
    columnHelper.accessor("created_at", {
        header: "Created at",
        cell: (info) => info.getValue(),
    }),
];

const RoleTable = ({ roles }: { roles: Role[] }) => {
    const table = useReactTable({
        data: roles,
        columns: columns,
        getCoreRowModel: getCoreRowModel(),
    });
    return (
        <>
            <table className="table-fixed">
                <thead>
                    {table.getHeaderGroups().map((headerGroup) => {
                        return (
                            <tr key={headerGroup.id}>
                                {headerGroup.headers.map(
                                    (
                                        header, // map over the headerGroup headers array
                                    ) => (
                                        <th
                                            key={header.id}
                                            colSpan={header.colSpan}
                                        >
                                            {header.isPlaceholder
                                                ? null
                                                : flexRender(
                                                      header.column.columnDef
                                                          .header,
                                                      header.getContext(),
                                                  )}
                                        </th>
                                    ),
                                )}
                            </tr>
                        );
                    })}
                </thead>
                <tbody>
                    {table.getRowModel().rows.map((row) => (
                        <tr key={row.id}>
                            {row.getVisibleCells().map((cell) => (
                                <td key={cell.id}>
                                    {flexRender(
                                        cell.column.columnDef.cell,
                                        cell.getContext(),
                                    )}
                                </td>
                            ))}
                        </tr>
                    ))}
                </tbody>
                <tfoot>
                    {table.getFooterGroups().map((footerGroup) => (
                        <tr key={footerGroup.id}>
                            {footerGroup.headers.map((header) => (
                                <th key={header.id} colSpan={header.colSpan}>
                                    {header.isPlaceholder
                                        ? null
                                        : flexRender(
                                              header.column.columnDef.footer,
                                              header.getContext(),
                                          )}
                                </th>
                            ))}
                        </tr>
                    ))}
                </tfoot>
            </table>
        </>
    );
};
