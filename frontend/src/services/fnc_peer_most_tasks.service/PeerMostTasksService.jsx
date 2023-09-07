import axios from "axios"

export const PeerMostTasksService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/fnc/peer_most_tasks")
        return response.data
    },
}