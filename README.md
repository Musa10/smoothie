# Smoothie 🚀

[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go&logoColor=white)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

**Smoothie** is a lightweight workflow engine built in **Go**. It executes workflows defined as JSON, fully compatible with **React Flow**, enabling easy integration with frontend visual workflow editors.

> Mix and run workflows smoothly with Go.

---

## Features

- Execute any workflow exported from React Flow.
- Supports sequential task execution (`start`, `task`, `end` nodes).
- Async workflow execution.
- Lightweight, Go-native backend engine.
- Simple HTTP API to start workflows.

---

## Architecture

React Flow JSON
│
▼ POST /workflows/run
Smoothie Backend (Go)
│ Executes nodes
▼
Console Logs


- **Workflow Engine**: Parses nodes and edges, executes tasks sequentially.
- **HTTP API**: Accepts React Flow JSON via POST requests.
- **In-memory Execution**: Tracks node status (`pending`, `running`, `completed`) without external DB (optional later).

---

## Getting Started

### Prerequisites

- Go 1.20+
- curl or Postman for testing

### Run the Backend

```bash
cd cmd/server
go run main.go


### Json Exmple
{
  "nodes": [
    {"id": "1", "type": "start", "data": {"label": "Start"}},
    {"id": "2", "type": "task", "task": "send_email", "data": {"label": "Send Email"}},
    {"id": "3", "type": "end", "data": {"label": "End"}}
  ],
  "edges": [
    {"from": "1", "to": "2"},
    {"from": "2", "to": "3"}
  ]
}
