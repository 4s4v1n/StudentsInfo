import axios from "axios"

export const LastFeastCameService = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/last_feast_came?date=${value.date}`)
        return response.data
    },
}