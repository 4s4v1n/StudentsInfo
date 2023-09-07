import axios from "axios"

export const ExportService = {
    async Export(data) {
        const response = await axios.get(`http://localhost:8080/api/v1/export/${data}`)
        return response.data
    },
}