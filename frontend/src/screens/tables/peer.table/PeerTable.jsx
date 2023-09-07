import React, { useState } from "react";
import Table from "../../../components/table/Table";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import { PeerService } from "../../../services/peer.service/Peer";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { ImportService } from "../../../services/import.service/Import";
import { ExportService } from "../../../services/export.service/Export";
import { saveAs } from 'file-saver';
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";

const PeerTable = () => {

    const queryClient = useQueryClient()


    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['peer'], () => PeerService.Get(),)

    const { mutate: mutateAdd, isLoading: isLoadingAdd, error: errorAdd } = useMutation(['add_peer'],
        (data) => PeerService.Add(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('peer')
            },
        })

    const { mutate: mutateDelete, isLoading: isLoadingDelete, error: errorDelete } = useMutation(
        (selectedValue) => PeerService.Delete(selectedValue),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('peer')
            },
        })

    const { mutate: mutateUpdate, isLoading: isLoadingUpdate, error: errorUpdate } = useMutation(
        (data) => PeerService.Update(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('peer')
            },
        })

    const { mutate: mutateImport, isLoading: isLoadingImport, error: errorImport } = useMutation(
        (data) => ImportService.Import("peers", data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('peer')
            },
        })

    const { mutate: mutateExport, isLoading: isLoadingExport, error: errorExport } = useMutation(
        (data) => ExportService.Export(data),
        {
            onSuccess: (data) => {
                const blob = new Blob([data], { type: 'text/csv' });
                saveAs(blob, 'peer.csv');
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
            <TableHeader data={data} setPattern={setPattern} field={"All Peers"} mutateAdd={mutateAdd} mutateImport={mutateImport} mutateExport={mutateExport} exportName={"peers"} />
            <Table data={data} pattern={pattern} mutateDelete={mutateDelete} mutateUpdate={mutateUpdate} />
        </>
    );
}

export default PeerTable;
