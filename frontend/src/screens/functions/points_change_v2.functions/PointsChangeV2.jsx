import { useState } from "react";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import Table from "../../../components/table/Table";
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";
import { useQuery } from "react-query";
import { PointsChangeV2Service } from "../../../services/fnc_points_change_v2.service/PointsChangeV2Service";


const PointsChangeV2 = () => {

    const describe = "Calculate the change in the number of peer points of each peer using the table returned by the first function from Part 3"

    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['fnc_points_change_v2'], () => PointsChangeV2Service.Get(),)

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

export default PointsChangeV2;
