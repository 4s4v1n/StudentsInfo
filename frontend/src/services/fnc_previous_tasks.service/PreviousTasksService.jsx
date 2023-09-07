import axios from "axios"

export const PreviousTasksService = {
    async Get() {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/previous_tasks`)
        return response.data
    },
}