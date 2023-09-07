import React, { useState } from "react";
import Table from "../../../components/table/Table";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import { VerterService } from "../../../services/verter.service/Verter";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { ImportService } from "../../../services/import.service/Import";
import { ExportService } from "../../../services/export.service/Export";
import { saveAs } from 'file-saver';
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";

const VerterTable = () => {

    const queryClient = useQueryClient()


    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['verter'], () => VerterService.Get(),)

    const { mutate: mutateAdd, isLoading: isLoadingAdd, error: errorAdd } = useMutation(['add_verter'],
        (data) => VerterService.Add(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('verter')
            },
        })

    const { mutate: mutateDelete, isLoading: isLoadingDelete, error: errorDelete } = useMutation(
        (selectedValue) => VerterService.Delete(selectedValue),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('verter')
            },
        })

    const { mutate: mutateUpdate, isLoading: isLoadingUpdate, error: errorUpdate } = useMutation(
        (data) => VerterService.Update(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('verter')
            },
        })

    const { mutate: mutateImport, isLoading: isLoadingImport, error: errorImport } = useMutation(
        (data) => ImportService.Import("verters", data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('verter')
            },
        })

    const { mutate: mutateExport, isLoading: isLoadingExport, error: errorExport } = useMutation(
        (data) => ExportService.Export(data),
        {
            onSuccess: (data) => {
                const blob = new Blob([data], { type: 'text/csv' });
                saveAs(blob, 'verter.csv');
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
            <TableHeader data={data} setPattern={setPattern} field={"All Verter"} mutateAdd={mutateAdd} mutateImport={mutateImport} mutateExport={mutateExport} exportName={"verter"} />
            <Table data={data} pattern={pattern} mutateDelete={mutateDelete} mutateUpdate={mutateUpdate} />
        </div>
    );
}

export default VerterTable;