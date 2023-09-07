import axios from "axios"

export const PeersDontLeaveService = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/peers_dont_leave?date=${value.date}`)
        return response.data
    },
}