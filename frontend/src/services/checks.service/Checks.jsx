import axios from "axios"

export const ChecksService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/check")

        return response.data
    },

    async Add(data) {
        return await axios.post("http://localhost:8080/api/v1/check", data)
    },

    async Delete(data) {
        return await axios.delete(`http://localhost:8080/api/v1/check/${data}`)
    },

    async Update(data) {
        return await axios.patch("http://localhost:8080/api/v1/check", data)
    },
}