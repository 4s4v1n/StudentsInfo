import axios from "axios"

export const PeerMostXpService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/fnc/peer_most_xp")
        return response.data
    },
}