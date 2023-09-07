import { useState } from "react";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import Table from "../../../components/table/Table";
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";
import { useQuery } from "react-query";
import { PeerXpSumService } from "../../../services/fnc_peer_xp_sum.service/PeerXpSumService";


const PeerXpSum = () => {

    const desribe = "Determine the total amount of XP gained by each peer"

    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['fnc_peer_xp_sum'], () => PeerXpSumService.Get(),)

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
                    field={desribe}
                />
                <Table data={[]} pattern={pattern} />
            </>
        )
    }

    return (
        <>
            <Sidebar />
            <TableHeader data={data} setPattern={setPattern} field={desribe} />
            <Table data={data} pattern={pattern} />
        </>
    );
}

export default PeerXpSum;