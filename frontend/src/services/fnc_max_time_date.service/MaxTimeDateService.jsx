import axios from "axios"

export const MaxTimeDateService = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/max_time_date?date=${value.date}`)
        return response.data
    },
}