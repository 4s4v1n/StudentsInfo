import axios from "axios"

export const SuccessfulDaysService = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/successful_days?n=${value.n}`)
        return response.data
    },
}