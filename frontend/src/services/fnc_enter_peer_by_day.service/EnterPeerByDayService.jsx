import axios from "axios"

export const EnterPeerByDayService = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/enter_peer_by_day?n=${value.n}&m=${value.m}`)
        return response.data
    },
}