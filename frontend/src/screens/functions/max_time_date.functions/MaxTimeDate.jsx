import { useState } from "react";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import Table from "../../../components/table/Table";
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";
import { useMutation, useQueryClient } from "react-query";
import FunctionModal from "../../../components/modal/function.modal/FunctionModal";
import { MaxTimeDateService } from "../../../services/fnc_max_time_date.service/MaxTimeDateService";

const MaxTimeDate = () => {

    const describe = "Find the peer who spent the longest amount of time on campus today"

    const queryClient = useQueryClient();

    const [funcModalActive, setFuncModalActive] = useState(true);
    const [data, setData] = useState(null);
    const [pattern, setPattern] = useState("");

    const { mutate, isLoading, error } = useMutation(
        (values) => MaxTimeDateService.Get(values),
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

    const queries = ["date"];

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

export default MaxTimeDate;