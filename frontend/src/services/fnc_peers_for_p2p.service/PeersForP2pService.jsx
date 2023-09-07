import axios from "axios"

export const PeersForP2pService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/fnc/peers_for_p2p")
        return response.data
    },
}