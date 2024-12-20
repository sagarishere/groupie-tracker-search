# Groupie Tracker Visualizer

## Audit

<https://github.com/01-edu/public/blob/master/subjects/groupie-tracker/visualizations/audit.md>

---

## How to run ( with Docker ) - preferred method

```bash
    docker-compose up
```

go to <http://localhost>

---

## How to run without docker

```bash
    cd backend
    go run .
```

Before running the frontend, make sure to change the `API_URL` in `frontend/stores/band.js` at line number 23 from
`http://localhost/data` to `http://localhost:8080/data`

```bash
    cd frontend
    npm install
    npm run build && npm run start
```

go to <http://localhost:3000>

---

## For frequent users

toggle dark-mode/light-mode via keyboard shortcut `ctrl + shift + t`

## Authors

- [Sagar Yadav](https://github.com/sagarishere)
- Göran : reverse proxy
