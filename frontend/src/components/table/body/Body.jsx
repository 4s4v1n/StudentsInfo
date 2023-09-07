import React from "react";

const Body = ({ data, setSelectedValueUpdate, setSelectedValueDelete, setDeleteModalActive, setUpdateModalActive, filteredData, isExistButton }) => {

    const handleRemoveClick = (value) => {
        setSelectedValueDelete(value);
        setDeleteModalActive(true);
    }

    const handleUpdateClick = (value) => {
        setSelectedValueUpdate(value);
        setUpdateModalActive(true);
    }

    return (
        <>
            <tbody>
                {
                    data.length ? (
                        filteredData.map((value => (
                            <tr
                                class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                                {
                                    Object.values(value).map((val => (
                                        typeof (val) == "boolean" ? (
                                            val === true ? (
                                                <td class="px-6 py-4">
                                                    <div class="flex items-center">
                                                        <div class="h-2.5 w-2.5 rounded-full bg-blue-500 mr-2"></div>
                                                    </div>
                                                </td>
                                            ) : (
                                                <td class="px-6 py-4">
                                                    <div class="flex items-center">
                                                        <div class="h-2.5 w-2.5 rounded-full bg-red-500 mr-2"></div>
                                                    </div>
                                                </td>
                                            )
                                        ) : (
                                            <td key={val} class="px-6 py-4">
                                                {val === null ? "-" : val}
                                            </td>
                                        )
                                    )))
                                }
                                {
                                    isExistButton === true ?
                                        <td class="flex items-center gap-x-3 whitespace-nowrap px-6 py-4">
                                            <button onClick={() => { handleUpdateClick(Object.values(value)[0]) }} class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">Update</button>
                                            <button onClick={() => { handleRemoveClick(Object.values(value)[0]) }} class="text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">Remove</button>
                                        </td> : null
                                }
                            </tr>
                        )))
                    ) : <div>
                        Data empty
                    </div>
                }
            </tbody>
        </>
    );
};

export default Body;