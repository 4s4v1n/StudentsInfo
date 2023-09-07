import axios from "axios"

class TimeTracking {
    constructor(data) {
        this.peer = data.peer
        this.date = data.date
        this.time = data.time
        this.state = parseInt(data.state, 10);
        if (data.hasOwnProperty('id')) {
            this.id = data.id
        }
    }
}

export const TimeTrackingService = {

    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/time_tracking")
        return response.data
    },

    async Add(data) {
        const timeTracking = new TimeTracking(data)
        return await axios.post("http://localhost:8080/api/v1/time_tracking", timeTracking)
    },

    async Delete(data) {
        return await axios.delete(`http://localhost:8080/api/v1/time_tracking/${data}`)
    },

    async Update(data) {
        const timeTracking = new TimeTracking(data)
        return await axios.patch("http://localhost:8080/api/v1/time_tracking", timeTracking)
    },
}