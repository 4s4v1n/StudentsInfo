import React from "react";
import Logo from "./logo/Logo";
import Items from "./items/Items";
import { Link } from "react-router-dom";

const Sidebar = () => {
  return (
    <div>
      <aside id="sidebar-multi-level-sidebar" class="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0" aria-label="Sidebar">
        <div class="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
          <Link to={"/"}>
            <Logo />
          </Link>
          <Items />
        </div>
      </aside>
    </div>
  )
}

export default Sidebar;
