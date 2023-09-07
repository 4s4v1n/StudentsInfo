import React, { useState } from "react";
import Header from "./header/Header";
import UpdateModal from "../modal/update_modal/UpdateModal";
import DeleteModal from "../modal/delete_modal/DeleteModal";
import Body from "./body/Body";

const Table = ({ data, pattern, mutateDelete, mutateUpdate }) => {

    const [updateModalActive, setUpdateModalActive] = useState(false)
    const [deleteModalActive, setDeleteModalActive] = useState(false)
    const [selectedValueUpdate, setSelectedValueUpdate] = useState(null);
    const [selectedValueDelete, setSelectedValueDelete] = useState(null);

    const filteredData = data.filter((val) => {
        if (!pattern) return true;
        return Object.values(val).some((value) =>
            value && value.toString().toLowerCase().includes(pattern.toLowerCase())
        );
    });


    const isExistButton = () => {
        return (mutateDelete !== undefined && mutateUpdate !== undefined) ? true : false;
    }

    return (
        <div>
            <div class="p-4 sm:ml-64 relative overflow-x-auto shadow-md sm:rounded-lg">
                <table class="w-full text-base text-left text-gray-500 dark:text-gray-400">
                    <thead class="text-sm text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <Header data={data} isExistButton={isExistButton()} />
                    </thead>
                    <Body data={data} setSelectedValueUpdate={setSelectedValueUpdate} setSelectedValueDelete={setSelectedValueDelete} setDeleteModalActive={setDeleteModalActive} setUpdateModalActive={setUpdateModalActive} filteredData={filteredData} isExistButton={isExistButton()} />
                </table>
                <UpdateModal active={updateModalActive} setActive={setUpdateModalActive} value={data} mutateUpdate={mutateUpdate} selectedValue={selectedValueUpdate} />
                <DeleteModal active={deleteModalActive} setActive={setDeleteModalActive} mutateDelete={mutateDelete} selectedValue={selectedValueDelete} />
            </div>
        </div>
    );
}

export default Table;