import React from "react";

const Header = ({ data, isExistButton }) => {
    return (
        <>
            {data.length ? (
                <tr>
                    {Object.keys(data[0]).map((key => (
                        <th key={key} scope="col" class="px-6 py-3">
                            {key}
                        </th>
                    )))}
                    {
                        isExistButton === true ?
                            <th scope="col" class="px-6 py-3">
                                Edit
                            </th> : null
                    }
                </tr>
            ) : (null)
            }
        </>
    );
};

export default Header;