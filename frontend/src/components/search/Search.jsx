import { Label, TextInput } from "flowbite-react";
import React from "react";

const Search = ({ setPattern }) => {
    return (
        <div>
            <form class="lg:pr-3">
                <Label htmlFor="search" className="sr-only">
                    Search
                </Label>
                <div class="relative lg:w-64 xl:w-96">
                    <TextInput
                        id="search"
                        name="search"
                        placeholder="Search for data"
                        onChange={(event) => setPattern(event.target.value)}
                    />
                </div>
            </form>
        </div>
    );
};

export default Search;