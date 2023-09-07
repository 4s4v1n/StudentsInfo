import axios from "axios"

class Task { 
    constructor(title, parent_task, max_xp) { 
      this.title = title 
      this.parent_task = parent_task 
      this.max_xp = parseInt(max_xp, 10); 
    }
  }

export const TaskService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/task")
        return response.data
    },

    async Add(data) {
        const task = new Task(data.title, data.parent_task, data.max_xp)
        return await axios.post("http://localhost:8080/api/v1/task", task)
    },

    async Delete(data) {
        return await axios.delete(`http://localhost:8080/api/v1/task/${data}`)
    },

    async Update(data) {
        const task = new Task(data.title, data.parent_task, data.max_xp)
        return await axios.patch("http://localhost:8080/api/v1/task", task)
    },
}