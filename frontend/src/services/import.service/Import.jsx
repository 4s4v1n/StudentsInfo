import axios from "axios"

export const ImportService = {
    async Import(value, data) {
        return await axios.post(`http://localhost:8080/api/v1/import/${value}`, data)
    },
}