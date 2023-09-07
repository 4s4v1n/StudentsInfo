import React, { useState } from "react";
import { useMutation, useQueryClient } from "react-query";
import ErrorResponse from "../error/ErrorResponse";
import Loading from "../../components/loading/Loading";
import Sidebar from "../../components/sidebar/Sidebar";
import TextArea from "../../components/text_area/TextArea";
import Table from "../../components/table/Table";
import TableHeader from "../../components/headers/table_header/TableHeader";
import { QueryRowService } from "../../services/query_row.service/QueryRowService";

const RawQuery = () => {
    const [pattern, setPattern] = useState('');
    const [data, setData] = useState(null);
    const queryClient = useQueryClient();

    const { isLoading, error, mutate } = useMutation(
        "raw_query",
        QueryRowService.Add,
        {
            onSuccess: (value) => {
                setData(value);
                queryClient.invalidateQueries("raw_query");
            },
        }
    );

    if (isLoading) {
        return (
            <>
                <Sidebar />
                <Loading />
            </>
        );
    }

    if (error) {
        return <ErrorResponse error={error} />;
    }

    return (
        <>
            <Sidebar />
            <TableHeader setPattern={setPattern} field={"Raw Query"} />
            <TextArea mutate={mutate} />
            {data && <Table data={data} pattern={pattern} />}
        </>
    );
};

export default RawQuery;