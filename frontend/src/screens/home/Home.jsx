import React from "react";
import Sidebar from "../../components/sidebar/Sidebar";

const Home = () => {
  return (
    <div>
      <Sidebar />
      <div class="flex flex-col items-center justify-center">
        <h1 class="text-9xl mt-20 mb-20 font-semibold dark:text-blue text-blue-500">Team</h1>
      </div>
      <div class="p-4 sm:ml-64">
        <h1 class="mb-10 text-5xl font-semibold dark:text-white">Telvina - Backend developer</h1>
        <h1 class="text-5xl font-semibold dark:text-white">Adough - Frontend developer</h1>
      </div>
    </div>
  )
}

export default Home;

