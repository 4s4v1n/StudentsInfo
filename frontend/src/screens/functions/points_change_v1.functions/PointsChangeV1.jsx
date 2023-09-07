import { useState } from "react";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import Table from "../../../components/table/Table";
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";
import { useQuery } from "react-query";
import { PointsChangeV1Service } from "../../../services/fnc_points_change_v1.service/PointsChangeV1Service";


const PointsChangeV1 = () => {

    const describe = "Calculate the change in the number of peer points of each peer using the TransferredPoints table"

    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['fnc_points_change_v1'], () => PointsChangeV1Service.Get(),)

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

export default PointsChangeV1;
