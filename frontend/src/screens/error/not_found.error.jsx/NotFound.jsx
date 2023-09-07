import React from "react";
import { Link } from "react-router-dom";

const ErrorNotFound = ({ error }) => {
    return (
        <div class="bg-gray-200 w-full px-16 md:px-0 h-screen flex items-center justify-center">
            <div class="bg-white border border-gray-200 flex flex-col items-center justify-center px-4 md:px-8 lg:px-24 py-8 rounded-lg shadow-2xl">
                <div class="text-6xl md:text-7xl lg:text-9xl font-bold tracking-wider text-gray-300">404</div>
                <div class="text-2xl md:text-3xl lg:text-5xl font-bold tracking-wider text-gray-500 mt-4">Page Not Found</div>
                <Link to="/">
                    <button type="button" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800 my-6">Return Home</button>
                </Link>
            </div>
        </div>
    );
}

export default ErrorNotFound;