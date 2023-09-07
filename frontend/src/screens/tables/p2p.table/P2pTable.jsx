import React, { useState } from "react";
import Table from "../../../components/table/Table";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import { P2pService } from "../../../services/p2p.service/P2p";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { ImportService } from "../../../services/import.service/Import";
import { ExportService } from "../../../services/export.service/Export";
import { saveAs } from 'file-saver';
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";

const P2pTable = () => {

    const queryClient = useQueryClient()


    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['p2p'], () => P2pService.Get(),)

    const { mutate: mutateAdd, isLoading: isLoadingAdd, error: errorAdd } = useMutation(['add_p2p'],
        (data) => P2pService.Add(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('p2p')
            },
        })

    const { mutate: mutateDelete, isLoading: isLoadingDelete, error: errorDelete } = useMutation(
        (selectedValue) => P2pService.Delete(selectedValue),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('p2p')
            },
        })

    const { mutate: mutateUpdate, isLoading: isLoadingUpdate, error: errorUpdate } = useMutation(
        (data) => P2pService.Update(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('p2p')
            },
        })

    const { mutate: mutateImport, isLoading: isLoadingImport, error: errorImport } = useMutation(
        (data) => ImportService.Import("p2ps", data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('p2p')
            },
        })

    const { mutate: mutateExport, isLoading: isLoadingExport, error: errorExport } = useMutation(
        (data) => ExportService.Export(data),
        {
            onSuccess: (data) => {
                const blob = new Blob([data], { type: 'text/csv' });
                saveAs(blob, 'p2p.csv');
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
            <TableHeader data={data} setPattern={setPattern} field={"All P2P"} mutateAdd={mutateAdd} mutateImport={mutateImport} mutateExport={mutateExport} exportName={"p2p"} />
            <Table data={data} pattern={pattern} mutateDelete={mutateDelete} mutateUpdate={mutateUpdate} />
        </div>
    );
}

export default P2pTable;