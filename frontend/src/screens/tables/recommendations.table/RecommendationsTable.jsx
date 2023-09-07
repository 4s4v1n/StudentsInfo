import React, { useState } from "react";
import Table from "../../../components/table/Table";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import { RecommendationService } from "../../../services/recommendations.servirse/Recommendations";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { ImportService } from "../../../services/import.service/Import";
import { ExportService } from "../../../services/export.service/Export";
import { saveAs } from 'file-saver';
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";

const RecommendationsTable = () => {

    const queryClient = useQueryClient()


    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['recommendation'], () => RecommendationService.Get(),)

    const { mutate: mutateAdd, isLoading: isLoadingAdd, error: errorAdd } = useMutation(['add_recommendation'],
        (data) => RecommendationService.Add(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('recommendation')
            },
        })

    const { mutate: mutateDelete, isLoading: isLoadingDelete, error: errorDelete } = useMutation(
        (selectedValue) => RecommendationService.Delete(selectedValue),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('recommendation')
            },
        })

    const { mutate: mutateUpdate, isLoading: isLoadingUpdate, error: errorUpdate } = useMutation(
        (data) => RecommendationService.Update(data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('recommendation')
            },
        })

    const { mutate: mutateImport, isLoading: isLoadingImport, error: errorImport } = useMutation(
        (data) => ImportService.Import("recommendations", data),
        {
            onSuccess: () => {
                queryClient.invalidateQueries('recommendation')
            },
        })

    const { mutate: mutateExport, isLoading: isLoadingExport, error: errorExport } = useMutation(
        (data) => ExportService.Export(data),
        {
            onSuccess: (data) => {
                const blob = new Blob([data], { type: 'text/csv' });
                saveAs(blob, 'recommendation.csv');
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
            <TableHeader data={data} setPattern={setPattern} field={"All Recommendations"} mutateAdd={mutateAdd} mutateImport={mutateImport} mutateExport={mutateExport} exportName={"recommendations"} />
            <Table data={data} pattern={pattern} mutateDelete={mutateDelete} mutateUpdate={mutateUpdate} />
        </>
    );
}

export default RecommendationsTable;