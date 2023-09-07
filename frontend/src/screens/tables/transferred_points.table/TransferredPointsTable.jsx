import React, { useState } from "react";
import Table from "../../../components/table/Table";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import { TransferredPointsService } from "../../../services/transferred_points.service/TransferredPoints";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { ImportService } from "../../../services/import.service/Import";
import { ExportService } from "../../../services/export.service/Export";
import { saveAs } from 'file-saver';
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";

const TransferredPointsTable = () => {

    const queryClient = useQueryClient()


    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['transferred_points'], () => TransferredPointsService.Get(),)

    const { mutate: mutateAdd, isLoading: isLoadingAdd, error: errorAdd } = useMutation(['add_transferred_points'],
        (data) => TransferredPointsService.Add(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('transferred_points')
            },
        })

    const { mutate: mutateDelete, isLoading: isLoadingDelete, error: errorDelete } = useMutation(
        (selectedValue) => TransferredPointsService.Delete(selectedValue),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('transferred_points')
            },
        })

    const { mutate: mutateUpdate, isLoading: isLoadingUpdate, error: errorUpdate } = useMutation(
        (data) => TransferredPointsService.Update(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('transferred_points')
            },
        })

    const { mutate: mutateImport, isLoading: isLoadingImport, error: errorImport } = useMutation(
        (data) => ImportService.Import("transferred_pointss", data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('transferred_points')
            },
        })

    const { mutate: mutateExport, isLoading: isLoadingExport, error: errorExport } = useMutation(
        (data) => ExportService.Export(data),
        {
            onSuccess: (data) => {
                const blob = new Blob([data], { type: 'text/csv' });
                saveAs(blob, 'transferred_points.csv');
            },
        })

    if (isLoading || isLoadingAdd || isLoadingUpdate || isLoadingDelete || isLoadingExport || isLoadingImport) return (
        <div>
            <Sidebar />
            <Loading />
        </div>
    )

    const errors = [errorAdd, errorDelete, errorExport, errorImport, errorUpdate, error];

    for (let i = 0; i < errors.length; i++) {
      if (errors[i]) {
        return (<ErrorResponse error={errors[i]} />)
      }
    }


    return (
        <>
            <Sidebar />
            <TableHeader data={data} setPattern={setPattern} field={"All Transferred Points"} mutateAdd={mutateAdd} mutateImport={mutateImport} mutateExport={mutateExport} exportName={"transferred_points"} />
            <Table data={data} pattern={pattern} mutateDelete={mutateDelete} mutateUpdate={mutateUpdate} />
        </>
    );
}

export default TransferredPointsTable;