# 🧭 Task Tracker CLI

A simple command-line task tracker written in **Golang**, allowing you to manage tasks directly from your terminal.  
You can **add**, **update**, **delete**, and **mark tasks** as `todo`, `in-progress`, or `done`.  
All tasks are stored persistently in a local JSON file — no external libraries required!

---

## 🚀 Features

- ✅ Add new tasks  
- ✏️ Update existing tasks  
- ❌ Delete tasks  
- 🔄 Mark tasks as **in-progress** or **done**  
- 📋 List all tasks or filter by status (`todo`, `in-progress`, `done`)  
- 💾 Stores data locally in `tasks.json`  
- 🛡️ Graceful error handling for invalid commands and missing IDs  

---

## 🧩 Task Properties

Each task in the JSON file includes:

| Property   | Description |
|-------------|-------------|
| `id`        | Unique identifier (integer) |
| `description` | Short description of the task |
| `status`    | One of `todo`, `in-progress`, or `done` |
| `createdAt` | Date and time when the task was created |
| `updatedAt` | Date and time when the task was last updated |

Example JSON structure:

```json
[
  {
    "id": 1,
    "description": "Buy groceries",
    "status": "todo",
    "createdAt": "2025-11-10T18:32:00",
    "updatedAt": "2025-11-10T18:32:00"
  }
]
