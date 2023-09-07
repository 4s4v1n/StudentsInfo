import axios from "axios"

export const MoreThenTimePeerService = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/more_then_time_peer?time=${value.time}`)
        return response.data
    },
}