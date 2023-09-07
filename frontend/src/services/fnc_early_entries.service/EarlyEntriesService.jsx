import axios from "axios"

export const EarlyEntriesService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/fnc/early_entries")
        return response.data
    },
}