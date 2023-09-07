import { useState } from "react";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import Table from "../../../components/table/Table";
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";
import { useQuery } from "react-query";
import { PreviousTasksService } from "../../../services/fnc_previous_tasks.service/PreviousTasksService";


const PreviousTasks = () => {

    const describe = "Using recursive common table expression, output the number of preceding tasks for each task"

    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['fnc_previous_tasks'], () => PreviousTasksService.Get(),)

    if (isLoading) return (
        <div>
            <Sidebar />
            <Loading />
        </div>
    )


    if (error) {
        return (<ErrorResponse error={error} />)
    }

    if (data === null) {
        return (
            <>
                <Sidebar />
                <TableHeader
                        data={[]}
                        setPattern={setPattern}
                        field={describe} 
                    />
                <Table data={[]} pattern={pattern} />
            </>
        )
    }

    return (
        <>
            <Sidebar />
            <TableHeader data={data} setPattern={setPattern} field={describe} />
            <Table data={data} pattern={pattern} />
        </>
    );
}

export default PreviousTasks;
