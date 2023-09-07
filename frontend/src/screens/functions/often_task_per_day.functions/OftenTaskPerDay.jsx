import { useState } from "react";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import Table from "../../../components/table/Table";
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";
import { useQuery } from "react-query";
import { OftenTaskPerDayService } from "../../../services/fnc_often_task_per_day.service/OftenTaskPerDayService";


const OftenTaskPerDay = () => {

    const describe = "Find the most frequently checked task for each day"

    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['fnc_often_task_per_day'], () => OftenTaskPerDayService.Get(),)

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

export default OftenTaskPerDay;
