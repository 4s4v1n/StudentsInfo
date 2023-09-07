import axios from "axios"

export const TimePeerByTimeService = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/time_peer_by_time?time=${value.time}&n=${value.n}`)
        return response.data
    },
}