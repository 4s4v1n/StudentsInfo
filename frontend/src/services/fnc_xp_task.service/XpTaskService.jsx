import axios from "axios"

export const XpTaskService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/fnc/xp_task")
        return response.data
    },
}