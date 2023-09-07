import axios from "axios"

export const MostFriendlyService = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/most_friendly?n=${value.n}`)
        return response.data
    },
}