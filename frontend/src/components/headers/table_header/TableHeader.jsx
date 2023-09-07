import { useState } from "react";
import Search from "../../search/Search";
import { RxPlus } from "react-icons/rx";
import { RxShare2 } from "react-icons/rx";
import { RxDownload } from "react-icons/rx";
import AddModal from "../../modal/add_modal/AddModal";
import Papa from 'papaparse';

const TableHeader = ({ data, setPattern, field, mutateAdd, mutateImport, mutateExport, exportName }) => {
    const [addModalActive, setAddModalActive] = useState(false)

    const handleFileDownload = async (event) => {
        const file = event.target.files[0];
        Papa.parse(file, {
            delimiter: ";",
            header: true,
            dynamicTyping: true,
            skipEmptyLines: true,
            complete: (results) => {
                mutateImport(results.data)
                event.target.value = null;
            },
        });
    };

    function importFile() {
        document.getElementById("fileImport").click();
    }

    const ExportData = data => {
        mutateExport(data)
    }

    return (
        <>
            {
                data ? (
                    <>
                        <div class="p-4 sm:ml-64 -mb-4">
                            <h2 class="text-2xl font-semibold dark:text-white">{field}</h2>
                        </div>
                        <div class="flex justify-between p-4 sm:ml-64">
                            <Search setPattern={setPattern} />
                            <div class="flex justify-end">
                                {
                                    mutateAdd ?
                                        <div>
                                            <button onClick={() => setAddModalActive(true)} type="button" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center mr-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                                                <RxPlus class="w-5 h-5 fill-white mr-3" />
                                                Add
                                            </button>
                                            <AddModal active={addModalActive} setActive={setAddModalActive} value={data} mutateAdd={mutateAdd} />
                                        </div> : null
                                }
                                {
                                    mutateImport ?
                                        <div>
                                            <button
                                                class="text-blue-700 hover:text-white border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center mr-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                                                onClick={() => { importFile(); }}>
                                                <RxDownload class="w-5 h-5 fill-white mr-3" />
                                                Import
                                            </button>
                                            <input id="fileImport" type="file" class="hidden" accept=".csv" onChange={handleFileDownload} />
                                        </div> : null
                                }

                                {
                                    mutateExport ?
                                        <div>
                                            <button
                                                class="text-blue-700 hover:text-white border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center mr-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                                                onClick={() => { ExportData(exportName) }}>
                                                <RxShare2 class="w-5 h-5 fill-white mr-3" />
                                                Export
                                            </button>
                                        </div> : null
                                }
                            </div>
                        </div >

                    </>
                ) : (
                    <>
                        <div class="p-4 sm:ml-64 -mb-4">
                            <h2 class="text-2xl font-semibold dark:text-white">{field}</h2>
                        </div>
                        <div class="flex justify-between p-4 sm:ml-64">
                            <Search setPattern={setPattern} />
                        </div >
                    </>
                )
            }
        </>
    );
}


export default TableHeader