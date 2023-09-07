import axios from "axios"

export const SuccessAtBirthdayService = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/success_at_birthday`)
        return response.data
    },
}