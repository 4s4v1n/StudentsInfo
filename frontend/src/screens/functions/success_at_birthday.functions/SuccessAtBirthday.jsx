import { useState } from "react";
import TableHeader from "../../../components/headers/table_header/TableHeader";
import Sidebar from "../../../components/sidebar/Sidebar";
import Table from "../../../components/table/Table";
import Loading from "../../../components/loading/Loading";
import ErrorResponse from "../../error/ErrorResponse";
import { useQuery } from "react-query";
import { SuccessAtBirthdayService } from "../../../services/fnc_success_at_birthday.service/SuccessAtBirthdayService";


const SuccessAtBirthday = () => {

    const describe = "Determine the percentage of peers who have ever successfully passed a check on their birthday"

    const [pattern, setPattern] = useState('')

    const { data, isLoading, error } = useQuery(['fnc_success_at_birthday'], () => SuccessAtBirthdayService.Get(),)

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

export default SuccessAtBirthday;
