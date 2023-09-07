import React from "react";
import { Tables } from "./table.data";
import { Link } from "react-router-dom";

const ListItem = () => {

    return (
        <>
            <ul class="py-2 space-y-2">
                {
                    Tables.map((table => (
                        <Link key={table.Name} to={table.Route}>
                            <li>
                                <div key={table.Name} class="flex items-center w-full p-2 text-gray-900 transition duration-75 rounded-lg pl-11 group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700 text-sm">{table.Name}</div>
                            </li>
                        </Link>
                    )))
                }
            </ul>
        </>
    )
}

export default ListItem;