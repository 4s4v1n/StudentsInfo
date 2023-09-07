import axios from "axios"

export const TransferredPointsService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/fnc/transferred_points")
        return response.data
    },
}