import axios from "axios"

export const SuccessFailureChecksService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/fnc/success_failure_checks")
        return response.data
    },
}