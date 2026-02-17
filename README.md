# Task
## 📝 To-do CLI written in Go
This app was made as a beginner project
# Task — emoji overload edition 🎉📝🚀

Welcome to the most emoji-friendly tiny to-do CLI you'll ever meet. This repo contains a minimal Go-based command-line task manager intended for learning, tinkering, and smiling while you code. 😄

--

Why this project exists
-: Simple learning project for Go beginners 🐣
-: Small, focused codebase — easy to read and modify 🔍
-: Because TODOs deserve happiness too 🎈

--

Table of contents
- **What it is**
- **Install**
- **Usage examples**
- **Commands**
- **File format & storage**
- **Development & contributing**
- **Fun ideas / Roadmap**

--

What it is
-----------

`task` is a tiny command-line to-do list app written in Go. It stores tasks as JSON on disk and exposes basic operations like add, edit, delete, toggle (done/undone), list, and clear. All commands are simple and fast — perfect for demos, learning, or personal use. 🧰✨

Install
-------

Prerequisites
- Install Go: https://go.dev/doc/install ✅

Clone & build

```bash
git clone https://github.com/dacixn/task.git
cd task
go build .
```

Run

```bash
./task
```

Usage examples (quick start) 🚦

- Add a task:

```bash
./task add "Buy groceries 🥦🍞"
```

- List tasks:

```bash
./task list
# 1. Buy groceries 🥦🍞
```

- Mark task done/undone:

```bash
./task done 1    # toggle completion for task 1 ✅↔️⬜️
```

- Edit task text:

```bash
./task edit 1 "Buy groceries and cat food 🛒🐱"
```

- Delete task:

```bash
./task del 1
```

- Clear all tasks (interactive confirmation):

```bash
./task clear
# prompts: Clear task file? (y/N):
```

Commands reference 🧭

- `help` — show help text
- `add [text]` — add a task
- `edit <id> <text>` — edit a task
- `del <id>` — delete a task by id
- `done <id>` — toggle task completion
- `list` — list all tasks
- `clear` — clear all tasks (asks for confirmation)

File format & storage 📦

Tasks are stored in JSON at `~/.task.json` by default. Example content:

```json
[ { "text": "Buy milk 🥛", "done": false } ]
```

This file is small, human-readable, and easy to manipulate if you want to do something fancy.

Development & contributing 🛠️

- The code is intentionally tiny and straightforward — perfect for modification.
- Feel free to add features, open pull requests, or just hack locally.

Suggested development steps

1. Fork the repo 🍴
2. Make changes 📐
3. Run `go build` and test locally 🧪
4. Open a PR with a friendly description ✍️

Fun ideas / Roadmap 🌈

- Add timestamps & due dates ⏰
- Add categories or tags 🏷️
- Add colored, pretty output (terminal UI) 🎨
- Sync across devices (experimental) ☁️🔁

Troubleshooting & notes 🔎

- First run: no `~/.task.json` file exists — that's fine. The app will create it when you add or save tasks. 🌱
- If you see JSON errors, check `~/.task.json` and remove/repair malformed content. 🧰

License & attribution 🧾

This tiny project is free to use for learning and tinkering. No heavy license attached — treat it kindly and share improvements. 🤝

Final words
-----------

Thanks for checking out this emoji-infused tiny to-do app! Whether you're learning Go, experimenting with CLI tools, or just collecting cute task managers, I hope this brings a smile to your workflow. 😺🚀🎉

Go nuts — add emojis, features, and flair. This repo is your playground. 🎡🎨
