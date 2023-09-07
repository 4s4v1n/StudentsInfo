import { useState } from "react";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import Table from "../../../components/table/Table";
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";
import { useMutation, useQueryClient } from "react-query";
import FunctionModal from "../../../components/modal/function.modal/FunctionModal";
import { EnterPeerByDayService } from "../../../services/fnc_enter_peer_by_day.service/EnterPeerByDayService";

const EnterPeerByDay = () => {

    const describe = "Determine the peers who left the campus more than M times during the last N days"

    const queryClient = useQueryClient();

    const [funcModalActive, setFuncModalActive] = useState(true);
    const [data, setData] = useState(null);
    const [pattern, setPattern] = useState("");

    const { mutate, isLoading, error } = useMutation(
        (values) => EnterPeerByDayService.Get(values),
        {
            onSuccess: (result) => {
                setData(result);
                queryClient.invalidateQueries("check");
            },
        }
    );

    if (isLoading) {
        return (
            <div>
                <Sidebar />
                <Loading />
            </div>
        );
    }

    const queries = ["n", "m"];

    if (error) {
        return <ErrorResponse error={error} />;
    }

    return (
        <>
            <Sidebar />
            <FunctionModal
                active={funcModalActive}
                setActive={setFuncModalActive}
                value={queries}
                mutateFunction={mutate}
            />
            {
                data ? (
                    <div>
                        <TableHeader
                            data={data}
                            setPattern={setPattern}
                            field={describe}
                        />
                        <Table data={data} pattern={pattern} />
                    </div>
                ) :
                    <>
                        <TableHeader
                            data={[]}
                            setPattern={setPattern}
                            field={describe}
                        />
                        <Table data={[]} pattern={pattern} />
                    </>
            }
        </>
    );
};

export default EnterPeerByDay;