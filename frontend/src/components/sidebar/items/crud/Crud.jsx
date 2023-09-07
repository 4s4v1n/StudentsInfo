import React, { useState } from "react";
import { RxTable } from "react-icons/rx";
import { RxChevronDown } from "react-icons/rx";
import ListItem from "./list_item/ListItem";



const Crud = () => {
  const [isOpen, setIsOoen] = useState(false)
  return (
    <>
      <li>
        <button type="button" onClick={
          () => {
            setIsOoen((prev) => !prev)
          }
        } class="flex items-center w-full p-2 text-gray-900 transition duration-75 rounded-lg group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700">
          <RxTable class="w-6 h-6 fill-gray-700" />
          <span class="flex-1 ml-3 text-left whitespace-nowrap" sidebar-toggle-item>CRUD</span>
          <RxChevronDown class="w-6 h-6 fill-gray-700" />
        </button>
        {isOpen && (
          <ListItem/>
          )}
      </li>
    </>
  )
}

export default Crud;