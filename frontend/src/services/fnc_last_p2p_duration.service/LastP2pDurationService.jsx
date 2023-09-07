import axios from "axios"

export const LastP2pDurationService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/fnc/last_p2p_duration")
        return response.data
    },
}