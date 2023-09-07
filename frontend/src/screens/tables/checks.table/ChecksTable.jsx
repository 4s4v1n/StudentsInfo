import React, { useState } from "react";
import Table from "../../../components/table/Table";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import { ChecksService } from "../../../services/checks.service/Checks";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { ImportService } from "../../../services/import.service/Import";
import { ExportService } from "../../../services/export.service/Export";
import { saveAs } from 'file-saver';
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";

const CheckTable = () => {

    const queryClient = useQueryClient()

    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['check'], () => ChecksService.Get(),)

    const { mutate: mutateAdd, isLoading: isLoadingAdd, error: errorAdd } = useMutation(['add_check'],
        (data) => ChecksService.Add(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('check')
            },
        })

    const { mutate: mutateDelete, isLoading: isLoadingDelete, error: errorDelete } = useMutation(
        (selectedValue) => ChecksService.Delete(selectedValue),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('check')
            },
        })

    const { mutate: mutateUpdate, isLoading: isLoadingUpdate, error: errorUpdate } = useMutation(
        (data) => ChecksService.Update(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('check')
            },
        })

    const { mutate: mutateImport, isLoading: isLoadingImport, error: errorImport } = useMutation(
        (data) => ImportService.Import("checks", data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('check')
            },
        })

    const { mutate: mutateExport, isLoading: isLoadingExport, error: errorExport } = useMutation(
        (data) => ExportService.Export(data),
        {
            onSuccess: (data) => {
                const blob = new Blob([data], { type: 'text/csv' });
                saveAs(blob, 'check.csv');
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
            return (
                <div>
                    <ErrorResponse error={errors[i]} />
                </div>
            )
        }
    }

    return (
        <div>
            <Sidebar />
            <TableHeader data={data} setPattern={setPattern} field={"All Check"} mutateAdd={mutateAdd} mutateImport={mutateImport} mutateExport={mutateExport} exportName={"checks"} />
            <Table data={data} pattern={pattern} mutateDelete={mutateDelete} mutateUpdate={mutateUpdate} />
        </div>
    );
}

export default CheckTable;