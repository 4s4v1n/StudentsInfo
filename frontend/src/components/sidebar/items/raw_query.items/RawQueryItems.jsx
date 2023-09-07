import React from "react";
import { Link } from "react-router-dom";
import { RxPencil2 } from "react-icons/rx";


const RawQueriesItems = () => {
    return (
        <>
            <li>
                <Link key="Raw Query" to="/raw_query">
                <button type="button"
                    class="flex items-center w-full p-2 text-gray-900 transition duration-75 rounded-lg group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700">
                    <RxPencil2 class="w-6 h-6 fill-gray-700" />
                    <span class="flex-1 ml-3 text-left whitespace-nowrap" sidebar-toggle-item>Raw Query</span>
                </button>
                </Link>
            </li>
        </>
    )
}

export default RawQueriesItems;