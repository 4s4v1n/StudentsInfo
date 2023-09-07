import { useState } from "react";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import Table from "../../../components/table/Table";
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";
import { useMutation, useQueryClient } from "react-query";
import FunctionModal from "../../../components/modal/function.modal/FunctionModal";
import { StatisticBlockService } from "../../../services/fnc_statistic_block.service/StatisticBlockService";

const StatisticBlock = () => {

    const describe = "Determine the percentage of peers who: Started block 1, Started block 2, Started both"

    const queryClient = useQueryClient();

    const [funcModalActive, setFuncModalActive] = useState(true);
    const [data, setData] = useState(null);
    const [pattern, setPattern] = useState("");

    const { mutate, isLoading, error } = useMutation(
        (values) => StatisticBlockService.Get(values),
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

    const queries = ["block1", "block2"];

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
                        <Sidebar />
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

export default StatisticBlock;