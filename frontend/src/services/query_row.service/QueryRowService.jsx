import axios from "axios"

export const QueryRowService = {
    async Add(data) {
        const response = await axios.post("http://localhost:8080/api/v1/raw_query", data)
        return response.data
    },
}