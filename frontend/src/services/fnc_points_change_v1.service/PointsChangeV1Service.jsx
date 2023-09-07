import axios from "axios"

export const PointsChangeV1Service = {
    async Get() {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/points_change_v1`)
        return response.data
    },
}