import axios from "axios"

export const OftenTaskPerDayService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/fnc/often_task_per_day")
        return response.data
    },
}