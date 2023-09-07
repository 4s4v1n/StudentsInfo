import axios from "axios"

export const PeerXpSumService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/fnc/peer_xp_sum")
        return response.data
    },
}