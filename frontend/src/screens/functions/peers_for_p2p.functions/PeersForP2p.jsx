import { useState } from "react";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import Table from "../../../components/table/Table";
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";
import { useQuery } from "react-query";
import { PeersForP2pService } from "../../../services/fnc_peers_for_p2p.service/PeersForP2pService";


const PeersForP2p = () => {

    const describe = "Determine which peer each student should go to for a check"

    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['fnc_peers_for_p2p'], () => PeersForP2pService.Get(),)

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

export default PeersForP2p;