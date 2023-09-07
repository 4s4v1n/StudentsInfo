import React, { useState } from "react";

const TextArea = ({ mutate }) => {

    const [data, setData] = useState({});

    const CreateData = data => {
        mutate(data)
    }

    return (
        <div class="p-4 sm:ml-64 relative overflow-x-auto shadow-md sm:rounded-lg">
            <textarea
                rows="4"
                class="block mb-4 p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="Write sql query..."
                onChange={(e) => setData({"expression": e.target.value})}>
            </textarea>
            <div>
                <button
                    onClick={() => { CreateData(data); }}
                    type="button"
                    class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center mr-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                    Execute
                </button>
            </div>
        </div >
    );

};

export default TextArea;