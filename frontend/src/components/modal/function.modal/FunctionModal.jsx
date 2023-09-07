import { TextInput } from "flowbite-react";
import React, { useState } from "react";
import { RxCross2 } from "react-icons/rx";

const FunctionModal = ({ active, setActive, value, mutateFunction }) => {

    const [formData, setFormData] = useState({});

    const FuctionData = () => {
        mutateFunction(formData)
    }

    return (
        <>
            {
                active ? (
                    <div class="flex justify-center items-center overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none">
                        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"></div>
                        <div class="relative w-full max-w-md max-h-full">
                            <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
                                <button onClick={() => setActive(false)} type="button" class="absolute top-3 right-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-800 dark:hover:text-white" data-modal-hide="authentication-modal">
                                    <RxCross2 class="w-5 h-5 fill-white" />
                                </button>
                                <div class="px-6 py-6 lg:px-8">
                                    <h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">Set Value</h3>
                                    <form class="space-y-6" action="#">
                                        {
                                            value.map((key) => (
                                                <div>
                                                    <label class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">{key}</label>
                                                    <TextInput
                                                        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                                                        placeholder={key}
                                                        onChange={(e) => setFormData({ ...formData, [key]: e.target.value })}>
                                                    </TextInput>
                                                </div>
                                            ))
                                        }
                                        <button onClick={() => { FuctionData(); setActive(false) }} type="submit" class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Send</button>
                                    </form>
                                </div>
                            </div>
                        </div >
                    </div >
                ) : null
            }
        </>
    );
};

export default FunctionModal;