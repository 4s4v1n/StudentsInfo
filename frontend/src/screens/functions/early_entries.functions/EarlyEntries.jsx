import { useState } from "react";
import { useQuery } from "react-query";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import Table from "../../../components/table/Table";
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";
import { EarlyEntriesService } from "../../../services/fnc_early_entries.service/EarlyEntriesService";

const EarlyEntries = () => {

    const describe = "Determine for each month the percentage of early entries"
    
    const [pattern, setPattern] = useState('');

    const { data, isLoading, error } = useQuery(['fnc_early_entries'], () => EarlyEntriesService.Get());

    if (isLoading) {
        return (
            <div>
                <Sidebar />
                <Loading />
            </div>
        )
    }

    if (error) {
        return (<ErrorResponse error={error} />);
    }

    if (data === null) {
        return (
            <>
                <Sidebar />
                <TableHeader data={[]} setPattern={setPattern} field={describe} />
                <Table data={[]} pattern={pattern} />
            </>
        );
    }

    return (
        <>
            <Sidebar />
            <TableHeader data={data} setPattern={setPattern} field={describe} />
            <Table data={data} pattern={pattern} />
        </>
    );
};

export default EarlyEntries;
