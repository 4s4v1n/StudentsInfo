import React from "react";
import Functions from "./functions/Functions";
import Crud from "./crud/Crud";
import RawQueriesItems from "./raw_query.items/RawQueryItems";

const Items = () => {
  return (
    <>
      <ul class="space-y-2 font-medium">
        <Crud />
        <Functions />
        <RawQueriesItems />
      </ul>
    </>
  )
}

export default Items;